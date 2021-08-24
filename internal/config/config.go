package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

const (
	defaultHTTPPort = "8000"
	localEnv        = "local"
)

type (
	Config struct {
		Env      string
		HTTP     HTTPConfig
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	}

	HTTPConfig struct {
		Port string
	}
)

func InitConfig(configDir string) (*Config, error) {
	setDefaults()

	if err := parseFile(configDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := populateEnv(&cfg); err != nil {
		return nil, err
	}

	fmt.Printf("%+v",cfg)

	return &cfg, nil
}

func parseFile(directory, env string) error {
	viper.AddConfigPath(directory)
	viper.SetConfigName("common")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == localEnv {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func setDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
}

func populateEnv(cfg *Config) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if dbPass, exists := os.LookupEnv("DB_PASSWORD"); exists {
		cfg.Postgres.Password = dbPass
	}

	return nil
}
