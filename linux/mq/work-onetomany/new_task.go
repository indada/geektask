package main

import (
	"log"
	"os"
	"strings"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error,msg string)  { //打印错误
	if err != nil {
		log.Fatalf("%s : %s",msg,err)
	}
}
func main()  {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err,"Failed to connect to RabbitMQ !")
	defer conn.Close()

	ch,err := conn.Channel()
	failOnError(err,"Failed to open a channel !")
	defer ch.Close()

	q,err := ch.QueueDeclare(
		"test_queue", //name
		true, //持久化
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:	[]byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
func bodyFrom(args []string) string {
	var s string
	if len(args) <2 || os.Args[1] == "" {
		s = "dddd"
	}else {
		s = strings.Join(args[1:],".")
	}

	return s
}