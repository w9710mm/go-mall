package model

//品牌表
type PmsBrand struct {
	Id                  int64   `gorm:"column:id" json:"id" `                                       //是否可空:NO
	Name                *string `gorm:"column:name" json:"name" `                                   //是否可空:YES
	FirstLetter         *string `gorm:"column:first_letter" json:"first_letter" `                   //是否可空:YES 首字母
	Sort                *int    `gorm:"column:sort" json:"sort" `                                   //是否可空:YES
	FactoryStatus       *int    `gorm:"column:factory_status" json:"factory_status" `               //是否可空:YES 是否为品牌制造商：0->不是；1->是
	ShowStatus          *int    `gorm:"column:show_status" json:"show_status" `                     //是否可空:YES
	ProductCount        *int    `gorm:"column:product_count" json:"product_count" `                 //是否可空:YES 产品数量
	ProductCommentCount *int    `gorm:"column:product_comment_count" json:"product_comment_count" ` //是否可空:YES 产品评论数量
	Logo                *string `gorm:"column:logo" json:"logo" `                                   //是否可空:YES 品牌logo
	BigPic              *string `gorm:"column:big_pic" json:"big_pic" `                             //是否可空:YES 专区大图
	BrandStory          *string `gorm:"column:brand_story" json:"brand_story" `                     //是否可空:YES 品牌故事
}

func (*PmsBrand) TableName() string {
	return "pms_brand"
}
