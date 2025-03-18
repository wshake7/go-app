package domain

import "go-app/domain/model"

type PermissionCreateRequest struct {
	Name         string `json:"name"`
	Expression   string `json:"expression"`
	Resource     string `json:"resource"`
	ResourceType int    `json:"resourceType"`
}

func (receiver *PermissionCreateRequest) ToModel() *model.Permission {
	return &model.Permission{
		Name:         receiver.Name,
		Expression:   receiver.Expression,
		Resource:     receiver.Resource,
		ResourceType: receiver.ResourceType,
	}
}

type AclCreateRequest struct {
	UserId       int64 `json:"userId"`
	PermissionId int64 `json:"permissionId"`
}

func (receiver *AclCreateRequest) ToModel() *model.Acl {
	return &model.Acl{
		UserId:       receiver.UserId,
		PermissionId: receiver.PermissionId,
	}
}
