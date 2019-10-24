package main

import (
	"go-demos/socket/chatserver/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen("127.0.0.1:8080")
	s.Start()
}
