package controller

import (
	"github.com/gin-gonic/gin"
	"go-app/bootstrap"
)

type LoginController struct {
	Env *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {

}
