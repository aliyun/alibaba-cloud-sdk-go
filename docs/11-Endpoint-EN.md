[←  Package Management](10-Package-Management-EN.md) | Endpoint[(中文)](11-Endpoint-CN.md) | [Home →](../README.md)
***

## Endpoint

> Endpoint is the domain name of the service API. like `ecs.cn-hangzhou.aliyuncs.com`

### Search Endpoint

[Related source code](../sdk/endpoints/resolver.go)

1. User custom defined

`User custom defined` is the highest-priority logic to search endpoint and allows you to specify endpoint specifics directly.

```go
// Global effect
endpoints.AddEndpointMapping("<RegionID>", "<Product>", "<Endpoint>")

// Only works for the current request
request := ecs.CreateDescribeRegionsRequest()
request.Domain = "<Endpoint>"

// Only works for the current client
client, _ := ecs.NewClientWithAccessKey("<RegionID>", "<AccessKeyID>", "<AccessKeySecret>")
client.Domain = "<Endpoint>"
```

2. Endpoint Splicing Rules

`Endpoint Splicing Rules` does not take effect until the vpc network is enabled or the product SDK has an Endpoint data file.
Endpoint Data File Example : ([Ecs Endpoint Data File](../services/ecs/endpoint.go))

```go
// Public network request(default)
client.Network = "public"; // Since the default value of `productNetwork` is `public`, there is no need to configure `productNetwork` by default.

// Share-domain request
client.Network = "share";

// Ipv6 network request
client.Network = "ipv6";

// Proxy network request
client.Network = "proxy";

// Internal network request
client.Network = "inner";

// Ipv4/Ipv6 Dual Stack
client.Network = "dualstack";

// Vpc network request
client.Network = "vpc";
```

3. Search endpoint from the [endpoints_config.go](../sdk/endpoints/endpoints_config.go) endpoint data file in the Go SDK Core
Internal operation, no additional configuration required.

4. Request Location Service API to get `Endpoint` from the remote end.

Requires product SDK with `ServiceCode`.

***
[←  Package Management](10-Package-Management-EN.md) | Endpoint[(中文)](11-Endpoint-CN.md) | [Home →](../README.md)