package domain

type OrderParam struct {
	MemberReceiveAddressId int64
	CouponId               int64
	UseIntegration         int64
	PayType                int
	SourceType             int
	CartIds                []int64
}
