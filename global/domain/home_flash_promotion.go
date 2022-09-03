package domain

import "time"

type HomeFlashPromotion struct {
	StartTime     time.Time
	EndTime       time.Time
	NextStartTime time.Time
	NextEndTime   time.Time
	ProductList   []FlashPromotionProduct
}
