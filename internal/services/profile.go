package services

import "github.com/SerjLeo/mlf_backend/internal/models"

type ProfileService struct {
	repo *Repository
}

func NewProfileService(repo *Repository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateUserProfile(userId int, name string) (int, error) {
	return s.repo.ProfileRepo.CreateProfile(userId, name)
}

func (s *ProfileService) UpdateProfile(input *models.UpdateProfileInput, userId int) (*models.FullProfile, error) {
	if err := s.repo.ProfileRepo.UpdateProfile(input, userId); err != nil {
		return nil, err
	}

	return s.GetUserProfile(userId)
}

func (s *ProfileService) GetUserProfile(userId int) (*models.FullProfile, error) {
	return s.repo.ProfileRepo.GetUserProfile(userId)
}
