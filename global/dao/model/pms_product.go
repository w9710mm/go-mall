package model

import "time"

//商品信息
type PmsProduct struct {
	Id                         int        `gorm:"column:id" json:"id" `                                                       //是否可空:NO
	BrandId                    *int       `gorm:"column:brand_id" json:"brand_id" `                                           //是否可空:YES
	ProductCategoryId          *int       `gorm:"column:product_category_id" json:"product_category_id" `                     //是否可空:YES
	FeightTemplateId           *int       `gorm:"column:feight_template_id" json:"feight_template_id" `                       //是否可空:YES
	ProductAttributeCategoryId *int       `gorm:"column:product_attribute_category_id" json:"product_attribute_category_id" ` //是否可空:YES
	Name                       string     `gorm:"column:name" json:"name" `                                                   //是否可空:NO
	Pic                        *string    `gorm:"column:pic" json:"pic" `                                                     //是否可空:YES
	ProductSn                  string     `gorm:"column:product_sn" json:"product_sn" `                                       //是否可空:NO 货号
	DeleteStatus               *int       `gorm:"column:delete_status" json:"delete_status" `                                 //是否可空:YES 删除状态：0->未删除；1->已删除
	PublishStatus              *int       `gorm:"column:publish_status" json:"publish_status" `                               //是否可空:YES 上架状态：0->下架；1->上架
	NewStatus                  *int       `gorm:"column:new_status" json:"new_status" `                                       //是否可空:YES 新品状态:0->不是新品；1->新品
	RecommandStatus            *int       `gorm:"column:recommand_status" json:"recommand_status" `                           //是否可空:YES 推荐状态；0->不推荐；1->推荐
	VerifyStatus               *int       `gorm:"column:verify_status" json:"verify_status" `                                 //是否可空:YES 审核状态：0->未审核；1->审核通过
	Sort                       *int       `gorm:"column:sort" json:"sort" `                                                   //是否可空:YES 排序
	Sale                       *int       `gorm:"column:sale" json:"sale" `                                                   //是否可空:YES 销量
	Price                      *float64   `gorm:"column:price" json:"price" `                                                 //是否可空:YES
	PromotionPrice             *float64   `gorm:"column:promotion_price" json:"promotion_price" `                             //是否可空:YES 促销价格
	GiftGrowth                 *int       `gorm:"column:gift_growth" json:"gift_growth" `                                     //是否可空:YES 赠送的成长值
	GiftPoint                  *int       `gorm:"column:gift_point" json:"gift_point" `                                       //是否可空:YES 赠送的积分
	UsePointLimit              *int       `gorm:"column:use_point_limit" json:"use_point_limit" `                             //是否可空:YES 限制使用的积分数
	SubTitle                   *string    `gorm:"column:sub_title" json:"sub_title" `                                         //是否可空:YES 副标题
	Description                *string    `gorm:"column:description" json:"description" `                                     //是否可空:YES 商品描述
	OriginalPrice              *float64   `gorm:"column:original_price" json:"original_price" `                               //是否可空:YES 市场价
	Stock                      *int       `gorm:"column:stock" json:"stock" `                                                 //是否可空:YES 库存
	LowStock                   *int       `gorm:"column:low_stock" json:"low_stock" `                                         //是否可空:YES 库存预警值
	Unit                       *string    `gorm:"column:unit" json:"unit" `                                                   //是否可空:YES 单位
	Weight                     *float64   `gorm:"column:weight" json:"weight" `                                               //是否可空:YES 商品重量，默认为克
	PreviewStatus              *int       `gorm:"column:preview_status" json:"preview_status" `                               //是否可空:YES 是否为预告商品：0->不是；1->是
	ServiceIds                 *string    `gorm:"column:service_ids" json:"service_ids" `                                     //是否可空:YES 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   *string    `gorm:"column:keywords" json:"keywords" `                                           //是否可空:YES
	Note                       *string    `gorm:"column:note" json:"note" `                                                   //是否可空:YES
	AlbumPics                  *string    `gorm:"column:album_pics" json:"album_pics" `                                       //是否可空:YES 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                *string    `gorm:"column:detail_title" json:"detail_title" `                                   //是否可空:YES
	DetailDesc                 *string    `gorm:"column:detail_desc" json:"detail_desc" `                                     //是否可空:YES
	DetailHtml                 *string    `gorm:"column:detail_html" json:"detail_html" `                                     //是否可空:YES 产品详情网页内容
	DetailMobileHtml           *string    `gorm:"column:detail_mobile_html" json:"detail_mobile_html" `                       //是否可空:YES 移动端网页详情
	PromotionStartTime         *time.Time `gorm:"column:promotion_start_time" json:"promotion_start_time" `                   //是否可空:YES 促销开始时间
	PromotionEndTime           *time.Time `gorm:"column:promotion_end_time" json:"promotion_end_time" `                       //是否可空:YES 促销结束时间
	PromotionPerLimit          *int       `gorm:"column:promotion_per_limit" json:"promotion_per_limit" `                     //是否可空:YES 活动限购数量
	PromotionType              *int       `gorm:"column:promotion_type" json:"promotion_type" `                               //是否可空:YES 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  *string    `gorm:"column:brand_name" json:"brand_name" `                                       //是否可空:YES 品牌名称
	ProductCategoryName        *string    `gorm:"column:product_category_name" json:"product_category_name" `                 //是否可空:YES 商品分类名称
}

func (*PmsProduct) TableName() string {
	return "pms_product"
}
