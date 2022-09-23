package mq
import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url 格式 amqp://账号:密码@rabbitmq服务器地址:端口号/Virtual Host
//格式在golang语言当中是固定不变的
const MQURL = "amqp://guest:guest@localhost:5672/"

type RabbitMQ struct {
	conn *amqp.Connection //amqp.Dial()连接 需要引入amqp包 https://learnku.com/articles/44185教会你如何引用amqp包
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//链接信息
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queuename string,exchange string,key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName:queuename,Exchange:exchange,Key:key,Mqurl:MQURL}
	var err error
	//创建rabbitmq连接
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)  //通过amqp.Dial()方法去链接rabbitmq服务端
	rabbitmq.failOnErr(err,"创建连接错误!")  //调用我们自定义的failOnErr()方法去处理异常错误信息
	rabbitmq.channel,err = rabbitmq.conn.Channel() //链接上rabbitmq之后通过rabbitmq.conn.Channel()去设置channel信道
	rabbitmq.failOnErr(err,"获取channel失败!")
	return rabbitmq
}

//断开channel和connection
//为什么要断开channel和connection 因为如果不断开他会始终使用和占用我们的channel和connection 断开是为了避免资源浪费
func (r *RabbitMQ) Destory() {
	r.channel.Close()    //关闭信道资源
	r.conn.Close()       //关闭链接资源
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error,message string) {
	if err != nil {
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

//简单模式step：1.创建简单模式下的rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//simple模式下交换机为空因为会默认使用rabbitmq默认的default交换机而不是真的没有 bindkey绑定建key也是为空的
	//特别注意：simple模式是最简单的rabbitmq的一种模式 他只需要传递queue队列名称过去即可  像exchange交换机会默认使用default交换机  绑定建key的会不必要传
	return NewRabbitMQ(queueName,"","")
}

//简单模式step:2.简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列,如果队列不存在，则会自动创建，如果队列存在则跳过创建直接使用  这样的好处保障队列存在，消息能发送到队列当中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		false,
		//是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
		false,
		//其他的属性，没有则直接诶传入空即可 nil  nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.发送消息到队列当中
	r.channel.Publish(
		//交换机 simple模式下默认为空 我们在上边已经赋值为空了  虽然为空 但其实也是在用的rabbitmq当中的default交换机运行
		r.Exchange,
		//队列的名称
		r.QueueName,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		//要发送的消息
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		})
}
//简单模式step:3.简单模式下消费者代码
func (r *RabbitMQ) ConsumeSimple() {

}
