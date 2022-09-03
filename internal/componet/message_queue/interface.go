package message_queue

type OrderMQ interface {
	CancelOrderSender(orderId int64, delayTimes int64)
	CancelOrderReceiver()
}
