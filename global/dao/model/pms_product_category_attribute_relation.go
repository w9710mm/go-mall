package model

//产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type PmsProductCategoryAttributeRelation struct {
	Id                 int `gorm:"column:id" json:"id" `                                     //是否可空:NO
	ProductCategoryId  int `gorm:"column:product_category_id" json:"product_category_id" `   //是否可空:YES
	ProductAttributeId int `gorm:"column:product_attribute_id" json:"product_attribute_id" ` //是否可空:YES
}

func (*PmsProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}
