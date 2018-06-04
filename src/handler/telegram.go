package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"tg.notify/src/model"
)

// Telegram 架構
type Telegram struct {
	sem *chan struct{}
	orm *orm.Ormer
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
func (t *Telegram) SetSem(s *chan struct{}) {
	t.sem = s
}

// SetOrm 設置 Orm
func (t *Telegram) SetOrm(o *orm.Ormer) {
	t.orm = o
}

// SendMessage 傳送訊息
func (t *Telegram) SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		reqBody []byte
		res     []byte
		chanID  interface{}
		msg     string
		slient  bool
	)
	// 紀錄
	defer func() {
		nr := strings.NewReplacer("\t", "", "\n", "")
		header, _ := json.Marshal(r.Header)
		apilog := model.APILogs{
			URI:     r.RequestURI,
			ReqData: nr.Replace(string(reqBody)),
			ResData: string(res),
			Headers: string(header),
		}
		model.APILogsAdd(*t.orm, apilog)
	}()

	reqBody, e := ioutil.ReadAll(r.Body)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}

	// REQUEST 轉 JSON
	reqData, e := simplejson.NewJson(reqBody)
	if e != nil {
		glog.Error(e)
	}

	// 取得 頻道ＩＤ
	code, e := reqData.Get("code").String()
	if e != nil {
		res = APIErrorResponse{Ok: false, Message: "code error."}.Response(&w)
		return
	}
	channel := os.Getenv("TG_CHAN_" + code)
	if channel == "" {
		res = APIErrorResponse{Ok: false, Message: "code error."}.Response(&w)
		return
	}
	chanID, e = strconv.Atoi(channel)
	if e != nil {
		chanID = channel
	}

	// 取得 訊息內容
	msg, e = reqData.Get("message").String()
	if e != nil || msg == "" {
		res = APIErrorResponse{Ok: false, Message: "message error."}.Response(&w)
		return
	}

	// 取得 安靜狀態
	slient, _ = reqData.Get("slient").Bool()

	res = APIResponse{Ok: true}.Response(&w)

	tgBody, _ := json.Marshal(&MessageBody{
		ChatID:  chanID,
		Message: msg,
		Slient:  slient,
	})

	*t.sem <- struct{}{}
	go t.Send(tgBody)
}

// Send 寄出訊息
func (t *Telegram) Send(json []byte) {
	var resp *http.Response

	defer func() {
		_, e := io.Copy(ioutil.Discard, resp.Body)
		if e != nil {
			glog.Error(e)
		}

		e = resp.Body.Close()
		if e != nil {
			glog.Error(e)
		}

		<-*t.sem
	}()

	req, e := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("TG_BOT_TOKEN")+"/sendMessage", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	if e != nil {
		glog.Error(e)
	}

	client := &http.Client{}
	resp, e = client.Do(req)
	if e != nil {
		glog.Error(e)
	}
}
