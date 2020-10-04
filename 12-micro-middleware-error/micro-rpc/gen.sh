#!/bin/bash

cd proto && \
protoc --micro_out=. --go_out=. prodservice.proto && \
protoc-go-inject-tag --input=./prodservice.pb.go && \
cd -
