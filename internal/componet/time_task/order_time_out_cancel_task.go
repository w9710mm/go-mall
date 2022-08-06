package time_task

import (
	"github.com/robfig/cron/v3"
	"mall/global/log"
	"mall/internal/service"
)

type orderTimeOutCancelTask struct {
	omsPortalOrderService service.OmsPortalOrderService
	c                     *cron.Cron
}

func NewOrderTimeOutCancelTask(orderService service.OmsPortalOrderService) TimeTask {
	return &orderTimeOutCancelTask{
		c:                     cron.New(cron.WithSeconds()),
		omsPortalOrderService: orderService}
}
func (O *orderTimeOutCancelTask) StartTask() {
	_, err := O.c.AddFunc("0 0/10 * ? * ?", O.cancelTimeOutOrder)
	if err != nil {
		panic("init cancel order task failed")
	}
	O.c.Start()
}

func (O *orderTimeOutCancelTask) StopTask() {
	O.c.Stop()
}
func (O *orderTimeOutCancelTask) cancelTimeOutOrder() {
	order, err := O.omsPortalOrderService.CancelTimeOutOrder()
	if err != nil {
		log.Logger.Debug("cancel order failed:" + err.Error())
		return
	}
	log.Logger.Info("Cancel order and release locked inventory according to SKU number, number of canceled orders: ",
		log.Any("number", order))
}
