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

type umsMemberController struct {
}

var UmsMemberController = new(umsMemberController)

// Register godoc
// @Summary 注册
// @Description 注册
// @Tags 用户接口
// @ID v1/UmsMemberController/Register
// @Accept  json
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Param telephone query string true "telephone"
// @Param authCode query string true "authCode"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/register [post]
func (C umsMemberController) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	telephone := c.Query("telephone")
	authCode := c.Query("authCode")
	err := umsMemberService.Register(username, password, telephone, authCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg(err.Error()))
		panic(err)
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("success"))
}

// Login godoc
// @Summary 登录
// @Description 登录
// @Tags 用户接口
// @ID v1/UmsMemberController/Login
// @Accept  json
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/login [post]
func (C umsMemberController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token, err := umsMemberService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.UnauthorizedMsg(err.Error()))
		panic(err)
		return
	}
	tokenMap := make(map[string]string, 2)
	tokenMap["token"] = token
	tokenMap["tokenHead"] = tokenHead
	c.JSON(http.StatusOK, response.SuccessMsg(tokenMap))
}

// Info godoc
// @Summary 获取会员信息
// @Description 获取会员信息
// @Tags 用户接口
// @ID v1/UmsMemberController/Info
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/info [GET]
func (C umsMemberController) Info(c *gin.Context) {
	//TODO 权限管理 casbin
	tokenString := c.GetHeader(tokenHead)
	member, err := umsMemberService.GetCurrentMember(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.FailedMsg(err.Error()))
		panic(err)
		return
	}
	c.JSON(http.StatusOK, member)
}

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
	code := umsMemberService.GenerateAuthCode(telephone)

	log.Logger.Info("Generate telephone auth code",
		zap.String("telephone", telephone),
		zap.String("code", code))
	c.JSON(http.StatusOK, response.SuccessMsg(code))
}

// UpdatePassword godoc
// @Summary 会员修改密码
// @Description 会员修改密码
// @Tags 用户接口
// @ID v1/UmsMemberController/UpdatePassword
// @Accept  json
// @Produce  json
// @Param telephone query string true "telephone"
// @Param authCode query string true "authCode"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/updatePassword [post]
func UpdatePassword(c *gin.Context) {
	telephone := c.Query("telephone")
	authCode := c.Query("authCode")
	verify, err := umsMemberService.VerifyAuthCode(telephone, authCode)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.FailedMsg("error"))
		panic(err)
	}

	if !verify {
		c.JSON(http.StatusOK, response.FailedMsg("failed"))
		return
	}
	log.Logger.Info("check auth code success,telephone",
		zap.String(telephone, authCode))
	c.JSON(http.StatusOK, response.SuccessMsg("success"))
}

// RefreshToken godoc
// @Summary 刷新token
// @Description 刷新token
// @Tags 用户接口
// @ID v1/UmsMemberController/RefreshToken
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/refreshToken [GET]
func (C umsMemberController) RefreshToken(c *gin.Context) {
	tokenString := c.GetHeader(tokenHeader)
	refreshToken, err := umsMemberService.RefreshToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, response.UnauthorizedMsg("token is expire"))
		panic(err)
	}
	tokenMap := make(map[string]string, 2)
	tokenMap["token"] = refreshToken
	tokenMap["tokenHead"] = tokenHead
	c.JSON(http.StatusOK, response.SuccessMsg(tokenMap))

}
