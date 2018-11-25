package main

import (
	"context"
	"log"
	"mygrpc/inf"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = "41005"
)

// Data struct
type Data struct{}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// 起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	log.Printf("grpc server in: %s\n", port)
}

// GetUser method
func (t *Data) GetUser(ctx context.Context, request *inf.UserRq) (response *inf.UserRp, err error) {
	response = &inf.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}
	return response, err
}
