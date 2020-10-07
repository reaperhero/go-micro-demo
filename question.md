# 问题汇总

- 安装说明

go-micro 依赖的部分包尚未适配go1.15版本

go get github.com/micro/micro/v2
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v2/cmd/protoc-gen-micro
go get -u github.com/favadi/protoc-go-inject-tag

- undefined: resolver.BuildOption 

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
