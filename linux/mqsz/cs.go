package main
import (
	"fmt"
	"github.com/streadway/amqp"
	"reflect"
)
func main() {
	conn, err := amqp.Dial("amqp://dada:123456@192.168.153.139:5672/aka") //拨号连接
	//判断amqp.Dial()
	dd := reflect.TypeOf(conn)
	fmt.Println("amqp.Dial type: ", dd) //type:  Demo struct
	fmt.Println(err) //type:  Demo struct
	defer conn.Close() //结束关闭
}
