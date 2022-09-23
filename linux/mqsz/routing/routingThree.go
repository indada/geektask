//路由模式消费者二
package main

import "mqsz/rbtmqcs"

func main(){
	rabbitmqTwo := rbtmqcs.NewRabbitMQRouting("exHxb","xiaobai_two")
	rabbitmqTwo.RecieveRouting()
}