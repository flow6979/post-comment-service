package config

import (
	"fmt"
	"os"
	"path/filepath"

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

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current directory: %w", err)
	}

	// path construction
	configPath := filepath.Join(currentDir, "internal", "config", "config.yaml")

	// path validation
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %s", configPath)
	}

	fmt.Printf("Using config file: %s\n", configPath)

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	cfg := &Config{
		ServerAddress: viper.GetString("server.address"),
		DatabaseURL:   viper.GetString("database.url"),
		JWTSecret:     viper.GetString("jwt.secret"),
	}

	// logger.Info.Printf("Loaded configuration:\n"+
	// 	"ServerAddress: %s\n"+
	// 	"DatabaseURL: %s\n"+
	// 	"JWTSecret: %s\n",
	// 	cfg.ServerAddress,
	// 	cfg.DatabaseURL,
	// 	cfg.JWTSecret)

	return cfg, nil
}
