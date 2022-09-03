package test

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"testing"
	"time"
)

const MQURL = "amqp://mall:mall@127.0.0.1:5672//mall"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
	mqurl     string
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s\n", err, message)
		panic(fmt.Sprintf("%s:%s\n", err, message))
	}
}
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {

	rabbitMQ := &RabbitMQ{
		queueName: queueName,
		exchange:  exchange,
		key:       key,
		mqurl:     MQURL,
	}

	dial, err := amqp.Dial(rabbitMQ.mqurl)
	rabbitMQ.failOnErr(err, "创建连接失败")
	rabbitMQ.conn = dial

	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "获取通道失败")

	return rabbitMQ
}

func (r *RabbitMQ) destory() {
	r.channel.Close()
	r.conn.Close()
}

func NewSimpleRabbitMQ(queueName string) *RabbitMQ {

	return NewRabbitMQ(queueName, "", "")
}

func (r *RabbitMQ) Publish(message string) {

	_, err := r.channel.QueueDeclare(
		r.queueName,
		//是否持久化
		true,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		panic(err)
	}

	r.channel.Publish(
		r.exchange,
		r.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

func (r *RabbitMQ) Consumer() {

	_, err := r.channel.QueueDeclare(r.queueName, true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	//2、接收消息
	msgs, err := r.channel.Consume(
		r.queueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messagees,To exit press CTRL+C")

	<-forever

}

func TestProdue(t *testing.T) {
	rabbitMQ := NewSimpleRabbitMQ("mall.order.cancel")

	for i := 0; i < 100000; i++ {
		time.Sleep(100 * time.Millisecond)
		rabbitMQ.Publish("新消息 " + strconv.Itoa(i))
		fmt.Println("发送成功")
	}

}
func TestConsumer(t *testing.T) {

	rabbitMQ := NewSimpleRabbitMQ("mall.order.cancel")

	rabbitMQ.Consumer()

}
