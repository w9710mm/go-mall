package domain

type EsProductAttributeValue struct {
	Id                 *int64  `gorm:"column:attr_id" json:"attr_id" `
	ProductAttributeID *int64  `gorm:"column:attr_product_attribute_id" json:"attr_product_attribute_id" `
	Value              *string `gorm:"column:attr_value" json:"attr_value" `
	Type               *string `gorm:"column:attr_type" json:"attr_type" `
	Name               *string `gorm:"column:attr_name" json:"attr_name" `

	//model.PmsProductAttribute
	//model.PmsProductAttributeValue
	//Id                 int64
	//ProductAttributeID int64
	//Value              string
	//Type               string
	//Name               string
}
