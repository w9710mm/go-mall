package model

import "time"

//订单表
type OmsOrder struct {
	Id                    int        `gorm:"column:id" json:"id" `                                           //是否可空:NO 订单id
	MemberId              int        `gorm:"column:member_id" json:"member_id" `                             //是否可空:NO
	CouponId              *int       `gorm:"column:coupon_id" json:"coupon_id" `                             //是否可空:YES
	OrderSn               *string    `gorm:"column:order_sn" json:"order_sn" `                               //是否可空:YES 订单编号
	CreateTime            *time.Time `gorm:"column:create_time" json:"create_time" `                         //是否可空:YES 提交时间
	MemberUsername        *string    `gorm:"column:member_username" json:"member_username" `                 //是否可空:YES 用户帐号
	TotalAmount           *float64   `gorm:"column:total_amount" json:"total_amount" `                       //是否可空:YES 订单总金额
	PayAmount             *float64   `gorm:"column:pay_amount" json:"pay_amount" `                           //是否可空:YES 应付金额（实际支付金额）
	FreightAmount         *float64   `gorm:"column:freight_amount" json:"freight_amount" `                   //是否可空:YES 运费金额
	PromotionAmount       *float64   `gorm:"column:promotion_amount" json:"promotion_amount" `               //是否可空:YES 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     *float64   `gorm:"column:integration_amount" json:"integration_amount" `           //是否可空:YES 积分抵扣金额
	CouponAmount          *float64   `gorm:"column:coupon_amount" json:"coupon_amount" `                     //是否可空:YES 优惠券抵扣金额
	DiscountAmount        *float64   `gorm:"column:discount_amount" json:"discount_amount" `                 //是否可空:YES 管理员后台调整订单使用的折扣金额
	PayType               *int       `gorm:"column:pay_type" json:"pay_type" `                               //是否可空:YES 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            *int       `gorm:"column:source_type" json:"source_type" `                         //是否可空:YES 订单来源：0->PC订单；1->app订单
	Status                *int       `gorm:"column:status" json:"status" `                                   //是否可空:YES 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             *int       `gorm:"column:order_type" json:"order_type" `                           //是否可空:YES 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       *string    `gorm:"column:delivery_company" json:"delivery_company" `               //是否可空:YES 物流公司(配送方式)
	DeliverySn            *string    `gorm:"column:delivery_sn" json:"delivery_sn" `                         //是否可空:YES 物流单号
	AutoConfirmDay        *int       `gorm:"column:auto_confirm_day" json:"auto_confirm_day" `               //是否可空:YES 自动确认时间（天）
	Integration           *int       `gorm:"column:integration" json:"integration" `                         //是否可空:YES 可以获得的积分
	Growth                *int       `gorm:"column:growth" json:"growth" `                                   //是否可空:YES 可以活动的成长值
	PromotionInfo         *string    `gorm:"column:promotion_info" json:"promotion_info" `                   //是否可空:YES 活动信息
	BillType              *int       `gorm:"column:bill_type" json:"bill_type" `                             //是否可空:YES 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            *string    `gorm:"column:bill_header" json:"bill_header" `                         //是否可空:YES 发票抬头
	BillContent           *string    `gorm:"column:bill_content" json:"bill_content" `                       //是否可空:YES 发票内容
	BillReceiverPhone     *string    `gorm:"column:bill_receiver_phone" json:"bill_receiver_phone" `         //是否可空:YES 收票人电话
	BillReceiverEmail     *string    `gorm:"column:bill_receiver_email" json:"bill_receiver_email" `         //是否可空:YES 收票人邮箱
	ReceiverName          string     `gorm:"column:receiver_name" json:"receiver_name" `                     //是否可空:NO 收货人姓名
	ReceiverPhone         string     `gorm:"column:receiver_phone" json:"receiver_phone" `                   //是否可空:NO 收货人电话
	ReceiverPostCode      *string    `gorm:"column:receiver_post_code" json:"receiver_post_code" `           //是否可空:YES 收货人邮编
	ReceiverProvince      *string    `gorm:"column:receiver_province" json:"receiver_province" `             //是否可空:YES 省份/直辖市
	ReceiverCity          *string    `gorm:"column:receiver_city" json:"receiver_city" `                     //是否可空:YES 城市
	ReceiverRegion        *string    `gorm:"column:receiver_region" json:"receiver_region" `                 //是否可空:YES 区
	ReceiverDetailAddress *string    `gorm:"column:receiver_detail_address" json:"receiver_detail_address" ` //是否可空:YES 详细地址
	Note                  *string    `gorm:"column:note" json:"note" `                                       //是否可空:YES 订单备注
	ConfirmStatus         *int       `gorm:"column:confirm_status" json:"confirm_status" `                   //是否可空:YES 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int        `gorm:"column:delete_status" json:"delete_status" `                     //是否可空:NO 删除状态：0->未删除；1->已删除
	UseIntegration        *int       `gorm:"column:use_integration" json:"use_integration" `                 //是否可空:YES 下单时使用的积分
	PaymentTime           *time.Time `gorm:"column:payment_time" json:"payment_time" `                       //是否可空:YES 支付时间
	DeliveryTime          *time.Time `gorm:"column:delivery_time" json:"delivery_time" `                     //是否可空:YES 发货时间
	ReceiveTime           *time.Time `gorm:"column:receive_time" json:"receive_time" `                       //是否可空:YES 确认收货时间
	CommentTime           *time.Time `gorm:"column:comment_time" json:"comment_time" `                       //是否可空:YES 评价时间
	ModifyTime            *time.Time `gorm:"column:modify_time" json:"modify_time" `                         //是否可空:YES 修改时间
}

func (*OmsOrder) TableName() string {
	return "oms_order"
}
