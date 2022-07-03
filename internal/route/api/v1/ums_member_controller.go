package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mall/common/response"
	"mall/global/log"
	"mall/internal/service"
	"net/http"
)

var umsMemberService = service.UmsMemberService

// GetAuthCode godoc
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 用户接口
// @ID v1/UmsMemberController/GetAuthCode
// @Accept  json
// @Produce  json
// @Param telephone query string true "telephone"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/getAuthCode [get]
func GetAuthCode(c *gin.Context) {

	telephone := c.Query("telephone")
	code, err := umsMemberService.GetAuthCode(telephone)
	if err != nil {
		log.Logger.Error("Save authCode failed,telephone:" + telephone)
		c.JSON(http.StatusInternalServerError, response.FailedMsg("0"))
		return
	}
	log.Logger.Info("Generate telephone auth code",
		zap.String("telephone", telephone),
		zap.Int("code", code))
	c.JSON(http.StatusOK, response.SuccessMsg(code))
}

// GetAuthCode godoc
// @Summary 确认验证码
// @Description 确认验证码
// @Tags 用户接口
// @ID v1/UmsMemberController/UpdatePassword
// @Accept  json
// @Produce  json
// @Param telephone query string true "telephone"
// @Param authcode query string true "authcode"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/verifyAuthCode [post]
func UpdatePassword(c *gin.Context) {
	telephone := c.Query("telephone")
	authcode := c.Query("authcode")
	verify, err := umsMemberService.VerifyAuthCode(telephone, authcode)
	if err != nil {
		log.Logger.Error("check auth code error,telephone",
			zap.String(telephone, authcode))
		c.JSON(http.StatusInternalServerError, response.FailedMsg("error"))
		return
	}

	if !verify {
		c.JSON(http.StatusOK, response.FailedMsg("failed"))
		return
	}
	log.Logger.Info("check auth code success,telephone",
		zap.String(telephone, authcode))
	c.JSON(http.StatusOK, response.SuccessMsg("success"))
}
