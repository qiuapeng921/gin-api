package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	conn, err := onOpen(c.Writer, c.Request)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			onClone(conn)
			break
		}
		log.Println(string(msg))
		if string(msg) == "PING" {
			err = conn.WriteMessage(msgType, []byte("DONG"))
			if err != nil {
				break
			}
		}
		// todo：业务操作
		err = onMessage(conn, msgType, string(msg))
		if err != nil {
			break
		}
	}
}

func onOpen(response http.ResponseWriter, request *http.Request) (conn *websocket.Conn, err error) {
	conn, err = wsUpGrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println("websocket upgrade err:", err.Error())
		return
	}
	_ = conn.WriteMessage(websocket.TextMessage, []byte("welcome"))
	return
}

func onMessage(conn *websocket.Conn, msgType int, data string) (err error) {
	err = conn.WriteMessage(msgType, []byte(data))
	return
}

func onClone(conn *websocket.Conn) {
	conn.CloseHandler()
	log.Println("用户下线")
}