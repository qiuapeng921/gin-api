package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

var connection *amqp.Connection

func InitRabbitMq() {
	var err error
	host := os.Getenv("RABBITMQ_HOST")
	user := os.Getenv("RABBITMQ_USERNAME")
	pass := os.Getenv("RABBITMQ_PASSWORD")
	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s", user, pass, host)
	// 创建连接
	connection, err = amqp.Dial(rabbitUrl)
	if err != nil {
		panic(fmt.Sprintf("rabbitmq初始化失败 错误信息：%s", err.Error()))
	}
	fmt.Println("rabbitMq连接成功")
}

// 关闭RabbitMQ连接
func Close() {
	var err error
	err = connection.Close()
	if err != nil {
		fmt.Println("MQ链接关闭失败")
	}
}

func getChannel(exchangeName, queueName string) (channel *amqp.Channel, queue amqp.Queue, err error) {
	channel, err = connection.Channel()
	if err != nil {
		return
	}
	// exchangeType = "direct", "Exchange type - direct|fanout|topic|x-custom")
	err = channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Printf("MQ注册交换机失败:%s \n", err)
		return
	}

	// 创建队列
	queue, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return
	}
	// 队列绑定
	err = channel.QueueBind(queueName, queue.Name, exchangeName, true, nil)
	if err != nil {
		fmt.Printf("MQ绑定队列失败:%s \n", err)
		return
	}
	return
}

// 生产队列消息
func Publish(exchangeName, queueName, data string) error {
	channel, queue, err := getChannel(exchangeName, queueName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer channel.Close()
	// 发送任务消息
	err = channel.Publish(exchangeName, queue.Name, false, false, amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		Body:            []byte(data),
		DeliveryMode:    amqp.Transient,
		Priority:        0,
	})
	if err != nil {
		fmt.Printf("MQ任务发送失败:%s \n", err)
		return err
	}
	return nil
}

// 消费队列消息
func Consumer(exchangeName, queueName string) (
	channel *amqp.Channel,
	message <-chan amqp.Delivery,
	err error) {

	var queue amqp.Queue
	channel, _, err = getChannel(exchangeName, queueName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 获取消费通道 确保rabbitmq会一个一个发消息
	_ = channel.Qos(1, 0, true)

	message, err = channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		_ = fmt.Errorf("获取队列 %s 的消费通道失败: %s", queueName, err.Error())
		return
	}
	return
}
