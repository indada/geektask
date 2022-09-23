package main

import "mqsz/rbtmq"

func main()  {
	rabitmq := rbtmq.NewRabbitMQPubSub("newExchangeName")
	rabitmq.RecieveSub()
}
