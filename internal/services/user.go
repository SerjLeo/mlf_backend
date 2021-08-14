package services

import "github.com/SerjLeo/mlf_backend/internal/repository"

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}