package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	securityGroupId = ""
	flag            = false
)

func Test_ECS_CreateSecurityGroupWithRPCrequestWithXMLWithNestingparameters(t *testing.T) {
	request := ecs.CreateCreateSecurityGroupRequest()
	request.SetContentType("XML")
	tag := ecs.CreateSecurityGroupTag{
		Key:   "test",
		Value: "test",
	}
	request.Tag = &[]ecs.CreateSecurityGroupTag{tag}
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.CreateSecurityGroup(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.SecurityGroupId) > 0)
	securityGroupId = response.SecurityGroupId
}

func Test_ECS_DescribeSecurityGroupsWithRPCrequestWithXMLWithNestingparameters(t *testing.T) {
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.SetContentType("XML")
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.DescribeSecurityGroups(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	for _, securitygroup := range response.SecurityGroups.SecurityGroup {
		if securitygroup.SecurityGroupId == securityGroupId {
			flag = true
			break
		}
	}
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, flag)
	flag = false

}

func Test_ECS_ModifyImageAttributeWithRPCrequestWithXML(t *testing.T) {
	request := ecs.CreateModifySecurityGroupAttributeRequest()
	request.SetContentType("XML")
	request.SecurityGroupName = "testintegration" + strings.Split(os.Getenv("TRAVIS_JOB_NUMBER"), ".")[0]
	request.SecurityGroupId = securityGroupId
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ModifySecurityGroupAttribute(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_ECS_DeleteImageWithRPCrequestWithXMLWithNestingparameters(t *testing.T) {
	request := ecs.CreateDeleteSecurityGroupRequest()
	request.SetContentType("XML")
	request.SecurityGroupId = securityGroupId
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.DeleteSecurityGroup(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	securityGroupId = ""
}

func Test_ECS_CreateSecurityGroupWithRPCrequestWithJSONWithNestingparameters(t *testing.T) {
	request := ecs.CreateCreateSecurityGroupRequest()
	request.SetContentType("JSON")
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.CreateSecurityGroup(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.SecurityGroupId) > 0)
	securityGroupId = response.SecurityGroupId
}

func Test_ECS_DescribeSecurityGroupsWithRPCrequestWithJSONWithNestingparameters(t *testing.T) {
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.SetContentType("JSON")
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.DescribeSecurityGroups(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	for _, securitygroup := range response.SecurityGroups.SecurityGroup {
		if securitygroup.SecurityGroupId == securityGroupId {
			flag = true
			break
		}
	}
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, flag)
	flag = false

}

func Test_ECS_ModifyImageAttributeWithRPCrequestWithJSON(t *testing.T) {
	request := ecs.CreateModifySecurityGroupAttributeRequest()
	request.SetContentType("JSON")
	request.SecurityGroupName = "testintegration" + strings.Split(os.Getenv("TRAVIS_JOB_NUMBER"), ".")[0]
	request.SecurityGroupId = securityGroupId
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ModifySecurityGroupAttribute(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_ECS_DeleteImageWithRPCrequestWithJSONWithNestingparameters(t *testing.T) {
	request := ecs.CreateDeleteSecurityGroupRequest()
	request.SetContentType("JSON")
	request.SecurityGroupId = securityGroupId
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.DeleteSecurityGroup(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	securityGroupId = ""
}

func Test_ECS_DescribeRegionsWithRPCrequestWith(t *testing.T) {
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
	client, err := rds.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := rds.CreateDescribeDBInstancesRequest()
	response, err := client.DescribeDBInstances(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_CDN_DescribeCdnDomainDetailWithRPCrequest(t *testing.T) {
	client, err := cdn.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := cdn.CreateDescribeRefreshTasksRequest()
	response, err := client.DescribeRefreshTasks(request)
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


func mockServer(status int, json string) (server *httptest.Server) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(json))
		return
	}))
	return ts
}

func Test_ListRolesWithRPCrequestWithoutRedirecting(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "HTTP"
	ts := mockServer(404, `{"Code": "YouMessedSomethingUp"}`)
	defer ts.Close()
	domain := strings.Replace(ts.URL, "http://", "",1)
	request.Domain = domain
	response, err := client.DescribeRegions(request)
	assert.NotNil(t, err)
	assert.Equal(t, 404, response.GetHttpStatus())
	assert.Equal(t, "{\"Code\": \"YouMessedSomethingUp\"}", response.GetHttpContentString())
}

//func Test_DescribeClusterDetailWithROArequest(t *testing.T) {
//	client, err := cs.NewClientWithAccessKey("default", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
//	assert.Nil(t, err)
//	request := cs.CreateDescribeClusterDetailRequest()
//	request.SetDomain("cs.aliyuncs.com")
//	request.QueryParams["RegionId"] = "default"
//	request.Method = "GET"
//	response, err := client.DescribeClusterDetail(request)
//	assert.Nil(t, err)
//	assert.True(t, response.IsSuccess())
//	assert.Nil(t, response.GetHttpHeaders())
//}
