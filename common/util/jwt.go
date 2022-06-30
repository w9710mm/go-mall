package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall/common/response"
	"mall/global/log"
	"net/http"
	"time"
)

var jwtKey = []byte(viper.GetString("server.jwt.key"))
var jwtExpiration = viper.GetUint("server.jwt.expiration")

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(username string) (tokenString string, err error) {
	claim := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(generateExpirationDate()), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),               // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),               // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(jwtKey)
	return tokenString, err
}
func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("server.jwt.key")), nil
	}
}

func ParseToken(tokenss string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &Claims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}

func generateExpirationDate() time.Time {
	return time.Now().Add(time.Second * time.Duration(jwtExpiration))
}

func GetUserNameFromToken(token string) string {
	claims, err := ParseToken(token)
	if err != nil {
		log.Logger.Debug("JWT authentication failed")
		return ""
	}
	return claims.Username
}

func ValidateToken(token string, name string) bool {
	return name == GetUserNameFromToken(token)
}

func isTokenExpired(token string) bool {
	claims, err := ParseToken(token)
	if err != nil {
		return false
	}

	return claims.VerifyExpiresAt(time.Now(), false)
}

func RefreshToken(token string) (string, error) {
	claims, err := ParseToken(token)
	if err != nil {
		log.Logger.Debug("token is invalid", zap.String("token", token))
		return "", errors.New("token is invalid:" + token)
	}
	if claims.VerifyIssuedAt(time.Now().Add(30*time.Minute), false) {
		return token, nil
	}
	return GenerateToken(claims.Username)
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, response.FailedMsg("unauthorized"))
			c.Abort()
			return
		}
		_, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FailedMsg(err.Error()))
			c.Abort()
			return
		}
		c.Next()

	}
}
