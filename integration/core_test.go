package integration

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/stretchr/testify/assert"
)

func Test_DescribeRegionsWithCommonRequestWithRPCrequest(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithSTStoken(t *testing.T) {
	assumeresponse, err := createAssumeRole()
	assert.Nil(t, err)
	credential := assumeresponse.Credentials
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithStsToken("cn-hangzhou", credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithHTTPS(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	request.SetScheme("HTTPS")
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithUnicodeSpecificParams(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	request.SetContent([]byte("sdk&-杭&&&州-test"))
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithError(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "Describe"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	request.SetContent([]byte("sdk&-杭&&&州-test"))
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	_, err = client.ProcessCommonRequest(request)
	realerr := err.(errors.Error)
	assert.Equal(t, "InvalidParameter", realerr.ErrorCode())
	assert.Equal(t, "The specified parameter \"Action or Version\" is not valid.", realerr.Message())
}

//func Test_DescribeClusterDetailWithCommonRequestWithROArequest(t *testing.T) {
//	client, err := sdk.NewClientWithAccessKey("default", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
//	assert.Nil(t, err)
//	request := requests.NewCommonRequest()
//	request.Method = "GET"
//	request.Domain = "cs.aliyuncs.com"
//	request.Version = "2015-12-15"
//	request.PathPattern = "/clusters/[ClusterId]"
//	request.QueryParams["RegionId"] = "default"
//	request.TransToAcsRequest()
//	response, err := client.ProcessCommonRequest(request)
//	assert.Nil(t, err)
//	assert.True(t, response.IsSuccess())
//}

//func Test_DescribeClusterDetailWithCommonRequestWithROArequestWithHTTPS(t *testing.T) {
//	request := requests.NewCommonRequest()
//	request.Version =  "2015-12-15"
//	request.Product = "CS"
//	request.Method = requests.GET
//	request.ApiName = "DescribeClusterDetail"
//	request.PathPattern = "/clusters/[ClusterId]"
//	request.SetDomain("cs.aliyuncs.com")
//	request.Scheme = "HTTPS"
//	request.TransToAcsRequest()
//	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
//	assert.Nil(t, err)
//	response, err := client.ProcessCommonRequest(request)
//	assert.Nil(t, err)
//	assert.True(t, response.IsSuccess())
//}
