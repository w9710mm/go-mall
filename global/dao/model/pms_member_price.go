package model

//商品会员价格表
type PmsMemberPrice struct {
	Id              int      `gorm:"column:id" json:"id" `                               //是否可空:NO
	ProductId       *int     `gorm:"column:product_id" json:"product_id" `               //是否可空:YES
	MemberLevelId   *int     `gorm:"column:member_level_id" json:"member_level_id" `     //是否可空:YES
	MemberPrice     *float64 `gorm:"column:member_price" json:"member_price" `           //是否可空:YES 会员价格
	MemberLevelName *string  `gorm:"column:member_level_name" json:"member_level_name" ` //是否可空:YES
}

func (*PmsMemberPrice) TableName() string {
	return "pms_member_price"
}
