package server

import (
	"fmt"
	"gin-api/app/utility/app"
	"gin-api/app/utility/system"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	defaultUpGrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		// 取消ws跨域校验
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	fd   int
	bind Binder
)

type WebSocketServer struct {
	*websocket.Conn
	Send
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{nil, Send{&bind}}
}

func (s *WebSocketServer) Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := defaultUpGrader.Upgrade(w, r, nil)
	app.Error("upgrade websocket", err)
	fd++
	s.Bind(fd, conn)

	s.Conn = conn
	requestFd := fd

	s.onOpen(r, fd)

	var frame Frame
	for {
		_, msg, err := s.Conn.ReadMessage()
		if err != nil {
			s.onClose(requestFd)
			break
		}
		log.Println(string(msg))
		frame.SetFd(requestFd)
		frame.SetData(string(msg))
		s.onMessage(&frame)
	}
	return
}

func (s *WebSocketServer) onOpen(r *http.Request, fd int) {
	log.Println("OnOpen", r.URL.Query())
	s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d 加入聊天室", fd), fd)
	return
}

func (s *WebSocketServer) onMessage(frame *Frame) {
	log.Println("OnMessage", frame.fd, frame.data)
	if frame.data == "PING" {
		s.Push(frame.fd, "DONG")
		return
	}
	var receiver ReceiverMessage
	err := system.JsonToStruct(frame.data, &receiver)
	if err != nil {
		s.Push(frame.fd, "消息类型错误")
		return
	}

	switch receiver.MessageType {
	case UserChat:
		var sender SenderMessage
		sender.SenderId = frame.fd
		sender.MessageType = receiver.MessageType
		sender.Data = receiver.Data
		s.PushMessage(receiver.ReceiverId, sender)
	case GroupChat:
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 在群里说:%s", frame.fd, frame.data), frame.fd)
	default:
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 对大家说:%s", frame.fd, frame.data), frame.fd)
	}
	return
}

func (s *WebSocketServer) onClose(fd int) {
	log.Println("OnClose", fd)
	s.UnBind(fd)
	return
}
