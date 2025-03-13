package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	Port    string `toml:“port`
	Timeout string `toml:“timeout`
	DB      *struct {
		Driver string `toml:"driver"`
		Host   string `toml:"host"`
		Port   string `toml:"port"`
		User   string `toml:"username"`
		Pwd    string `toml:"password"`
		Name   string `toml:"name"`
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
