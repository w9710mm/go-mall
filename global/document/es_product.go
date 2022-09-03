package document

import (
	"github.com/shopspring/decimal"
	"mall/global/domain"
)

type EsProduct struct {
	//model.PmsProduct
	Id                  int                              `json:"id"`
	ProductSn           string                           `json:"productSn"`
	BrandId             int                              `json:"brandId"`
	BrandName           string                           `json:"brandName"`
	ProductCategoryId   int                              `json:"productCategoryId"`
	ProductCategoryName string                           `json:"productCategoryName"`
	Pic                 string                           `json:"pic"`
	Name                string                           `json:"name"`
	SubTitle            string                           `json:"subTitle"`
	Keywords            string                           `json:"keywords"`
	Price               decimal.Decimal                  `json:"price"`
	Sale                int                              `json:"sale"`
	NewStatus           int                              `json:"newStatus"`
	RecommandStatus     int                              `json:"recommandStatus"`
	Stock               int                              `json:"stock"`
	PromotionType       int                              `json:"promotionType"`
	Sort                int                              `json:"sort"`
	AttrValueList       []domain.EsProductAttributeValue `json:"attrValueList"`
}

func (p EsProduct) GetType() string {
	return "product"
}

func (p EsProduct) GetIndex() string {
	return "pms"
}
