package service

import (
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"mall/global/domain"
)

type memberAttentionService struct {
	umsMemberService UmsMemberService
}

func NewMemberAttentionService(service UmsMemberService) MemberAttentionService {
	return &memberAttentionService{umsMemberService: service}
}

func (m memberAttentionService) Add(attention domain.MemberBrandAttention) int {
	//TODO implement me
	panic("implement me")
}

func (m memberAttentionService) Delete(brandId int64) int {
	//TODO implement me
	panic("implement me")
}

func (m memberAttentionService) List(pageNum int, pageSize int) paginator.Page[domain.MemberBrandAttention] {
	//TODO implement me
	panic("implement me")
}

func (m memberAttentionService) Detail(brandId int64) domain.MemberBrandAttention {
	//TODO implement me
	panic("implement me")
}

func (m memberAttentionService) Clear() {
	//TODO implement me
	panic("implement me")
}
