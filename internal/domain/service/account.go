package service

import (
	"go-grpc-tests/internal/domain"
	"go-grpc-tests/internal/domain/repository"
	"go-grpc-tests/internal/entity"
)

type Account struct {
	accountRepo repository.Account
}

func (s *Account) Deposit(wallet string, amount float32) error {
	if amount < 0 {
		return entity.ErrAmountCannotBeNegative
	}
	err := s.accountRepo.Deposit(wallet, amount)
	return err
}

var _ domain.Account = &Account{}

func NewAccountService(accountRepo repository.Account) *Account {
	return &Account{accountRepo}
}
