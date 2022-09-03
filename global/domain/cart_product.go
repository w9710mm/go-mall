package domain

import "mall/global/model"

type CartProduct struct {
	model.PmsProduct
	ProductAttributeList []model.PmsProductAttribute `gorm:"foreignKey:ProductAttributeCategoryId"`
	SkuStockList         []model.PmsSkuStock         `gorm:"foreignKey:ProductId"`
}
