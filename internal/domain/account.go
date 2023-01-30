package domain

import "context"

type Account interface {
	Deposit(context.Context, string, float32) error
}
