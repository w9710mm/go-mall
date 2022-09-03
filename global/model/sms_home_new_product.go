package model

//新鲜好物表
type SmsHomeNewProduct struct {
	Id              int     `gorm:"column:id" json:"id" `                             //是否可空:NO
	ProductId       *int    `gorm:"column:product_id" json:"product_id" `             //是否可空:YES
	ProductName     *string `gorm:"column:product_name" json:"product_name" `         //是否可空:YES
	RecommendStatus *int    `gorm:"column:recommend_status" json:"recommend_status" ` //是否可空:YES
	Sort            *int    `gorm:"column:sort" json:"sort" `                         //是否可空:YES
}

func (*SmsHomeNewProduct) TableName() string {
	return "sms_home_new_product"
}
