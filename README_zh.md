# 阿里云开发者 Go 工具套件
[![Build Status](https://travis-ci.org/aliyun/alibaba-cloud-sdk-go.svg?branch=master)](https://travis-ci.org/aliyun/alibaba-cloud-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/aliyun/alibaba-cloud-sdk-go)](https://goreportcard.com/report/github.com/aliyun/alibaba-cloud-sdk-go)
[![codecov](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go)

欢迎使用阿里云开发者工具套件（SDK）。阿里云 Go SDK 让您不用复杂编程即可访问云服务器、云监控等多个阿里云服务。这里向您介绍如何获取阿里云 Go SDK 并开始调用。

如果您在使用 SDK 的过程中遇到任何问题，欢迎前往[阿里云 SDK 问答社区](https://yq.aliyun.com/tags/type_ask-tagid_23350)提问，提问前请阅读[提问引导](https://help.aliyun.com/document_detail/93957.html)。亦可在当前 GitHub 提交 Issues。

## 环境准备
1. 要使用阿里云 Go SDK，您需要一个云账号以及一对 `Access Key ID` 和 `Access Key Secret`。 请在阿里云控制台中的 [AccessKey 管理页面](https://usercenter.console.aliyun.com/?spm=5176.doc52740.2.3.QKZk8w#/manage/ak) 上创建和查看您的 AccessKey，或者联系您的系统管理员
2. 要使用阿里云 SDK 访问某个产品的 API ，您需要事先在[阿里云控制台](https://home.console.aliyun.com/?spm=5176.doc52740.2.4.QKZk8w)中开通这个产品。

## SDK 获取和安装

使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
```

如果您使用了 glide 管理依赖，您也可以使用 glide 来安装阿里云GO SDK

```sh
$ glide get github.com/aliyun/alibaba-cloud-sdk-go
```

另外，阿里云 Go SDK 也会发布在 https://develop.aliyun.com/tools/sdk#/go 这个地址。

## 开始调用

以下这个代码示例向您展示了调用阿里云 GO SDK 的 3 个主要步骤：

1. 创建 Client 实例
2. 创建 API 请求并设置参数
3. 发起请求并处理异常

```go
package main

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
    "fmt"
)

func main() {
    // 创建ecsClient实例
    ecsClient, err := ecs.NewClientWithAccessKey(
        "<your-region-id>",             // 您的可用区ID
        "<your-access-key-id>",         // 您的Access Key ID
        "<your-access-key-secret>")     // 您的Access Key Secret
    if err != nil {
        // 异常处理
        panic(err)
    }
    // 创建API请求并设置参数
    request := ecs.CreateDescribeInstancesRequest()
    request.PageSize = "10"
    // 发起请求并处理异常
    response, err := ecsClient.DescribeInstances(request)
    if err != nil {
        // 异常处理
        panic(err)
    }
    fmt.Println(response)
}
```

在创建Client实例时，您需要填写3个参数：`Region ID`、`Access Key ID`和`Access Key Secret`。`Access Key ID`和`Access Key Secret`可以从控制台获得；而`Region ID`可以从[地域列表](https://help.aliyun.com/document_detail/40654.html?spm=5176.doc52740.2.8.FogWrd)中获得


## Keepalive
阿里云 Go SDK 底层使用 Go 语言原生的 `net/http` 收发请求，因此配置方式与 `net/http`相同，您可以通过 config 直接将配置传递给底层的 httpClient

```go
httpTransport := http.Transport{
    // set http client options
}
config := sdk.NewConfig()
            .WithHttpTransport(&httpTransport)
            .WithTimeout(timeout)
ecsClient, err := ecs.NewClientWithOptions(config)
```

## 并发请求

* 因 Go 语言的并发特性，我们建议您在应用层面控制 SDK 的并发请求。
* 为了方便您的使用，我们也提供了可直接使用的并发调用方式，相关的并发控制由 SDK 内部实现。

### 开启 SDK Client 的并发功能

```go
// 最大并发数
poolSize := 2
// 可缓存的最大请求数
maxTaskQueueSize := 5

// 在创建时开启异步功能
config := sdk.NewConfig()
            .WithEnableAsync(true)
            .WithGoRoutinePoolSize(poolSize)            // 可选，默认5
            .WithMaxTaskQueueSize(maxTaskQueueSize)     // 可选，默认1000
ecsClient, err := ecs.NewClientWithOptions(config)

// 也可以在client初始化后再开启
client.EnableAsync(poolSize, maxTaskQueueSize)
```

##### 发起异步调用
阿里云Go SDK支持两种方式的异步调用：

1. 使用channel作为返回值
    ```go
    responseChannel, errChannel := client.FooWithChan(request)

    // this will block
    response := <-responseChannel
    err = <-errChannel
    ```

2. 使用 callback 控制回调

    ```go
    blocker := client.FooWithCallback(request, func(response *FooResponse, err error) {
        // handle the response and err
    })

    // blocker 为(chan int)，用于控制同步，返回1为成功，0为失败
    // 在<-blocker返回失败时，err依然会被传入的callback处理
    result := <-blocker
    ```

## 泛化调用接口(CommonAPI)

### 什么是 CommonAPI

CommonAPI 是阿里云 SDK 推出的，泛用型的 API 调用方式。CommonAPI 具有以下几个特点：
1. 轻量：只需 Core 包即可发起调用，无需下载安装各产品线 SDK。
2. 简便：无需更新 SDK 即可调用最新发布的 API。
3. 快速迭代

### 开始使用

CommonAPI，需要配合相应的 API 文档使用，以查询 API 的相关信息。

您可以在 [文档中心](https://help.aliyun.com/?spm=5176.8142029.388261.173.23896dfaav2hEF) 查询到所有产品的 API 文档。

发起一次 CommonAPI 请求，需要您查询到以下几个参数：
* 域名(domain)：即该产品的通用访问域名，一般可以在"调用方式"页查看到
* API 版本(version)：即该 API 的版本号，以’YYYY-MM-DD’的形式表现，一般可以在"公共参数"页面查到
* 接口名称(apiName)：即该 API 的名称

我们以 ECS 产品的 [DescribeInstanceStatus API](https://help.aliyun.com/document_detail/25505.html?spm=5176.doc25506.6.820.VbHnW6) 为例

```go
package main

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func main() {

	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "DescribeInstanceStatus"

	request.QueryParams["PageNumber"] = "1"
	request.QueryParams["PageSize"] = "30"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}

	fmt.Print(response.GetHttpContentString())
}
```
