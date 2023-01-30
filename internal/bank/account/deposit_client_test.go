package account

import (
	"context"
	"errors"
	"fmt"
	pb "go-grpc-tests/pkg/proto/bank/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"testing"
	"time"
)

type mockDepositServer struct {
	pb.UnimplementedDepositServiceServer
}

func (*mockDepositServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	if req.GetAmount() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot deposit %v", req.GetAmount())
	}

	return &pb.DepositResponse{Ok: true}, nil
}

func TestDepositClient_Deposit(t *testing.T) {
	tests := []struct {
		name   string
		amount float32
		res    bool
		err    error
	}{
		{
			"invalid request with negative amount",
			-1.11,
			false,
			fmt.Errorf("grpc: InvalidArgument, cannot deposit %v", -1.11),
		},
		{
			"valid request with non negative amount",
			0.00,
			true,
			nil,
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := NewDepositClient(conn, time.Second).Deposit(context.Background(), tt.amount)

			if response != tt.res {
				t.Error("error: expected", tt.res, "received", response)
			}

			if err != nil && errors.Is(err, tt.err) {
				t.Error("error: expected", tt.err, "received", err)
			}
		})
	}
}
