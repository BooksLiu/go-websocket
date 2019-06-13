package handler

import (
	"errors"
	"github.com/gorilla/websocket"
)

type WorkRouter struct {
	router map[int32]Handler
}

func (wr WorkRouter) GetWorker(action int32) (Handler, error) {
	if h, ok := wr.router[action]; !ok {
		return nil, errors.New("can not find handler")
	} else {
		return h, nil
	}
}

type Handler interface {
	Work(string, *websocket.Conn) (string, error)
}

var (
	Router  map[int32]Handler
	ConnMap map[string]*websocket.Conn
)

const (
	RegisterAction = 1
)

func init() {
	Router = make(map[int32]Handler)
	ConnMap = make(map[string]*websocket.Conn)
	// 添加handler
	Router[RegisterAction] = &RegisterHandler{}
}
