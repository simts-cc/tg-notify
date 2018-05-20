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

	"github.com/golang/glog"

	"github.com/bitly/go-simplejson"
	"github.com/julienschmidt/httprouter"
)

// MessageBody 訊息內容
type MessageBody struct {
	ChatID  interface{} `json:"chat_id,string"`
	Message string      `json:"text,string"`
	Slient  bool        `json:"disable_notification,bool"`
}

func (body *MessageBody) send() {
	json, e := json.Marshal(body)
	if e != nil {
		glog.Error(e)
	}

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
	defer func() {
		_, e = io.Copy(ioutil.Discard, resp.Body)
		if e != nil {
			glog.Error(e)
		}
		e = resp.Body.Close()
		if e != nil {
			glog.Error(e)
		}
	}()
}

// SendMessage 傳送訊息
func SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	tgBody := MessageBody{
		chanID,
		msg,
		slient,
	}

	go tgBody.send()
}
