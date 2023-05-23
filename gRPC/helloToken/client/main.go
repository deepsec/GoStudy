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
	Address = "127.0.0.1:50444"
	OpenTLS = true
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "I love you and you love me",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func main() {
	var err error
	var opts []grpc.DialOption

	if OpenTLS {
		creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "localhost")
		if err != nil {
			grpclog.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := hello.NewHelloClient(conn)
	req := &hello.HelloRequest{Name: "Ilove you and you love me, 我想超越这平凡的生活，注定现在暂时漂泊,  with Token .........."}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	grpclog.Println(res.Message)
	fmt.Println(res.Message)
}
