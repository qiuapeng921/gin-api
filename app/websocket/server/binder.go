package server

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Binder struct {
	mutex  sync.RWMutex
	client map[int]*websocket.Conn
}

// 绑定客户
func (b *Binder) Bind(fd int, conn *websocket.Conn) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.client == nil {
		b.client = make(map[int]*websocket.Conn)
	}
	b.client[fd] = conn
	log.Println("在线用户：", b.client)
}

// 解绑客户
func (b *Binder) UnBind(fd int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	delete(b.client, fd)
	log.Println("下线用户：", b.client)
}

// 获取所有在线客户端
func (b *Binder) GetAll() map[int]*websocket.Conn {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.client
}

// 获取所有现在的fd
func (b *Binder) GetAllFd() []int {
	var fds []int
	for fd, _ := range b.client {
		fds = append(fds, fd)
	}
	return fds
}

// 获取客户连接
func (b *Binder) GetConnByFd(fd int) *websocket.Conn {
	conn, ok := b.client[fd]
	if !ok {
		log.Printf("%d 不在线", fd)
	}
	return conn
}

// 获取客户fd
func (b *Binder) GetFdByConn(conn *websocket.Conn) int {
	var fd int
	for f, c := range b.GetAll() {
		if c == conn {
			fd = f
			break
		}
	}
	if fd == 0 {
		log.Printf("%p 不在线", conn)
	}
	return fd
}