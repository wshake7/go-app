package domain

import "go-app/domain/model"

type AclPermission struct {
	model.Acl  `xorm:"extends"`
	Expression string
}
