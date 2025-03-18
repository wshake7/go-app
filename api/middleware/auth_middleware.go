package middleware

import (
	"github.com/gin-gonic/gin"
	"go-app/domain"
	"go-app/internal/utils"
	"net/http"
	"strings"
)

func AuthHandler(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := utils.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse(err.Error()))
					c.Abort()
					return
				}
				c.Set("userId", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse(err.Error()))
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse("Not authorized"))
		c.Abort()
	}
}
