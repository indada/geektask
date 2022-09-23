//路由模式消费者一
package main

import "mqsz/rbtmqcs"

func main(){
	rabbitmqOne := rbtmqcs.NewRabbitMQRouting("exHxb","xiaobai_one")
	rabbitmqOne.RecieveRouting()
}