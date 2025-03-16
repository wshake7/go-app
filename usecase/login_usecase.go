package usecase

import (
	"github.com/syndtr/goleveldb/leveldb/errors"
	"go-app/domain"
	"go-app/domain/model"
	"go-app/internal/utils"
	"time"
	"xorm.io/xorm"
)

type loginRepository struct {
	engine *xorm.Engine
}

func NewLoginUseCase(engine *xorm.Engine) domain.LoginUseCase {
	return loginRepository{engine: engine}
}

func (lr loginRepository) CreateAccessToken(user *model.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (lr loginRepository) CreateRefreshToken(user *model.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

func (lr loginRepository) Create(user *model.User) error {
	_, err := lr.engine.InsertOne(user)
	return err
}

func (lr loginRepository) GetByAccount(account string) (model.User, error) {
	var user model.User
	one, err := lr.engine.Where("account = ?", account).Get(&user)
	if !one {
		return user, errors.New("用户不存在")
	}
	return user, err
}

func (lr loginRepository) ExtractIDFromToken(refreshToken string, refreshTokenSecret string) (id string, err error) {
	return utils.ExtractIDFromToken(refreshToken, refreshTokenSecret)
}

func (lr loginRepository) GetUserById(id string) (model.User, error) {
	var user model.User
	one, err := lr.engine.ID(id).Get(&user)
	if !one {
		return user, errors.New("用户不存在")
	}
	return user, err
}
