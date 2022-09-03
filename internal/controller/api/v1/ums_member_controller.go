package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall/common/response"
	"mall/common/util"
	"mall/global/log"
	"mall/internal/controller/api"
	"mall/internal/middleware"
	"mall/internal/service"
	"net/http"
	"strings"
)

type UmsMemberController struct {
	umsMemberService      service.UmsMemberService
	umsMemberCacheService service.UmsMemberCacheService
	tokenHead             string
	tokenHeader           string
}

func NewUmsMemberController(memberService service.UmsMemberService, cacheService service.UmsMemberCacheService) api.Controller {
	return &UmsMemberController{
		umsMemberService:      memberService,
		umsMemberCacheService: cacheService,
		tokenHead:             viper.GetString("server.jwt.tokenHead"),
		tokenHeader:           viper.GetString("server.jwt.tokenHeader"),
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
		//panic(err)
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
// @Security JWT
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/info [GET]
func (C *UmsMemberController) Info(c *gin.Context) {
	//TODO 权限管理 casbin
	//tokenString := c.GetHeader(C.tokenHead)
	value, exists := c.Get("clamis")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.FailedMsg("token is failed"))
		return
	}
	claims := value.(*util.Claims)
	member, err := C.umsMemberCacheService.GetMember(claims.Username)
	if err != nil {
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
// @Security JWT
// @Param telephone query string true "telephone"
// @Param authCode query string true "authCode"
// @Param password query string true "password"
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/updatePassword [post]
func (C *UmsMemberController) UpdatePassword(c *gin.Context) {
	telephone := c.Query("telephone")
	authCode := c.Query("authCode")
	password := c.Query("password")
	err := C.umsMemberService.UpdatePassword(telephone, password, authCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedMsg("error"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("success"))
}

// RefreshToken godoc
// @Summary 刷新token
// @Description 刷新token
// @Tags 用户接口
// @ID v1/UmsMemberController/RefreshToken
// @Accept  json
// @Produce  json
// @Security JWT
// @Success 200 {object} response.ResponseMsg "success"
// @Failure 500 {object} response.ResponseMsg "failure"
// @Router /sso/refreshToken [GET]
func (C *UmsMemberController) RefreshToken(c *gin.Context) {
	authString := c.Request.Header.Get(C.tokenHeader)

	tokenString := strings.Fields(authString)
	if len(tokenString) != 2 || tokenString[0] != "Bearer" || tokenString[1] == "" {
		c.JSON(http.StatusUnauthorized, response.UnauthorizedMsg("token is expire"))
		return
	}

	refreshToken, err := C.umsMemberService.RefreshToken(tokenString[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.UnauthorizedMsg("generate token error"))
		return
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
	api.GET("/getAuthCode", C.GetAuthCode)
	api.Use(middleware.JWTAuth())
	api.GET("/info", C.Info)
	api.POST("/updatePassword", C.UpdatePassword)
	api.GET("/refreshToken", C.RefreshToken)

}
