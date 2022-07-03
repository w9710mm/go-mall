package model

//商品限时购与商品关系表
type SmsFlashPromotionProductRelation struct {
	Id                      int     `gorm:"column:id" json:"id" `                                                 //是否可空:NO 编号
	FlashPromotionId        int     `gorm:"column:flash_promotion_id" json:"flash_promotion_id" `                 //是否可空:YES
	FlashPromotionSessionId int     `gorm:"column:flash_promotion_session_id" json:"flash_promotion_session_id" ` //是否可空:YES 编号
	ProductId               int     `gorm:"column:product_id" json:"product_id" `                                 //是否可空:YES
	FlashPromotionPrice     float64 `gorm:"column:flash_promotion_price" json:"flash_promotion_price" `           //是否可空:YES 限时购价格
	FlashPromotionCount     int     `gorm:"column:flash_promotion_count" json:"flash_promotion_count" `           //是否可空:YES 限时购数量
	FlashPromotionLimit     int     `gorm:"column:flash_promotion_limit" json:"flash_promotion_limit" `           //是否可空:YES 每人限购数量
	Sort                    int     `gorm:"column:sort" json:"sort" `                                             //是否可空:YES 排序
}

func (*SmsFlashPromotionProductRelation) TableName() string {
	return "sms_flash_promotion_product_relation"
}
