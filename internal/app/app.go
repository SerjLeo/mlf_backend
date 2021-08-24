package app

import (
	"github.com/SerjLeo/mlf_backend/internal/config"
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"log"
)

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	db, err := postgres.NewPostgresDB(cfg.Postgres)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	repos := repository.NewPostgresRepository(db)
	service := services.NewService(services.ServiceDependencies{
		Repo: repos,
	})
	handler := handlers.NewHandler(service)

	server := models.NewServer(cfg.HTTP.Port, handler.InitRoutes())

	if err = server.RunServer(); err != nil {
		log.Fatal(err.Error())
	}
}