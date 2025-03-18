package usecase

import "xorm.io/xorm"

type UserUseCase struct {
	*xorm.Engine
}
