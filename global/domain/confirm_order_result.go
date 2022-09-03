package domain

import (
	"github.com/shopspring/decimal"
	"mall/global/model"
)

/**
 *@author:
 *@date:2022/8/22
**/

type ConfirmOrderResult struct {
	CartPromotionItemList     []CartPromotionItem
	MemberReceiveAddressList  []model.UmsMemberReceiveAddress
	CouponHistoryDetailList   []SmsCouponHistoryDetail
	IntegrationConsumeSetting model.UmsIntegrationConsumeSetting
	MemberIntegration         int64
	CalcAmount                CalcAmount
}

type CalcAmount struct {
	TotalAmount     decimal.Decimal
	FreightAmount   decimal.Decimal
	PromotionAmount decimal.Decimal
	PayAmount       decimal.Decimal
}
