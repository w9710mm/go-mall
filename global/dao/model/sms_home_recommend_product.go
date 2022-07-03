package model

//人气推荐商品表
type SmsHomeRecommendProduct struct {
	Id              int    `gorm:"column:id" json:"id" `                             //是否可空:NO
	ProductId       int    `gorm:"column:product_id" json:"product_id" `             //是否可空:YES
	ProductName     string `gorm:"column:product_name" json:"product_name" `         //是否可空:YES
	RecommendStatus int    `gorm:"column:recommend_status" json:"recommend_status" ` //是否可空:YES
	Sort            int    `gorm:"column:sort" json:"sort" `                         //是否可空:YES
}

func (*SmsHomeRecommendProduct) TableName() string {
	return "sms_home_recommend_product"
}
