package websockets

import (
	"gin-api/app/utility/app"
	"gin-api/app/utility/system"
	"github.com/gorilla/websocket"
	"log"
)

type Send struct {
	*Binder
}

func (s *Send) PushMessage(fd int, message SenderMessage) {
	conn, ok := s.client[fd]
	if !ok {
		log.Println("不在线", fd)
		return
	}
	err := conn.WriteMessage(websocket.TextMessage, []byte(system.StructToJson(message)))
	app.Error("push message", err)
}

// Push 发送消息给指定用户
func (s *Send) Push(fd int, message string) {
	conn, ok := s.client[fd]
	if !ok {
		log.Println("不在线", fd)
		return
	}
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	app.Error("push message", err)
}

// SendToAll 发送消息给所有用户
func (s *Send) SendToAll(message string) {
	for fd, _ := range s.GetAll() {
		s.Push(fd, message)
	}
}

// SendToSome 发送消息给用户
func (s *Send) SendToSome(fds []int, message string, exclude ...int) {
	if len(fds) < 0 {
		log.Println("fds为空发送结束")
		return
	}
	for _, fd := range fds {
		if len(exclude) > 0 {
			if InArrayInt(fd, exclude) {
				continue
			}
		}
		s.Push(fd, message)
	}
}

func InArrayInt(item int, items []int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
