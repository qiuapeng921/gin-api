package service

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
)

// Message 消息结构体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	From string `json:"from"`
}

// 连接绑定用户
var clientUser = make(map[*websocket.Conn]string)

// 用户绑定连接
var userClient = make(map[string]*websocket.Conn)

// 绑定用户连接
func BindUser(user string, conn *websocket.Conn) {
	// 判断用户是否在线 在线则推送下线通知
	if beforeUser, ok := userClient[user]; ok {
		_ = beforeUser.WriteMessage(websocket.TextMessage, []byte("别处登录"))
		_ = beforeUser.Close()
	}
	if len(userClient) > 0 {
		for _, client := range userClient {
			_ = client.WriteMessage(1, []byte(user+"上线啦"))
		}
	}
	userClient[user] = conn
	clientUser[conn] = user
	fmt.Println("=============userClient================", userClient)
	fmt.Println("=============clientUser================", clientUser)
}

// 获取指定用户的连接
func GetUser(user string) (conn *websocket.Conn, err error) {
	var ok bool
	if conn, ok = userClient[user]; !ok {
		err = errors.New(user + "不存在")
		return nil, err
	}
	return conn, nil
}

// 删除用户绑定的链接
func RemoveUser(conn *websocket.Conn) bool {
	if user, ok := clientUser[conn]; ok {
		delete(userClient, user)
		delete(clientUser, conn)
		fmt.Println("=============userClient================", userClient)
		fmt.Println("=============clientUser================", clientUser)
		for _, client := range userClient {
			_ = client.WriteMessage(1, []byte(user+"下线"))
		}
		return true
	} else {
		fmt.Println(user + "不存在")
		return false
	}
}