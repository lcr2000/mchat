package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Env      string
	Address  string
	HTTPPort string
	TCPPort  string
}

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Init config fail, err=%v", err)
	}

	Cfg = &Config{
		Env:      viper.GetString("env"),
		Address:  viper.GetString("server.address"),
		HTTPPort: viper.GetString("server.http_port"),
		TCPPort:  viper.GetString("server.tcp_port"),
	}
}
