package domain

import "mall/global/model"

/**
 *@author:
 *@date:2022/8/17
**/

type PromotionProduct struct {
	model.PmsProduct
	SkuStockList             []model.PmsSkuStock             `gorm:"foreignKey:ProductId"`
	ProductLadderList        []model.PmsProductLadder        `gorm:"foreignKey:ProductId"`
	productFullReductionList []model.PmsProductFullReduction `gorm:"foreignKey:ProductId"`
}
