package domain

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	NickName string `json:"nick_name"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}
