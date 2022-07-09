package model

import "time"

//订单退货申请
type OmsOrderReturnApply struct {
	Id               int        `gorm:"column:id" json:"id" `                                 //是否可空:NO
	OrderId          *int       `gorm:"column:order_id" json:"order_id" `                     //是否可空:YES 订单id
	CompanyAddressId *int       `gorm:"column:company_address_id" json:"company_address_id" ` //是否可空:YES 收货地址表id
	ProductId        *int       `gorm:"column:product_id" json:"product_id" `                 //是否可空:YES 退货商品id
	OrderSn          *string    `gorm:"column:order_sn" json:"order_sn" `                     //是否可空:YES 订单编号
	CreateTime       *time.Time `gorm:"column:create_time" json:"create_time" `               //是否可空:YES 申请时间
	MemberUsername   *string    `gorm:"column:member_username" json:"member_username" `       //是否可空:YES 会员用户名
	ReturnAmount     *float64   `gorm:"column:return_amount" json:"return_amount" `           //是否可空:YES 退款金额
	ReturnName       *string    `gorm:"column:return_name" json:"return_name" `               //是否可空:YES 退货人姓名
	ReturnPhone      *string    `gorm:"column:return_phone" json:"return_phone" `             //是否可空:YES 退货人电话
	Status           *int       `gorm:"column:status" json:"status" `                         //是否可空:YES 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	HandleTime       *time.Time `gorm:"column:handle_time" json:"handle_time" `               //是否可空:YES 处理时间
	ProductPic       *string    `gorm:"column:product_pic" json:"product_pic" `               //是否可空:YES 商品图片
	ProductName      *string    `gorm:"column:product_name" json:"product_name" `             //是否可空:YES 商品名称
	ProductBrand     *string    `gorm:"column:product_brand" json:"product_brand" `           //是否可空:YES 商品品牌
	ProductAttr      *string    `gorm:"column:product_attr" json:"product_attr" `             //是否可空:YES 商品销售属性：颜色：红色；尺码：xl;
	ProductCount     *int       `gorm:"column:product_count" json:"product_count" `           //是否可空:YES 退货数量
	ProductPrice     *float64   `gorm:"column:product_price" json:"product_price" `           //是否可空:YES 商品单价
	ProductRealPrice *float64   `gorm:"column:product_real_price" json:"product_real_price" ` //是否可空:YES 商品实际支付单价
	Reason           *string    `gorm:"column:reason" json:"reason" `                         //是否可空:YES 原因
	Description      *string    `gorm:"column:description" json:"description" `               //是否可空:YES 描述
	ProofPics        *string    `gorm:"column:proof_pics" json:"proof_pics" `                 //是否可空:YES 凭证图片，以逗号隔开
	HandleNote       *string    `gorm:"column:handle_note" json:"handle_note" `               //是否可空:YES 处理备注
	HandleMan        *string    `gorm:"column:handle_man" json:"handle_man" `                 //是否可空:YES 处理人员
	ReceiveMan       *string    `gorm:"column:receive_man" json:"receive_man" `               //是否可空:YES 收货人
	ReceiveTime      *time.Time `gorm:"column:receive_time" json:"receive_time" `             //是否可空:YES 收货时间
	ReceiveNote      *string    `gorm:"column:receive_note" json:"receive_note" `             //是否可空:YES 收货备注
}

func (*OmsOrderReturnApply) TableName() string {
	return "oms_order_return_apply"
}
