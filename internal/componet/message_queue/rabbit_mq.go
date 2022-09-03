package message_queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"mall/global/config"
	"mall/global/log"
)

type RabbitMQ struct {
	// 连接
	conn *amqp.Connection
	// 频道
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 连接信息
	MqUrl string
}

func newRabbitMQ(queuename string, exchange string, key string) *RabbitMQ {
	MQconfig := config.GetConfig().RabbitMQ
	mqurl := fmt.Sprintf("amqp://%s:%s@%s:%d//%s", MQconfig.Username, MQconfig.Password, MQconfig.Host, MQconfig.Port, MQconfig.VirtualHost)
	rabbitmq := &RabbitMQ{QueueName: queuename, Exchange: exchange, Key: key, MqUrl: mqurl}
	var err error
	//创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "创建连接错误！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 定义错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Logger.Fatal(err.Error(), log.String("msg", message))

		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//simple model
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return newRabbitMQ(queueName, "", "")
}

//subscription model
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitMQ := newRabbitMQ("", exchangeName, "")
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.MqUrl)
	rabbitMQ.failOnErr(err, "failed to connect rabbitmq!")

	//get channel
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "failed to open a channel!")
	return rabbitMQ
}

//简单模式Step:2、简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//1、申请队列，如果队列存在就跳过，不存在创建
	//优点：保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		//队列名称
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除 当最后一个消费者断开连接之后，是否把消息从队列中删除
		false,
		//是否具有排他性 true表示自己可见 其他用户不能访问
		false,
		//是否阻塞 true表示要等待服务器的响应
		false,
		//额外数学系
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2.发送消息到队列中
	r.channel.Publish(
		//默认的Exchange交换机是default,类型是direct直接类型
		r.Exchange,
		//要赋值的队列名称
		r.QueueName,
		//如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息还给发送者
		false,
		//消息
		amqp.Publishing{
			//类型
			ContentType: "text/plain",
			//消息
			Body: []byte(message),
		})
}

func NewSimpleRabbitMQ(queueName string) *RabbitMQ {

	return newRabbitMQ(queueName, "", "")
}
