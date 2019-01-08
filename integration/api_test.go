package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func Test_ECS_DescribeRegionsWithRPCrequest(t *testing.T) {
	request := ecs.CreateDescribeRegionsRequest()
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_RDS_DescribeDBInstancesWithRPCrequest(t *testing.T) {
	client, err := rds.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"), )
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := rds.CreateDescribeDBInstancesRequest()
	response, err := client.DescribeDBInstances(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_CDN_DescribeCdnDomainDetailWithRPCrequest(t *testing.T) {
	client, err := cdn.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"), )
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := cdn.CreateDescribeCdnCertificateDetailRequest()
	request.CertName = "sdk-test"
	response, err := client.DescribeCdnCertificateDetail(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_RAM_ListRolesWithRPCrequest(t *testing.T) {
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ram.CreateListRolesRequest()
	request.Scheme = "HTTPS"
	response, err := client.ListRoles(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_SLB_DescribeRegionsWithRPCrequest(t *testing.T) {
	client, err := slb.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := slb.CreateDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_VPC_DescribeRegionsWithUnicodeSpecificParams(t *testing.T) {
	client, err := vpc.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := vpc.CreateDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

//func Test_DescribeRegionsWithHTTPS(t *testing.T) {
//	request := ecs.CreateDescribeRegionsRequest()
//	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
//	assert.Nil(t, err)
//	response, err := client.DescribeRegions(request)
//	assert.Nil(t, err)
//	assert.True(t, response.IsSuccess())
//	assert.Equal(t, 36, len(response.RequestId))
//	assert.True(t, len(response.Regions.Region) > 0)
//}
//
//func Test_DescribeRegionsWithSTStoken(t *testing.T) {
//	assumeresponse, err := createAssumeRole()
//	assert.Nil(t, err)
//	credential := assumeresponse.Credentials
//	request := ecs.CreateDescribeRegionsRequest()
//	client, err := ecs.NewClientWithStsToken("cn-hangzhou", credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
//	response, err := client.DescribeRegions(request)
//	assert.Nil(t, err)
//	assert.True(t, response.IsSuccess())
//	assert.Equal(t, 36, len(response.RequestId))
//	assert.True(t, len(response.Regions.Region) > 0)
//}
