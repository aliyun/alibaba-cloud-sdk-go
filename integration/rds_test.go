package integration

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/stretchr/testify/assert"
)

func Test_DescribeDBInstances(t *testing.T) {
	client, err := rds.NewClientWithAccessKey(
		"cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"),
	)
	assert.Nil(t, err)
	assert.NotNil(t, client)

	request := rds.CreateDescribeDBInstancesRequest()
	response, err := client.DescribeDBInstances(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.RequestId)
	// {"Items":{"DBInstance":[]},"TotalRecordCount":0,"PageNumber":1,"RequestId":"AB26A264-72DA-4924-A8BC-93891B35D8D0","PageRecordCount":0}
}
