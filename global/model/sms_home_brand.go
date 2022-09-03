package model

//首页推荐品牌表
type SmsHomeBrand struct {
	Id              int     `gorm:"column:id" json:"id" `                             //是否可空:NO
	BrandId         *int    `gorm:"column:brand_id" json:"brand_id" `                 //是否可空:YES
	BrandName       *string `gorm:"column:brand_name" json:"brand_name" `             //是否可空:YES
	RecommendStatus *int    `gorm:"column:recommend_status" json:"recommend_status" ` //是否可空:YES
	Sort            *int    `gorm:"column:sort" json:"sort" `                         //是否可空:YES
}

func (*SmsHomeBrand) TableName() string {
	return "sms_home_brand"
}
