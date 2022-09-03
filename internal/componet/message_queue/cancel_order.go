package message_queue

import (
	"encoding/binary"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"mall/global/log"
	"mall/internal/service"
	"strconv"
)

type OrderRabbitMQ struct {
	portalOrderService service.OmsPortalOrderService
	rb                 *RabbitMQ
}

func (r *OrderRabbitMQ) CancelOrderSender(orderId int64, delayTimes int64) {
	//1、申请队列，如果队列存在就跳过，不存在创建
	//优点：保证队列存在，消息能发送到队列中
	_, err := r.rb.channel.QueueDeclare(
		//队列名称
		r.rb.QueueName,
		//是否持久化
		true,
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

	var message []byte
	binary.PutVarint(message, orderId)
	r.rb.channel.Publish(
		//默认的Exchange交换机是default,类型是direct直接类型
		r.rb.Exchange,
		//要赋值的队列名称
		r.rb.QueueName,
		//如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息还给发送者
		false,
		//消息
		amqp.Publishing{
			//类型
			ContentType: "text/plain",
			//消息.

			Body:       message,
			Expiration: strconv.FormatInt(delayTimes, 10),
		})
}

func NewCancelOrderSender(mq *RabbitMQ, orderService service.OmsPortalOrderService) OrderMQ {

	return &OrderRabbitMQ{rb: mq, portalOrderService: orderService}
}

func (r *OrderRabbitMQ) CancelOrderReceiver() {
	//1、申请队列，如果队列存在就跳过，不存在创建
	//优点：保证队列存在，消息能发送到队列中
	_, err := r.rb.channel.QueueDeclare(
		//队列名称
		r.rb.QueueName,
		//是否持久化
		true,
		//是否为自动删除 当最后一个消费者断开连接之后，是否把消息从队列中删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外数学系
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	msgs, err := r.rb.channel.Consume(
		r.rb.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)

	//启用协程处理
	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			varint, i := binary.Varint(d.Body)
			if i <= 0 {
				log.Logger.Info("process orderId occur error:", zap.Int64("order", varint))
				continue
			}
			log.Logger.Info("process orderId:", zap.Int64("order", varint))
			r.portalOrderService.CancelOrder(varint)
		}
	}()

	<-forever
}
