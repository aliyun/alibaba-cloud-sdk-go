[← 并发](8-Concurrent-CN.md) | 异步调用[(English)](9-Asynchronous-EN.md) | [包管理 →](10-Package-Management-CN.md)
***
## 异步调用

### 开启异步调用
Alibaba Cloud SDK for Go 支持两种方式开启异步调用：
注：开启异步调用之后需要先调用 `Shutdown()` 才能再次开启异步调用。

1. 初始化 `client` 的时候开启异步调用
   ```go
   import (
       "github.com/aliyun/alibaba-cloud-sdk-go/sdk"
       "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
       "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
   )

   c := sdk.NewConfig()
   c.EnableAsync = true // 异步任务开关
   c.GoRoutinePoolSize = 10 // 开启的协程数
   c.MaxTaskQueueSize = 20 // 单个协程的最大任务数
   c.Timeout = 10 * time.Second
   credential := credentials.NewAccessKeyCredential("acesskeyid", "accesskeysecret")
   client, err := ecs.NewClientWithOptions("regionid", c, credential)
   ```

2. 调用 `EnableAsync` 的时候开启异步调用  
   ```go
   // 第一个入参为开启的协程数
   // 第二个入参为单个协程的最大任务数
   // EnableAsync 只允许调用一次, 除非使用 Shutdown() 方法先关闭异步调用，再调用 EnableAsync
   client.EnableAsync(10, 20)
   ``` 

### 发起异步调用
Alibaba Cloud SDK for Go 支持两种方式的异步调用：

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

### 关闭异步调用
通过调用 `client` 的实例方法关闭异步调用并且关闭对应的协程
   ```go
   client.Shutdown()
   ```

***
[← 并发](8-Concurrent-CN.md) | 异步调用[(English)](9-Asynchronous-EN.md) | [包管理 →](10-Package-Management-CN.md)