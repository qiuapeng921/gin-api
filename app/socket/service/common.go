package service

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
)

// Message 消息结构体
type Message struct {
	ChatType string `json:"chat_type"`
	Data     string `json:"data"`
	Sender   string    `json:"sender"`
}

type UserClient struct {
	mutex sync.RWMutex
	// 绑定用户连接
	userClient map[string]*websocket.Conn
	// 连接绑定用户
	clientUser map[*websocket.Conn]string
}

// 绑定用户连接
func (u *UserClient) BindUser(user string, conn *websocket.Conn) {
	u.mutex.Lock()
	if u.userClient == nil {
		u.userClient = make(map[string]*websocket.Conn)
		u.clientUser = make(map[*websocket.Conn]string)
	}

	// 判断用户是否在线 在线则推送下线通知
	if beforeUser, ok := u.userClient[user]; ok {
		_ = beforeUser.WriteMessage(websocket.TextMessage, []byte("别处登录"))
		_ = beforeUser.Close()
	}
	if len(u.userClient) > 0 {
		for _, client := range u.userClient {
			_ = client.WriteMessage(1, []byte(user+"上线啦"))
		}
	}
	u.userClient[user] = conn
	u.clientUser[conn] = user
	fmt.Println("=============上线打印================")
	fmt.Println("=============userClient================", u.userClient)
	fmt.Println("=============clientUser================", u.clientUser)
	u.mutex.Unlock()
}

// 获取指定用户的连接
func (u *UserClient) GetUser(user string) (conn *websocket.Conn, err error) {
	var ok bool
	if conn, ok = u.userClient[user]; !ok {
		err = errors.New(user + "不存在")
		return nil, err
	}
	return conn, nil
}

// 删除用户绑定的链接
func (u *UserClient) RemoveUser(conn *websocket.Conn) {
	u.mutex.Lock()
	if user, ok := u.clientUser[conn]; ok {
		delete(u.userClient, user)
		delete(u.clientUser, conn)
		fmt.Println("=============下线打印================")
		fmt.Println("=============userClient================", u.userClient)
		fmt.Println("=============clientUser================", u.clientUser)
		for _, client := range u.userClient {
			_ = client.WriteMessage(1, []byte(user+"下线啦"))
		}
	} else {
		fmt.Println(user + "不存在")
	}
	u.mutex.Unlock()
}

func (u *UserClient) SendToUser(userId string, message string) bool {
	u.mutex.Lock()
	if client, ok := u.userClient[userId]; ok {
		if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			logrus.Println("发送消息给【" + userId + "】失败" + err.Error())
			return false
		}
	}
	u.mutex.Unlock()
	return true
}

// 发送消息给指定的用户
func (u *UserClient) SendToSomeUser(userIds []string, message string) {
	for _, userId := range userIds {
		if _, ok := u.userClient[userId]; !ok {
			continue
		}
		u.SendToUser(userId, message)
	}
}

// 发送消息到群组
func (u *UserClient) SendToGroup(groupId int, message string) bool {
	u.mutex.Lock()
	u.mutex.Unlock()
	return true
}
