package model

//优惠券和产品分类关系表
type SmsCouponProductCategoryRelation struct {
	Id                  int    `gorm:"column:id" json:"id" `                                       //是否可空:NO
	CouponId            int    `gorm:"column:coupon_id" json:"coupon_id" `                         //是否可空:YES
	ProductCategoryId   int    `gorm:"column:product_category_id" json:"product_category_id" `     //是否可空:YES
	ProductCategoryName string `gorm:"column:product_category_name" json:"product_category_name" ` //是否可空:YES 产品分类名称
	ParentCategoryName  string `gorm:"column:parent_category_name" json:"parent_category_name" `   //是否可空:YES 父分类名称
}

func (*SmsCouponProductCategoryRelation) TableName() string {
	return "sms_coupon_product_category_relation"
}
