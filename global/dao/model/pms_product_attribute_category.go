package model

//产品属性分类表
type PmsProductAttributeCategory struct {
	Id             int    `gorm:"column:id" json:"id" `                           //是否可空:NO
	Name           string `gorm:"column:name" json:"name" `                       //是否可空:YES
	AttributeCount int    `gorm:"column:attribute_count" json:"attribute_count" ` //是否可空:YES 属性数量
	ParamCount     int    `gorm:"column:param_count" json:"param_count" `         //是否可空:YES 参数数量
}

func (*PmsProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}
