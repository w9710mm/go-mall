package domain

import (
	"github.com/shopspring/decimal"
	"mall/global/model"
)

type FlashPromotionProduct struct {
	model.PmsProduct
	FlashPromotionPrice decimal.Decimal
	FlashPromotionCount int
	FlashPromotionLimit int
}
