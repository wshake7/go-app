package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

type Application struct {
	Env         *Env
	DBEngine    *xorm.Engine
	RedisClient *redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DBEngine = NewDataBase(app.Env)
	app.RedisClient = NewRedis(app.Env)
	return *app
}

func (app *Application) Close() {
	err := app.DBEngine.Close()
	err = app.RedisClient.Close()
	if err != nil {
		log.Fatal("app close error", err)
	}
}
