package server

import "time"

type Config struct {
	Host         string        `yaml:"host" validate:"required"`
	Port         string        `yaml:"port" validate:"required"`
	AllowOrigins string        `yaml:"allow_origins" validate:"required"`
	AllowHeaders string        `yaml:"allow_headers" validate:"required"`
	StopTimeout  time.Duration `yaml:"stop_timeout" validate:"required"`
}
