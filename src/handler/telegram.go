package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/golang/glog"

	"github.com/julienschmidt/httprouter"
)

// Telegram 架構
type Telegram struct {
	sem chan struct{}
}

// MessageBody 訊息內容
type MessageBody struct {
	ChatID  interface{} `json:"chat_id,string"`
	Message string      `json:"text,string"`
	Slient  bool        `json:"disable_notification,bool"`
}

// NewTelegram Telegram 實體
func NewTelegram() *Telegram {
	return &Telegram{}
}

// SetSem 設置 Semaphore
func (t *Telegram) SetSem(sem *chan struct{}) {
	t.sem = *sem
}

// SendMessage 傳送訊息
func (t *Telegram) SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.sem <- struct{}{}

	b, e := ioutil.ReadAll(r.Body)
	defer func() {
		_, e = io.Copy(ioutil.Discard, r.Body)
		if e != nil {
			glog.Error(e)
		}
		e = r.Body.Close()
		if e != nil {
			glog.Error(e)
		}
	}()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}

	// REQUEST 轉 JSON
	req, e := simplejson.NewJson(b)
	if e != nil {
		glog.Error(e)
	}

	// 取得 頻道ＩＤ
	code, e := req.Get("code").String()
	if e != nil {
		glog.Error(e)
	}
	channel := os.Getenv("TG_CHAN_" + code)
	if channel == "" {
		fmt.Fprint(w, `{"ok":false,"message":"code error."}`)
		return
	}
	var chanID interface{}
	chanID, e = strconv.Atoi(channel)
	if e != nil {
		chanID = channel
	}

	// 取得 訊息內容
	msg, e := req.Get("message").String()
	if e != nil {
		glog.Error(e)
	}

	// 取得 安靜狀態
	slient, e := req.Get("slient").Bool()
	if e != nil {
		glog.Error(e)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"ok":true}`)

	json, e := json.Marshal(&MessageBody{
		chanID,
		msg,
		slient,
	})
	if e != nil {
		glog.Error(e)
	}

	go t.Send(json)
}

// Send 寄出訊息
func (t *Telegram) Send(json []byte) {
	defer func() {
		<-t.sem
	}()

	req, e := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("TG_BOT_TOKEN")+"/sendMessage", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	if e != nil {
		glog.Error(e)
	}

	client := &http.Client{}
	resp, e := client.Do(req)
	if e != nil {
		glog.Error(e)
	}

	_, e = io.Copy(ioutil.Discard, resp.Body)
	if e != nil {
		glog.Error(e)
	}

	e = resp.Body.Close()
	if e != nil {
		glog.Error(e)
	}
}
