package model

import "time"

//订单操作历史记录
type OmsOrderOperateHistory struct {
	Id          int       `gorm:"column:id" json:"id" `                     //是否可空:NO
	OrderId     int       `gorm:"column:order_id" json:"order_id" `         //是否可空:YES 订单id
	OperateMan  string    `gorm:"column:operate_man" json:"operate_man" `   //是否可空:YES 操作人：用户；系统；后台管理员
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time" `   //是否可空:YES 操作时间
	OrderStatus int       `gorm:"column:order_status" json:"order_status" ` //是否可空:YES 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string    `gorm:"column:note" json:"note" `                 //是否可空:YES 备注
}

func (*OmsOrderOperateHistory) TableName() string {
	return "oms_order_operate_history"
}
