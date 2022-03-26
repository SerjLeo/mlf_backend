package config

import (
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
		Env           string
		JWTSignKey    string
		HashSalt      string
		TemplatesPath string
		HTTP          HTTPConfig
		Postgres      PostgresConfig
		SMTP          SMTPConfig
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

	if err := populateEnv(&cfg); err != nil {
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

func populateEnv(cfg *Config) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if dbPass, exists := os.LookupEnv("POSTGRES_PASSWORD"); exists {
		cfg.Postgres.Password = dbPass
	}

	if dbUsername, exists := os.LookupEnv("POSTGRES_USER"); exists {
		cfg.Postgres.Username = dbUsername
	}

	if dbName, exists := os.LookupEnv("POSTGRES_DB"); exists {
		cfg.Postgres.DBName = dbName
	}

	if hashSalt, exists := os.LookupEnv("HASH_SALT"); exists {
		cfg.HashSalt = hashSalt
	}

	if signKey, exists := os.LookupEnv("JWT_SIGN_KEY"); exists {
		cfg.JWTSignKey = signKey
	}

	if smtpHost, exists := os.LookupEnv("SMTP_HOST"); exists {
		cfg.SMTP.Host = smtpHost
	}

	if smtpPort, exists := os.LookupEnv("SMTP_PORT"); exists {
		cfg.SMTP.Port = smtpPort
	}

	if smtpPass, exists := os.LookupEnv("SMTP_PASSWORD"); exists {
		cfg.SMTP.Password = smtpPass
	}

	if smtpFrom, exists := os.LookupEnv("SMTP_FROM"); exists {
		cfg.SMTP.From = smtpFrom
	}

	return nil
}
