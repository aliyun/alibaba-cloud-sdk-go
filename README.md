# Alibaba Cloud Go Software Development Kit

[![Build Status](https://travis-ci.org/aliyun/alibaba-cloud-sdk-go.svg?branch=master)](https://travis-ci.org/aliyun/alibaba-cloud-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/aliyun/alibaba-cloud-sdk-go)](https://goreportcard.com/report/github.com/aliyun/alibaba-cloud-sdk-go)
[![codecov](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_shield)

See [中文文档](./README_zh.md)

The Alibaba Cloud Go Software Development Kit (SDK) allows you to access Alibaba Cloud services such as Elastic Compute Service (ECS), Server Load Balancer (SLB), and CloudMonitor. You can access Alibaba Cloud services without the need to handle API related tasks, such as signing and constructing your requests.

This document introduces how to obtain and call Alibaba Cloud Go SDK.

If you have any problem while using Go SDK, please [submit an issue](https://github.com/aliyun/alibaba-cloud-sdk-go/issues/new).

## Prerequisites

- To use Alibaba Cloud Go SDK, you must have an Alibaba Cloud account as well as an AccessKey.

	The AccessKey is required when initializing `AcsClient`. You can create an AccessKey in the Alibaba Cloud console. For more information, see [Create an AccessKey](https://usercenter.console.aliyun.com/?spm=5176.doc52740.2.3.QKZk8w#/manage/ak).

	> **Note:** To increase the security of your account, we recommend that you use the AccessKey of the RAM user to access Alibaba Cloud services.

- To use Alibaba Cloud Go SDK to access the APIs of a product, you must first activate the product on the [Alibaba Cloud console](https://home.console.aliyun.com/?spm=5176.doc52740.2.4.QKZk8w) if required.


## Installation

Run the following commands to install Alibaba Cloud Go SDK:

```powershell
glide get github.com/aliyun/alibaba-cloud-sdk-go
```
## Using

The following code example shows the three main steps to use Alibaba Cloud Go SDK:

1. Create the client.
2. Create an API request and set parameters.
3. Initiate the request and handle exceptions.

```go
package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func main() {
	// Create an ECS client
	ecsClient, err := ecs.NewClientWithAccessKey(
		"<your-region-id>",         // Your Region ID
		"<your-access-key-id>",     // Your AccessKey ID
		"<your-access-key-secret>") // Your AccessKey Secret
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	// Create an API request and set parameters
	request := ecs.CreateDescribeInstancesRequest()
	// Set the request.PageSize to "10"
	request.PageSize = requests.NewInteger(10)
	// Initiate the request and handle exceptions
	response, err := ecsClient.DescribeInstances(request)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	fmt.Println(response)
}
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_large)
