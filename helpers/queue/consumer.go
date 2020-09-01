package queue

import "fmt"

func PopMessage() {

	channel, message, err := Consumer("test", "test")
	if err != nil {
		fmt.Println("获取队列失败" + err.Error())
	}
	defer channel.Close()
	// 使用callback消费数据
	for msg := range message {
		// 当接收者消息处理失败的时候，
		// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
		// 通过重试可以成功的操作，那么这个时候是需要重试的
		// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
		//for !receiver.OnReceive(msg.Body) {
		//	log.Warnf("receiver 数据处理失败，将要重试")
		//	time.Sleep(1 * time.Second)
		//}
		fmt.Println("接受", string(msg.Body))
		// 确认收到本条消息, multiple必须为false
		if err := msg.Ack(false); err != nil {
			fmt.Println("rabbit确认消息失败 错误信息:" + err.Error())
		}

		break
	}
}
