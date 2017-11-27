# 阿里云Go SDK
欢迎使用阿里云开发者工具套件（SDK）。阿里云Go SDK让您不用复杂编程即可访问云服务器、云监控等多个阿里云服务。这里向您介绍如何获取阿里云Go SDK并开始调用。

## 环境准备
1. 要使用阿里云Go SDK，您需要一个云账号以及一对`Access Key ID`和`Access Key Secret`。 请在阿里云控制台中的[AccessKey管理页面](https://usercenter.console.aliyun.com/?spm=5176.doc52740.2.3.QKZk8w#/manage/ak)上创建和查看您的Access Key，或者联系您的系统管理员
2. 要使用阿里云SDK访问某个产品的API，您需要事先在[阿里云控制台](https://home.console.aliyun.com/?spm=5176.doc52740.2.4.QKZk8w)中开通这个产品。

## SDK获取和安装

使用`go get`下载安装SDK

```
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
```

如果您使用了glide管理依赖，您也可以使用glide来安装阿里云GO SDK

```
glide get github.com/aliyun/alibaba-cloud-sdk-go
```

另外，阿里云Java SDK也会发布在 https://develop.aliyun.com/tools/sdk#/go 这个地址。

## 开始调用
以下这个代码示例向您展示了调用阿里云GO SDK的3个主要步骤：

1. 创建Client实例
2. 创建API请求并设置参数
3. 发起请求并处理异常

```go
import(
	github.com/aliyun/aliyun-openapi-go-sdk/services/ecs
	github.com/aliyun/aliyun-openapi-go-sdk/services/rds
)

// 创建ecsClient实例
ecsClient, err := ecs.NewClientWithAccessKey(
	"<your-region-id>", 			// 您的可用区ID
	"<your-access-key-id>", 		// 您的Access Key ID
	"<your-access-key-secret>")		// 您的Access Key Secret
if err != nil {
	// 异常处理
}

// 创建API请求并设置参数
request := ecs.CreateDescribeInstancesRequest()
request.PageSize = "10"

// 发起请求并处理异常
response, err := ecsClient.DescribeInstances(request)
if err != nil {
	// 异常处理
}
```

在创建Client实例时，您需要填写3个参数：`Region ID`、`Access Key ID`和`Access Key Secret`。`Access Key ID`和`Access Key Secret`可以从控制台获得；而`Region ID`可以从[地域列表](https://help.aliyun.com/document_detail/40654.html?spm=5176.doc52740.2.8.FogWrd)中获得


## Keepalive
阿里云Go SDK底层使用Go语言原生的`net/http`收发请求，因此配置方式与`net/http`相同，您可以通过config直接将配置传递给底层的httpClient
```go
httpTransport := http.Transport{
	// set http client options
}
config := sdk.NewConfig()
            .WithHttpTransport(httpTransport)
            .WithTimeout(timeout)
ecsClient, err := ecs.NewClientWithOptions(config)

```

## 并发请求

* 因Go语言的并发特性，我们建议您在应用层面控制SDK的并发请求。
* 为了方便您的使用，我们也提供了可直接使用的并发调用方式，相关的并发控制由SDK内部实现。

##### 开启SDK Client的并发功能
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

2. 使用callback控制回调
    
    ```go
    blocker := client.FooWithCallback(request, func(response *FooResponse, err error) {
    		// handle the response and err
    	})
 	
    // blocker 为(chan int)，用于控制同步，返回1为成功，0为失败
    // 在<-blocker返回失败时，err依然会被传入的callback处理
    result := <-blocker
    ```
    
## 泛化调用接口(CommonApi)
阿里云Go SDK提供了一个特殊的"泛化调用接口"，该接口具有以下特点：

* 轻量、简便：只需核心包(github.com/aliyun/alibaba-cloud-sdk-go/sdk)即可发起调用
* 灵活：无需更新，即可调用全阿里云的任意API，即使是最新发布的

具体的使用方法，请参考官方文档
