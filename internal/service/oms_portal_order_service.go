package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
	"github.com/wxnacy/wgo/arrays"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"gorm.io/gorm"
	"mall/global/config"
	"mall/global/dao/mapper"
	"mall/global/domain"
	"mall/global/model"
	"mall/internal/componet/message_queue"
	"strconv"
	"strings"
	time "time"
)

type omsPortalOrderService struct {
	portalOrderDao              mapper.PortalOrderDao
	memberService               UmsMemberService
	CartItemService             OmsCartItemService
	memberReceiveAddressService UmsMemberReceiveAddressService
	memberCouponService         UmsMemberCouponService
	redisKeyOrderId             string
	redisDATABASE               string
	orderMQ                     message_queue.OrderMQ
	redisdb                     *redis.Client
	db                          *gorm.DB
}

func NewOmsPortalOrderService(m mapper.PortalOrderDao, memberService UmsMemberService,
	memberReceiveAddressService UmsMemberReceiveAddressService, CartItemService OmsCartItemService,
	couponService UmsMemberCouponService, orderMQ message_queue.OrderMQ,
	redisdb *redis.Client, db *gorm.DB) OmsPortalOrderService {
	return &omsPortalOrderService{
		portalOrderDao:              m,
		memberService:               memberService,
		CartItemService:             CartItemService,
		memberReceiveAddressService: memberReceiveAddressService,
		memberCouponService:         couponService,
		redisKeyOrderId:             config.GetConfig().Redis.Key.OrderId,
		redisDATABASE:               strconv.Itoa(config.GetConfig().Redis.DB),
		orderMQ:                     orderMQ,
		redisdb:                     redisdb,
		db:                          db,
	}
}

func (s *omsPortalOrderService) CancelTimeOutOrder() (count int, err error) {

	var orderSetting model.OmsOrderSetting
	s.db.First(&orderSetting)
	timeOutOrders, err := s.portalOrderDao.GetTimeOutOrders(*orderSetting.NormalOrderOvertime)

	if err != nil {
		return 0, err
	}
	if timeOutOrders == nil && len(timeOutOrders) == 0 {
		return count, nil
	}
	ids := make([]int64, len(timeOutOrders))
	for i, order := range timeOutOrders {
		ids[i] = order.Id
	}

	err = s.portalOrderDao.UpdateOrderStatus(ids, 4)
	if err != nil {
		return 0, err
	}
	for _, order := range timeOutOrders {
		s.portalOrderDao.ReleaseSkuStockLock(order.OrderItemList)
		s.UpdateCouponStatus(order.CouponId, order.MemberId, 0)
		if order.UseIntegration != 0 {
			var user model.UmsMember
			s.db.First(&user, order.MemberId)
			user.Integration = user.Integration + order.Integration
			s.db.Save(&user)
		}
	}
	count = len(timeOutOrders)
	return
}

func (s omsPortalOrderService) UpdateCouponStatus(couponId int64, memberId int64, useStatus int) {
	if couponId == 0 {
		return
	}
	if useStatus == 0 {
		useStatus = 1
	} else {
		useStatus = 0
	}

	var coupons []model.SmsCouponHistory
	s.db.Model(&model.SmsCouponHistory{}).Where(&model.SmsCouponHistory{MemberId: memberId,
		CouponId: couponId, UseStatus: useStatus,
	}).Find(&coupons)

	if coupons != nil && len(coupons) != 0 {
		coupon := coupons[0]
		t := time.Now()
		coupon.UseTime = &t
		coupon.UseStatus = useStatus
		s.db.Save(&coupon)
	}

}
func (s *omsPortalOrderService) GenerateConfirmOrder(cartIds []int64, member model.UmsMember) (result domain.ConfirmOrderResult) {
	promotion := s.CartItemService.ListPromotion(member.Id, cartIds)
	result.CartPromotionItemList = promotion

	addresses := s.memberReceiveAddressService.List(member.Id)
	result.MemberReceiveAddressList = addresses

	couponHistoryDetails := s.memberCouponService.ListCart(promotion, member.Id, 1)
	result.CouponHistoryDetailList = couponHistoryDetails

	result.MemberIntegration = member.Integration

	var integrationConsumeSetting model.UmsIntegrationConsumeSetting
	s.db.Model(&model.UmsIntegrationConsumeSetting{}).Where(1).First(&integrationConsumeSetting)
	result.IntegrationConsumeSetting = integrationConsumeSetting
	s.CalCartPromotion(promotion)
	amount := s.CalcCartAmount(promotion)
	result.CalcAmount = amount
	return
}

