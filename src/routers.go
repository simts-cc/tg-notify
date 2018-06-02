package tg

import (
	"os"
	"strconv"

	"github.com/golang/glog"

	mux "github.com/julienschmidt/httprouter"
	"tg.notify/src/handler"
)

// Router 路由結構
type Router struct {
	Method  string
	Pattern string
	Handle  mux.Handle
}

// Routers 路由組
type Routers []Router

// NewRouter 路由實體
func NewRouter() *mux.Router {
	semCount, e := strconv.Atoi(os.Getenv("SERVER_SEM_COUNT"))
	if e != nil {
		glog.Fatal(e)
	}
	sem := make(chan struct{}, semCount)
	t := handler.NewTelegram()
	t.SetSem(&sem)

	routers := Routers{
		Router{
			Method:  "POST",
			Pattern: "/v1/message",
			Handle:  t.SendMessage,
		},
	}

	router := mux.New()
	for _, route := range routers {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}
