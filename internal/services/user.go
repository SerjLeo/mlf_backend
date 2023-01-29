package services

import (
	"errors"
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/cache"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"github.com/SerjLeo/mlf_backend/pkg/templates"
	generatePassword "github.com/sethvargo/go-password/password"
	"time"
)

type UserService struct {
	repo            *Repository
	tokenManager    auth.TokenManager
	hashGenerator   password.HashGenerator
	mailManager     MailManager
	templateManager templates.TemplateManager
	cache           *cache.Cache
	env             string
}

func NewUserService(
	repo *Repository,
	tokenManager auth.TokenManager,
	hashGenerator password.HashGenerator,
	mailManager MailManager,
	templateManager templates.TemplateManager,
	cache *cache.Cache,
	env string,
) *UserService {
	return &UserService{
		env:             env,
		repo:            repo,
		tokenManager:    tokenManager,
		hashGenerator:   hashGenerator,
		mailManager:     mailManager,
		cache:           cache,
		templateManager: templateManager,
	}
}

func (s *UserService) Create(input *models.CreateUserInput) (string, error) {
	hashedPassword, err := s.hashGenerator.EncodeString(input.Password)
	if err != nil {
		return "", err
	}
	input.Password = hashedPassword

	id, err := s.repo.UserRepo.CreateUser(input)
	if err != nil {
		return "", err
	}

	//Here we should send confirmation email
	if s.env == "local" {
		return s.tokenManager.GenerateToken(id, time.Hour*60)
	}

	err = s.mailManager.SendEmail(
		input.Email,
		"",
		"",
	)

	if err != nil {
		return "", err
	}

	return s.tokenManager.GenerateToken(id, time.Hour*60)
}

func (s *UserService) CreateUserByEmail(email string) (string, error) {
	pass, err := generatePassword.Generate(10, 2, 2, false, false)
	if err != nil {
		return "", errors.New("error while generating password")
	}

	passHash, err := s.hashGenerator.EncodeString(pass)
	if err != nil {
		return "", err
	}

	input := models.CreateUserInput{
		Email:    email,
		Password: passHash,
	}

	id, err := s.repo.UserRepo.CreateUser(&input)
	if err != nil {
		return "", err
	}

	_, err = s.repo.ProfileRepo.CreateProfile(id, input.Name)
	if err != nil {
		return "", err
	}

	//Here we should send confirmation + pass via email
	fmt.Println(pass)

	if s.env == "local" {
		return s.tokenManager.GenerateToken(id, time.Hour*60)
	}

	err = s.mailManager.SendEmail(
		input.Email,
		"",
		"",
	)

	if err != nil {
		return "", err
	}

	return s.tokenManager.GenerateToken(id, time.Hour*60)
}

func (s *UserService) SignIn(email, password string) (string, error) {
	passHash, err := s.hashGenerator.EncodeString(password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.UserRepo.AuthenticateUser(email, passHash)
	if err != nil {
		return "", err
	}

	return s.tokenManager.GenerateToken(user.Id, time.Hour*60)
}

func (s *UserService) CheckUserToken(token string) (int, error) {
	claims, err := s.tokenManager.ParseToken(token)
	if err != nil {
		return 0, err
	}

	_, err = s.repo.GetUserById(claims.UserId)
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

	return s.mailManager.SendEmail(
		"sergejleontev111@gmail.com",
		"Email from mail manager",
		body,
	)
}
