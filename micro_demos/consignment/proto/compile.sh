#!/bin/bash

protoc --proto_path=. --micro_out=. --go_out=. $1