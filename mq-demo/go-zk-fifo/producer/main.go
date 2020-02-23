package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nladuo/go-zk-fifo/fifo"
)

var (
	hosts    = []string{"117.51.148.112:2181"}
	basePath = "/fifo"
	fifoData = []byte("the fifo data")
	prefix   = "seq-"
)

func produce(f *fifo.DistributedFIFO) {
	for {
		size, err := f.Size()
		if err != nil {
			panic(err)
		}
		if size < 100 {
			data := "data---->" + strconv.FormatInt(time.Now().UnixNano(), 10)
			fmt.Println("Put : ", data)
			f.Put([]byte(data))
		}
	}
}

func main() {
	err := fifo.EstablishZkConn(hosts)
	if err != nil {
		panic(err)
	}
	myfifo := fifo.NewFifo(basePath, fifoData, prefix)
	for i := 0; i < 3; i++ {
		go produce(myfifo)
	}

	ch := make(chan int)
	<-ch
	fifo.CloseZkConn()
}
