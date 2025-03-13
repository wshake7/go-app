package bootstrap

import (
	"log"
	"xorm.io/xorm"
)

func NewDataBase(env *Env) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		log.Fatal("Database can't be loaded: ", err)
	}
	return engine
}
