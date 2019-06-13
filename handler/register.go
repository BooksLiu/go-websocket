package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type RegisterHandler struct {
	Action int32
}

type RegisterMessage struct {
	UserId int64 `json:"user_id"`
	AppId  int32 `json:"app_id"`
}

func (h *RegisterHandler) Work(reqStr string, conn *websocket.Conn) (string, error) {
	rm := RegisterMessage{}
	json.Unmarshal([]byte(reqStr), &rm)
	// 注册conn到 map中
	key := fmt.Sprintf("%v_%v", rm.AppId, rm.UserId)
	ConnMap[key] = conn
	// todo 查询用户相关的消息 写回客户端
	return "", nil
}
