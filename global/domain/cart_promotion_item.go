package domain

import (
	"github.com/shopspring/decimal"
	"mall/global/model"
)

type CartPromotionItem struct {
	model.OmsCartItem
	PromotionMessage string
	ReduceAmount     decimal.Decimal
	RealStock        int
	Integration      int
	Growth           int
}
