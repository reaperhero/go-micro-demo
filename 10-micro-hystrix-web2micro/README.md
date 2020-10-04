# hystrix

#### hystrix 1秒超时配置

```
curl --location --request POST 'localhost:9000/v1/prods' \
--header 'Content-Type: application/json' \
--data-raw '{"size":1}'

超时响应
{
    "status": "hystrix: timeout"
}
```


#### 服务降级

```
curl --location --request POST 'localhost:9000/v1/prods' 
--header 'Content-Type: application/json' \
--data-raw '{"size":1}'

{
    "data": [
        {
            "pid": 20,
            "pname": "defaultname20"
        },
        {
            "pid": 21,
            "pname": "defaultname21"
        }
    ]
}
```