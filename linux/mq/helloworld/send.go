package main

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") //拨号连接
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close() //结束关闭

	ch, err := conn.Channel() //打开通道
	failOnError(err, "Failed to open a channel")
	defer ch.Close() //通道关闭

	q, err := ch.QueueDeclare( //队列声明
		"hello", // name 名称
		false,   // durable 是否持久化，如果持久化，mq重启后队列还在
		false,   // delete when unused 未使用时删除 自动删除，队列不再使用时是否自动删除此队列，如果将此参数和exclusive参数设置为true就可以实现临时队列（队列不用了就自动删除）
		false,   // exclusive 是否独占连接，队列只允许在该连接中访问，如果connection连接关闭队列则自动删除,如果将此参数设置true可用于临时队列的创建
		false,   // no-wait 不，等等
		nil,     // arguments 参数，可以设置一个队列的扩展参数，比如：可设置存活时间
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello Dada!"
	err = ch.Publish( //发布
		"",     // exchange 交换机 如果不指定将使用mq的默认交换机
		q.Name, // routing key 交换机根据路由key来将消息转发到指定的队列，如果使用默认交换机，routingKey设置为队列的名称
		false,  // mandatory 强制性
		false,  // immediate 立即
		amqp.Publishing{ //消息内容
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
