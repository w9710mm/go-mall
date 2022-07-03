package model

//订单中所包含的商品
type OmsOrderItem struct {
	Id                int     `gorm:"column:id" json:"id" `                                   //是否可空:NO
	OrderId           int     `gorm:"column:order_id" json:"order_id" `                       //是否可空:YES 订单id
	OrderSn           string  `gorm:"column:order_sn" json:"order_sn" `                       //是否可空:YES 订单编号
	ProductId         int     `gorm:"column:product_id" json:"product_id" `                   //是否可空:YES
	ProductPic        string  `gorm:"column:product_pic" json:"product_pic" `                 //是否可空:YES
	ProductName       string  `gorm:"column:product_name" json:"product_name" `               //是否可空:YES
	ProductBrand      string  `gorm:"column:product_brand" json:"product_brand" `             //是否可空:YES
	ProductSn         string  `gorm:"column:product_sn" json:"product_sn" `                   //是否可空:YES
	ProductPrice      float64 `gorm:"column:product_price" json:"product_price" `             //是否可空:YES 销售价格
	ProductQuantity   int     `gorm:"column:product_quantity" json:"product_quantity" `       //是否可空:YES 购买数量
	ProductSkuId      int     `gorm:"column:product_sku_id" json:"product_sku_id" `           //是否可空:YES 商品sku编号
	ProductSkuCode    string  `gorm:"column:product_sku_code" json:"product_sku_code" `       //是否可空:YES 商品sku条码
	ProductCategoryId int     `gorm:"column:product_category_id" json:"product_category_id" ` //是否可空:YES 商品分类id
	Sp1               string  `gorm:"column:sp1" json:"sp1" `                                 //是否可空:YES 商品的销售属性
	Sp2               string  `gorm:"column:sp2" json:"sp2" `                                 //是否可空:YES
	Sp3               string  `gorm:"column:sp3" json:"sp3" `                                 //是否可空:YES
	PromotionName     string  `gorm:"column:promotion_name" json:"promotion_name" `           //是否可空:YES 商品促销名称
	PromotionAmount   float64 `gorm:"column:promotion_amount" json:"promotion_amount" `       //是否可空:YES 商品促销分解金额
	CouponAmount      float64 `gorm:"column:coupon_amount" json:"coupon_amount" `             //是否可空:YES 优惠券优惠分解金额
	IntegrationAmount float64 `gorm:"column:integration_amount" json:"integration_amount" `   //是否可空:YES 积分优惠分解金额
	RealAmount        float64 `gorm:"column:real_amount" json:"real_amount" `                 //是否可空:YES 该商品经过优惠后的分解金额
	GiftIntegration   int     `gorm:"column:gift_integration" json:"gift_integration" `       //是否可空:YES
	GiftGrowth        int     `gorm:"column:gift_growth" json:"gift_growth" `                 //是否可空:YES
	ProductAttr       string  `gorm:"column:product_attr" json:"product_attr" `               //是否可空:YES 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

func (*OmsOrderItem) TableName() string {
	return "oms_order_item"
}
