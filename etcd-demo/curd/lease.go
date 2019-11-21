package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func lease1() {
	var (
		config                clientv3.Config
		client                *clientv3.Client
		err                   error
		putResp               *clientv3.PutResponse
		getResp               *clientv3.GetResponse
		leaseResp, leaseResp1 *clientv3.LeaseGrantResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := context.TODO()

	if leaseResp, err = client.Grant(ctx, 10); err != nil {
		fmt.Println(err)
		return
	}

	// 用client也可以设置key，kv是client的一个结构，因此可以使用其方法
	if putResp, err = client.Put(ctx, "/cron/lock/job1", "ok", clientv3.WithLease(leaseResp.ID)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(putResp.Header.Revision)

	// 由协程来帮自动续租，每秒一次
	if keepAliveChan, err := client.KeepAlive(ctx, leaseResp.ID); err != nil {
		fmt.Println(err)
		return
	} else {
		go func() {
			for {
				select {
				case resp := <-keepAliveChan:
					if resp == nil {
						fmt.Println("续租失败")
						return
					}
					fmt.Println("续租成功")
				}
			}
		}()
	}

	k := 8
	for k != 0 {
		if getResp, err = client.Get(ctx, "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(getResp.Count)
		time.Sleep(2 * time.Second)
		k--
	}

	go func() {
		for {
			client.Put(ctx, "/cron/watch/job1", "I am watch job1")
			time.Sleep(1 * time.Second)
			client.Delete(ctx, "/cron/watch/job1")
			time.Sleep(1 * time.Second)
		}
	}()

	ctxWithTimeout, cancelFunc := context.WithCancel(context.TODO())
	wch := client.Watch(ctxWithTimeout, "/cron/watch/job1", clientv3.WithRev(getResp.Header.Revision))
	tt := time.After(10 * time.Second)
	go func() {
		select {
		case <-tt:
			cancelFunc()
		}
	}()

	for resp := range wch {
		for _, res := range resp.Events {
			fmt.Println(res.Type, string(res.Kv.Key), string(res.Kv.Value))
		}
	}

	ctxWithTimeout1, cancelFunc1 := context.WithCancel(context.TODO())
	// 做分布式锁相关,执行事务
	// 建立租约、用租约抢key，抢到后执行业务逻辑，抢失败返回。函数退出时要defer吧租约关闭
	client.Grant(ctxWithTimeout1, 10)
	if leaseResp1, err = client.Grant(ctx, 10); err != nil {
		fmt.Println(err)
		return
	}

	// defer逻辑可以保证租约被清理，防止长期占用key
	defer client.Revoke(ctx, leaseResp1.ID)
	defer cancelFunc1()
	if keepAliveChan1, err := client.KeepAlive(ctxWithTimeout1, leaseResp1.ID); err != nil {
		fmt.Println(err)
		return
	} else {
		go func() {
			for {
				<-keepAliveChan1
			}
		}()
	}

	// 打开下面可以看锁已经被抢占的情况
	client.Put(ctx, "/cron/txn/job1", "I GET FIRST", clientv3.WithLease(leaseResp1.ID))
}
