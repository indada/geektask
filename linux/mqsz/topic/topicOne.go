package main
//topic主题模式生产者
import (
	"fmt"
	"mqsz/rbtmqcs"
	"strconv"
	"time"
	)

func main(){
	rabbitmqOne := rbtmqcs.NewRabbitMQTopic("hxbExc","huxiaobai.one")
	rabbitmqTwo := rbtmqcs.NewRabbitMQTopic("hxbExc","huxiaobai.two.cs")
	for i:=0;i<=10;i++{
		rabbitmqOne.PublishTopic("hello huxiaobai one" + strconv.Itoa(i))
		rabbitmqTwo.PublishTopic("hello huxiaobai two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}