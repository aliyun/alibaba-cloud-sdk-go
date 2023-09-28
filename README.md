English | [简体中文](README-CN.md)

<p align="center">
<a href=" https://www.alibabacloud.com"><img src="https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg"></a>
</p>

<h1 align="center">Alibaba Cloud SDK for Go</h1>

[![Go](https://github.com/aliyun/alibaba-cloud-sdk-go/actions/workflows/go.yml/badge.svg)](https://github.com/aliyun/alibaba-cloud-sdk-go/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go/graph/badge.svg?token=kHbylWc7aV)](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=shield&issueType=license)](https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_shield&issueType=license)

Alibaba Cloud SDK for Go allows you to access Alibaba Cloud services such as Elastic Compute Service (ECS), Server Load Balancer (SLB), and CloudMonitor. You can access Alibaba Cloud services without the need to handle API related tasks, such as signing and constructing your requests.

This document introduces how to obtain and call [Alibaba Cloud SDK for Go][SDK].

## Troubleshoot

[Troubleshoot](https://troubleshoot.api.aliyun.com/?source=github_sdk) Provide OpenAPI diagnosis service to help developers locate quickly and provide solutions for developers through `RequestID` or `error message`.

## Online Demo

[Alibaba Cloud OpenAPI Developer Portal][open-api-portal] provides the ability to call the cloud product OpenAPI online, and dynamically generate SDK Example code and quick retrieval interface, which can significantly reduce the difficulty of using the cloud API.

## Requirements

- It's necessary for you to make sure your system meet the [Requirements][Requirements], such as installing a Go environment which is new than 1.13.x.

## Installation

Use `go get` to install SDK：

```sh
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
```

## Quick Examples

Before you begin, you need to sign up for an Alibaba Cloud account and retrieve your [Credentials](https://usercenter.console.aliyun.com/#/manage/ak).

### Create Client

```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk"

func main() {
  client, err := sdk.NewClientWithAccessKey("REGION_ID", "ACCESS_KEY_ID", "ACCESS_KEY_SECRET")
  if err != nil {
    // Handle exceptions
    panic(err)
  }
}
```

### ROA Request

```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

func main() {
  request := requests.NewCommonRequest()        // Make a common request
  request.Method = "GET"                        // Set request method
  request.Product = "CS"                        // Specify product
  request.Domain = "cs.aliyuncs.com"            // Location Service will not be enabled if the host is specified. For example, service with aCertification     type-Bearer Token should be specified
  request.Version = "2015-12-15"                // Specify product version
  request.PathPattern = "/clusters/[ClusterId]" // Specify path rule with ROA-style
  request.Scheme = "https"                      // Set request scheme. Default: http
  request.ApiName = "DescribeCluster"           // Specify product interface
  request.QueryParams["ClusterId"] = "123456"   // Assign values to parameters in the path
  request.QueryParams["RegionId"] = "region_id" // Specify the requested regionId, if not specified, use the client regionId, then default regionId
  request.TransToAcsRequest()                   // Trans commonrequest to acsRequest, which is used by client.
}
```

### RPC Request

```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

func main() {
  request := requests.NewCommonRequest()                // Make a common request
  request.Method = "POST"                               // Set request method
  request.Product = "Ecs"                               // Specify product
  request.Domain = "ecs.aliyuncs.com"                   // Location Service will not be enabled if the host is specified. For example, service with a   Certification type-Bearer Token should be specified
  request.Version = "2014-05-26"                        // Specify product version
  request.Scheme = "https"                              // Set request scheme. Default: http
  request.ApiName = "CreateInstance"                    // Specify product interface
  request.QueryParams["InstanceType"] = "ecs.g5.large"  // Assign values to parameters in the path
  request.QueryParams["RegionId"] = "region_id"         // Specify the requested regionId, if not specified, use the client regionId, then default regionId
  request.TransToAcsRequest()                           // Trans commonrequest to acsRequest, which is used by client.
}
```

## Documentation

- [Requirements](docs/0-Requirements-EN.md)
- [Installation](docs/1-Installation-EN.md)
- [Client & Credentials](docs/2-Client-EN.md)
- [SSL Verify](docs/3-Verify-EN.md)
- [Proxy](docs/4-Proxy-EN.md)
- [Timeout](docs/5-Timeout-EN.md)
- [Debug](docs/6-Debug-EN.md)
- [Logger](docs/7-Logger-EN.md)
- [Concurrent](docs/8-Concurrent-EN.md)
- [Asynchronous Call](docs/9-Asynchronous-EN.md)
- [Package Management](docs/10-Package-Management-EN.md)
- [Endpoint](docs/11-Endpoint-EN.md)

## Issues

[Opening an Issue][issue], Issues not conforming to the guidelines may be closed immediately.

## Contribution

Please make sure to read the [Contributing Guide](CONTRIBUTING.md) before making a pull request.

## References

- [Alibaba Cloud Regions & Endpoints][endpoints]
- [Alibaba Cloud OpenAPI Developer Portal][open-api-portal]
- [Go][go]
- [Latest Release][latest-release]

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_large)

[SDK]: https://github.com/aliyun/alibaba-cloud-sdk-go
[issue]: https://github.com/aliyun/alibaba-cloud-sdk-go/issues/new
[open-api-portal]: https://api.aliyun.com/
[latest-release]: https://github.com/aliyun/alibaba-cloud-sdk-go/releases
[go]: https://golang.org/dl/
[endpoints]: https://developer.aliyun.com/endpoints
[Requirements]: docs/0-Requirements-EN.md
