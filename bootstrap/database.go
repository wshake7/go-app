package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"go-app/domain/model"
	"log"
	"xorm.io/xorm"
)

func NewDataBase(env *Env) *xorm.Engine {
	db := env.DB
	datasource := db.User + ":" + db.Pwd + "@(" + db.Host + ":" + db.Port + ")/" + db.Name + "?charset=utf8"
	engine, err := xorm.NewEngine(db.Driver, datasource)
	if err != nil {
		log.Fatal("Database can't be loaded: ", err)
	}
	sync(engine)
	return engine
}

func sync(engine *xorm.Engine) {
	model.SyncUser(engine)
	model.SyncAcl(engine)
	model.SyncPermission(engine)
}
