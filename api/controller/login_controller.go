package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go-app/bootstrap"
	"go-app/domain"
	"go-app/domain/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginController struct {
	Env          *bootstrap.Env
	RedisClient  *redis.Client
	LoginUseCase domain.LoginUseCase
}

func (lc *LoginController) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
	user, err := lc.LoginUseCase.GetByAccount(req.Account)

	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("密码错误"))
		return
	}
	accessToken, err := lc.LoginUseCase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, time.Duration(lc.Env.AccessTokenExpiryHour)*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	refreshToken, err := lc.LoginUseCase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, time.Duration(lc.Env.RefreshTokenExpiryHour)*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	loginResponse := domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, loginResponse)
}

func (lc *LoginController) Signup(c *gin.Context) {
	var req domain.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
		return
	}
	req.Password = string(encryptedPassword)

	err = lc.LoginUseCase.Create(&model.User{
		NickName: "用户" + req.Account,
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, domain.Response{Msg: "注册成功"})
}

func (lc *LoginController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
		return
	}

	id, err := lc.LoginUseCase.ExtractIDFromToken(request.RefreshToken, lc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse("user not found"))
		return
	}

	user, err := lc.LoginUseCase.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse("user not found"))
		return
	}
	accessToken, err := lc.LoginUseCase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, time.Duration(lc.Env.AccessTokenExpiryHour)*time.Hour)
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
		return
	}

	refreshToken, err := lc.LoginUseCase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, time.Duration(lc.Env.RefreshTokenExpiryHour)*time.Hour)
	if err != nil {
		c.JSON(http.StatusOK, domain.NewErrorResponse(err.Error()))
		return
	}

	refreshTokenResponse := domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
