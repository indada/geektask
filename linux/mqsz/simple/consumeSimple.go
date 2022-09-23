package main
import (
	"mqsz/rbtmqcs"
)

func main() {
	rabbitmq := rbtmqcs.NewRabbitMQSimple("queuetwo")
	rabbitmq.ConsumeSimple()
}
