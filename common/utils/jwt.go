package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUtils struct{}

var Jwt = new(JwtUtils)

type UserAuthClaims struct {
	UserId uint //用户id
	jwt.StandardClaims
}

// 生成token
func (j *JwtUtils) GenerateToken(claim UserAuthClaims, expireTime time.Time) (string, error) {
	claim.StandardClaims = jwt.StandardClaims{
		//签名时间
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expireTime.Unix(),
		Issuer:    "GoBlog",
		Subject:   "GoBlogAdmin",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//签名
	var secreKey = []byte("RRInjFYM5H6CwHwCTKGVclICroaEbRo9")

	return token.SignedString(secreKey)
}

// 解析token
func (j *JwtUtils) ParseToken(tokenStr string) (*jwt.Token, error) {
	var secreKey = []byte("RRInjFYM5H6CwHwCTKGVclICroaEbRo9")

	// 解析token
	var claim UserAuthClaims
	return jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return secreKey, nil
	})
}
