package document

import (
	"mall/global/dao/domain"
	"mall/global/dao/model"
)

type EsProduct struct {
	model.PmsProduct
	AttrValueList []domain.EsProductAttributeValue
}
