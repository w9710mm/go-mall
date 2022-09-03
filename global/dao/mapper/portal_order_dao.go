package mapper

import (
	"fmt"
	"gorm.io/gorm"
	"mall/global/domain"
	"mall/global/model"
)

type portalOrderDao struct {
	db *gorm.DB
}

func (m *portalOrderDao) GetDetail(orderId int64) (orderDetail domain.OmsOrderDetail) {
	m.db.Model(&model.OmsOrder{}).Where(orderId).Preload("OrderItemList").First(&orderDetail)
	return
}

func (m *portalOrderDao) UpdateSkuStock(itemList []model.OmsOrderItem) int64 {
	updateSql := "set stock = case id "
	ids := make([]int64, len(itemList))
	for i, item := range itemList {
		updateSql += fmt.Sprintf(" when %d then stock - %d ", item.ProductSkuId, item.ProductQuantity)
		ids[i] = item.ProductSkuId
	}
	updateSql += "end , lock_stock =case id "

	return m.db.Model(&model.PmsSkuStock{}).Updates(updateSql).Where(ids).RowsAffected
}

func NewPortalOrderDao(db *gorm.DB) PortalOrderDao {
	return &portalOrderDao{db: db}
}

func (m *portalOrderDao) GetTimeOutOrders(time int) (timeOutOrders []domain.OmsOrderDetail, err error) {
	err = m.db.Model(&model.OmsOrder{}).Where(" status =0  and"+
		" create_time <= date_add(NOW(), INTERVAL -? MINUTE)", time).
		Preload("OrderItemList").Find(&timeOutOrders).Error
	return
}

func (m *portalOrderDao) UpdateOrderStatus(ids []int64, status int) error {

	return m.db.Model(model.OmsOrder{}).Where("id in ?", ids).Updates(model.OmsOrder{Status: status}).Error
}

func (m *portalOrderDao) ReleaseSkuStockLock(list []model.OmsOrderItem) error {
	var ids = make([]int, len(list))
	exp := "case id "
	for i, item := range list {
		exp = exp + fmt.Sprintf(" when %d then lock_stock - %d", item.ProductId, item.ProductQuantity)
		ids[i] = item.Id
	}
	exp = exp + " end  "

	return m.db.Model(&model.PmsSkuStock{}).Where(ids).
		Update("lock_stock", exp).Error

}
