package repository

import (
	"github.com/jmoiron/sqlx"
	"go-grpc-tests/internal/domain/repository"
	"go-grpc-tests/internal/entity"
)

type AccountRepository struct {
	db *sqlx.DB
}

func (r *AccountRepository) Deposit(wallet string, amount float32) (err error) {
	rows, err := r.db.Exec(`UPDATE main.accounts SET amount=amount+? WHERE wallet=?;`, amount, wallet)
	if err != nil {
		return entity.ErrInternalError
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return entity.ErrInternalError
	}

	if rowsAff == 0 {
		return entity.ErrWalletNotFound
	}

	return nil
}

var _ repository.Account = &AccountRepository{}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db}
}
