package main

import (
	"fmt"
	"net"

	"deepsec/gRPC/proto/hello"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

const (
	Address = "127.0.0.1:50444"
)

type helloService struct {
	hello.UnimplementedHelloServer
}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "No token info")
	}
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "I love you and you love me" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token auth info invalied: appid:%s, appkey:%s\n", appid, appkey)
	}
	resp := new(hello.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s, token info: appid=%s, appkey=%s", in.Name, appid, appkey)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		grpclog.Fatalf("failed to generate credentials %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	hello.RegisterHelloServer(s, HelloService)
	grpclog.Println("Listen on " + Address + " with TLS + Token")
	fmt.Println("Listen on " + Address + " with TLS + Token")

	s.Serve(listen)
}
