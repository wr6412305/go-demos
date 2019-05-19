package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	proto "go-demos/micro_demos/consignment/proto"

	"github.com/micro/go-micro"
)

const (
	defaultFile = "consignment.json"
)

func parseFile(fileName string) (*proto.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *proto.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Create new greeter client
	client := proto.NewShippingService("go.micro.srv.consignment", service.Client())

	// 在命令行中指定新的货物信息 json 文件
	infoFile := defaultFile
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v\n", err)
	}

	rsp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		fmt.Println(err)
	}
	if rsp == nil {
		fmt.Println("response is nil.")
	}

	log.Printf("created: %t\n", rsp.Created)
}
