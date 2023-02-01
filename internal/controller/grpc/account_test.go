package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"go-grpc-tests/internal/domain/service/mocks"
	"go-grpc-tests/internal/entity"
	pb "go-grpc-tests/pkg/proto/bank/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func dialer(accountController *AccountController) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, accountController)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestDepositServer_Deposit(t *testing.T) {
	tests := []struct {
		name    string
		wallet  string
		amount  float32
		res     *pb.DepositResponse
		errCode codes.Code
		errMsg  error
	}{
		{
			"invalid request with negative amount",
			"1",
			-1.11,
			nil,
			codes.InvalidArgument,
			entity.ErrAmountCannotBeNegative,
		},
		{
			"valid request with non negative amount",
			"1",
			0.00,
			&pb.DepositResponse{Ok: true},
			codes.OK,
			nil,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			accountService := mocks.NewMockAccount(ctrl)
			accountController := NewAccountContoller(accountService)
			accountService.EXPECT().Deposit(tt.wallet, tt.amount).Return(tt.errMsg).MaxTimes(1)

			conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(accountController)))
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			client := pb.NewDepositServiceClient(conn)

			request := &pb.DepositRequest{Amount: tt.amount, Wallet: tt.wallet}

			response, err := client.Deposit(ctx, request)

			if err != nil {
				if err != tt.errMsg {
					t.Error("response: expected", tt.res.GetOk(), "received", response.GetOk())
				}
			}
		})
	}
}
