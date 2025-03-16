package bootstrap

import (
	"go-app/domain/model"
	"log"
	"xorm.io/xorm"
)

func Sync(engine *xorm.Engine) {
	err := engine.Sync(new(model.User))

	if err != nil {
		log.Fatal("Database can't be synced: ", err)
	}
}
