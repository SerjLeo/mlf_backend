package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

type MigrationConfig struct {
	Host           string
	User           string
	Port           string
	Password       string
	DBName         string
	MigrationsPath string
}

func Run(config MigrationConfig) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	m, err := migrate.New(
		fmt.Sprintf("file:///%s/%s", path, config.MigrationsPath),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DBName),
	)

	if err != nil {
		return err
	}

	m.Up()

	return nil
}
