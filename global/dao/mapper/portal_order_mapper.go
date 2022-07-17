package mapper

import (
	"fmt"
	"mall/global/dao"
	"mall/global/dao/domain"
	"mall/global/dao/model"
)

type portalOrderMapper struct {
}

var PortalOrderMapper = new(portalOrderMapper)

func (m *portalOrderMapper) GetTimeOutOrders(time int) (timeOutOrders []domain.OmsOrderDetail, err error) {
	err = dao.DB.Model(&model.OmsOrder{}).Where(" status =0  and"+
		" create_time <= date_add(NOW(), INTERVAL -? MINUTE)", time).
		Preload("OrderItemList").Find(&timeOutOrders).Error
	return
}

func (m *portalOrderMapper) UpdateOrderStatus(ids []int, status int) error {

	return dao.DB.Model(model.OmsOrder{}).Where("id in ?", ids).Updates(model.OmsOrder{Status: &status}).Error
}

func (m *portalOrderMapper) ReleaseSkuStockLock(list []model.OmsOrderItem) error {
	var ids = make([]int, len(list))
	exp := "case id "
	for i, item := range list {
		exp = exp + fmt.Sprintf(" when %d then lock_stock - %d", item.ProductId, item.ProductQuantity)
		ids[i] = item.Id
	}
	exp = exp + " end  "

	return dao.DB.Model(&model.PmsSkuStock{}).Where(ids).
		Update("lock_stock", exp).Error

}
