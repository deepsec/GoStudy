package main

import (
	"fmt"
	"net"

	"deepsec/gRPC/proto/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50443"
)

type helloService struct {
	hello.UnimplementedHelloServer
}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	creds, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	if err != nil {
		grpclog.Fatalf("failed to generate credentials %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	hello.RegisterHelloServer(s, HelloService)
	grpclog.Println("Listen on " + Address + " with TLS")

	s.Serve(listen)
}
