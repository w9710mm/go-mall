package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mall/common/response"
	"mall/internal/dto"
	"mall/internal/service"
	"net/http"
)

var umsAdminService = service.UmsAdminService
var tokenHead = viper.GetString("server.jwt.tokenHead")
var tokenHeader = viper.GetString("server.jwt.tokenHeader")

func register(c *gin.Context) {

}

func Login(c *gin.Context) {
	var loginParam dto.UmsAdminParam
	c.ShouldBind(&loginParam)
	umsAdminService.Login(loginParam.Username, loginParam.Password)
}

func GetPermissionList() {

}

func RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get(tokenHeader)
	newToken, err := umsAdminService.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusOK, response.FailedMsg("token is out of date"))
		return
	}
	var tokenMap map[string]string
	tokenMap["token"] = newToken
	tokenMap["tokenHead"] = tokenHead
	c.JSON(http.StatusOK, response.SuccessMsg(tokenMap))

}
