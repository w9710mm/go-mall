package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall/common/response"
	"mall/global/log"
	"mall/internal/controller/api"
	"mall/internal/service"
	"net/http"
)

type UmsMemberController struct {
	umsMemberService service.UmsMemberService
	tokenHead        string
	tokenHeader      string
}

func NewUmsMemberController(memberService service.UmsMemberService) api.Controller {
	return &UmsMemberController{
		umsMemberService: memberService,
		tokenHead:        viper.GetString("server.jwt.tokenHead"),
		tokenHeader:      viper.GetString("server.jwt.tokenHeader"),
	}
}

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
func (C *UmsMemberController) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	telephone := c.Query("telephone")
	authCode := c.Query("authCode")
	err := C.umsMemberService.Register(username, password, telephone, authCode)
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
func (C *UmsMemberController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token, err := C.umsMemberService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.UnauthorizedMsg(err.Error()))
		panic(err)
		return
	}
	tokenMap := make(map[string]string, 2)
	tokenMap["token"] = token
	tokenMap["tokenHead"] = C.tokenHead
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
func (C *UmsMemberController) Info(c *gin.Context) {
	//TODO 权限管理 casbin
	tokenString := c.GetHeader(C.tokenHead)
	member, err := C.umsMemberService.GetCurrentMember(tokenString)
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
func (C *UmsMemberController) GetAuthCode(c *gin.Context) {

	telephone := c.Query("telephone")
	code := C.umsMemberService.GenerateAuthCode(telephone)

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
func (C *UmsMemberController) UpdatePassword(c *gin.Context) {
	telephone := c.Query("telephone")
	authCode := c.Query("authCode")
	verify, err := C.umsMemberService.VerifyAuthCode(telephone, authCode)
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
func (C *UmsMemberController) RefreshToken(c *gin.Context) {
	tokenString := c.GetHeader(C.tokenHeader)
	refreshToken, err := C.umsMemberService.RefreshToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, response.UnauthorizedMsg("token is expire"))
		panic(err)
	}
	tokenMap := make(map[string]string, 2)
	tokenMap["token"] = refreshToken
	tokenMap["tokenHead"] = C.tokenHead
	c.JSON(http.StatusOK, response.SuccessMsg(tokenMap))

}
func (C *UmsMemberController) Name() string {
	//TODO implement me
	return "UmsMemberController"
}

func (C *UmsMemberController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/register", C.Register)
	api.POST("/login", C.Login)
	api.GET("/info", C.Info)
	api.GET("/getAuthCode", C.GetAuthCode)
	api.POST("/updatePassword", C.UpdatePassword)
	api.GET("/refreshToken", C.RefreshToken)

}
