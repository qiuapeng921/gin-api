package service

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var socket UserClient

func GetClient() *UserClient {
	return &socket
}

// Message 消息结构体
type Message struct {
	ChatType string `json:"chat_type"`
	Data     string `json:"data"`
	Receive  int    `json:"receive"`
}

type UserClient struct {
	mutex sync.RWMutex
	// 绑定用户连接
	userClient map[int]*websocket.Conn
	// 连接绑定用户
	clientUser map[*websocket.Conn]int
}

// 绑定用户连接
func (u *UserClient) BindUser(userId int, conn *websocket.Conn) {
	u.mutex.Lock()
	if u.userClient == nil {
		u.userClient = make(map[int]*websocket.Conn)
		u.clientUser = make(map[*websocket.Conn]int)
	}

	// 判断用户是否在线 在线则推送下线通知
	if beforeUser, ok := u.userClient[userId]; ok {
		_ = beforeUser.WriteMessage(websocket.TextMessage, []byte("别处登录"))
		_ = beforeUser.Close()
	}
	if len(u.userClient) > 0 {
		for _, client := range u.userClient {
			_ = client.WriteMessage(1, []byte(fmt.Sprintf("%d 上线啦", userId)))
		}
	}
	u.userClient[userId] = conn
	u.clientUser[conn] = userId
	fmt.Println("=============上线打印================")
	fmt.Println("=============userClient================", u.userClient)
	fmt.Println("=============clientUser================", u.clientUser)
	u.mutex.Unlock()
}

// 获取指定用户的连接
func (u *UserClient) GetUser(userId int) (conn *websocket.Conn, err error) {
	var ok bool
	if conn, ok = u.userClient[userId]; !ok {
		err = errors.New(fmt.Sprintf("%d 不存在", userId))
		return nil, err
	}
	return conn, nil
}

// 删除用户绑定的链接
func (u *UserClient) RemoveUser(conn *websocket.Conn) {
	u.mutex.Lock()
	if userId, ok := u.clientUser[conn]; ok {
		delete(u.userClient, userId)
		delete(u.clientUser, conn)
		fmt.Println("=============下线打印================")
		fmt.Println("=============userClient================", u.userClient)
		fmt.Println("=============clientUser================", u.clientUser)
		for _, client := range u.userClient {
			_ = client.WriteMessage(1, []byte(fmt.Sprintf("%d 下线啦", userId)))
		}
	} else {
		fmt.Println(fmt.Sprintf("%d 不在线", userId))
	}
	u.mutex.Unlock()
}

func (u *UserClient) SendToUser(userId int, message string) bool {
	u.mutex.Lock()
	if client, ok := u.userClient[userId]; ok {
		if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("发送消息给 %d ,错误信息：%s", userId, err.Error())
			return false
		}
	}
	u.mutex.Unlock()
	return true
}

// 发送消息给指定的用户
func (u *UserClient) SendToSomeUser(userIds []int, message string) {
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
