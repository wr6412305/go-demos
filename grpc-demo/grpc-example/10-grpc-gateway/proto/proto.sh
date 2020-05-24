#!/bin/bash

# 生成http.pb.go
protoc --go_out=plugins=grpc:./ ./google/api/http.proto

# 生成annotations.pb.go
protoc --proto_path=. --go_out=plugins=grpc:./ ./google/api/annotations.proto

# 生成simple.validator.pb.go和simple.pb.go
protoc --proto_path=${GOPATH}/src --proto_path=. --govalidators_out=. --go_out=plugins=grpc:. ./simple.proto

# 生成simple.pb.gw.go
protoc --proto_path=${GOPATH}/src --proto_path=. --grpc-gateway_out=logtostderr=true:. ./simple.proto

# 生成simple.swagger.json
protoc --proto_path=${GOPATH}/src --proto_path=. --swagger_out=logtostderr=true:./ ./simple.proto
