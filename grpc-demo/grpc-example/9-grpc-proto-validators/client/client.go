package main

import (
	"context"
	"log"

	"grpcprotovalidators/client/auth"
	pb "grpcprotovalidators/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Address 连接地址
const Address string = "127.0.0.1:8080"

var grpcClient pb.SimpleClient

func main() {
	// 从输入的证书文件中为客户端构造TLS凭证
	creds, err := credentials.NewClientTLSFromFile("../tls/server.pem", "ljs")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	// 构建Token
	token := auth.Token{
		Value: "bearer grpc.auth.token",
		// Value: "basic grpc.auth.token",
		// Value: "bearer grpc.auth.token1",
	}

	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	route()
}

// route 调用服务端Route方法
func route() {
	// 创建发送结构体
	req := pb.InnerMessage{
		SomeInteger: 99,
		SomeFloat:   -1.1,
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
