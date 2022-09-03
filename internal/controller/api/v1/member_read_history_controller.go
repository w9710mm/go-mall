package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/common/util"
	"mall/global/config"
	"mall/global/domain"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
	"strconv"
)

type MemberReadHistoryController struct {
	memberReadHistoryService service.MemberReadHistoryService
	umsMemberService         service.UmsMemberService
	umsMemberCacheService    service.UmsMemberCacheService
	tokenHeader              string
}

func NewMemberReadHistoryController(memberReadHistoryService service.MemberReadHistoryService,
	umsMemberService service.UmsMemberService, umsMemberCacheService service.UmsMemberCacheService) api.Controller {
	return &MemberReadHistoryController{
		memberReadHistoryService: memberReadHistoryService,
		umsMemberService:         umsMemberService,
		umsMemberCacheService:    umsMemberCacheService,
		tokenHeader:              config.GetConfig().Server.Jwt.TokenHeader,
	}
}

// Create godoc
// @Summary 创建浏览记录
// @Description 创建一个浏览记录
// @Tags 用户浏览记录接口
// @ID v1/MemberReadHistoryController/Create
// @Accept  json
// @Produce  json
// @Security JWT
// @Param MemberReadHistory body domain.MemberReadHistory true "MemberReadHistory"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /member/readHistory/create [post]
func (C *MemberReadHistoryController) Create(c *gin.Context) {
	var history domain.MemberReadHistory
	err := c.ShouldBind(&history)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, response.FailedMsg(err.Error()))
		return
	}
	value, exists := c.Get("clamis")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.FailedMsg("token is failed"))
		return
	}
	claims := value.(*util.Claims)
	member, err := C.umsMemberCacheService.GetMember(claims.Username)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, response.UnauthorizedMsg(err.Error()))
		return
	}
	err = C.memberReadHistoryService.Create(member, history)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("OK"))
}

// Delete godoc
// @Summary  删除浏览记录
// @Description 根据ids清空出浏览记录
// @Tags 用户浏览记录接口
// @ID v1/MemberReadHistoryController/Delete
// @Accept  json
// @Produce  json
// @Security JWT
// @Param ids query []string true "history_ids"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /member/readHistory/delete [post]
func (C *MemberReadHistoryController) Delete(c *gin.Context) {
	values := c.QueryArray("ids")
	count, err := C.memberReadHistoryService.Delete(values)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(count))

}

// Clear godoc
// @Summary  清空浏览记录
// @Description 清空浏览记录
// @Tags 用户浏览记录接口
// @ID v1/MemberReadHistoryController/Clear
// @Accept  json
// @Produce  json
// @Security JWT
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /member/readHistory/clear [post]
func (C *MemberReadHistoryController) Clear(c *gin.Context) {

	value, exists := c.Get("clamis")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.FailedMsg("token is failed"))
		return
	}
	claims := value.(*util.Claims)
	member, err := C.umsMemberCacheService.GetMember(claims.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.UnauthorizedMsg(err.Error()))
		return
	}
	resut, err := C.memberReadHistoryService.Clear(member.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(resut))
}

// List godoc
// @Summary  获取浏览记录列表
// @Description 分页获取浏览记录列表
// @Tags 用户浏览记录接口
// @ID v1/MemberReadHistoryController/List
// @Accept  json
// @Produce  json
// @Security JWT
// @Param pageNum query int false "page number" default(0)
// @Param pageSize query int false "page size"  default(5)
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /member/readHistory/list [get]
func (C *MemberReadHistoryController) List(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	value, exists := c.Get("clamis")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.FailedMsg("token is failed"))
		return
	}
	claims := value.(*util.Claims)
	member, err := C.umsMemberCacheService.GetMember(claims.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.UnauthorizedMsg(err))
		return
	}

	page, err := C.memberReadHistoryService.List(member.Id, pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(page))
}

func (C *MemberReadHistoryController) RegisterRoute(api *gin.RouterGroup) {

	api.POST("/create", C.Create)
	api.POST("/delete", C.Delete)
	api.POST("/clear", C.Clear)
	api.GET("/list", C.List)

}

func (C *MemberReadHistoryController) Name() string {
	return "MemberReadHistoryController"
}
