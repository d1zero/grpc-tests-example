package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go-grpc-tests/internal/entity"
	"testing"
)

func TestAccountRepository_Deposit(t *testing.T) {
	tests := []struct {
		name   string
		wallet string
		amount float32
		errMsg error
	}{
		{
			"test ok",
			"1",
			1.2,
			nil,
		},
		{
			"non existing wallet",
			"2",
			1.2,
			entity.ErrWalletNotFound,
		},
	}

	db, err := sqlx.Open("sqlite3", "../../db/account.db")
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		t.Error(err.Error())
		return
	}

	accountRepo := NewAccountRepository(db)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			_, err = db.Exec(`DELETE FROM main.accounts WHERE true`)
			if err != nil {
				t.Error("unable to clear db")
				return
			}

			_, err = db.Exec(`INSERT INTO main.accounts (id, wallet, amount) VALUES (1, 1, 1.2)`)
			if err != nil {
				t.Error("unable to fill db")
				return
			}

			err = accountRepo.Deposit(tc.wallet, tc.amount)
			if err != tc.errMsg {
				t.Error("error message: expected", tc.errMsg, "received", err)
			}

			_, err = db.Exec(`DELETE FROM main.accounts WHERE true`)
			if err != nil {
				t.Error("unable to clear db")
				return
			}
		})
	}
}
