package services

import (
	"errors"
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/cache"
	"github.com/SerjLeo/mlf_backend/pkg/email"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"github.com/SerjLeo/mlf_backend/pkg/templates"
	generatePassword "github.com/sethvargo/go-password/password"
	"time"
)

type UserService struct {
	repo            *repository.Repository
	tokenManager    auth.TokenManager
	hashGenerator   password.HashGenerator
	mailManager     email.MailManager
	templateManager templates.TemplateManager
	cache           *cache.Cache
}

func NewUserService(
	repo *repository.Repository,
	tokenManager auth.TokenManager,
	hashGenerator password.HashGenerator,
	mailManager email.MailManager,
	templateManager templates.TemplateManager,
	cache *cache.Cache,
) *UserService {
	return &UserService{
		repo:            repo,
		tokenManager:    tokenManager,
		hashGenerator:   hashGenerator,
		mailManager:     mailManager,
		cache:           cache,
		templateManager: templateManager,
	}
}

func (s *UserService) Create(user models.User) (string, error) {
	hashedPassword, err := s.hashGenerator.EncodeString(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	id, err := s.repo.User.Create(user)
	if err != nil {
		return "", err
	}

	//Here we should send confirmation email

	return s.tokenManager.GenerateToken(id, time.Hour*60)
}

func (s *UserService) CreateByEmail(email string) (string, error) {
	pass, err := generatePassword.Generate(10, 2, 2, false, false)
	if err != nil {
		return "", errors.New("error while generating password")
	}

	passHash, err := s.hashGenerator.EncodeString(pass)
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:    email,
		Password: passHash,
	}

	id, err := s.repo.User.Create(user)
	if err != nil {
		return "", err
	}

	//Here we should send confirmation + pass via email
	fmt.Println(pass)

	return s.tokenManager.GenerateToken(id, time.Hour*60)
}

func (s *UserService) SignIn(email, password string) (string, error) {
	passHash, err := s.hashGenerator.EncodeString(password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.User.GetUser(email, passHash)
	if err != nil {
		return "", err
	}

	return s.tokenManager.GenerateToken(user.UserId, time.Hour*60)
}

func (s *UserService) CheckUserToken(token string) (int, error) {
	claims, err := s.tokenManager.ParseToken(token)
	if err != nil {
		return 0, err
	}

	return claims.UserId, nil
}

func (s *UserService) SendTestEmail() error {
	body, err := s.templateManager.ExecuteTemplateToString(
		s.cache.Templates["confirmEmail.html"],
		templates.ConfirmEmailData{
			Host:        "MyLocalFinancier.com",
			ConfirmLink: "https://youtube.com",
		})
	if err != nil {
		return err
	}

	return s.mailManager.SendEmail(email.SendInput{
		To:      "sergejleontev111@gmail.com",
		Subject: "Email from mail manager",
		Body:    body,
	})
}
