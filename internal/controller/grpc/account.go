package grpc

import (
	"context"
	"go-grpc-tests/internal/domain"
	pb "go-grpc-tests/pkg/proto/bank/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountController struct {
	pb.UnimplementedDepositServiceServer
	AccountService domain.Account
}

func (c *AccountController) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	result := &pb.DepositResponse{}

	am := req.GetAmount()
	wall := req.GetWallet()

	err := c.AccountService.Deposit(wall, am)
	if err != nil {
		return result, status.Error(codes.InvalidArgument, err.Error())
	}

	result.Ok = true
	return result, nil
}

func NewAccountContoller(accountService domain.Account) *AccountController {
	return &AccountController{
		AccountService: accountService,
	}
}
