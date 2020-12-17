package websocket

import (
	"fmt"
	"gin-api/app/utility/system"
	"gin-api/app/utility/websockets"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewWebsocket(ctx *gin.Context) {
	var handle Handle
	server := websockets.NewServer(handle)
	server.Upgrade(ctx.Writer, ctx.Request)
}

type Handle struct {
}

func (h Handle) OnOpen(s *websockets.Server, r *http.Request, fd int) {
	log.Println("OnOpen", r.URL.Query())
	s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d 加入聊天室", fd), fd)
}

func (h Handle) OnMessage(s *websockets.Server, frame *websockets.Frame) {
	log.Println("OnMessage", frame.GetFd(), frame.GetData())
	if frame.GetData() == "PING" {
		s.Push(frame.GetFd(), "DONG")
		return
	}
	var receiver websockets.ReceiverMessage
	err := system.JsonToStruct(frame.GetData(), &receiver)
	if err != nil {
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 说:%s", frame.GetFd(), frame.GetData()), frame.GetFd())
		return
	}

	switch receiver.MessageType {
	case websockets.UserChat:
		var sender websockets.SenderMessage
		sender.SenderId = frame.GetFd()
		sender.MessageType = receiver.MessageType
		sender.Data = receiver.Data
		s.PushMessage(receiver.ReceiverId, sender)
	case websockets.GroupChat:
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 在群里说:%s", frame.GetFd(), frame.GetData()), frame.GetFd())
	default:
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 对大家说:%s", frame.GetFd(), frame.GetData()), frame.GetFd())
	}
}

func (h Handle) OnClose(s *websockets.Server, fd int) {
	log.Println("OnClose", fd)
	s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 下线", fd), fd)
}
