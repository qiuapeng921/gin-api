package websocket

import (
	"fmt"
	"gin-api/app/utility/socket"
	"gin-api/app/utility/system"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewWebsocket(ctx *gin.Context) {
	var handle Handle
	server := socket.NewServer(handle)
	server.Upgrade(ctx.Writer, ctx.Request)
}

type Handle struct{
	
}

func (h Handle) OnOpen(s *socket.Server,r *http.Request, fd int) {
	log.Println("OnOpen", r.URL.Query())
	s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d 加入聊天室", fd), fd)
}

func (h Handle) OnMessage(s *socket.Server,frame *socket.Frame) {
	log.Println("OnMessage", frame.GetFd(), frame.GetData())
	if frame.GetData() == "PING" {
		s.Push(frame.GetFd(), "DONG")
		return
	}
	var receiver socket.ReceiverMessage
	err := system.JsonToStruct(frame.GetData(), &receiver)
	if err != nil {
		s.Push(frame.GetFd(), "消息类型错误")
		return
	}

	switch receiver.MessageType {
	case socket.UserChat:
		var sender socket.SenderMessage
		sender.SenderId = frame.GetFd()
		sender.MessageType = receiver.MessageType
		sender.Data = receiver.Data
		s.PushMessage(receiver.ReceiverId, sender)
	case socket.GroupChat:
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 在群里说:%s", frame.GetFd(), frame.GetData()), frame.GetFd())
	default:
		s.SendToSome(s.GetAllFd(), fmt.Sprintf("用户:%d, 对大家说:%s", frame.GetFd(), frame.GetData()), frame.GetFd())
	}
}

func (h Handle) OnClose(s *socket.Server,fd int) {
	log.Println("333333333333333")
	log.Println("OnClose", fd)
	s.UnBind(fd)
}