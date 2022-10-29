package bot

import (
	"github.com/SerjLeo/mlf_backend/internal/config"
	"github.com/SerjLeo/mlf_backend/internal/handlers/tgBot"
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
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func RunBot(configPath string) {
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

	templateManager := templates.NewStandardTemplatesManager(cfg.TemplatesPath)
	tmpls, err := templateManager.ParseTemplates()
	if err != nil {
		log.Fatal(err.Error())
	}
	appCache := cache.NewCache(tmpls)

	repos := repository.NewPostgresRepository(db)
	service := services.NewService(services.ServiceDependencies{
		Repo:            repos,
		TokenManager:    tokenManager,
		HashGenerator:   password.NewSHA1Hash(cfg.HashSalt),
		MailManager:     smtpSender,
		Cache:           appCache,
		TemplateManager: templateManager,
		ColorManager:    colors.NewColorWorker(),
	})
	handler := tgBot.NewBotRouter(service)

	botApi, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	bot := models.NewTgBot(botApi, handler)

	if err = bot.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
