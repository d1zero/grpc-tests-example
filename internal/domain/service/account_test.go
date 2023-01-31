package service

import (
	"github.com/golang/mock/gomock"
	"go-grpc-tests/internal/domain/repository/mocks"
	"go-grpc-tests/internal/entity"
	"testing"
)

func TestDepositServer_Deposit(t *testing.T) {
	tests := []struct {
		name   string
		amount float32
		wallet string
		errMsg error
	}{
		{
			"invalid request with negative amount",
			-22,
			"2",
			entity.ErrAmountCannotBeNegative,
		},
		{
			"valid request with non negative amount",
			0.00,
			"1",
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountRepo := mocks.NewMockAccount(ctrl)
			accountService := NewAccountService(accountRepo)
			accountRepo.EXPECT().Deposit(tc.wallet, tc.amount).Return(tc.errMsg).MaxTimes(1)

			err := accountService.Deposit(tc.wallet, tc.amount)

			if err != tc.errMsg {
				t.Error("error message: expected", tc.errMsg, "received", err)
			}
		})
	}
}
