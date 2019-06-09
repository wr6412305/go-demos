package main

import (
	"context"
	hello "go-demos/micro-demos/grpc/greeter/srv/proto"
	"log"
	"time"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

// Say ...
type Say struct{}

// Hello ...
func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
