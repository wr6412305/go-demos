package main

import (
	"context"
	"flag"
	"log"
	"time"

	"go-demos/rpcx-demo/service"

	"github.com/smallnest/rpcx/client"
)

var (
	addr1 = flag.String("addr1", "tcp@localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "tcp@localhost:9981", "server2 address")
)

func main() {
	flag.Parse()

	option := client.DefaultOption
	// if failed 5 times, return error immediately, and will try to connect after 30 seconds
	option.GenBreaker = func() client.Breaker {
		return client.NewConsecCircuitBreaker(5, 30*time.Second)
	}

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	option.Retries = 10
	xclient := client.NewXClient("Arith", client.Failfast, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &service.Args{
		A: 10,
		B: 20,
	}

	for i := 0; i < 100; i++ {
		reply := &service.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v", err)
			continue
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Second)
	}
}
