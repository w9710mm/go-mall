package domain

import "mall/global/model"

type HomeContentResult struct {
	AdvertiseList      []model.SmsHomeAdvertise
	BranList           []model.PmsBrand
	HomeFlashPromotion HomeFlashPromotion
	NewProductList     []model.PmsProduct
	HotProductList     []model.PmsProduct
	SubjectList        []model.CmsSubject
}