func (s *omsPortalOrderService) CalcCartAmount(cartPromotionList []domain.CartPromotionItem) (calc domain.CalcAmount) {
	calc.FreightAmount = decimal.NewFromInt(0)
	totalAmount := decimal.NewFromInt(0)
	promotionAmount := decimal.NewFromInt(0)
	for _, item := range cartPromotionList {
		totalAmount = totalAmount.Add(item.Price.Mul(decimal.NewFromInt32(int32(item.Quantity))))
		promotionAmount = promotionAmount.Add(item.ReduceAmount.Mul(decimal.NewFromInt32(int32(item.Quantity))))
	}
	calc.TotalAmount = totalAmount
	calc.PromotionAmount = promotionAmount
	calc.PayAmount = totalAmount.Sub(promotionAmount)
	return
}
func (s *omsPortalOrderService) GenerateOrder(param domain.OrderParam, member model.UmsMember) (result map[string]interface{},
	err error) {
	var oderItemList []model.OmsOrderItem
	cartPromotionItemList := s.CartItemService.ListPromotion(member.Id, param.CartIds)
	for _, item := range cartPromotionItemList {
		var orderItem model.OmsOrderItem
		orderItem.ProductId = item.ProductId
		orderItem.ProductName = item.ProductName
		orderItem.ProductPic = item.ProductPic
		orderItem.ProductAttr = item.ProductAttr
		orderItem.ProductBrand = item.ProductBrand
		orderItem.ProductSn = item.ProductSn
		orderItem.ProductPrice = item.Price
		orderItem.ProductQuantity = item.Quantity
		orderItem.ProductSkuId = item.ProductSkuId
		orderItem.ProductSkuCode = item.ProductSkuCode
		orderItem.ProductCategoryId = item.ProductCategoryId
		orderItem.PromotionAmount = item.ReduceAmount
		orderItem.PromotionName = item.PromotionMessage
		orderItem.GiftIntegration = item.Integration
		orderItem.GiftGrowth = item.Growth
		oderItemList = append(oderItemList, orderItem)
	}
	if s.HasStock(cartPromotionItemList) {
		err = errors.New("there is not enough stock to place orders")
		return
	}

	if param.CouponId == 0 {
		for _, item := range oderItemList {
			item.CouponAmount = decimal.NewFromInt(0)
		}
	} else {
		couponHistoryDetail := s.GetUseCoupon(cartPromotionItemList, member.Id, param.CouponId)
		if couponHistoryDetail.Id == 0 {
			err = errors.New("can't use this coupon")
			return
		}
		s.HandleCouponAmount(oderItemList, couponHistoryDetail)
	}
	if param.UseIntegration == 0 {
		for _, orderItem := range oderItemList {
			orderItem.IntegrationAmount = decimal.NewFromInt(0)
		}
	} else {
		total := s.CalcTotalAmount(oderItemList)
		integrationAmount := s.GetUseIntegrationAmount(param.UseIntegration, total, member, param.CouponId != 0)
		if integrationAmount.Equal(decimal.NewFromInt32(0)) {
			err = errors.New("integration can`t use")
			return
		} else {
			for _, item := range oderItemList {
				perAmount := item.ProductPrice.DivRound(total, 3).Mul(integrationAmount)
				item.IntegrationAmount = perAmount
			}
		}
	}
	s.HandleRealAmount(oderItemList)
	s.LockStock(cartPromotionItemList)
	var order model.OmsOrder
	order.DiscountAmount = decimal.NewFromInt(0)
	order.TotalAmount = s.CalcTotalAmount(oderItemList)
	order.FreightAmount = decimal.NewFromInt(0)
	order.PromotionAmount = s.CalcPromotionAmount(oderItemList)
	order.PromotionInfo = s.GetOrderPromotionInfo(oderItemList)
	if param.CouponId == 0 {
		order.CouponAmount = decimal.NewFromInt(0)
	} else {
		order.CouponId = param.CouponId
		order.CouponAmount = s.CalcCouponAmount(oderItemList)
	}
	if param.UseIntegration == 0 {
		order.Integration = 0
		order.IntegrationAmount = decimal.NewFromInt(0)
	} else {
		order.Integration = param.UseIntegration
		order.IntegrationAmount = s.CalcIntegrationAmount(oderItemList)
	}
	order.PayAmount = s.CalcPayAmount(order)
	order.MemberId = member.Id
	var t = time.Now()
	order.CreateTime = &t
	order.MemberUsername = member.Username
	order.PayType = param.PayType
	order.SourceType = param.SourceType
	order.Status = 0
	order.OrderType = 0
	address := s.memberReceiveAddressService.GetItem(member.Id, param.MemberReceiveAddressId)
	order.ReceiverName = address.Name
	order.ReceiverPhone = address.PhoneNumber
	order.ReceiverPostCode = address.PostCode
	order.ReceiverProvince = address.Province
	order.ReceiverCity = address.City
	order.ReceiverRegion = address.Region
	order.ReceiverDetailAddress = address.DetailAddress
	order.ConfirmStatus = 0
	order.DeleteStatus = 0
	order.Integration = s.CalcGifIntegration(oderItemList)
	order.Growth = s.CalcGifTGrowth(oderItemList)
	order.OrderSn = s.GenrateOrderSn(order)
	var orderSettings model.OmsOrderSetting
	s.db.Save(&orderSettings).First(&orderSettings)
	order.AutoConfirmDay = orderSettings.ConfirmOvertime
	s.db.Save(&order)
	for _, item := range oderItemList {
		item.OrderId = order.Id
		order.OrderSn = order.OrderSn
	}
	s.db.Create(&oderItemList)
	if param.CouponId != 0 {
		s.UpdateCouponStatus(param.CouponId, member.Id, 1)
	}
	if param.UseIntegration != 0 {
		order.UseIntegration = param.UseIntegration
		s.memberService.UpdateIntegration(member.Id, member.Integration-param.UseIntegration)
	}
	s.DeleteCartItemList(cartPromotionItemList, member.Id)
	s.SendDelayMessageCancelOrder(order.Id)
	result["order"] = order
	result["orderItemList"] = oderItemList
	return
}
func (s *omsPortalOrderService) DeleteCartItemList(cartPromotionItemList []domain.CartPromotionItem, memberId int64) {
	var ids []int64
	for _, item := range cartPromotionItemList {
		ids = append(ids, item.Id)
	}
	s.CartItemService.Delete(memberId, ids)
}
func (s *omsPortalOrderService) GenrateOrderSn(order model.OmsOrder) string {
	var sb strings.Builder

	date := time.Now().Format("20060102")
	key := fmt.Sprintf("%s:%s%s", s.redisDATABASE, s.redisKeyOrderId, date)
	incr, _ := s.redisdb.Incr(context.Background(), key).Result()
	sb.WriteString(date)
	sb.WriteString(fmt.Sprintf("%02d", order.SourceType))
	sb.WriteString(fmt.Sprintf("%02d", order.PayType))
	incrStr := strconv.FormatInt(incr, 10)
	if len(incrStr) <= 6 {
		sb.WriteString(fmt.Sprintf("%06d", incr))
	} else {
		sb.WriteString(incrStr)
	}
	return sb.String()
}
func (s *omsPortalOrderService) CalcGifTGrowth(orderItemList []model.OmsOrderItem) int {
	sum := 0
	for _, item := range orderItemList {
		sum = sum + item.GiftGrowth*item.ProductQuantity
	}
	return sum
}
func (s *omsPortalOrderService) CalcGifIntegration(orderItemList []model.OmsOrderItem) int64 {
	var sum int64
	for _, item := range orderItemList {
		sum += int64(item.GiftIntegration) * int64(item.ProductQuantity)
	}
	return sum
}
func (s *omsPortalOrderService) CalcIntegrationAmount(orderItemList []model.OmsOrderItem) decimal.Decimal {
	var integrationAmount = decimal.NewFromInt(0)
	for _, item := range orderItemList {
		if !item.IntegrationAmount.Equal(decimal.NewFromInt32(0)) {
			integrationAmount = integrationAmount.Add(item.IntegrationAmount.Mul(decimal.NewFromInt32(int32(item.ProductQuantity))))
		}
	}
	return integrationAmount
}

