package v1

import (
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/global/config"
	"mall/global/dao/dto"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
)

type UmsAdminController struct {
	tokenHead       string
	tokenHeader     string
	umsAdminService service.UmsAdminService
}

func NewUmsAdminController(adminService service.UmsAdminService) api.Controller {
	return &UmsAdminController{
		tokenHead:       config.GetConfig().Server.Jwt.TokenHead,
		tokenHeader:     config.GetConfig().Server.Jwt.TokenHeader,
		umsAdminService: adminService,
	}

}

func (C *UmsAdminController) register(c *gin.Context) {

}

func (C *UmsAdminController) Login(c *gin.Context) {
	var loginParam dto.UmsAdminParam
	c.ShouldBind(&loginParam)
	C.umsAdminService.Login(loginParam.Username, loginParam.Password)
}

func GetPermissionList() {

}

func (C *UmsAdminController) RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get(C.tokenHeader)
	newToken, err := C.umsAdminService.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusOK, response.FailedMsg("token is out of date"))
		return
	}
	var tokenMap map[string]string
	tokenMap["token"] = newToken
	tokenMap["tokenHead"] = C.tokenHead
	c.JSON(http.StatusOK, response.SuccessMsg(tokenMap))

}

func (C *UmsAdminController) Name() string {
	//TODO implement me
	return "UmsAdminController"
}

func (C *UmsAdminController) RegisterRoute(api *gin.RouterGroup) {
	//api.GET("/getAuthCode", C.GetAuthCode)
	//api.POST("/verifyAuthCode", C.UpdatePassword)
}
