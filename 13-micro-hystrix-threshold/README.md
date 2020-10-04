# hystrix

#### 熔断配置

```
configA := hystrix.CommandConfig{
    Timeout:                1000, // 超时时间 单位毫秒
    RequestVolumeThreshold: 5,    // 请求数量
    ErrorPercentThreshold:  50,   // 错误百分比
    SleepWindow:            5000, // 尝试正常请求时间 单位毫秒 默认为5秒
}
```

