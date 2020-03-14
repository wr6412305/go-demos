#!/bin/bash

go test -v -run="TestGrpcClient" *.go
# 或者直接
# go test -v -run="TestGrpcClient"

# sleep 1s 因为服务端限制了每秒访问一次
sleep 1s

go test -v -run="TestGrpc"
