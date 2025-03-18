package middleware

import (
	"github.com/gin-gonic/gin"
	"go-app/domain"
	"log"
	"net/http"
)

func RecoverHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//打印错误
				log.Println(err)
				c.JSON(http.StatusOK, domain.ErrorResponse(err.(string)))
				c.Abort()
			}
		}()
		c.Next()
	}
}
