package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	Port                   string
	Timeout                string
	DB                     db
	Redis                  rdb
	AccessTokenSecret      string
	AccessTokenExpiryHour  int64
	RefreshTokenSecret     string
	RefreshTokenExpiryHour int64
}

type db struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pwd    string
	Name   string
}

type rdb struct {
	Host string
	Port string
	User string
	Pwd  string
	DB   int
}

const configPath = "./config/config.toml"

func NewEnv() *Env {
	env := Env{
		Redis: rdb{
			DB: 0,
		},
	}
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &env
}
