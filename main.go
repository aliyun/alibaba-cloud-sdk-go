package main

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func main() {
	input := ecs.CreateDescribeRegionsRequest()
	client, err := ecs.NewClientWithAccessKey("cn-beijing", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		fmt.Printf("Failed to NewClientWithAccessKey! Error: %s\n", err.Error())
	}
	output, err := client.DescribeRegions(input)
	if err != nil {
		msg := fmt.Sprintf("Failed to DescribeRegions.!")
		fmt.Printf("%s\nError: %s\n", msg, err.Error())
		return
	}

	fmt.Printf("Regions:%s", output.Regions.Region)
}

// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
// 	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
// )

// func main() {

// 	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	request := requests.NewCommonRequest()
// 	// request.Domain = "ecs.aliyuncs.com"
// 	request.Version = "2014-05-26"
// 	request.ApiName = "DescribeInstanceStatus"

// 	request.QueryParams["PageNumber"] = "1"
// 	request.QueryParams["PageSize"] = "30"

// 	response, err := client.ProcessCommonRequest(request)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Print(response.GetHttpContentString())
// }
