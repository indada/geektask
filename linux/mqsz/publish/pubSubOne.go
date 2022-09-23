package main
import (
	"fmt"
	"mqsz/rbtmq"
	"strconv"
	"time"
)

func main()  {
	rabitmq := rbtmq.NewRabbitMQPubSub("newExchangeName")
	rabitmq.PublishPub("start dada!!!")

	for i := 0; i < 100; i++ {
		rabitmq.PublishPub("订阅模式生成第 "+strconv.Itoa(i)+" 条数据")
		fmt.Println("订阅模式生成第 "+strconv.Itoa(i)+" 条数据")
		time.Sleep(1*time.Second)
	}

}
