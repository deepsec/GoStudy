package main

import (
	"fmt"

	"deepsec/gRPC/proto/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50443"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "localhost")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := hello.NewHelloClient(conn)
	req := &hello.HelloRequest{Name: "Ilove you and you love me, 我想超越这平凡的生活，注定现在暂时漂泊"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	grpclog.Println(res.Message)
	fmt.Println(res.Message)
}
