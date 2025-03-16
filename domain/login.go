package domain

import (
	"go-app/domain/model"
	"time"
)

type SignupRequest struct {
	Account  string `json:"account" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type LoginRequest struct {
	Account  string `json:"account" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUseCase interface {
	Create(user *model.User) error
	CreateAccessToken(user *model.User, secret string, expiry time.Duration) (accessToken string, err error)
	CreateRefreshToken(user *model.User, secret string, expiry time.Duration) (refreshToken string, err error)
	GetByAccount(account string) (model.User, error)
	ExtractIDFromToken(refreshToken string, refreshTokenSecret string) (id string, err error)
	GetUserById(id string) (model.User, error)
}
