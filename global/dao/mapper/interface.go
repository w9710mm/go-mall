package mapper

import (
	"mall/global/document"
	"mall/global/domain"
	"mall/global/model"
)

type EsProductMapper interface {
	GetAllEsProductList(int64) ([]document.EsProduct, error)
}

type HomeDao interface {
	GetFlashProductList(flashPromotionId int, sessionId int) []domain.FlashPromotionProduct
	GetRecommendBrandList() (brandList []model.PmsBrand)
	GetNewProductList(offset int, limit int) (brandList []model.PmsProduct)
	GetHotProductList(offset int, limit int) (brandList []model.PmsProduct)
	GetRecommendSubjectList(offset int, limit int) (cmsSubjectList []model.CmsSubject)
}

type PortalOrderDao interface {
	GetDetail(orderId int64) domain.OmsOrderDetail
	UpdateSkuStock(itemList []model.OmsOrderItem) int64
	GetTimeOutOrders(minute int) ([]domain.OmsOrderDetail, error)
	UpdateOrderStatus(ids []int64, status int) error
	ReleaseSkuStockLock(itemList []model.OmsOrderItem) error
}

type PortalProductDao interface {
	GetCartProduct(id int64) domain.CartProduct
	GetPromotionProductList(ids []int64) []domain.PromotionProduct
	GetAvailableCouponList(productId int64, productCategoryId int64) []model.SmsCoupon
}

type SmsCouponHistoryDao interface {
	GetDetailList(memberId int64) []domain.SmsCouponHistoryDetail
	GetCouponList(memberId int64, useStatus int) []model.SmsCoupon
}
