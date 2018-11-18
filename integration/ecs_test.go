package integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func Test_DescribeRegions(t *testing.T) {
	request := ecs.CreateDescribeRegionsRequest()
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}
