package service

import (
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"mall/global/dao/dto"
	"mall/global/dao/nosql"
	"mall/global/dao/repository"
	"mall/global/document"
	"mall/global/domain"
	"mall/global/model"
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
	Create(model.UmsMember, domain.MemberReadHistory) error
	Delete([]string) (int64, error)
	List(memberID int64, pageNum int, pageSize int) (
		p *nosql.PaginatedData[domain.MemberReadHistory], err error)
	Clear(int64) (int64, error)
}

type OmsPortalOrderService interface {
	GenerateConfirmOrder(cartIds []int64, member model.UmsMember) (result domain.ConfirmOrderResult)
	GenerateOrder(param domain.OrderParam, member model.UmsMember) (result map[string]interface{}, err error)
	PaySuccess(orderId int64, payType int) int64
	CancelTimeOutOrder() (int, error)
	CancelOrder(orderId int64)
	SendDelayMessageCancelOrder(orderId int64)
	ConfirmReceiveOrder(orderId int64, memberId int64) error
	List(memberId int64, status int, pageNum int, pageSize int) paginator.Page[domain.OmsOrderDetail]
	Detail(orderId int64) domain.OmsOrderDetail
	DeleteOrder(orderId int64, memberId int64) error
	UpdateCouponStatus(couponId int64, memberId int64, useStatus int)
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
	SetMember(model.UmsMember) error
	DelMember(int64)
	GetMember(string) (model.UmsMember, error)
	SetAuthCode(string, string)
	GetAuthCode(string) (string, error)
}

type UmsMemberService interface {
	GetById(int64) model.UmsMember
	Register(string, string, string, string) error
	GenerateAuthCode(string) string
	UpdatePassword(telephone string, password string, authCode string) (err error)
	GetCurrentMember(string) (model.UmsMember, error)
	UpdateIntegration(int64, int64)
	VerifyAuthCode(string, string) (bool, error)
	RefreshToken(string) (string, error)
	Login(string, string) (string, error)
}

type HomeService interface {
	Content() domain.HomeContentResult
	RecommendProductList(pageSize int, pageNum int) (error, paginator.Page[model.PmsProduct])
	GetProductCateList(parentId int64) []model.PmsProductCategory
	GetSubjectList(cateId int64, pageNum int, pageSize int) (error, paginator.Page[model.CmsSubject])
	HotProductList(pageNum int, pageSize int) []model.PmsProduct
	NewProductList(pageNum int, pageSize int) []model.PmsProduct
}

type MemberAttentionService interface {
	Add(attention domain.MemberBrandAttention) int

	Delete(brandId int64) int

	List(pageNum int, pageSize int) paginator.Page[domain.MemberBrandAttention]

	Detail(brandId int64) domain.MemberBrandAttention

	Clear()
}

type OmsCartItemService interface {
	//TODO transactional
	Add(cartItem model.OmsCartItem, member model.UmsMember) int64

	List(memberId int64) []model.OmsCartItem

	ListPromotion(memberId int64, cartIds []int64) []domain.CartPromotionItem

	UpdateQuantity(id int64, memberId int64, quantity int) int64

	Delete(memberId int64, ids []int64) int64

	GetCartProduct(productId int64) domain.CartProduct

	UpdateAttr(item model.OmsCartItem) int64

	Clear(MemberId int64) int64
}

type UmsMemberReceiveAddressService interface {
	Add(address model.UmsMemberReceiveAddress)
	Delete(id int64, memberId int64) int64
	//TODO transcational
	Update(id int64, address model.UmsMemberReceiveAddress, memberId int64) int64

	List(memberId int64) []model.UmsMemberReceiveAddress
	GetItem(memberId int64, id int64) model.UmsMemberReceiveAddress
}

type UmsMemberCouponService interface {
	//TODO transcational
	Add(couponId int64, member model.UmsMember) (err error)

	ListHistory(useStatus int, memberId int64) []model.SmsCouponHistory

	ListCart(cartItem []domain.CartPromotionItem, memberId int64, t int) []domain.SmsCouponHistoryDetail

	ListByProduct(productId int64) []model.SmsCoupon

	List(useStatus int, memberId int64) []model.SmsCoupon
}

type OmsPromotionService interface {
	CalcCartPromotion(cartItemList []model.OmsCartItem) []domain.CartPromotionItem
}
