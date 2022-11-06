package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

type MigrationConfig struct {
	Host           string
	Port           string
	DBName         string
	MigrationsPath string
}

func Run(config MigrationConfig) error {
	m, err := migrate.New(
		fmt.Sprintf("file:///%s", config.MigrationsPath),
		fmt.Sprintf("postgres://%s:%s/%s?sslmode=disable", config.Host, config.Port, config.DBName))

	if err != nil {
		return err
	}

	m.Up()

	return nil
}
