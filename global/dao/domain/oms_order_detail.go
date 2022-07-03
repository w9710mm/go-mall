package domain

import "mall/global/dao/model"

type OmsOrderDetail struct {
	model.OmsOrder
	OrderItemList []model.OmsOrderItem
}
