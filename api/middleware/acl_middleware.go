package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-app/internal/utils"
	"go-app/usecase"
)

func AclMiddleware(aclUseCase *usecase.AclUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		list, err := aclUseCase.PermissionList()
		utils.Panic(err)
		fmt.Println("permission list", list)
		for _, permission := range list {
			if path == permission.Resource {
				fmt.Println("permission", permission.Expression)
				return
			}
		}
		ctx.Next()
	}
}
