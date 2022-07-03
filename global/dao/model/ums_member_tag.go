package model

//用户标签表
type UmsMemberTag struct {
	Id                int     `gorm:"column:id" json:"id" `                                   //是否可空:NO
	Name              string  `gorm:"column:name" json:"name" `                               //是否可空:YES
	FinishOrderCount  int     `gorm:"column:finish_order_count" json:"finish_order_count" `   //是否可空:YES 自动打标签完成订单数量
	FinishOrderAmount float64 `gorm:"column:finish_order_amount" json:"finish_order_amount" ` //是否可空:YES 自动打标签完成订单金额
}

func (*UmsMemberTag) TableName() string {
	return "ums_member_tag"
}
