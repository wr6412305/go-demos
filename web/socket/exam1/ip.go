package main

import (
	"fmt"
	"net"
)

func ipOp() {
	ip := "127.0.0.1"
	// ParseIP(s string) IP函数会把一个IPv4或者IPv6的地址转化成IP类型
	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}

	ip = "2002:c0e8:82e7:0:0:0:c0e8:82e7"
	addr = net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}
}
