package account

import (
	"context"
	"fmt"
	pb "go-grpc-tests/pkg/proto/bank/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, &DepositServer{})

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
		amount  float32
		res     *pb.DepositResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"invalid request with negative amount",
			-1.11,
			nil,
			codes.InvalidArgument,
			fmt.Sprintf("cannot deposit %v", -1.11),
		},
		{
			"valid request with non negative amount",
			0.00,
			&pb.DepositResponse{Ok: true},
			codes.OK,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewDepositServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.DepositRequest{Amount: tt.amount}

			response, err := client.Deposit(ctx, request)

			if response != nil {
				if response.GetOk() != tt.res.GetOk() {
					t.Error("response: expected", tt.res.GetOk(), "received", response.GetOk())
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
