package services

import "github.com/SerjLeo/mlf_backend/internal/models"

type CurrencyService struct {
	repo *Repository
}

func NewCurrencyService(repo *Repository) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) GetCurrenciesList() ([]models.Currency, error) {
	return s.repo.CurrencyRepo.GetCurrencyList()
}
