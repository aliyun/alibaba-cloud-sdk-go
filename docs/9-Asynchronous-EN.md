[← Concurrent](8-Concurrent-EN.md) | Asynchronous Call[(中文)](9-Asynchronous-CN.md) | [Package Management →](10-Package-Management-EN.md)
***
## Asynchronous Call

### Open asynchronous call
Alibaba Cloud Go SDK supports opening asynchronous calls in two ways：
Note: After opening the asynchronous call, you need to call `Shutdown()` before you can start the asynchronous call again

1. Enable asynchronous call when initializing `client`
   ```go
   import (
       "github.com/aliyun/alibaba-cloud-sdk-go/sdk"
       "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
       "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
   )

   c := sdk.NewConfig()
   c.EnableAsync = true // Asynchronous task switch
   c.GoRoutinePoolSize = 10 // Number of goroutines
   c.MaxTaskQueueSize = 20 // Maximum number of tasks for a single goroutine
   c.Timeout = 10 * time.Second
   credential := credentials.NewAccessKeyCredential("acesskeyid", "accesskeysecret")
   client, err := ecs.NewClientWithOptions("regionid", c, credential)
   ```

2. Enable asynchronous call when calling `EnableAsync`
   ```go
   // the first parameter is the maximum number of goroutines
   // the second parameter is the maximum number of tasks for a single goroutine
   // EnableAsync is only allowed to be called once, unless the Shutdown() method is used to first close the asynchronous call, and then call EnableAsync
   client.EnableAsync(10, 20)
   ```

### Make an asynchronous call
Alibaba Cloud Go SDK supports asynchronous calls in two ways：

1. Using channel as return values
    ```go
    responseChannel, errChannel := client.FooWithChan(request)

    // this will block
    response := <-responseChannel
    err = <-errChannel
    ```

2. Use callback to control the callback

    ```go
    blocker := client.FooWithCallback(request, func(response *FooResponse, err error) {
        // handle the response and err
    })

    // blocker which is type of (chan int)，is used to control synchronization，when returning 1 means success，and returning 0 means failure.
    // When <-blocker returns failure，err also will be handled by afferent callback.
    result := <-blocker
    ```

### Close asynchronous call
Use the function of `client` to close asynchronous call and the goroutines
   ```go
   client.Shutdown()
   ```

***
[← Concurrent](8-Concurrent-EN.md) | Asynchronous Call[(中文)](9-Asynchronous-CN.md) | [Package Management →](10-Package-Management-EN.md)