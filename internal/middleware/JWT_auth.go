package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mall/common/response"
	"mall/common/util"
	"mall/global/log"
	"net/http"
)

var JWTHeader = viper.GetString("server.jwt.tokenHeader")
var E *casbin.Enforcer

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		method := c.Request.Method
		for _, u := range CasbinExclude {
			if u.Url == uri && method == u.Method {
				c.Next()
			}
		}
		tokenString := c.Request.Header.Get(JWTHeader)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, response.FailedMsg("unauthorized"))
			c.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenString)
		log.Logger.Info("user access url",
			log.Any("user", claims.Username),
			log.String("url", uri))
		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, response.FailedMsg(err))
			c.Abort()
			return
		}

		//TODO 访问控制
		//access,err:=E.

		c.Next()
	}
}

type UrlInfo struct {
	Url    string
	Method string
}

var CasbinExclude = []UrlInfo{
	{Url: "/swagger/index.html", Method: "GET"},
	{Url: "/favicon.ico", Method: "GET"},
}
