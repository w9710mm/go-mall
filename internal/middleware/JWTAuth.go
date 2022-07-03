package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mall/common/response"
	"mall/common/util"
	"net/http"
)

var JWTHeader = viper.GetString("server.jwt.tokenHeader")
var E *casbin.Enforcer

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get(JWTHeader)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, response.FailedMsg("unauthorized"))
			c.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, response.FailedMsg(err))
			c.Abort()
			return
		}
		//access,err:=E.
		//claims.Username

		c.Next()
	}
}
