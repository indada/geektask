package main
import (
	"mqsz/rbtmqcs"
)

func main() {
	rabbitmq := rbtmqcs.NewRabbitMQSimple("queuework")
	rabbitmq.ConsumeSimple()
}