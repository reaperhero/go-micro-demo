## micro工具

- 

micro server

- 通过micro工具箱请求后端服务
```
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379

micro logs api.mudu.com.app 

micro call api.mudu.com.app UsecaseService.Call '{"id": 123}' 通过micro工具箱请求后端服务
```

- 运行micro api

```
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
export MICRO_API_HANDLER=rpc
export MICRO_API_NAMESPACE=api.mudu.com.app
micro api

curl -X POST "http://localhost:8080/UsecaseService/Call" -d '{"id": 1}'
{"data":"test1"}
```