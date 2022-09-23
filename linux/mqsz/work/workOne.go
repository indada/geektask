package main

import (
	"fmt"
	"mqsz/rbtmqcs"
	"strconv"
)

func main()  {
	rabitmq := rbtmqcs.NewRabbitMQSimple("queuework")
	for i := 0; i < 1000; i++ {
		//strconv.Itoa(i) 将整形转为字符串
		rabitmq.PublishSimple("hello dd "+strconv.Itoa(i))
	}
	fmt.Println("发送ok！")
}
