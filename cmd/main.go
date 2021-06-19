package main

import (
	"context"
	"log"
	"net"
	"sample/grpc/proto/sample/greet"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	grpcSrv := grpc.NewServer()
	greet.RegisterGreetServer(grpcSrv, &server{})
	reflection.Register(grpcSrv)
	log.Printf("GreetServer is running!")
	grpcSrv.Serve(listener)
}

type server struct{}

func (s *server) Hello(ctx context.Context, req *greet.HelloReq) (*greet.HelloRes, error) {
	return &greet.HelloRes{
		Hello: "Hello!! " + req.Name,
	}, nil
}
