package model

//优惠券和产品的关系表
type SmsCouponProductRelation struct {
	Id          int     `gorm:"column:id" json:"id" `                     //是否可空:NO
	CouponId    int64   `gorm:"column:coupon_id" json:"coupon_id" `       //是否可空:YES
	ProductId   int64   `gorm:"column:product_id" json:"product_id" `     //是否可空:YES
	ProductName *string `gorm:"column:product_name" json:"product_name" ` //是否可空:YES 商品名称
	ProductSn   *string `gorm:"column:product_sn" json:"product_sn" `     //是否可空:YES 商品编码
}

func (*SmsCouponProductRelation) TableName() string {
	return "sms_coupon_product_relation"
}
