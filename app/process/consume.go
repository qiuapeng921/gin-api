package process

import (
	"fmt"
	"gin-api/helpers/queue"
	"time"
)

func InitConsume() {
	channel, message, err := queue.Consumer("test", "test")
	defer channel.Close()
	if err != nil {
		fmt.Println("获取队列失败" + err.Error())
	}
	for {
		select {
		case msg := <-message:
			fmt.Println("接受", string(msg.Body))
			// 确认收到本条消息, multiple必须为false
			if err := msg.Ack(false); err != nil {
				fmt.Println("rabbit确认消息失败 错误信息:" + err.Error())
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
