package controller

import (
	"github.com/gin-gonic/gin"
	"go-app/bootstrap"
	"go-app/domain"
	"go-app/internal/utils"
	"go-app/usecase"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthController struct {
	*bootstrap.Env
	usecase.AuthUseCase
}

const duration = time.Hour

func (receiver *AuthController) Login(c *gin.Context) {
	var req domain.LoginRequest
	err := c.ShouldBindJSON(&req)
	utils.Panic(err)
	user, err := receiver.AuthUseCase.GetByAccount(req.Account)
	utils.Panic(err)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse("password error"))
		return
	}
	accessToken, err := receiver.AuthUseCase.CreateAccessToken(user, receiver.Env.AccessTokenSecret, time.Duration(receiver.Env.AccessTokenExpiryHour)*duration)
	utils.Panic(err)

	refreshToken, err := receiver.AuthUseCase.CreateRefreshToken(user, receiver.Env.RefreshTokenSecret, time.Duration(receiver.Env.RefreshTokenExpiryHour)*duration)
	utils.Panic(err)

	loginResponse := domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, &domain.Response{Data: loginResponse})
}

func (receiver *AuthController) Signup(c *gin.Context) {
	var req domain.SignupRequest
	err := c.ShouldBindJSON(&req)
	utils.Panic(err)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	utils.Panic(err)

	req.Password = string(encryptedPassword)
	user := req.ToModel()
	user.NickName = "用户" + req.Account
	err = receiver.AuthUseCase.Create(user)
	utils.Panic(err)

	c.JSON(http.StatusOK, domain.Response{Msg: "signup success"})
}

func (receiver *AuthController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	utils.Panic(err)

	id, err := receiver.AuthUseCase.ExtractIDFromToken(request.RefreshToken, receiver.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Code: -10, Msg: "refreshToken error"})
		return
	}

	user, err := receiver.AuthUseCase.GetUserById(id)
	utils.Panic(err)

	accessToken, err := receiver.AuthUseCase.CreateAccessToken(user, receiver.Env.AccessTokenSecret, time.Duration(receiver.Env.AccessTokenExpiryHour)*duration)
	utils.Panic(err)

	refreshToken, err := receiver.AuthUseCase.CreateRefreshToken(user, receiver.Env.RefreshTokenSecret, time.Duration(receiver.Env.RefreshTokenExpiryHour)*duration)
	utils.Panic(err)

	refreshTokenResponse := domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, &domain.Response{Data: refreshTokenResponse})
}
