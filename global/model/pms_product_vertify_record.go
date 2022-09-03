package model

import "time"

//商品审核记录
type PmsProductVertifyRecord struct {
	Id         int        `gorm:"column:id" json:"id" `                   //是否可空:NO
	ProductId  *int       `gorm:"column:product_id" json:"product_id" `   //是否可空:YES
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time" ` //是否可空:YES
	VertifyMan *string    `gorm:"column:vertify_man" json:"vertify_man" ` //是否可空:YES 审核人
	Status     *int       `gorm:"column:status" json:"status" `           //是否可空:YES
	Detail     *string    `gorm:"column:detail" json:"detail" `           //是否可空:YES 反馈详情
}

func (*PmsProductVertifyRecord) TableName() string {
	return "pms_product_vertify_record"
}
