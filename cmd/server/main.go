package main

import (
	"go-grpc-tests/internal/bank/account"
	pb "go-grpc-tests/pkg/proto/bank/account"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, &account.DepositServer{})

	log.Fatalln(server.Serve(listener))
}
