package config

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/v2"
)

type ApplicationConfig struct {
	Database DatabaseConfig `koanf:"database" validate:"required"`
}

type DatabaseConfig struct {
	Host     string `koanf:"host" validate:"required"`
	Port     int    `koanf:"port" validate:"required"`
	User     string `koanf:"user" validate:"required"`
	Password string `koanf:"password" validate:"required"`
	Name     string `koanf:"name" validate:"required"`
}

var prefix = "RAVEN_"

func Load() *ApplicationConfig {
	k := koanf.New(".")

	err := k.Load(env.Provider(".", env.Opt{
		Prefix: prefix,
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, prefix)), "_", ".")

			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)
	if err != nil {
		log.Fatalf("error loading application config: %v", err)
	}

	config := &ApplicationConfig{}

	err = k.Unmarshal("", config)
	if err != nil {
		log.Fatalf("error unmarshalling application config: %v", err)
	}

	validate := validator.New()

	err = validate.Struct(config)
	if err != nil {
		log.Fatalf("error validating application config: %v", err)
	}

	return config
}
