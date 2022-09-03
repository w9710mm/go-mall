package mapper

import (
	"gorm.io/gorm"
	"mall/global/domain"
	"mall/global/model"
)

/**
 *@author:
 *@date:2022/8/17
**/

type portalProductDao struct {
	db *gorm.DB
}

func (d *portalProductDao) GetCartProduct(id int64) (cart domain.CartProduct) {
	d.db.Model(&model.PmsProduct{}).Where(id).Preload("ProductAttributeList", " type =0", "order by sort desc ").
		Preload("SkuStockList").Select("id ,`name`, sub_title, price,pic,product_attribute_category_id,stock").Find(&cart)
	return
}

func (d *portalProductDao) GetPromotionProductList(ids []int64) (product []domain.PromotionProduct) {
	d.db.Model(&model.PmsProduct{}).Where(ids).Preload("SkuStockList").
		Preload("ProductLadderList").Preload("productFullReductionList").
		Select("id, `name`, promotion_type, gift_growth,gift_point ").Find(&product)
	return
}

func (d *portalProductDao) GetAvailableCouponList(productId int64, productCategoryId int64) (coupon []model.SmsCoupon) {
	d.db.Raw("  SELECT *"+
		"      FROM sms_coupon"+
		"        WHERE use_type = 0"+
		"          AND start_time < NOW()"+
		"          AND end_time > NOW()"+
		"       UNION"+
		"        ("+
		"            SELECT c.*"+
		"            FROM sms_coupon_product_category_relation cpc"+
		"                     LEFT JOIN sms_coupon c ON cpc.coupon_id = c.id"+
		"            WHERE c.use_type = 1"+
		"              AND c.start_time < NOW()"+
		"              AND c.end_time > NOW()"+
		"              AND cpc.product_category_id = ?"+
		"        )"+
		"        UNION"+
		"       ("+
		"            SELECT c.*"+
		"            FROM sms_coupon_product_relation cp"+
		"                     LEFT JOIN sms_coupon c ON cp.coupon_id = c.id"+
		"            WHERE c.use_type = 2"+
		"              AND c.start_time < NOW()"+
		"              AND c.end_time > NOW()"+
		"              AND cp.product_id = ?"+
		"        )", productCategoryId, productId).Find(&coupon)
	return
}

func NewPortalProductDao(db *gorm.DB) PortalProductDao {
	return &portalProductDao{db: db}
}
