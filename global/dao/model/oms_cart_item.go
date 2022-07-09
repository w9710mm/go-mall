package model

import "time"

//购物车表
type OmsCartItem struct {
	Id                int        `gorm:"column:id" json:"id" `                                   //是否可空:NO
	ProductId         *int       `gorm:"column:product_id" json:"product_id" `                   //是否可空:YES
	ProductSkuId      *int       `gorm:"column:product_sku_id" json:"product_sku_id" `           //是否可空:YES
	MemberId          *int       `gorm:"column:member_id" json:"member_id" `                     //是否可空:YES
	Quantity          *int       `gorm:"column:quantity" json:"quantity" `                       //是否可空:YES 购买数量
	Price             *float64   `gorm:"column:price" json:"price" `                             //是否可空:YES 添加到购物车的价格
	Sp1               *string    `gorm:"column:sp1" json:"sp1" `                                 //是否可空:YES 销售属性1
	Sp2               *string    `gorm:"column:sp2" json:"sp2" `                                 //是否可空:YES 销售属性2
	Sp3               *string    `gorm:"column:sp3" json:"sp3" `                                 //是否可空:YES 销售属性3
	ProductPic        *string    `gorm:"column:product_pic" json:"product_pic" `                 //是否可空:YES 商品主图
	ProductName       *string    `gorm:"column:product_name" json:"product_name" `               //是否可空:YES 商品名称
	ProductSubTitle   *string    `gorm:"column:product_sub_title" json:"product_sub_title" `     //是否可空:YES 商品副标题（卖点）
	ProductSkuCode    *string    `gorm:"column:product_sku_code" json:"product_sku_code" `       //是否可空:YES 商品sku条码
	MemberNickname    *string    `gorm:"column:member_nickname" json:"member_nickname" `         //是否可空:YES 会员昵称
	CreateDate        *time.Time `gorm:"column:create_date" json:"create_date" `                 //是否可空:YES 创建时间
	ModifyDate        *time.Time `gorm:"column:modify_date" json:"modify_date" `                 //是否可空:YES 修改时间
	DeleteStatus      *int       `gorm:"column:delete_status" json:"delete_status" `             //是否可空:YES 是否删除
	ProductCategoryId *int       `gorm:"column:product_category_id" json:"product_category_id" ` //是否可空:YES 商品分类
	ProductBrand      *string    `gorm:"column:product_brand" json:"product_brand" `             //是否可空:YES
	ProductSn         *string    `gorm:"column:product_sn" json:"product_sn" `                   //是否可空:YES
	ProductAttr       *string    `gorm:"column:product_attr" json:"product_attr" `               //是否可空:YES 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

func (*OmsCartItem) TableName() string {
	return "oms_cart_item"
}
