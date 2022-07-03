package model

//帮助分类表
type CmsHelpCategory struct {
	Id         int    `gorm:"column:id" json:"id" `                   //是否可空:NO
	Name       string `gorm:"column:name" json:"name" `               //是否可空:YES
	Icon       string `gorm:"column:icon" json:"icon" `               //是否可空:YES 分类图标
	HelpCount  int    `gorm:"column:help_count" json:"help_count" `   //是否可空:YES 专题数量
	ShowStatus int    `gorm:"column:show_status" json:"show_status" ` //是否可空:YES
	Sort       int    `gorm:"column:sort" json:"sort" `               //是否可空:YES
}

func (*CmsHelpCategory) TableName() string {
	return "cms_help_category"
}
