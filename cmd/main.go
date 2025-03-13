package main

import (
	"github.com/gin-gonic/gin"
	"go-app/api/route"
	"go-app/bootstrap"
	"log"
)

func main() {
	app := bootstrap.App()
	defer app.Close()
	env := app.Env
	engine := app.Engine
	r := gin.Default()

	route.Setup(env, engine, r)

	err := r.Run(":" + env.Port)
	if err != nil {
		log.Fatal("Server can't be started: ", err)
	}
}