func (s *omsPortalOrderService) CalcPayAmount(order model.OmsOrder) decimal.Decimal {
	return order.TotalAmount.Add(order.FreightAmount).
		Sub(order.PromotionAmount).Sub(order.CouponAmount).
		Sub(order.IntegrationAmount)
}
func (s *omsPortalOrderService) GetOrderPromotionInfo(orderItemList []model.OmsOrderItem) string {
	var builder strings.Builder
	for _, item := range orderItemList {
		builder.WriteString(item.PromotionName)
		builder.WriteString(";")
	}
	return builder.String()
}
func (s *omsPortalOrderService) CalcCouponAmount(orderItemList []model.OmsOrderItem) decimal.Decimal {
	var couponAmount = decimal.NewFromInt(0)
	for _, item := range orderItemList {
		if !item.CouponAmount.Equal(decimal.NewFromInt32(0)) {
			couponAmount = couponAmount.Add(item.CouponAmount.Mul(decimal.NewFromInt32(int32(item.ProductQuantity))))
		}
	}
	return couponAmount
}
func (s *omsPortalOrderService) CalcPromotionAmount(orderItemList []model.OmsOrderItem) decimal.Decimal {
	var promotionAomunt = decimal.NewFromInt(0)
	for _, item := range orderItemList {
		if !item.PromotionAmount.Equal(decimal.NewFromInt32(0)) {
			promotionAomunt = promotionAomunt.Add(item.PromotionAmount).Mul(decimal.NewFromInt32(int32(item.ProductQuantity)))
		}
	}
	return promotionAomunt
}
func (s *omsPortalOrderService) HandleRealAmount(orderItemList []model.OmsOrderItem) {
	for _, item := range orderItemList {
		realAmount := item.ProductPrice.Sub(item.PromotionAmount).Sub(item.CouponAmount).Sub(item.IntegrationAmount)
		item.RealAmount = realAmount
	}
}
func (s *omsPortalOrderService) LockStock(items []domain.CartPromotionItem) {
	for _, item := range items {
		var skuStock model.PmsSkuStock
		s.db.Model(&skuStock).Where(item.ProductSkuId).First(&skuStock)
		skuStock.LockStock = skuStock.LockStock + int64(item.Quantity)
		s.db.Save(&skuStock)
	}
}
func (s *omsPortalOrderService) GetUseIntegrationAmount(useIntegration int64, totalAmount decimal.Decimal,
	member model.UmsMember, hashCoupon bool) decimal.Decimal {
	zeroAmount := decimal.NewFromInt(0)
	if member.Integration < useIntegration {
		return zeroAmount
	}
	var integrationConsumeSetting model.UmsIntegrationConsumeSetting
	s.db.Model(&integrationConsumeSetting).Where(1).First(&integrationConsumeSetting)
	if hashCoupon && *integrationConsumeSetting.CouponStatus == 0 {
		return zeroAmount
	}
	if useIntegration < integrationConsumeSetting.UseUnit {
		return zeroAmount
	}
	integrationAmount := decimal.NewFromInt(useIntegration).DivRound(decimal.NewFromInt(integrationConsumeSetting.UseUnit), 2)
	maxPercent := decimal.NewFromInt(int64(*integrationConsumeSetting.MaxPercentPerOrder)).DivRound(decimal.NewFromInt32(100), 2)
	if integrationAmount.GreaterThan(totalAmount.Mul(maxPercent)) {
		return zeroAmount
	}
	return integrationAmount

}

