package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/SerjLeo/mlf_backend/internal/repository"
	"github.com/SerjLeo/mlf_backend/pkg/auth"
	"github.com/SerjLeo/mlf_backend/pkg/password"
	"time"
)

type UserService struct {
	repo          *repository.Repository
	tokenManager  auth.TokenManager
	hashGenerator password.HashGenerator
}

func NewUserService(
	repo *repository.Repository,
	tokenManager auth.TokenManager,
	hashGenerator password.HashGenerator,
) *UserService {
	return &UserService{repo: repo, tokenManager: tokenManager, hashGenerator: hashGenerator}
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

	return s.tokenManager.GenerateToken(id, time.Hour*60)
}

func (s *UserService) CreateByEmail(email string) (string, error) {
	return "", nil
}
