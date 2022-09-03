package domain

import "mall/global/model"

type SmsCouponHistoryDetail struct {
	model.SmsCouponHistory
	Coupon               model.SmsCoupon                          `gorm:"foreignKey:CouponId"`
	ProductRelationList  []model.SmsCouponProductRelation         `gorm:"foreignKey:Coupon.Id"`
	CategoryRelationList []model.SmsCouponProductCategoryRelation `gorm:"foreignKey:Coupon.Id"`
}
