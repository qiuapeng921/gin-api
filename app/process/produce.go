package process

import (
	"fmt"
	"gin-api/helpers/queue"
	"gin-api/helpers/system"
	"time"
)

func InitProduce() {
	for {
		data := system.GetCurrentDate()
		if err := queue.Publish("test", "test", data); err != nil {
			fmt.Println("推送消息失败", err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}