func (s *omsPortalOrderService) HasStock(item []domain.CartPromotionItem) bool {
	for _, promotionItem := range item {
		if promotionItem.RealStock <= 0 {
			return false
		}
	}
	return true
}
func (s *omsPortalOrderService) GetUseCoupon(cartPromotionItemList []domain.CartPromotionItem, memberId int64, couponId int64) (detail domain.SmsCouponHistoryDetail) {
	couponHistoryDetails := s.memberCouponService.ListCart(cartPromotionItemList, memberId, 1)
	for _, couponHistoryDetail := range couponHistoryDetails {
		if couponHistoryDetail.Coupon.Id == couponId {
			return couponHistoryDetail
		}
	}
	return domain.SmsCouponHistoryDetail{}
}

func (s *omsPortalOrderService) HandleCouponAmount(orderItemList []model.OmsOrderItem,
	couponHistoryDetail domain.SmsCouponHistoryDetail) {
	coupon := couponHistoryDetail.Coupon
	if *coupon.UseType == 0 {
		s.CalcPerCouponAmount(orderItemList, coupon)
	} else if *coupon.UseType == 1 {
		orderItems := s.GetCouponOrderItemByRelation(couponHistoryDetail, orderItemList, 0)
		s.CalcPerCouponAmount(orderItems, coupon)
	} else if *coupon.UseType == 2 {
		orderItems := s.GetCouponOrderItemByRelation(couponHistoryDetail, orderItemList, 1)
		s.CalcPerCouponAmount(orderItems, coupon)
	}
	return
}

