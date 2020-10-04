## micro工具

- 安装

```
go get github.com/micro/micro/v2
```

- 常用命令

```
micro web # web管理页面
micro new helloworld # 创建服务
micro run helloworld #运行服务
micro list services # 列出服务
# call a service
micro helloworld --name=Alice

# curl via the api
curl -d '{"name": "Alice"}' http://localhost:8080/helloworld
```

- 通过micro工具箱请求后端服务
```
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
micro call api.mudu.com.app UsecaseService.Call '{"id": 123}' 
{
        "data": "test123"
}
```

- micro api

```
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
export MICRO_API_HANDLER=rpc
export MICRO_API_NAMESPACE=api.mudu.com
micro api

curl -X POST "http://localhost:8080/UsecaseService/Call" -d '{"id": 1}'
{"data":"test1"}
```