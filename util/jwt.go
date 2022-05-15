package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("www.miniDouyin.com")
var str string

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成用户Token
func GenerateToken(id uint, username string, authority int) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour) //token有效期设置为7天
	claims := &Claims{                               //TODO 后续可以将token写入Redis并设置过期时间，而不用Web服务器进行设置
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "minidouyin",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}
