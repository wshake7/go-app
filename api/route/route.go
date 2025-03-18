package route

import (
	"github.com/gin-gonic/gin"
	"go-app/api/middleware"
	"go-app/bootstrap"
	"go-app/usecase"
)

func Setup(app *bootstrap.Application, engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(middleware.RecoverHandler())

	publicRouter := engine.Group("public")
	NewAuthRoute(app, publicRouter)

	privateRouter := engine.Group("api")
	aclUseCase := &usecase.AclUseCase{Engine: app.DBEngine, Client: app.RedisClient}
	privateRouter.Use(middleware.AclMiddleware(aclUseCase))
	privateRouter.Use(middleware.AuthHandler(app.Env.AccessTokenSecret))

	NewAclRoute(app, aclUseCase, privateRouter)
	NewUserRoute(app, privateRouter)
}
