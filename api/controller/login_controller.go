package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go-app/bootstrap"
	"go-app/internal/exception"
)

type LoginController struct {
	Env         *bootstrap.Env
	RedisClient *redis.Client
}

func (lc *LoginController) Login(c *gin.Context) {
	lc.RedisClient.Set(context.Background(), "key", "value", 0)
	result, err := lc.RedisClient.Get(context.Background(), "key").Result()
	lc.RedisClient.Del(context.Background(), "key")
	if err != nil {
		c.Error(exception.New(-1, "错误"))
		return
	}
	c.JSON(200, gin.H{
		"result": result,
	})
}
