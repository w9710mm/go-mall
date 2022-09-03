package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
	"strconv"
)

type OmsPortalOrderController struct {
	omsPortalOrderService service.OmsPortalOrderService
}

func NewOmsPortalOrderController(orderService service.OmsPortalOrderService) api.Controller {
	return &OmsPortalOrderController{omsPortalOrderService: orderService}
}

func (C *OmsPortalOrderController) GenerateConfirmOrder(c *gin.Context) {
	cartIds := c.QueryArray("cartIds")
	ids := make([]int64, len(cartIds))
	for i, id := range cartIds {
		ids[i], _ = strconv.ParseInt(id, 10, 64)
	}
	confirmOrderResult := C.omsPortalOrderService.GenerateConfirmOrder(ids)

	c.JSON(http.StatusOK, response.SuccessMsg(confirmOrderResult))
}

func (C *OmsPortalOrderController) PaySuccess(c *gin.Context) {
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)
	payType, _ := strconv.Atoi(c.Query("payType"))
	count := C.omsPortalOrderService.PaySuccess(orderId, payType)
	c.JSON(http.StatusOK, response.SuccessMsg(count))
}

func (C *OmsPortalOrderController) CancelTimeOutOrder(c *gin.Context) {
	order, err := C.omsPortalOrderService.CancelTimeOutOrder()
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(order))
}

func (C *OmsPortalOrderController) CancelOrder(c *gin.Context) {
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)

	C.omsPortalOrderService.SendDelayMessageCancelOrder(orderId)
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

func (C *OmsPortalOrderController) List(c *gin.Context) {
	status, _ := strconv.Atoi(c.Query("status"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	page := C.omsPortalOrderService.List(status, pageNum, pageSize)
	c.JSON(http.StatusOK, response.SuccessMsg(page))
}
func (C *OmsPortalOrderController) Detail(c *gin.Context) {
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)

	C.omsPortalOrderService.Detail(orderId)
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))

}

func (C *OmsPortalOrderController) CancelUserOrder(c *gin.Context) {
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)
	C.omsPortalOrderService.CancelOrder(orderId)
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))

}

func (C *OmsPortalOrderController) ConfirmReceiveOrder(c *gin.Context) {
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)
	C.omsPortalOrderService.DeleteOrder(orderId)
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))

}
func (o *OmsPortalOrderController) Name() string {
	return "OmsPortalOrderController"
}

func (o *OmsPortalOrderController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/generateConfirmOrder", o.GenerateConfirmOrder)
	group.POST("/generateOrder", o.GenerateConfirmOrder)
	group.POST("/paySuccess", o.PaySuccess)
	group.POST("/cancelTimeOutOrder", o.CancelTimeOutOrder)
	group.POST("/cancelOrder", o.CancelOrder)
	group.GET("/list", o.List)
	group.GET("/detail/:orderId", o.Detail)
	group.POST("/cancelUserOrder", o.CancelUserOrder)
	group.POST("/confirmReceiveOrder", o.ConfirmReceiveOrder)
	group.POST("/deleteOrder", o.Detail)
}
