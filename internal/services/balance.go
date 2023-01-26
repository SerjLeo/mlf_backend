package services

import "github.com/SerjLeo/mlf_backend/internal/models"

type BalanceService struct {
	repo *Repository
}

func NewBalanceService(repo *Repository) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetUserBalancesAmount(userId int) ([]models.BalanceOfCurrency, error) {
	return s.repo.BalanceRepo.GetUserBalancesAmount(userId)
}
