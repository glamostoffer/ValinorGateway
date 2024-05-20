package config

import (
	"flag"
	"github.com/glamostoffer/ValinorGateway/internal/delivery"
	"github.com/glamostoffer/ValinorGateway/internal/server"
	authclient "github.com/glamostoffer/ValinorProtos/auth"
	chatclient "github.com/glamostoffer/ValinorProtos/chat"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	StartTimeout time.Duration     `yaml:"start_timeout"`
	StopTimeout  time.Duration     `yaml:"stop_timeout"`
	Env          string            `yaml:"env"`
	HTTPServer   server.Config     `yaml:"http_server"`
	AuthCfg      authclient.Config `yaml:"auth"`
	ChatCfg      chatclient.Config `yaml:"chat"`
	RouteConfig  delivery.Config   `yaml:"route_config"`
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag.StringVar(&configPath, "config", "", "path to config file")
		flag.Parse()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
