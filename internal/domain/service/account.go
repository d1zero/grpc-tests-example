package service

import (
	"context"
	"go-grpc-tests/internal/domain"
	"go-grpc-tests/internal/domain/repository"
)

type Account struct {
	accountRepo repository.Account
}

func (s *Account) Deposit(ctx context.Context, wallet string, amount float32) error {
	return s.accountRepo.Deposit(ctx, wallet, amount)
}

var _ domain.Account = &Account{}

func NewAccountService(accountRepo repository.Account) *Account {
	return &Account{accountRepo}
}
