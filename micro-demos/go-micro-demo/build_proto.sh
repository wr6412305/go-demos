#!/bin/bash

protoc --micro_out=./src/share/pb --go_out=./src/share/pb ./proto/*.proto
