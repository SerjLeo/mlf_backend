package services

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/pkg/errors"
	"time"
)

type AccountService struct {
	repo *Repository
}

func NewAccountService(repo *Repository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(input *models.CreateAccountInput, userId int) (*models.AccountWithBalances, error) {
	if input.CurrencyId == 0 {
		profile, err := s.repo.GetUserProfile(userId)
		if err != nil {
			return nil, err
		}
		input.CurrencyId = profile.CurrencyId
	}

	accountId, err := s.repo.AccountRepo.CreateAccount(input.Name, input.CurrencyId, userId)
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

func (s *AccountService) GetAccounts(pagination models.PaginationParams, userId int) ([]models.AccountWithBalances, error) {

	accounts, err := s.repo.AccountRepo.GetUsersAccounts(userId, pagination)
	if err != nil {
		return nil, err
	}

	result := []models.AccountWithBalances{}

	for _, acc := range accounts {
		accWithBalances := models.AccountWithBalances{
			Id:        acc.Id,
			Name:      acc.Name,
			Suspended: acc.Suspended,
			IsDefault: acc.IsDefault,
			CreatedAt: acc.CreatedAt,
			UpdatedAt: acc.UpdatedAt,
			Balances:  []models.Balance{},
		}

		balances, err := s.repo.BalanceRepo.GetAccountBalances(userId, acc.Id)
		if err != nil {
			return nil, err
		}
		accWithBalances.Balances = *balances
		result = append(result, accWithBalances)
	}

	return result, nil
}

func (s *AccountService) GetAccountById(accountId, userId int) (*models.AccountWithBalances, error) {
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

func (s *AccountService) UpdateAccount(accountId, userId int, input *models.UpdateAccountInput) (*models.AccountWithBalances, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "not valid input")
	}
	input.UpdatedAt = time.Now().Format(time.RFC3339)

	err := s.repo.UpdateAccount(accountId, userId, input)
	if err != nil {
		return nil, errors.Wrap(err, "error while updating account")
	}

	return s.GetAccountById(accountId, userId)
}

func (s *AccountService) SoftDeleteAccount(accountId, userId int) error {
	return s.repo.SoftDeleteAccount(accountId, userId)
}
