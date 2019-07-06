package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 一个receiver，N个sender，receiver通过关闭一个额外的signal channel说“请停止发送”
// 这种场景比上一个要复杂一点。我们不能让receiver关闭data channel，因为这么做将会打破
// channel closing principle。但是我们可以让receiver关闭一个额外的
// signal channel来通知sender停止发送值

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// The first select is to try to exit the goroutine
				// as early as possible. In fact, it is not essential
				// for this specified example, so it can be omitted.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first branch in the
				// second select may be still not selected for some
				// loops if the send to dataCh is also unblocked.
				// But this is acceptable for this example, so the
				// first select block above can be omitted.
				value := rand.Intn(MaxRandomNumber)
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			// the receiver of the dataCh channel is
			// also the sender of the stopCh cahnnel.
			// It is safe to close the stop channel here.
			if value == MaxRandomNumber-1 {
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
