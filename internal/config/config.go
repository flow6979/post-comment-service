package config

import (
	"post-comment-service/internal/ports"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
	JWTSecret     string
	PostService   ports.PostService
	UserService   ports.UserService
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		ServerAddress: viper.GetString("server.address"),
		DatabaseURL:   viper.GetString("database.url"),
		JWTSecret:     viper.GetString("jwt.secret"),
	}

	return cfg, nil
}
