package socket

import (
	"gin-api/app/utility/app"
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

type WebSocketHandleInterface interface {
	OnOpen(s *Server,r *http.Request, fd int)
	OnMessage(s *Server,frame *Frame)
	OnClose(s *Server,fd int)
}


type Server struct {
	*websocket.Conn
	WebSocketHandleInterface
	Send
}

func NewServer(wh WebSocketHandleInterface) *Server {
	return &Server{nil, wh, Send{Binder: &bind}}
}

func (s *Server) Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := defaultUpGrader.Upgrade(w, r, nil)
	app.Error("upgrade websocket", err)
	fd++
	s.Bind(fd, conn)

	s.Conn = conn
	requestFd := fd

	s.OnOpen(s,r, fd)

	var frame Frame
	for {
		_, msg, err := s.Conn.ReadMessage()
		if err != nil {
			s.OnClose(s,requestFd)
			break
		}
		log.Println(string(msg))
		frame.SetFd(requestFd)
		frame.SetData(string(msg))
		s.OnMessage(s,&frame)
	}
	return
}