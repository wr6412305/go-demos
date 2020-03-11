package main

import (
	"fmt"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {

}

// NSQHandler ...
type NSQHandler struct {
}

// HandleMessage ...
func (*NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {
	url := "117.51.148.112:4150"
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("test", "nsq_to_file", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}
	}()

	waiter.Wait()
}
