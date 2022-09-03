package domain

/*
 * 搜索商品的关联信息，包括品牌名称，分类名称及属性

 */
type EsProductRelatedInfo struct {
	BrandNames           []string
	ProductCategoryNames []string
	ProductAttrs         []ProductAttr
}

type ProductAttr struct {
	AttrId     int
	AttrName   string
	AttrValues []string
}
