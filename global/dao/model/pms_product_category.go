package model

//产品分类
type PmsProductCategory struct {
	Id           int    `gorm:"column:id" json:"id" `                       //是否可空:NO
	ParentId     int    `gorm:"column:parent_id" json:"parent_id" `         //是否可空:YES 上机分类的编号：0表示一级分类
	Name         string `gorm:"column:name" json:"name" `                   //是否可空:YES
	Level        int    `gorm:"column:level" json:"level" `                 //是否可空:YES 分类级别：0->1级；1->2级
	ProductCount int    `gorm:"column:product_count" json:"product_count" ` //是否可空:YES
	ProductUnit  string `gorm:"column:product_unit" json:"product_unit" `   //是否可空:YES
	NavStatus    int    `gorm:"column:nav_status" json:"nav_status" `       //是否可空:YES 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int    `gorm:"column:show_status" json:"show_status" `     //是否可空:YES 显示状态：0->不显示；1->显示
	Sort         int    `gorm:"column:sort" json:"sort" `                   //是否可空:YES
	Icon         string `gorm:"column:icon" json:"icon" `                   //是否可空:YES 图标
	Keywords     string `gorm:"column:keywords" json:"keywords" `           //是否可空:YES
	Description  string `gorm:"column:description" json:"description" `     //是否可空:YES 描述
}

func (*PmsProductCategory) TableName() string {
	return "pms_product_category"
}
