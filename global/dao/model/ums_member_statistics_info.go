package model

import "time"

//会员统计信息
type UmsMemberStatisticsInfo struct {
	Id                  int       `gorm:"column:id" json:"id" `                                       //是否可空:NO
	MemberId            int       `gorm:"column:member_id" json:"member_id" `                         //是否可空:YES
	ConsumeAmount       float64   `gorm:"column:consume_amount" json:"consume_amount" `               //是否可空:YES 累计消费金额
	OrderCount          int       `gorm:"column:order_count" json:"order_count" `                     //是否可空:YES 订单数量
	CouponCount         int       `gorm:"column:coupon_count" json:"coupon_count" `                   //是否可空:YES 优惠券数量
	CommentCount        int       `gorm:"column:comment_count" json:"comment_count" `                 //是否可空:YES 评价数
	ReturnOrderCount    int       `gorm:"column:return_order_count" json:"return_order_count" `       //是否可空:YES 退货数量
	LoginCount          int       `gorm:"column:login_count" json:"login_count" `                     //是否可空:YES 登录次数
	AttendCount         int       `gorm:"column:attend_count" json:"attend_count" `                   //是否可空:YES 关注数量
	FansCount           int       `gorm:"column:fans_count" json:"fans_count" `                       //是否可空:YES 粉丝数量
	CollectProductCount int       `gorm:"column:collect_product_count" json:"collect_product_count" ` //是否可空:YES
	CollectSubjectCount int       `gorm:"column:collect_subject_count" json:"collect_subject_count" ` //是否可空:YES
	CollectTopicCount   int       `gorm:"column:collect_topic_count" json:"collect_topic_count" `     //是否可空:YES
	CollectCommentCount int       `gorm:"column:collect_comment_count" json:"collect_comment_count" ` //是否可空:YES
	InviteFriendCount   int       `gorm:"column:invite_friend_count" json:"invite_friend_count" `     //是否可空:YES
	RecentOrderTime     time.Time `gorm:"column:recent_order_time" json:"recent_order_time" `         //是否可空:YES 最后一次下订单时间
}

func (*UmsMemberStatisticsInfo) TableName() string {
	return "ums_member_statistics_info"
}
