package domain

import "mall/global/model"

type OmsOrderDetail struct {
	model.OmsOrder
	OrderItemList []model.OmsOrderItem `gorm:"foreignKey:OrderId"`
}
