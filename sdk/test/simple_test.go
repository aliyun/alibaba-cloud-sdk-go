package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
)

type TestConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	PublicKeyId     string
	PrivateKey      string
	RoleArn         string
	ChildAK         string
	ChildSecret     string
}

func getConfigFromEnv() *TestConfig {
	config := &TestConfig{
		AccessKeyId:     os.Getenv("ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("ACCESS_KEY_SECRET"),
		PublicKeyId:     os.Getenv("PUBLIC_KEY_ID"),
		PrivateKey:      os.Getenv("PRIVATE_KEY"),
		RoleArn:         os.Getenv("ROLE_ARN"),
		ChildAK:         os.Getenv("CHILD_AK"),
		ChildSecret:     os.Getenv("CHILD_SECRET"),
	}
	if config.AccessKeyId == "" {
		return nil
	} else {
		return config
	}
}

func TestSimple(t *testing.T){
	config := getConfigFromEnv()
	ecsClient, err := ecs.NewClientWithAccessKey("cn-hangzhou", config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	request := ecs.CreateDescribeRegionsRequest()
	response, err := ecsClient.DescribeRegions(request)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 200, response.GetHttpStatus())
	fmt.Print(response.GetHttpContentString())
}
