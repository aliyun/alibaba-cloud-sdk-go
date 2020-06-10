[← 包管理](10-Package-Management-CN.md) | 域名[(English)](11-Endpoint-EN.md) | [首页 →](../README-CN.md)
***

## Endpoint
Endpoint 是请求访问的域名

优先级：request > client > global > rules

```go
import (
    "fmt"

    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
    "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func main() {
    client, err := ecs.NewClientWithAccessKey("regionId", "AccessKeyId", "AccessKeySecret")
    if err != nil {
        fmt.Println(err)
    }
    request := ecs.CreateDescribeRegionsRequest()
    // 全局生效
    endpoints.AddEndpointMapping("regionId", "Ecs", "your endpoint")
    // 只对当前 request 生效
    request.Domain = "your endpoint"
    // 只对当前 client 生效
    client.Domain = "your endpoint"
    // 拼接规则设置网络,参数：share、 inner、 vpc、 public
    client.Network = "public";
    resp, err := client.DescribeRegions(request)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)
}
```

拼接规则方式需要产品端支持并确认规范后方可使用。

***
[← 包管理](10-Package-Management-CN.md) | 域名[(English)](11-Endpoint-EN.md) | [首页 →](../README-CN.md)