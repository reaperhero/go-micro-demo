
- 商品列表
```
curl --location --request POST 'localhost:9000/v1/prods' \
--header 'Content-Type: application/json' \
--data-raw '{"size":1}'

{
    "data": [
        {
            "pid": 20,
            "pname": "prodName20"
        },
        {
            "pid": 21,
            "pname": "prodName21"
        },
        {
            "pid": 22,
            "pname": "prodName22"
        }
    ]
}
```

- 商品详情

```
curl --location --request GET 'localhost:9000/v1/prods/1'

{
    "data": {
        "pid": 1,
        "pname": "测试商品"
    }
}
```