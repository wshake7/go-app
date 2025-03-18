package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"go-app/internal/utils"
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
	utils.Panic(err)
	err = app.RedisClient.Close()
	utils.Panic(err)
}
