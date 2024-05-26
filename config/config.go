package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Repository Repository `yaml:"repository"`
}

// NewConfig init config
func NewConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
