package app

import (
	"github.com/SerjLeo/mlf_backend/internal/config"
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/email/smtp"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"log"
)

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	db, err := postgres.NewPostgresDB(cfg.Postgres)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	tokenManager, err := auth.NewTokenManager(cfg.JWTSignKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	smtpSender, err := smtp.NewSMTPSender(cfg.SMTP.Host, cfg.SMTP.Password, cfg.SMTP.Port, cfg.SMTP.From)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewPostgresRepository(db)
	service := services.NewService(services.ServiceDependencies{
		Repo:          repos,
		TokenManager:  tokenManager,
		HashGenerator: password.NewSHA1Hash(cfg.HashSalt),
		MailManager:   smtpSender,
	})
	handler := handlers.NewHandler(service, tokenManager)

	server := models.NewServer(cfg.HTTP.Port, handler.InitRoutes())

	if err = server.RunServer(); err != nil {
		log.Fatal(err.Error())
	}
}
