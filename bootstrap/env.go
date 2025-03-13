package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	Port    string `toml:“port`
	Timeout string `toml:“timeout`
	DB      *struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
}

const configPath = "./config/config.toml"

func NewEnv() *Env {
	env := Env{}
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
