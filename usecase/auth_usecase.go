package usecase

import (
	"github.com/syndtr/goleveldb/leveldb/errors"
	"go-app/domain/model"
	"go-app/internal/utils"
	"time"
	"xorm.io/xorm"
)

type AuthUseCase struct {
	*xorm.Engine
}

func (receiver *AuthUseCase) CreateAccessToken(user *model.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (receiver *AuthUseCase) CreateRefreshToken(user *model.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

func (receiver *AuthUseCase) Create(user *model.User) error {
	_, err := receiver.Engine.InsertOne(user)
	return err
}

func (receiver *AuthUseCase) GetByAccount(account string) (*model.User, error) {
	var user model.User
	one, err := receiver.Engine.Where("account = ?", account).Get(&user)
	if !one {
		return &user, errors.New("用户不存在")
	}
	return &user, err
}

func (receiver *AuthUseCase) ExtractIDFromToken(refreshToken string, refreshTokenSecret string) (id string, err error) {
	return utils.ExtractIDFromToken(refreshToken, refreshTokenSecret)
}

func (receiver *AuthUseCase) GetUserById(id string) (*model.User, error) {
	var user model.User
	one, err := receiver.Engine.ID(id).Get(&user)
	if !one {
		return &user, errors.New("用户不存在")
	}
	return &user, err
}
