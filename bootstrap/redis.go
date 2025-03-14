package bootstrap

import "github.com/redis/go-redis/v9"

func NewRedis(env *Env) *redis.Client {
	r := env.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Pwd,
	})
	return client
}
