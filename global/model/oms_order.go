package model

import (
	"github.com/shopspring/decimal"
	"time"
)

//订单表
type OmsOrder struct {
	Id                    int64           `gorm:"column:id" json:"id" `                                           //是否可空:NO 订单id
	MemberId              int64           `gorm:"column:member_id" json:"member_id" `                             //是否可空:NO
	CouponId              int64           `gorm:"column:coupon_id" json:"coupon_id" `                             //是否可空:YES
	OrderSn               string          `gorm:"column:order_sn" json:"order_sn" `                               //是否可空:YES
	CreateTime            *time.Time      `gorm:"column:create_time" json:"create_time" `                         //是否可空:YES
	MemberUsername        *string         `gorm:"column:member_username" json:"member_username" `                 //是否可空:YES
	TotalAmount           decimal.Decimal `gorm:"column:total_amount" json:"total_amount" `                       //是否可空:YES
	PayAmount             decimal.Decimal `gorm:"column:pay_amount" json:"pay_amount" `                           //是否可空:YES
	FreightAmount         decimal.Decimal `gorm:"column:freight_amount" json:"freight_amount" `                   //是否可空:YES
	PromotionAmount       decimal.Decimal `gorm:"column:promotion_amount" json:"promotion_amount" `               //是否可空:YES
	IntegrationAmount     decimal.Decimal `gorm:"column:integration_amount" json:"integration_amount" `           //是否可空:YES
	CouponAmount          decimal.Decimal `gorm:"column:coupon_amount" json:"coupon_amount" `                     //是否可空:YES
	DiscountAmount        decimal.Decimal `gorm:"column:discount_amount" json:"discount_amount" `                 //是否可空:YES
	PayType               int             `gorm:"column:pay_type" json:"pay_type" `                               //是否可空:YES
	SourceType            int             `gorm:"column:source_type" json:"source_type" `                         //是否可空:YES
	Status                int             `gorm:"column:status" json:"status" `                                   //是否可空:YES
	OrderType             int             `gorm:"column:order_type" json:"order_type" `                           //是否可空:YES
	DeliveryCompany       *string         `gorm:"column:delivery_company" json:"delivery_company" `               //是否可空:YES
	DeliverySn            *string         `gorm:"column:delivery_sn" json:"delivery_sn" `                         //是否可空:YES
	AutoConfirmDay        *int            `gorm:"column:auto_confirm_day" json:"auto_confirm_day" `               //是否可空:YES
	Integration           int64           `gorm:"column:integration" json:"integration" `                         //是否可空:YES
	Growth                int             `gorm:"column:growth" json:"growth" `                                   //是否可空:YES
	PromotionInfo         string          `gorm:"column:promotion_info" json:"promotion_info" `                   //是否可空:YES
	BillType              *int            `gorm:"column:bill_type" json:"bill_type" `                             //是否可空:YES
	BillHeader            *string         `gorm:"column:bill_header" json:"bill_header" `                         //是否可空:YES
	BillContent           *string         `gorm:"column:bill_content" json:"bill_content" `                       //是否可空:YES
	BillReceiverPhone     *string         `gorm:"column:bill_receiver_phone" json:"bill_receiver_phone" `         //是否可空:YES
	BillReceiverEmail     *string         `gorm:"column:bill_receiver_email" json:"bill_receiver_email" `         //是否可空:YES
	ReceiverName          *string         `gorm:"column:receiver_name" json:"receiver_name" `                     //是否可空:YES
	ReceiverPhone         *string         `gorm:"column:receiver_phone" json:"receiver_phone" `                   //是否可空:YES
	ReceiverPostCode      *string         `gorm:"column:receiver_post_code" json:"receiver_post_code" `           //是否可空:YES
	ReceiverProvince      *string         `gorm:"column:receiver_province" json:"receiver_province" `             //是否可空:YES
	ReceiverCity          *string         `gorm:"column:receiver_city" json:"receiver_city" `                     //是否可空:YES
	ReceiverRegion        *string         `gorm:"column:receiver_region" json:"receiver_region" `                 //是否可空:YES
	ReceiverDetailAddress *string         `gorm:"column:receiver_detail_address" json:"receiver_detail_address" ` //是否可空:YES
	Note                  *string         `gorm:"column:note" json:"note" `                                       //是否可空:YES
	ConfirmStatus         int             `gorm:"column:confirm_status" json:"confirm_status" `                   //是否可空:YES
	DeleteStatus          int             `gorm:"column:delete_status" json:"delete_status" `                     //是否可空:YES
	UseIntegration        int64           `gorm:"column:use_integration" json:"use_integration" `                 //是否可空:YES
	PaymentTime           *time.Time      `gorm:"column:payment_time" json:"payment_time" `                       //是否可空:YES
	DeliveryTime          *time.Time      `gorm:"column:delivery_time" json:"delivery_time" `                     //是否可空:YES
	ReceiveTime           *time.Time      `gorm:"column:receive_time" json:"receive_time" `                       //是否可空:YES
	CommentTime           *time.Time      `gorm:"column:comment_time" json:"comment_time" `                       //是否可空:YES
	ModifyTime            *time.Time      `gorm:"column:modify_time" json:"modify_time" `                         //是否可空:YES
}

func (*OmsOrder) TableName() string {
	return "oms_order"
}
