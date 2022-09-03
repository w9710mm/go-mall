package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/global/domain"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
	"strconv"
)

type MemberAttentionController struct {
	memberAttentionService service.MemberAttentionService
}

func NewMemberAttentionController(memberAttentionService service.MemberAttentionService) api.Controller {
	return &MemberAttentionController{memberAttentionService: memberAttentionService}
}

func (C *MemberAttentionController) add(c *gin.Context) {
	var memberBrandAttention domain.MemberBrandAttention

	c.ShouldBind(&memberBrandAttention)
	count := C.memberAttentionService.Add(memberBrandAttention)
	if count < 1 {
		c.JSON(http.StatusBadRequest, response.FailedMsg(count))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(count))
}

func (C *MemberAttentionController) delete(c *gin.Context) {
	brandId, _ := strconv.ParseInt(c.Query("brandId"), 10, 64)
	count := C.memberAttentionService.Delete(brandId)
	if count < 1 {
		c.JSON(http.StatusBadRequest, response.FailedMsg(count))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(count))
}
func (C *MemberAttentionController) list(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	page := C.memberAttentionService.List(pageNum, pageSize)

	c.JSON(http.StatusOK, response.SuccessMsg(page))
}

func (C *MemberAttentionController) detail(c *gin.Context) {
	brandId, _ := strconv.ParseInt(c.Query("brandId"), 10, 64)
	detail := C.memberAttentionService.Detail(brandId)
	c.JSON(http.StatusOK, response.SuccessMsg(detail))
}

func (C *MemberAttentionController) clear(c *gin.Context) {
	C.memberAttentionService.Clear()
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

func (m *MemberAttentionController) Name() string {
	return "MemberAttentionController"
}

func (m *MemberAttentionController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/add", m.add)
	group.POST("/delete", m.delete)
	group.GET("/list", m.list)
	group.GET("/detail", m.detail)
	group.POST("/clear", m.clear)
}
