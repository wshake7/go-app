package model

import (
	"log"
	"xorm.io/xorm"
)

type Permission struct {
	Id           int64
	Name         string `xorm:"varchar(255) comment('权限名')"`
	Expression   string `xorm:"varchar(255) comment('权限表达式')"`
	Resource     string `xorm:"varchar(255) comment('资源')"`
	ResourceType int    `xorm:"comment('资源类型')"`
}

func SyncPermission(engine *xorm.Engine) {
	err := engine.Sync(new(Permission))
	if err != nil {
		log.Fatal("Database can't be synced: ", err)
	}
}
