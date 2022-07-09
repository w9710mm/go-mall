package model

import "time"

//
type PmsProductOperateLog struct {
	Id               int        `gorm:"column:id" json:"id" `                                   //是否可空:NO
	ProductId        *int       `gorm:"column:product_id" json:"product_id" `                   //是否可空:YES
	PriceOld         *float64   `gorm:"column:price_old" json:"price_old" `                     //是否可空:YES
	PriceNew         *float64   `gorm:"column:price_new" json:"price_new" `                     //是否可空:YES
	SalePriceOld     *float64   `gorm:"column:sale_price_old" json:"sale_price_old" `           //是否可空:YES
	SalePriceNew     *float64   `gorm:"column:sale_price_new" json:"sale_price_new" `           //是否可空:YES
	GiftPointOld     *int       `gorm:"column:gift_point_old" json:"gift_point_old" `           //是否可空:YES 赠送的积分
	GiftPointNew     *int       `gorm:"column:gift_point_new" json:"gift_point_new" `           //是否可空:YES
	UsePointLimitOld *int       `gorm:"column:use_point_limit_old" json:"use_point_limit_old" ` //是否可空:YES
	UsePointLimitNew *int       `gorm:"column:use_point_limit_new" json:"use_point_limit_new" ` //是否可空:YES
	OperateMan       *string    `gorm:"column:operate_man" json:"operate_man" `                 //是否可空:YES 操作人
	CreateTime       *time.Time `gorm:"column:create_time" json:"create_time" `                 //是否可空:YES
}

func (*PmsProductOperateLog) TableName() string {
	return "pms_product_operate_log"
}
