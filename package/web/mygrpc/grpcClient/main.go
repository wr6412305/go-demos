package main

import (
	"context"
	"log"
	"math/rand"
	"mygrpc/inf"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	wg sync.WaitGroup
)

const (
	networkType = "tcp"
	server      = "127.0.0.1"
	port        = "41005"
	parallel    = 50     // 连接并行度
	times       = 100000 // 每连接请求次数
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	// 并行请求
	for i := 0; i < int(parallel); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exe()
		}()
	}

	wg.Wait()
	log.Printf("time taken: %.2f ", time.Now().Sub(currTime).Seconds())
}

func exe() {
	// 建立连接
	conn, _ := grpc.Dial(server + ":" + port)
	defer conn.Close()
	client := inf.NewDataClient(conn)

	for i := 0; i < int(times); i++ {
		getUser(client)
	}
}

func getUser(client inf.DataClient) {
	var request inf.UserRq
	r := rand.Intn(parallel)
	request.Id = int32(r)

	// 调用远程方法
	response, _ := client.GetUser(context.Background(), &request)

	if id, _ := strconv.Atoi(strings.Split(response.Name, ":")[0]); id != r {
		log.Printf("response error %#v", response)
	}
}
