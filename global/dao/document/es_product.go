package document

import (
	"mall/global/dao/domain"
	"mall/global/dao/model"
)

type EsProduct struct {
	model.PmsProduct
	AttrValueList []domain.EsProductAttributeValue
}

func (p EsProduct) GetType() string {
	return "product"
}

func (p EsProduct) GetIndex() string {
	return "pms"
}
