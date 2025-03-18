package model

import (
	"log"
	"time"
	"xorm.io/xorm"
)

type Acl struct {
	Id           int64
	UserId       int64     `xorm:"notnull comment('用户ID') index(user_id-permission_id)"`
	PermissionId int64     `xorm:"varchar(255) notnull comment('权限ID') index(user_id-permission_id)"`
	CreateBy     int64     `xorm:"comment('权限授予人')"`
	CreateAt     time.Time `xorm:"comment('创建时间')"`
}

func SyncAcl(engine *xorm.Engine) {
	err := engine.Sync(new(Acl))
	if err != nil {
		log.Fatal("Database can't be synced: ", err)
	}
}
