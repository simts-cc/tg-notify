package tg

import (
	"flag"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
)

// Service 服務架構
type Service struct {
	Router *httprouter.Router
}

// NewService 服務實體
func NewService() *Service {
	flag.Parse()
	s := &Service{
		Router: NewRouter(),
	}
	return s
}

// Start 服務啟動
func (s *Service) Start() {
	defer func() {
		if err := recover(); err != nil {
			glog.Error(err)
		}
	}()

	port := ":" + os.Getenv("SERVER_PORT")

	http.ListenAndServe(port, s.Router)
}
