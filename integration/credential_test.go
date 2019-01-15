package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func Test_DescribeRegionsWithRPCrequestWithAK(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_DescribeRegionsWithRPCrequestWithSTStoken(t *testing.T) {
	assumeresponse, err := createAssumeRole()
	assert.Nil(t, err)
	credential := assumeresponse.Credentials
	client, err := ecs.NewClientWithStsToken("cn-hangzhou", credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_DescribeClusterDetailWithROArequestWithAK(t *testing.T) {
	client, err := cs.NewClientWithAccessKey("default", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateDescribeClusterDetailRequest()
	request.SetDomain("cs.aliyuncs.com")
	request.QueryParams["RegionId"] = "default"
	request.Method = "GET"
	response, err := client.DescribeClusterDetail(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

//func Test_DescribeRegionsWithRPCrequestWithArn(t *testing.T) {
//	assumeresponse, err := createAssumeRole()
//	assert.Nil(t, assumeresponse)
//	credential := assumeresponse.Credentials
//	assert.Nil(t, err)
//	client, err := ecs.NewClientWithRamRoleArn("cn-hangzhou", credential.AccessKeyId, credential.AccessKeySecret, assumeresponse.AssumedRoleUser.Arn,"testrole")
//	assert.Nil(t, err)
//	request := ecs.CreateDescribeRegionsRequest()
//	request.Scheme = "https"
//	response, err := client.DescribeRegions(request)
//	assert.Nil(t, err)
//	realerr := err.(errors.Error)
//	assert.Nil(t, realerr.OriginError().Error())
//	assert.Equal(t, 36, len(response.RequestId))
//}
