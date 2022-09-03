package model

//会员与产品分类关系表（用户喜欢的分类）
type UmsMemberProductCategoryRelation struct {
	Id                int  `gorm:"column:id" json:"id" `                                   //是否可空:NO
	MemberId          *int `gorm:"column:member_id" json:"member_id" `                     //是否可空:YES
	ProductCategoryId *int `gorm:"column:product_category_id" json:"product_category_id" ` //是否可空:YES
}

func (*UmsMemberProductCategoryRelation) TableName() string {
	return "ums_member_product_category_relation"
}
