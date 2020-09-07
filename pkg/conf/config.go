package conf

import (
	"github.com/caarlos0/env/v6"
)

type Struct struct {
	Bot Bot
}

func ParseEnv() (*Struct, error) {
	cfg := Struct{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
