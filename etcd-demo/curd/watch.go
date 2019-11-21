package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func watch1() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer cli.Close()
	fmt.Println("conn success")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, "name", "ljs")
	cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
		return
	}

	// watch
	for {
		rch := cli.Watch(context.Background(), "name")
		for resp := range rch {
			for k, v := range resp.Events {
				fmt.Println(k, v.Type, string(v.Kv.Key), string(v.Kv.Value))
			}
		}
	}
}
