package main

import (
	"context"
	"fmt"
	proto "go-demos/micro_demos/consignment/proto"
	"log"

	"github.com/micro/go-micro"
)

type service struct{}

// CreateConsignment ...
func (g *service) CreateConsignment(ctx context.Context, req *proto.Consignment, rsp *proto.Response) error {
	fmt.Printf("%+v\n", req)
	rsp = &proto.Response{Created: true, Consignment: req}
	return nil
}

func main() {
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
	)

	// 解析命令行参数
	server.Init()
	proto.RegisterShippingServiceHandler(server.Server(), &service{})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
