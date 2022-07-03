package model

//产品满减表(只针对同商品)
type PmsProductFullReduction struct {
	Id          int     `gorm:"column:id" json:"id" `                     //是否可空:NO
	ProductId   int     `gorm:"column:product_id" json:"product_id" `     //是否可空:YES
	FullPrice   float64 `gorm:"column:full_price" json:"full_price" `     //是否可空:YES
	ReducePrice float64 `gorm:"column:reduce_price" json:"reduce_price" ` //是否可空:YES
}

func (*PmsProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}
