package model

//存储产品参数信息的表
type PmsProductAttributeValue struct {
	Id                 int    `gorm:"column:id" json:"id" `                                     //是否可空:NO
	ProductId          int    `gorm:"column:product_id" json:"product_id" `                     //是否可空:YES
	ProductAttributeId int    `gorm:"column:product_attribute_id" json:"product_attribute_id" ` //是否可空:YES
	Value              string `gorm:"column:value" json:"value" `                               //是否可空:YES 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}

func (*PmsProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}
