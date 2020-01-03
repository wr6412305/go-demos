package main

import (
	"context"
	"flag"

	"go-demos/rpcx-demo/service"

	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

// Arith ...
type Arith struct{}

// Mul the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args service.Args, reply *service.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	blacklist := &serverplugin.BlacklistPlugin{
		Blacklist: map[string]bool{"127.0.0.1": true},
	}
	s.Plugins.Add(blacklist)
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
