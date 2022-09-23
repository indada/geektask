package main
import "mqsz/rbtmqcs"

func main(){
	//这里只是匹配到了huxiaobai.后边只能是一个单词的key 通过这个key去找绑定到交换机上的相应的队列
	rabbitmq := rbtmqcs.NewRabbitMQTopic("hxbExc","huxiaobai.*.cs")
	rabbitmq.RecieveTopic()
}
