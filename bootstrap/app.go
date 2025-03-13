package bootstrap

import (
	"log"
	"xorm.io/xorm"
)

type Application struct {
	Env    *Env
	Engine *xorm.Engine
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Engine = NewDataBase(app.Env)
	return *app
}

func (app *Application) Close() {
	err := app.Engine.Close()
	if err != nil {
		log.Fatal("Database can't be closed: ", err)
	}
}
