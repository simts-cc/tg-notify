package tg

import (
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

var routers = Routers{
	Router{
		Method:  "POST",
		Pattern: "/v1/message",
		Handle:  handler.SendMessage,
	},
}

// NewRouter 取得 API 路由
func NewRouter() *mux.Router {
	router := mux.New()
	for _, route := range routers {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}
