package app

import (
	"github.com/SerjLeo/mlf_backend/internal/config"
	"github.com/SerjLeo/mlf_backend/internal/handlers"
	"github.com/SerjLeo/mlf_backend/internal/handlers/http_1_1"
	"github.com/SerjLeo/mlf_backend/internal/migrations"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/internal/repository/postgres"
	"github.com/SerjLeo/mlf_backend/internal/services"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/cache"
	"github.com/SerjLeo/mlf_backend/pkg/colors"
	"github.com/SerjLeo/mlf_backend/pkg/email/smtp"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"github.com/SerjLeo/mlf_backend/pkg/templates"
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

	migrationsConfig := migrations.MigrationConfig{
		Host:           cfg.Postgres.Host,
		Port:           cfg.Postgres.Port,
		DBName:         cfg.Postgres.DBName,
		Password:       cfg.Postgres.Password,
		User:           cfg.Postgres.Username,
		MigrationsPath: cfg.MigrationsPath,
	}
	err = migrations.Run(migrationsConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	tokenManager, err := auth.NewTokenManager(cfg.JWTSignKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	smtpSender, err := smtp.NewSMTPSender(cfg.SMTP.Host, cfg.SMTP.Password, cfg.SMTP.Port, cfg.SMTP.From)
	if err != nil {
		log.Fatal(err.Error())
	}

	templateManager := templates.NewStandardTemplatesManager(cfg.TemplatesPath)
	tmpls, err := templateManager.ParseTemplates()
	if err != nil {
		log.Fatal(err.Error())
	}
	appCache := cache.NewCache(tmpls)

	repos := repository.NewPostgresRepository(db)
	service := services.NewService(services.ServiceDependencies{
		Env:             cfg.Env,
		Repo:            repos,
		TokenManager:    tokenManager,
		HashGenerator:   password.NewSHA1Hash(cfg.HashSalt),
		MailManager:     smtpSender,
		Cache:           appCache,
		TemplateManager: templateManager,
		ColorManager:    colors.NewColorWorker(),
	})
	var handler handlers.Handler = http_1_1.NewRequestHandler(service, cfg.Env)

	server := models.NewServer(cfg.HTTP.Port, handler.InitRoutes())

	if err = server.RunServer(); err != nil {
		log.Fatal(err.Error())
	}
}
