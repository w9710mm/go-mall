package service

import (
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"gorm.io/gorm"
	"mall/common/util"
	"mall/global/dao/mapper"
	"mall/global/domain"
	"mall/global/model"
	"time"
)

type homeService struct {
	db      *gorm.DB
	homeDao mapper.HomeDao
}

func NewHomeService(db *gorm.DB, dao mapper.HomeDao) HomeService {
	return &homeService{db: db, homeDao: dao}
}
func (h *homeService) Content() (result domain.HomeContentResult) {
	//获取首页广告
	result.AdvertiseList = h.getHomeAdvertiseList()
	//获取推荐品牌
	result.BranList = h.homeDao.GetRecommendBrandList()
	//获取秒杀信息
	result.HomeFlashPromotion = h.getHomeFlashPromotion()
	//获取新品推荐
	result.NewProductList = h.homeDao.GetNewProductList(0, 4)
	//获取人气推荐

	result.HotProductList = h.homeDao.GetHotProductList(0, 4)
	//获取推荐专题
	result.SubjectList = h.homeDao.GetRecommendSubjectList(0, 4)
	return
}

func (h *homeService) RecommendProductList(pageSize int, pageNum int) (error, paginator.Page[model.PmsProduct]) {
	p := paginator.Page[model.PmsProduct]{
		Pages:    int64(pageNum),
		PageSize: int64(pageSize),
	}
	tx := h.db.Model(&model.PmsProduct{}).Where(map[string]interface{}{"delete_status": 0, "publish_status": 1})
	err := p.SelectPages(tx)
	if err != nil {
		return err, p
	}
	return nil, p
}

func (h *homeService) GetProductCateList(parentId int64) (productCategories []model.PmsProductCategory) {
	h.db.Model(&model.PmsProductCategory{}).Where(map[string]interface{}{"show_status": 1, "parent_id": parentId}).
		Order("sort desc").Find(&productCategories)
	return
}

func (h *homeService) GetSubjectList(cateId int64, pageNum int, pageSize int) (error, paginator.Page[model.CmsSubject]) {
	p := paginator.Page[model.CmsSubject]{
		Pages:    int64(pageNum),
		PageSize: int64(pageSize),
	}
	tx := h.db.Model(&model.CmsSubject{}).Where(map[string]interface{}{"show_status": 1})
	if cateId != 0 {
		tx.Where(map[string]interface{}{"category_id": cateId})
	}
	err := p.SelectPages(tx)
	if err != nil {
		return err, p
	}
	return nil, p
}

func (h *homeService) HotProductList(pageNum int, pageSize int) []model.PmsProduct {
	offset := pageSize * (pageNum - 1)
	return h.homeDao.GetHotProductList(offset, pageSize)
}

func (h *homeService) NewProductList(pageNum int, pageSize int) []model.PmsProduct {
	offset := pageSize * (pageNum - 1)
	return h.homeDao.GetNewProductList(offset, pageSize)
}

func (h *homeService) getHomeAdvertiseList() (advertises []model.SmsHomeAdvertise) {

	h.db.Where(&model.SmsHomeAdvertise{Type: 1, Status: 1}).Order("sort desc").Find(&advertises)
	return
}

func (h *homeService) getHomeFlashPromotion() (homeFlashPromotion domain.HomeFlashPromotion) {
	now := time.Now()
	flashPromotions := h.getFlashPromotion(now)
	if len(flashPromotions) != 0 {
		sessions := h.getFlashPromotionSession(now)
		if len(sessions) != 0 {
			flashPromotion := &flashPromotions[0]
			session := &sessions[0]
			homeFlashPromotion.StartTime = *session.StartTime
			homeFlashPromotion.EndTime = *session.EndTime
			nextFlashPromotionSessions := h.getNextFlashPromotionSession(homeFlashPromotion.StartTime)
			if len(nextFlashPromotionSessions) != 0 {
				nextPromotionSession := nextFlashPromotionSessions[0]
				homeFlashPromotion.NextStartTime = *nextPromotionSession.StartTime
				homeFlashPromotion.NextEndTime = *nextPromotionSession.EndTime
			}
			productList := h.homeDao.GetFlashProductList(flashPromotion.Id, session.Id)
			homeFlashPromotion.ProductList = productList
		}
	}
	return
}

func (h *homeService) getFlashPromotion(date time.Time) (flashPromotionList []model.SmsFlashPromotion) {
	day := util.GetDateDayStart(date)
	h.db.Where("start_date < ? and end_date > ?", day, day).Where(&model.SmsFlashPromotion{Status: 1}).Find(flashPromotionList)
	return
}

func (h *homeService) getFlashPromotionSession(date time.Time) (flashPromotionSessionList []model.SmsFlashPromotionSession) {
	t := util.GetTimeStart(date)
	h.db.Where("start_time < ? and end_time > ? ", t, t).Find(&flashPromotionSessionList)
	return
}

func (h *homeService) getNextFlashPromotionSession(date time.Time) (flashPromotionSessionList []model.SmsFlashPromotionSession) {
	h.db.Where(" start_time > ?", date).Order(" start_time asc").Find(&flashPromotionSessionList)
	return
}
