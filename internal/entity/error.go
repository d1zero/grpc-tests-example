package entity

import "errors"

var (
	ErrAmountCannotBeNegative = errors.New("amount cannot be negative")
	ErrWalletNotFound         = errors.New("wallet not found")
	ErrInternalError          = errors.New("internal error")
)
