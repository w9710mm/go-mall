package model

//sku的库存
type PmsSkuStock struct {
	Id             int     `gorm:"column:id" json:"id" `                           //是否可空:NO
	ProductId      int     `gorm:"column:product_id" json:"product_id" `           //是否可空:YES
	SkuCode        string  `gorm:"column:sku_code" json:"sku_code" `               //是否可空:NO sku编码
	Price          float64 `gorm:"column:price" json:"price" `                     //是否可空:YES
	Stock          int     `gorm:"column:stock" json:"stock" `                     //是否可空:YES 库存
	LowStock       int     `gorm:"column:low_stock" json:"low_stock" `             //是否可空:YES 预警库存
	Sp1            string  `gorm:"column:sp1" json:"sp1" `                         //是否可空:YES 销售属性1
	Sp2            string  `gorm:"column:sp2" json:"sp2" `                         //是否可空:YES
	Sp3            string  `gorm:"column:sp3" json:"sp3" `                         //是否可空:YES
	Pic            string  `gorm:"column:pic" json:"pic" `                         //是否可空:YES 展示图片
	Sale           int     `gorm:"column:sale" json:"sale" `                       //是否可空:YES 销量
	PromotionPrice float64 `gorm:"column:promotion_price" json:"promotion_price" ` //是否可空:YES 单品促销价格
	LockStock      int     `gorm:"column:lock_stock" json:"lock_stock" `           //是否可空:YES 锁定库存
}

func (*PmsSkuStock) TableName() string {
	return "pms_sku_stock"
}
