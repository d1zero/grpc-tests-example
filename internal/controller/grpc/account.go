package grpc

import (
	"context"
	"go-grpc-tests/internal/domain"
	pb "go-grpc-tests/pkg/proto/bank/account"
)

type AccountController struct {
	pb.UnimplementedDepositServiceServer
	accountService domain.Account
}

func (c *AccountController) Deposit(ctx context.Context, req *pb.DepositRequest) (res *pb.DepositResponse, err error) {
	err = c.accountService.Deposit(ctx, req.GetWallet(), req.GetAmount())
	if err != nil {
		res.Ok = false
		return
	}
	res.Ok = true
	return
}

func NewAccountContoller(accountService domain.Account) *AccountController {
	return &AccountController{accountService: accountService}
}
