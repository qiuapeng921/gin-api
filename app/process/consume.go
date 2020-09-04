package process

import (
	"fmt"
	"gin-api/helpers/db"
	"gin-api/helpers/queue"
	"time"
)

func InitConsume() {
	for {
		channel, message, err := queue.Consumer("request", "request")
		if err != nil {
			fmt.Println("获取队列失败" + err.Error())
		}
		select {
		case msg := <-message:
			go db.EsClient.Insert("request", string(msg.Body))
			// 确认收到本条消息, multiple必须为false
			if err := msg.Ack(false); err != nil {
				fmt.Println("rabbit确认消息失败 错误信息:" + err.Error())
			}
		}
		channel.Close()
		time.Sleep(100 * time.Millisecond)
	}
}
