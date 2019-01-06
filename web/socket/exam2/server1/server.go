package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := "127.0.0.1:9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// set 2 minutes timeout
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// set maxium request length to 128B to prevent flood attack
	request := make([]byte, 128)
	// close connection before exit
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)
		readLen, err := reader.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:readLen-1])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			writer.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			daytime += "\n"
			writer.Write([]byte(daytime))
			writer.Flush()
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
