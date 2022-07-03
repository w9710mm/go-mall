package model

//订单设置表
type OmsOrderSetting struct {
	Id                  int `gorm:"column:id" json:"id" `                                       //是否可空:NO
	FlashOrderOvertime  int `gorm:"column:flash_order_overtime" json:"flash_order_overtime" `   //是否可空:YES 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int `gorm:"column:normal_order_overtime" json:"normal_order_overtime" ` //是否可空:YES 正常订单超时时间(分)
	ConfirmOvertime     int `gorm:"column:confirm_overtime" json:"confirm_overtime" `           //是否可空:YES 发货后自动确认收货时间（天）
	FinishOvertime      int `gorm:"column:finish_overtime" json:"finish_overtime" `             //是否可空:YES 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int `gorm:"column:comment_overtime" json:"comment_overtime" `           //是否可空:YES 订单完成后自动好评时间（天）
}

func (*OmsOrderSetting) TableName() string {
	return "oms_order_setting"
}
