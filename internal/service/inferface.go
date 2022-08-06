package service

import (
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"mall/global/dao/document"
	"mall/global/dao/domain"
	"mall/global/dao/dto"
	"mall/global/dao/model"
	"mall/global/dao/nosql"
	"mall/global/dao/repository"
)

type EsProductService interface {
	ImportAll() (int, error)
	DeleteById(int) error
	DeleteByList(ids []int) (int, error)
	Create(int64) (document.EsProduct, error)
	SearchByKeyword(string, int, int) (repository.Page[document.EsProduct], error)
	SearchByDetail(string, int64, int64, int, int, int) (repository.Page[document.EsProduct], error)
	Recommend(int64, int, int) (repository.Page[document.EsProduct], error)
	SearchRelatedInfo(string) (domain.EsProductRelatedInfo, error)
}

type MemberReadHistoryService interface {
	Create(model.UmsMember, domain.MemberReadHistory) (string, error)
	Delete([]string) (int64, error)
	List(int64, int, int) (*nosql.PaginatedData[domain.MemberReadHistory], error)
	Clear(int64) (int64, error)
}

type OmsPortalOrderService interface {
	CancelTimeOutOrder() (int, error)
	UpdateCouponStatus(int, int, int)
}

type PmsBrandService interface {
	CrateBrand(model.PmsBrand) int64
	UpdateBrand(int, model.PmsBrand) int64
	DeleteBrand(int) int64
	GetBrand(int) (model.PmsBrand, error)
	ListBrand(int, int) (paginator.Page[model.PmsBrand], error)
}

type UmsAdminService interface {
	Register(dto.UmsAdminParam) (model.UmsAdmin, error)
	Login(string, string)
	LoadUserByUsername(string)
	RefreshToken(string) (string, error)
}
type UmsMemberCacheService interface {
	SetMember(model.UmsMember)
	DelMember(int64)
	GetMember(string) (model.UmsMember, error)
	SetAuthCode(string, string)
	GetAuthCode(string) (string, error)
}

type UmsMemberService interface {
	GetById(int64) model.UmsMember
	Register(string, string, string, string) error
	GenerateAuthCode(string) string
	UpdatePassword(string, string, string) error
	GetCurrentMember(string) (model.UmsMember, error)
	UpdateIntegration(int64, int)
	VerifyAuthCode(string, string) (bool, error)
	RefreshToken(string) (string, error)
	Login(string, string) (string, error)
}
