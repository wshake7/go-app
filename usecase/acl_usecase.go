package usecase

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"go-app/domain"
	"go-app/domain/model"
	"go-app/internal/constants"
	"go-app/internal/utils"
	"xorm.io/xorm"
)

type AclUseCase struct {
	*xorm.Engine
	*redis.Client
}

func (receiver *AclUseCase) AclPermissionVerify(ctx *gin.Context, expression string) {
	userId, _ := ctx.Get("userId")

	var aclPermission domain.AclPermission

	one, err := receiver.Engine.Table("acl").
		//Select("acl.id,acl.permission_id,create_by,create_at,permission.expression").
		Join("INNER", "permission", "acl.permission_id = permission.id").
		Where("acl.user_id = ?", userId).
		And("permission.expression = ?", expression).
		Get(&aclPermission)
	utils.Panic(err)
	if !one {
		panic("您没有权限")
	}
}

func (receiver AclUseCase) CreatePermission(permission *model.Permission) error {
	insert, err := receiver.Engine.InsertOne(permission)
	if err != nil {
		return err
	}
	if insert == 0 {
		return errors.New("权限创建失败")
	}
	return err
}

func (receiver *AclUseCase) CreateAcl(acl *model.Acl) error {
	insert, err := receiver.Engine.InsertOne(acl)
	if err != nil {
		return err
	}
	if insert == 0 {
		return errors.New("权限访问关系创建失败")
	}
	return err
}

func (receiver *AclUseCase) PermissionList() ([]model.Permission, error) {
	var permissions []model.Permission
	result, err := receiver.Client.Get(context.Background(), constants.PERMISSION_LIST).Result()
	if err != nil || result == "" {
		err = receiver.Engine.Find(&permissions)
		return permissions, err
	}
	err = json.Unmarshal([]byte(result), &permissions)
	return permissions, err
}
