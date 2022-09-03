package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"mall/common/response"
	"mall/common/util"
	"mall/global/config"
	"mall/global/log"
	"net/http"
	"strings"
)

type JwtAuth struct {
	JwtHead   string
	JwtHeader string
	casbin    *casbin.Enforcer
}

var jwt = JwtAuth{
	JwtHead:   config.GetConfig().Server.Jwt.TokenHead,
	JwtHeader: config.GetConfig().Server.Jwt.TokenHeader,
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI

		//TODO authorization control
		//method := c.Request.Method
		//for _, u := range CasbinExclude {
		//	if u.Url == uri && method == u.Method {
		//		c.Next()
		//	}
		//}
		authString := c.Request.Header.Get(jwt.JwtHeader)

		tokenString := strings.Fields(authString)
		if len(tokenString) != 2 || tokenString[0] != "Bearer" || tokenString[1] == "" {
			c.JSON(http.StatusUnauthorized, response.FailedMsg("unauthorized"))
			c.Abort()
			return
		}

		token, claims, err := util.ParseToken(tokenString[1])

		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, response.FailedMsg(err))
			c.Abort()
			return
		}
		if !util.IsTokenExpired(claims) {
			c.JSON(http.StatusUnauthorized, response.FailedMsg("token is out date"))
			c.Abort()
			return
		}

		log.Logger.Info("user access url",
			log.Any("user", claims.Username),
			log.String("url", uri))
		c.Set("token", token)
		c.Set("clamis", claims)

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
