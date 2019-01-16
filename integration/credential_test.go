package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func Test_DescribeRegionsWithRPCrequestWithAK(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
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
	client, err := ecs.NewClientWithStsToken(os.Getenv("REGION_ID"), credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
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
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateDescribeClusterDetailRequest()
	request.SetDomain("cs.aliyuncs.com")
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.Method = "GET"
	response, err := client.DescribeClusterDetail(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_DescribeRegionsWithRPCrequestWithArn(t *testing.T) {
	subaccesskeyid, subaccesskeysecret, err := createAccessKey()
	assert.Nil(t, err)
	client, err := ecs.NewClientWithRamRoleArn(os.Getenv("REGION_ID"), subaccesskeyid, subaccesskeysecret, rolearn, "alice_test")
	assert.Nil(t, err)

	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	request.Domain = "ecs.aliyuncs.com"
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.Equal(t, 36, len(response.RequestId))
}
