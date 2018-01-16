package integration

import (
	"testing"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestCdnInstance(t *testing.T) {

	// init client
	config := getConfigFromEnv()
	cdnClient, err := cdn.NewClientWithAccessKey("cn-hangzhou", config.AccessKeyId, config.AccessKeySecret)
	assertErrorNil(t, err, "Failed to init client")
	fmt.Printf("Init client success\n")

	// getCdnStatus
	assertCdnStatus(t, cdnClient)

	// test travis if the key is hidden
	testKey := os.Getenv("TestKey")
	fmt.Println("test key : " + testKey)
}

func assertCdnStatus(t *testing.T, client *cdn.Client){
	fmt.Print("describing cdn service status...")
	request := cdn.CreateDescribeCdnServiceRequest()
	response, err := client.DescribeCdnService(request)
	assertErrorNil(t, err, "Failed to describing cdn service status")
	assert.Equal(t, 200, response.GetHttpStatus(), response.GetHttpContentString())
	assert.Equal(t, "PayByTraffic", response.InternetChargeType)
	fmt.Printf("ok(%d)!\n", response.GetHttpStatus())
}
