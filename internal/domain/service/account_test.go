package service

import (
	"github.com/golang/mock/gomock"
	"go-grpc-tests/internal/domain/repository/mocks"
	"go-grpc-tests/internal/entity"
	pb "go-grpc-tests/pkg/proto/bank/account"
	"testing"
)

func TestDepositServer_Deposit(t *testing.T) {
	tests := []struct {
		name   string
		amount float32
		wallet string
		res    *pb.DepositResponse
		errMsg error
	}{
		{
			"invalid request with negative amount",
			-22,
			"2",
			nil,
			entity.ErrAmountCannotBeNegative,
		},
		{
			"valid request with non negative amount",
			0.00,
			"1",
			&pb.DepositResponse{Ok: true},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountRepo := mocks.NewMockAccount(ctrl)
			accountService := NewAccountService(accountRepo)
			accountRepo.EXPECT().Deposit(tt.wallet, tt.amount).Return(tt.errMsg).MaxTimes(1)

			err := accountService.Deposit(tt.wallet, tt.amount)

			if err != tt.errMsg {
				t.Error("error message: expected", tt.errMsg, "received", err)
			}
		})
	}
}
