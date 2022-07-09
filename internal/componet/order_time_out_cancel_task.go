package componet

import (
	"github.com/robfig/cron/v3"
	"mall/global/log"
	"mall/internal/service"
)

var omsPortalOrderService = service.OmsPortalOrderService

func init() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 0/10 * ? * ?", cancelTimeOutOrder)
	if err != nil {
		panic("init cancel order task failed")
	}
	c.Start()
}

func cancelTimeOutOrder() {
	order, err := omsPortalOrderService.CancelTimeOutOrder()
	if err != nil {
		log.Logger.Debug("cancel order failed:" + err.Error())
		return
	}
	log.Logger.Info("Cancel order and release locked inventory according to SKU number, number of canceled orders: ",
		log.Any("number", order))
}
