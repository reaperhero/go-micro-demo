#!/bin/bash
cd proto
protoc --micro_out=. --go_out=. service.proto
cd -