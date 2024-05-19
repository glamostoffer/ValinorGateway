package delivery

import "time"

type Config struct {
	AccessTokenTTL time.Duration `yaml:"access_token_ttl"`
}
