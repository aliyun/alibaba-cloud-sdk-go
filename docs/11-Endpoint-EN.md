[←  Package Management](10-Package-Management-EN.md) | Endpoint[(中文)](11-Endpoint-CN.md) | [Home →](../README.md)
***

## Endpoint
The endpoint is the domain you want to access

Priority：request > client > global > rules

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
    // global
    endpoints.AddEndpointMapping("regionId", "Ecs", "your endpoint")
    // Only the current request is valid
    request.Domain = "your endpoint"
    // Only the current client is valid
    client.Domain = "your endpoint"
    // Splicing rules set the network; parameter:share、 inner、 vpc、 public
    client.Network = "public";
    resp, err := client.DescribeRegions(request)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)
}
```

Splicing rules need to be supported by the product side and conform to the specification before they can be used。

***
[←  Package Management](10-Package-Management-EN.md) | Endpoint[(中文)](11-Endpoint-CN.md) | [Home →](../README.md)