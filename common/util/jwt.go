package util

import (
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

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, Secret())
	return token, claims, err
}

func generateExpirationDate() time.Time {
	return time.Now().Add(time.Second * time.Duration(jwtExpiration))
}

func GetUserNameFromToken(tokenString string) string {
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return ""
	}
	return claims.Username
}

func ValidateToken(token string, name string) bool {
	return name == GetUserNameFromToken(token)
}

func isTokenExpired(tokenString string) bool {
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return false
	}

	return claims.VerifyExpiresAt(time.Now(), false)
}

func RefreshToken(tokenString string) (string, error) {
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return "", err
	}
	if claims.VerifyIssuedAt(time.Now().Add(30*time.Minute), false) {
		return tokenString, nil
	}
	return GenerateToken(claims.Username)
}
