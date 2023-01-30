package domain

type Account interface {
	Deposit(string, float32) error
}
