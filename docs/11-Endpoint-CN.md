[← 包管理](10-Package-Management-CN.md) | 域名[(English)](11-Endpoint-EN.md) | [首页 →](../README-CN.md)
***

## Endpoint

> Endpoint 是请求接口服务的网络域名，如 `ecs.cn-hangzhou.aliyuncs.com`

### Endpoint 寻址

[相关源码](../sdk/endpoints/resolver.go)

1. 用户自定义

`用户自定义`是优先级最高的寻址逻辑，可以直接指定 endpoint 的具体内容。

```go
// 全局生效
endpoints.AddEndpointMapping("<RegionID>", "<Product>", "<Endpoint>")

// 只对当前 Request 生效
request := ecs.CreateDescribeRegionsRequest()
request.Domain = "<Endpoint>"

// 只对当前 client 生效
client, _ := ecs.NewClientWithAccessKey("<RegionID>", "<AccessKeyID>", "<AccessKeySecret>")
client.Domain = "<Endpoint>"
```

2. Endpoint 拼接规则

在请求 vpc 网络时或产品 SDK 具有 Endpoint 数据文件时，当前寻址逻辑才会生效。
Endpoint 数据文件示例 ([Ecs Endpoint Data File](../services/ecs/endpoint.go))

```go
// 公网请求
client.Network = "public"; // 因为 `productNetwork` 的默认值为 `public`，所以默认情况下不需要配置 `productNetwork`

// 跨域请求
client.Network = "share";

// Ipv6 请求
client.Network = "ipv6";

// 代理请求
client.Network = "proxy";

// 内部请求
client.Network = "inner";

// Ipv4/Ipv6 双协议栈
client.Network = "dualstack";

// vpc 请求
client.Network = "vpc";
```

3. 根据 Go SDK Core 中的 [endpoints_config.go](../sdk/endpoints/endpoints_config.go) 数据文件进行寻址

内部操作，无需额外配置。

4. 请求 Location 服务接口，从远端获取

需要产品 SDK 具备 `ServiceCode`

***
[← 包管理](10-Package-Management-CN.md) | 域名[(English)](11-Endpoint-EN.md) | [首页 →](../README-CN.md)