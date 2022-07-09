package model

//优选专区
type CmsPrefrenceArea struct {
	Id         int     `gorm:"column:id" json:"id" `                   //是否可空:NO
	Name       *string `gorm:"column:name" json:"name" `               //是否可空:YES
	SubTitle   *string `gorm:"column:sub_title" json:"sub_title" `     //是否可空:YES
	Pic        *string `gorm:"column:pic" json:"pic" `                 //是否可空:YES 展示图片
	Sort       *int    `gorm:"column:sort" json:"sort" `               //是否可空:YES
	ShowStatus *int    `gorm:"column:show_status" json:"show_status" ` //是否可空:YES
}

func (*CmsPrefrenceArea) TableName() string {
	return "cms_prefrence_area"
}
