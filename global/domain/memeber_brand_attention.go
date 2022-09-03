package domain

import "time"

type MemberBrandAttention struct {
	Id             string
	MemberId       string
	MemberNickName string
	BrandId        string
	BrandName      string
	BrandLogo      string
	BrandCity      string
	CreateTime     time.Time
}
