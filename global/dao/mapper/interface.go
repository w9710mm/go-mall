package mapper

import (
	"mall/global/dao/document"
	"mall/global/dao/domain"
	"mall/global/dao/model"
)

type EsProductMapper interface {
	GetAllEsProductList(int64) ([]document.EsProduct, error)
}

type PortalOrderMapper interface {
	GetTimeOutOrders(int) ([]domain.OmsOrderDetail, error)
	UpdateOrderStatus([]int, int) error
	ReleaseSkuStockLock([]model.OmsOrderItem) error
}