func (s *omsPortalOrderService) CalCartPromotion(cartItemList []domain.CartPromotionItem) (amount domain.CalcAmount) {
	amount.FreightAmount = decimal.NewFromInt(0)
	totalAmount, promotionAmount := decimal.NewFromInt(0), decimal.NewFromInt32(0)
	for _, item := range cartItemList {
		totalAmount.Add(item.Price.Mul(decimal.NewFromInt(int64(item.Quantity))))
		promotionAmount = promotionAmount.Add(item.ReduceAmount.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}
	amount.TotalAmount = totalAmount
	amount.PromotionAmount = promotionAmount
	amount.PayAmount = totalAmount.Sub(promotionAmount)
	return
}

func (s *omsPortalOrderService) CalcPerCouponAmount(orderItemList []model.OmsOrderItem, coupon model.SmsCoupon) decimal.Decimal {
	totalAmount := s.CalcTotalAmount(orderItemList)
	for _, item := range orderItemList {
		couponAmount := item.ProductPrice.DivRound(totalAmount, 3).Mul(coupon.Amount)
		item.CouponAmount = couponAmount
	}
	return totalAmount
}
func (s *omsPortalOrderService) GetCouponOrderItemByRelation(couponHistoryDetail domain.SmsCouponHistoryDetail,
	orderItemList []model.OmsOrderItem, type_ int) (result []model.OmsOrderItem) {
	if type_ == 0 {
		var categoryIdList []int64
		for _, relation := range couponHistoryDetail.CategoryRelationList {
			categoryIdList = append(categoryIdList, relation.CouponId)
		}
		for _, item := range orderItemList {
			if -1 != arrays.ContainsInt(categoryIdList, item.ProductCategoryId) {
				result = append(result, item)
			} else {
				item.CouponAmount = decimal.NewFromInt(0)
			}
		}
	} else if type_ == 1 {
		var productIdList []int64
		for _, relation := range couponHistoryDetail.ProductRelationList {
			productIdList = append(productIdList, relation.ProductId)
		}
		for _, item := range orderItemList {
			if -1 != arrays.ContainsInt(productIdList, item.ProductId) {
				result = append(result, item)
			} else {
				item.CouponAmount = decimal.NewFromInt(0)
			}
		}
	}
	return
}
func (s *omsPortalOrderService) CalcTotalAmount(orderItemList []model.OmsOrderItem) decimal.Decimal {
	totalAmount := decimal.NewFromInt(0)
	for _, item := range orderItemList {
		totalAmount = totalAmount.Add(item.ProductPrice.Mul(decimal.NewFromInt32(int32(item.ProductQuantity))))
	}
	return totalAmount
}
func (s *omsPortalOrderService) PaySuccess(orderId int64, payType int) int64 {
	var order model.OmsOrder
	s.db.Model(&order).Where(orderId).First(&order)
	//order.Id=orderId
	order.Status = 1
	var t = time.Now()
	order.PaymentTime = &t
	order.PayType = payType
	s.db.Save(&order)
	orderDetail := s.portalOrderDao.GetDetail(orderId)
	count := s.portalOrderDao.UpdateSkuStock(orderDetail.OrderItemList)
	return count
}

func (s *omsPortalOrderService) CancelOrder(orderId int64) {
	var cancelOrder model.OmsOrder

	s.db.Model(&cancelOrder).Where(map[string]interface{}{"order_id": orderId, "status_id": 0, "delete_status": 0}).First(&cancelOrder)
	if cancelOrder.Id == 0 {
		cancelOrder.Status = 4
		s.db.Save(&cancelOrder)
		var orderItemList []model.OmsOrderItem
		s.db.Where(map[string]interface{}{"order_id": orderId}).Find(&orderItemList)
		s.portalOrderDao.ReleaseSkuStockLock(orderItemList)
		s.UpdateCouponStatus(cancelOrder.CouponId, cancelOrder.MemberId, 0)
		if cancelOrder.Integration != 0 {
			member := s.memberService.GetById(cancelOrder.MemberId)
			s.memberService.UpdateIntegration(cancelOrder.MemberId, member.Integration+cancelOrder.Integration)

		}
	}
}

func (s *omsPortalOrderService) SendDelayMessageCancelOrder(orderId int64) {
	var orderSetting model.OmsOrderSetting
	s.db.Model(&orderSetting).Where(1).First(&orderSetting)
	var deleTimes = int64(*orderSetting.NormalOrderOvertime * 60 * 1000)
	s.orderMQ.CancelOrderSender(orderId, deleTimes)
}

func (s *omsPortalOrderService) ConfirmReceiveOrder(orderId int64, memberId int64) error {
	var order model.OmsOrder
	s.db.Model(&order).Where(orderId).First(&order)
	if memberId != order.MemberId {
		return errors.New("can't confirm other's order")
	}
	if order.Status != 2 {
		return errors.New("this order have not send")
	}
	order.Status = 3
	order.ConfirmStatus = 1
	var t = time.Now()
	order.ReceiveTime = &t
	s.db.Save(&order)
	return nil
}

func (s *omsPortalOrderService) List(memberId int64, status int, pageNum int, pageSize int) paginator.Page[domain.OmsOrderDetail] {
	orders := paginator.Page[model.OmsOrder]{Pages: int64(pageNum), PageSize: int64(pageSize)}
	var opt map[string]interface{}
	opt["delete_status"] = 0
	opt["member_id"] = memberId

	if status != -1 {
		opt["status"] = status
	}
	tx := s.db.Model(&model.OmsOrder{}).Where(opt).Order("create_time desc")
	orders.SelectPages(tx)
	if len(orders.Data) == 0 {
		return paginator.Page[domain.OmsOrderDetail]{}
	}
	var orderIds []int64
	for _, order := range orders.Data {
		orderIds = append(orderIds, order.Id)
	}
	var orderItemList []model.OmsOrderItem
	s.db.Model(&model.OmsOrderItem{}).Where(orderIds).Find(&orderItemList)
	var orderItemMap map[int64][]model.OmsOrderItem
	for _, item := range orderItemList {
		if v, ok := orderItemMap[item.OrderId]; ok {
			v = append(v, item)
			orderItemMap[item.OrderId] = v
		}
		var v []model.OmsOrderItem
		orderItemMap[item.OrderId] = append(v, item)
	}
	var orderDetailItemList []domain.OmsOrderDetail
	for _, order := range orders.Data {
		var orderDetail domain.OmsOrderDetail
		copier.Copy(&orderDetail, &order)
		orderDetail.OrderItemList = orderItemMap[order.Id]
		orderDetailItemList = append(orderDetailItemList, orderDetail)
	}

	return paginator.Page[domain.OmsOrderDetail]{
		CurrentPage: orders.CurrentPage,
		PageSize:    orders.PageSize,
		Total:       orders.Total,
		Pages:       orders.Pages,
		Data:        orderDetailItemList,
	}

}

func (s *omsPortalOrderService) Detail(orderId int64) domain.OmsOrderDetail {
	var omsOrder model.OmsOrder
	s.db.Model(&omsOrder).Where(orderId).First(&omsOrder)
	var omsOrderItemList []model.OmsOrderItem
	s.db.Model(&model.OmsOrderItem{}).Where(map[string]interface{}{"order_id": orderId}).Find(&omsOrderItemList)
	var detail domain.OmsOrderDetail
	copier.Copy(&detail, omsOrder)
	detail.OrderItemList = omsOrderItemList
	return detail
}

func (s *omsPortalOrderService) DeleteOrder(orderId int64, memberId int64) error {
	var omsOrder model.OmsOrder
	s.db.Model(&omsOrder).Where(orderId).First(&omsOrder)
	if memberId != omsOrder.MemberId {
		return errors.New("can't delete other's order")
	}
	if omsOrder.Status == 3 || omsOrder.Status == 4 {
		omsOrder.DeleteStatus = 1
		s.db.Save(&omsOrder)
	} else {
		return errors.New("only delete order in complete or close")
	}
	return nil
}
