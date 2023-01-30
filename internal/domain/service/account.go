package service

import (
	"fmt"
	"go-grpc-tests/internal/domain"
	"go-grpc-tests/internal/domain/repository"
)

type Account struct {
	accountRepo repository.Account
}

func (s *Account) Deposit(wallet string, amount float32) error {
	if amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}
	err := s.accountRepo.Deposit(wallet, amount)
	return err
}

var _ domain.Account = &Account{}

func NewAccountService(accountRepo repository.Account) *Account {
	return &Account{accountRepo}
}
