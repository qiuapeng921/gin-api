package controller

import (
	"errors"
	"fmt"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// Message 消息结构体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	From string `json:"from"`
}

// 用户绑定用户连接id
var userClient = make(map[string]*websocket.Conn, 200)

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 循环处理消息数据
func WebSocketHandler(c *gin.Context) {
	conn, err := onOpen(c.Writer, c.Request)
	if err != nil {
		_ = conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		_ = conn.Close()
		c.Abort()
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
			continue
		}
		// todo：业务操作
		err = onMessage(conn, msgType, string(msg))
		if err != nil {
			break
		}
	}
}

// 初始化连接
func onOpen(response http.ResponseWriter, request *http.Request) (conn *websocket.Conn, err error) {
	conn, err = wsUpGrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println("websocket upgrade err:", err.Error())
		return
	}

	username := request.URL.Query().Get("username")
	if username == "" {
		return conn, errors.New("username不能为空")
	}
	// 将用户和用户连接Id绑定
	userClient[username] = conn

	fmt.Println("-------------userClient----------------------", userClient)

	_ = conn.WriteMessage(websocket.TextMessage, []byte("welcome"))
	return
}

// 收到消息处理
func onMessage(conn *websocket.Conn, msgType int, data string) (err error) {
	var message Message
	_ = system.JsonToStruct(data, &message)
	if message.Type == "chat" {
		from := message.From
		if fromClient, ok := userClient[from]; ok {
			_ = fromClient.WriteMessage(websocket.TextMessage, []byte("connId不存在"))
		} else {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(from+"不在线"))
		}
	} else {
		err = conn.WriteMessage(msgType, []byte(data))
	}
	return
}

// 连接断开
func onClone(conn *websocket.Conn) {
	conn.CloseHandler()
	log.Println("用户下线")
}
