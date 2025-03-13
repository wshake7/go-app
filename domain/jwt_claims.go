package domain

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
