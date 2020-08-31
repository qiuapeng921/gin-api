package socket

import (
	"gin-api/app/socket/service"
	"gin-api/helpers/system"
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

var socket service.UserClient

// 循环处理消息数据
func Handler(c *gin.Context) {
	conn, err := onOpen(c)
	if err != nil {
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
func onOpen(c *gin.Context) (conn *websocket.Conn, err error) {
	conn, err = wsUpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("websocket upgrade err:", err.Error())
		c.Abort()
		return
	}

	user := c.Query("user")
	if user == "" {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("user[不能为空]"))
		_ = conn.Close()
		c.Abort()
	}
	socket.BindUser(user, conn)

	_ = conn.WriteMessage(websocket.TextMessage, []byte("welcome"))
	return
}

// 收到消息处理
func onMessage(conn *websocket.Conn, msgType int, data string) (err error) {
	var message service.Message
	// 判断是否成功绑定消息到结构体
	if err := system.JsonToStruct(data, &message); err != nil {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("消息体格式错误"))
		return err
	}
	// 判断是否为私聊
	if message.Type == "chat" {
		// 判断发送用户是否在线 (在线则用接收方连接写消息给客户端,否则将返回消息给发送方用户不在线)
		if fromClient, err := socket.GetUser(message.From); err == nil {
			_ = fromClient.WriteMessage(websocket.TextMessage, []byte(message.Data))
		} else {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(message.From+"不在线"))
		}
	} else {
		err = conn.WriteMessage(msgType, []byte(data))
	}
	return nil
}

// 连接断开
func onClone(conn *websocket.Conn) {
	socket.RemoveUser(conn)
}
