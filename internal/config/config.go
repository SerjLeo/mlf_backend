package config

import (
	"github.com/spf13/viper"
	"os"
)

const (
	defaultHTTPPort = "8000"
	localEnv        = "local"
)

type (
	Config struct {
		Env           string
		JWTSignKey    string `yaml:"jwt_sign_key"`
		HashSalt      string `yaml:"hash_salt"`
		TemplatesPath string `yaml:"templates_path"`
		HTTP          HTTPConfig
		Postgres      PostgresConfig
		SMTP          SMTPConfig
		Bot           BotConfig
	}

	BotConfig struct {
		Token string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string `yaml:"ssl_mode"`
	}

	HTTPConfig struct {
		Port string
	}

	SMTPConfig struct {
		Host     string
		Port     string
		From     string
		Password string
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
