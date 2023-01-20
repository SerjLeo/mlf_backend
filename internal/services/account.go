package services

import "github.com/SerjLeo/mlf_backend/internal/models"

type AccountService struct {
	repo *Repository
}

func NewAccountService(repo *Repository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(name string, currencyId, userId int) (*models.AccountWithBalances, error) {
	accountId, err := s.repo.AccountRepo.CreateAccount(name, currencyId, userId)
	if err != nil {
		return nil, err
	}

	acc, err := s.repo.AccountRepo.GetAccountById(accountId, userId)
	if err != nil {
		return nil, err
	}

	accWithBalances := &models.AccountWithBalances{
		Id:        acc.Id,
		Name:      acc.Name,
		Suspended: acc.Suspended,
		IsDefault: acc.IsDefault,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
		Balances:  []models.Balance{},
	}

	balances, err := s.repo.BalanceRepo.GetAccountBalances(userId, accountId)
	if err != nil {
		return nil, err
	}

	accWithBalances.Balances = *balances

	return accWithBalances, nil
}
