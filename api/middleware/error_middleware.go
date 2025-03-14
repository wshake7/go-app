package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-app/internal/exception"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, errorItem := range c.Errors {
			err := errorItem.Err
			var businessError *exception.BusinessError
			if errors.As(err, &businessError) {
				c.JSON(http.StatusOK, gin.H{
					"code": businessError.Code,
					"msg":  businessError.Msg,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  err.Error(),
				})
			}
			return
		}
	}
}
