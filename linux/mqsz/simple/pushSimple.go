package main

import (
	"fmt"
	"mqsz/rbtmqcs"
)

func main()  {
	rabitmq := rbtmqcs.NewRabbitMQSimple("queuetwo")
	rabitmq.PublishSimple("hello dada!!!")
	fmt.Println("发送ok！")
}
