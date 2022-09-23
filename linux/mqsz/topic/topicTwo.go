//topic主题模式消费者一
package main

import "mqsz/rbtmqcs"

func main(){
	//#号表示匹配多个单词 也就是读取hxbExc交换机里面所有队列的消息
	rabbitmq := rbtmqcs.NewRabbitMQTopic("hxbExc","#")
	rabbitmq.RecieveTopic()
}