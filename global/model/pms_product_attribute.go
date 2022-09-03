package model

//商品属性参数表
type PmsProductAttribute struct {
	Id                         int     `gorm:"column:id" json:"id" `                                                       //是否可空:NO
	ProductAttributeCategoryId *int    `gorm:"column:product_attribute_category_id" json:"product_attribute_category_id" ` //是否可空:YES
	Name                       *string `gorm:"column:name" json:"name" `                                                   //是否可空:YES
	SelectType                 *int    `gorm:"column:select_type" json:"select_type" `                                     //是否可空:YES 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  *int    `gorm:"column:input_type" json:"input_type" `                                       //是否可空:YES 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  *string `gorm:"column:input_list" json:"input_list" `                                       //是否可空:YES 可选值列表，以逗号隔开
	Sort                       *int    `gorm:"column:sort" json:"sort" `                                                   //是否可空:YES 排序字段：最高的可以单独上传图片
	FilterType                 *int    `gorm:"column:filter_type" json:"filter_type" `                                     //是否可空:YES 分类筛选样式：1->普通；1->颜色
	SearchType                 *int    `gorm:"column:search_type" json:"search_type" `                                     //是否可空:YES 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              *int    `gorm:"column:related_status" json:"related_status" `                               //是否可空:YES 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus              *int    `gorm:"column:hand_add_status" json:"hand_add_status" `                             //是否可空:YES 是否支持手动新增；0->不支持；1->支持
	Type                       *int    `gorm:"column:type" json:"type" `                                                   //是否可空:YES 属性的类型；0->规格；1->参数
}

func (*PmsProductAttribute) TableName() string {
	return "pms_product_attribute"
}
