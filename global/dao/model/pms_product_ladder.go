package model

//产品阶梯价格表(只针对同商品)
type PmsProductLadder struct {
	Id        int      `gorm:"column:id" json:"id" `                 //是否可空:NO
	ProductId *int     `gorm:"column:product_id" json:"product_id" ` //是否可空:YES
	Count     *int     `gorm:"column:count" json:"count" `           //是否可空:YES 满足的商品数量
	Discount  *float64 `gorm:"column:discount" json:"discount" `     //是否可空:YES 折扣
	Price     *float64 `gorm:"column:price" json:"price" `           //是否可空:YES 折后价格
}

func (*PmsProductLadder) TableName() string {
	return "pms_product_ladder"
}
