package model

import "time"

//退货原因表
type OmsOrderReturnReason struct {
	Id         int       `gorm:"column:id" json:"id" `                   //是否可空:NO
	Name       string    `gorm:"column:name" json:"name" `               //是否可空:YES 退货类型
	Sort       int       `gorm:"column:sort" json:"sort" `               //是否可空:YES
	Status     int       `gorm:"column:status" json:"status" `           //是否可空:YES 状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES 添加时间
}

func (*OmsOrderReturnReason) TableName() string {
	return "oms_order_return_reason"
}
