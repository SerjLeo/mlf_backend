package postgres

import (
	"fmt"
	"github.com/SerjLeo/mlf_backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) UpdateTransaction(userId, transactionId int, input models.Transaction) (models.Transaction, error) {
	query := fmt.Sprintf(`
		UPDATE %s
		SET amount = $1, description = $2, type = $3, updated_at = $4
		WHERE user_id = $5 AND transaction_id = $6
		RETURNING *
	`, transactionTable)

	transaction := models.Transaction{}

	row := r.db.QueryRow(query, input.Amount, input.Description, input.TransactionType, input.UpdatedAt, userId, transactionId)
	err := row.Scan(&transaction.Id, &transaction.UserId, &transaction.Amount, &transaction.Description, &transaction.TransactionType, &transaction.CreatedAt, &transaction.UpdatedAt)
	return transaction, err
}

func (r *TransactionPostgres) DeleteTransaction(userId, transactionId int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id = $1 AND transaction_id = $2
	`, transactionTable)

	_, err := r.db.Exec(query, userId, transactionId)
	return err
}

func (r *TransactionPostgres) GetTransactions(userId int) ([]models.Transaction, error) {
	query := fmt.Sprintf(`
		SELECT transaction_id, amount, description, type, %s.name as currency, %s.created_at
		FROM %s INNER JOIN %s ON %s.currency_id=%s.currency_id
		WHERE user_id=$1
	`, currencyTable, transactionTable, transactionTable, currencyTable, transactionTable, currencyTable)
	transactions := []models.Transaction{}
	err := r.db.Select(&transactions, query, userId)
	return transactions, err
}

func (r *TransactionPostgres) GetTransactionById(userId, transactionId int) (models.Transaction, error) {
	query := fmt.Sprintf(`
		SELECT transaction_id, amount, description, type, %s.name as currency, %s.created_at
		FROM %s INNER JOIN %s ON %s.currency_id=%s.currency_id
		WHERE user_id=$1 AND transaction_id=$2
	`, currencyTable, transactionTable, transactionTable, currencyTable, transactionTable, currencyTable)
	transaction := models.Transaction{}
	err := r.db.Get(&transaction, query, userId, transactionId)
	return transaction, err
}

func (r *TransactionPostgres) CreateTransactionWithCategories(userId, balanceId int, newAmount float64, input models.CreateTransactionInput) (*models.Transaction, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	createTransactionQuery := fmt.Sprintf(`
		INSERT INTO %s (amount, description, type, user_id, currency_id)
		VALUES($1, $2, $3, $4, $5)
		RETURNING transaction_id, user_id, amount, description, type, currency_id, created_at, updated_at
	`, transactionTable)
	transaction := models.Transaction{}
	row := tx.QueryRow(createTransactionQuery, input.Amount, input.Description, input.TransactionType, userId, input.CurrencyId)
	err = row.Scan(&transaction.Id, &transaction.UserId, &transaction.Amount, &transaction.Description, &transaction.TransactionType, &transaction.CurrencyId, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(input.CategoriesIds) != 0 {
		createCategoryLinkQuery := fmt.Sprintf(`
			INSERT INTO %s (user_id, category_id, transaction_id)
			VALUES($1,$2,$3)
		`, transactionCategoryTable)

		for i := 0; i < len(input.CategoriesIds); i++ {
			_, err := tx.Exec(createCategoryLinkQuery, userId, input.CategoriesIds[i], transaction.Id)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	getTransactionCategoriesQuery := fmt.Sprintf(`
		SELECT %s.category_id, color, name FROM %s INNER JOIN %s ON %s.category_id=%s.category_id
		WHERE %s.user_id=$1 AND transaction_id=$2
	`, categoryTable, transactionCategoryTable, categoryTable, transactionCategoryTable, categoryTable, categoryTable)
	categories := []models.Category{}
	err = tx.Select(&categories, getTransactionCategoriesQuery, userId, transaction.Id)
	if err != nil {
		tx.Rollback()
		return &transaction, err
	}
	transaction.Categories = categories

	tx.Commit()
	return &transaction, nil
}

func (r *TransactionPostgres) CreateTransaction(userId int, input models.CreateTransactionInput) (*models.Transaction, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (amount, description, type, user_id, currency_id)
		VALUES($1, $2, $3, $4, $5)
		RETURNING transaction_id, user_id, amount, description, type, currency_id, created_at, updated_at
	`, transactionTable)
	transaction := models.Transaction{}

	row := r.db.QueryRow(query, input.Amount, input.Description, input.TransactionType, userId, input.CurrencyId)
	err := row.Scan(&transaction.Id, &transaction.UserId, &transaction.Amount, &transaction.Description, &transaction.TransactionType, &transaction.CurrencyId, &transaction.CreatedAt, &transaction.UpdatedAt)
	return &transaction, err
}

func (r *TransactionPostgres) AttachCategory(userId, transactionId int, categoryId int) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, category_id, transaction_id)
		VALUES($1,$2,$3)
	`, transactionCategoryTable)
	_, err := r.db.Exec(query, userId, categoryId, transactionId)
	return err
}

func (r *TransactionPostgres) DetachCategory(userId, transactionId int, categoryId int) error {
	fmt.Printf("%d, %d \n", transactionId, categoryId)
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id=$1 AND category_id=$2 AND transaction_id=$3
	`, transactionCategoryTable)
	_, err := r.db.Exec(query, userId, categoryId, transactionId)
	return err
}

func (r *TransactionPostgres) GetTransactionCategories(userId, transactionId int) ([]models.Category, error) {
	query := fmt.Sprintf(`
		SELECT %s.category_id, color, name FROM %s INNER JOIN %s ON %s.category_id=%s.category_id
		WHERE %s.user_id=$1 AND transaction_id=$2
	`, categoryTable, transactionCategoryTable, categoryTable, transactionCategoryTable, categoryTable, categoryTable)
	categories := []models.Category{}
	err := r.db.Select(&categories, query, userId, transactionId)
	return categories, err
}
