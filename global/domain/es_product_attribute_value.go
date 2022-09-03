package domain

type EsProductAttributeValue struct {
	Id                 *int64  `gorm:"column:attr_id" json:"id" `
	ProductAttributeID *int64  `gorm:"column:attr_product_attribute_id" json:"productAttributeId" `
	Value              *string `gorm:"column:attr_value" json:"value" `
	Type               *int    `gorm:"column:attr_type" json:"type" `
	Name               *string `gorm:"column:attr_name" json:"name" `

	//model.PmsProductAttribute
	//model.PmsProductAttributeValue
	//Id                 int64
	//ProductAttributeID int64
	//Value              string
	//Type               string
	//Name               string
}
