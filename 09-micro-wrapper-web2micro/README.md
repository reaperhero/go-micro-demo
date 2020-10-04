#### micro-middlewares-web2micro

- 启动micro-rpc，注册到etcd中，名称为ProdService
- web从etcd中获取ProdService到客户端接口，传入web的handler

```
请求
curl --location --request POST 'localhost:9000/v1/prods' \
--header 'Content-Type: application/json' \
--data-raw '{"size":1}'

响应
{
    "data": [
        {
            "pid": 100,
            "pname": "prodName100"
        }
    ]
}

服务端装饰器打的日志
[Log Wrapper] ctx: map[] service: ProdService method: ProdService.GetProdList
```