package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var jwtKey = []byte(viper.GetString("server.jwt.key"))
var jwtExpiration = viper.GetUint("server.jwt.expiration")

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func MakeToken(username string) (tokenString string, err error) {
	claim := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(jwtExpiration))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                 // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                 // 生效时间
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
