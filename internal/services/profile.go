package services

import "github.com/SerjLeo/mlf_backend/internal/models"

type ProfileService struct {
	repo *Repository
}

func NewProfileService(repo *Repository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetUserProfile(userId int) (*models.FullProfile, error) {
	return s.repo.ProfileRepo.GetUserProfile(userId)
}
