package mapper

import (
	"gorm.io/gorm"
	"mall/global/domain"
	"mall/global/model"
)

type homeDao struct {
	db *gorm.DB
}

func NewHomeDao(db *gorm.DB) HomeDao {
	return &homeDao{db: db}
}

func (h *homeDao) GetFlashProductList(flashPromotionId int, sessionId int) (promotions []domain.FlashPromotionProduct) {
	h.db.Table("    sms_flash_promotion_product_relation pr ").Joins("   LEFT JOIN pms_product p ON pr.product_id = p.id ").
		Where("   pr.flash_promotion_id = ?  AND pr.flash_promotion_session_id =? ", flashPromotionId, sessionId).
		Select("    pr.flash_promotion_price,    pr.flash_promotion_count,   pr.flash_promotion_limit,  p.*").Scan(&promotions)

	return
}

func (d *homeDao) GetRecommendBrandList() (brandList []model.PmsBrand) {
	d.db.Table("  sms_home_brand hb ").Select(" b.* ").Joins("   LEFT JOIN pms_brand b ON hb.brand_id = b.id ").
		Where("     hb.recommend_status = ?  AND b.show_status = ? ", 1, 1).Scan(&brandList)
	return
}
func (d *homeDao) GetNewProductList(offset int, limit int) (brandList []model.PmsProduct) {
	d.db.Select("p.*").Table(" sms_home_new_product hp").
		Joins("  LEFT JOIN pms_product p ON hp.product_id = p.id").
		Where(" hp.recommend_status = ? AND p.publish_status = ?", 1, 1).Order("  hp.sort DESC ").
		Offset(offset).Limit(limit).Scan(&brandList)
	return
}

func (d *homeDao) GetHotProductList(offset int, limit int) (brandList []model.PmsProduct) {
	d.db.Select("p.*").Table(" sms_home_recommend_product hp").
		Joins("  LEFT JOIN pms_product p ON hp.product_id = p.id").
		Where(" hp.recommend_status = ? AND p.publish_status = ?", 1, 1).Order("  hp.sort DESC ").
		Offset(offset).Limit(limit).Scan(&brandList)
	return
}

func (d *homeDao) GetRecommendSubjectList(offset int, limit int) (cmsSubjectList []model.CmsSubject) {
	d.db.Select("s.*").Table(" sms_home_recommend_subject hs").
		Joins("  LEFT JOIN cms_subject s ON hs.subject_id = s.id").
		Where(" hs.recommend_status = ? AND s.show_status = ?", 1, 1).Order("  hs.sort DESC ").
		Offset(offset).Limit(limit).Scan(&cmsSubjectList)
	return
}
