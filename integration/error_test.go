package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/stretchr/testify/assert"

	"os"
	"strings"
	"testing"
)

func Test_DescribeRegionsWithParameterError(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "Describe"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	_, err = client.ProcessCommonRequest(request)
	realerr := err.(errors.Error)
	assert.Equal(t, "InvalidParameter", realerr.ErrorCode())
	assert.Equal(t, "The specified parameter \"Action or Version\" is not valid.", realerr.Message())
}


func Test_DescribeRegionsWithUnreachableError(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("www.ecs.ali.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Equal(t, 0, response.GetHttpStatus())
	assert.True(t, strings.Contains(err.Error(), "www.ecs.ali.com"))
}

//func Test_DescribeRegionsWithTimeout(t *testing.T) {
//	rolename, _, err := createRole(os.Getenv("USER_ID"))
//	credentail := credentials.NewEcsRamRoleCredential(rolename)
//	request := ecs.CreateDescribeRegionsRequest()
//	request.SetDomain("www.aliyun-hangzhou.com")
//	config := &sdk.Config{
//		Timeout: 1,
//	}
//	client, err := ecs.NewClientWithOptions("cn-hangzhou", config, credentail)
//	assert.Nil(t, err)
//	_, err = client.DescribeRegions(request)
//}

//func Test_DescribeRegionsWithTimeOutError(t *testing.T) {
//
//}
