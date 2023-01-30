package repository

import (
	"fmt"
	"go-grpc-tests/internal/domain/repository"
)

type AccountRepository struct {
	balances map[string]float32
}

func (r *AccountRepository) Deposit(wallet string, amount float32) error {
	_, ok := r.balances[wallet]
	if !ok {
		return fmt.Errorf("wallet not found")
	}

	r.balances[wallet] += amount
	return nil
}

var _ repository.Account = &AccountRepository{}

func NewAccountRepository(balances map[string]float32) *AccountRepository {
	return &AccountRepository{balances}
}
