package main

import (
	"flag"
	"log"

	"go-demos/socket/chatserver/client"
	"go-demos/socket/chatserver/tui"
)

func main() {
	address := flag.String("server", "127.0.0.1:8080", "Which server to connect to")
	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	tui.StartUI(client)
}
