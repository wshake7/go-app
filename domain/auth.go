package domain

import "go-app/domain/model"

type SignupRequest struct {
	Account  string `json:"account" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

func (receiver *SignupRequest) ToModel() *model.User {
	return &model.User{
		Account:  receiver.Account,
		Password: receiver.Password,
	}
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
