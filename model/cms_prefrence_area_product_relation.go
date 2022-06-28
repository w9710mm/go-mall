package model

//优选专区和产品关系表
type CmsPrefrenceAreaProductRelation struct {
	Id              int  `gorm:"column:id" json:"id" `                               //是否可空:NO
	PrefrenceAreaId *int `gorm:"column:prefrence_area_id" json:"prefrence_area_id" ` //是否可空:YES
	ProductId       *int `gorm:"column:product_id" json:"product_id" `               //是否可空:YES
}

func (*CmsPrefrenceAreaProductRelation) TableName() string {
	return "cms_prefrence_area_product_relation"
}
