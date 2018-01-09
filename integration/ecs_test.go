package integration

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

const (
	InstanceDefaultTimeout = 120
	DefaultWaitForInterval = 10

	Running = "Running"
	Stopped = "Stopped"
	Deleted = "Deleted"
)

// create -> start -> stop -> delete
func TestEcsInstance(t *testing.T) {

	// init client
	config := getConfigFromEnv()
	ecsClient, err := ecs.NewClientWithAccessKey("cn-hangzhou", config.AccessKeyId, config.AccessKeySecret)
	assertErrorNil(t, err, "Failed to init client")
	fmt.Printf("Init client success\n")

	// get demo instance attributes
	param := getDemoInstanceAttributes(t, ecsClient)

	// create
	instanceId := createInstance(t, ecsClient, param)

	// defer wait for deleted
	defer waitForInstance(t, ecsClient, instanceId, Deleted, 120)

	// defer delete
	defer deleteInstance(t, ecsClient, instanceId)

	// wait
	waitForInstance(t, ecsClient, instanceId, Stopped, 60)

	// start
	startInstance(t, ecsClient, instanceId)

	// wait
	waitForInstance(t, ecsClient, instanceId, Running, 120)

	// stop
	stopInstance(t, ecsClient, instanceId)

	// wait
	waitForInstance(t, ecsClient, instanceId, Stopped, 600)
}

func getDemoInstanceAttributes(t *testing.T, client *ecs.Client) *ecs.DescribeInstanceAttributeResponse {
	fmt.Print("trying to get demo instance...")
	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.InstanceId = os.Getenv("DEMO_ECS_INSTANCE_ID")
	response, err := client.DescribeInstanceAttribute(request)
	assertErrorNil(t, err, "Failed to get demo instance attributes")
	assert.Equal(t, 200, response.GetHttpStatus(), response.GetHttpContentString())
	fmt.Println("success!")
	return response
}

func createInstance(t *testing.T, client *ecs.Client, param *ecs.DescribeInstanceAttributeResponse) (instanceId string) {
	fmt.Print("creating instance...")
	request := ecs.CreateCreateInstanceRequest()
	request.ImageId = param.ImageId
	request.SecurityGroupId = param.SecurityGroupIds.SecurityGroupId[0]
	request.InstanceType = "ecs.t1.small"
	response, err := client.CreateInstance(request)
	assertErrorNil(t, err, "Failed to create instance")
	assert.Equal(t, 200, response.GetHttpStatus(), response.GetHttpContentString())
	instanceId = response.InstanceId
	fmt.Printf("success(%d)! instanceId = %s\n", response.GetHttpStatus(), instanceId)
	return
}

func startInstance(t *testing.T, client *ecs.Client, instanceId string) {
	fmt.Printf("starting instance(%s)...", instanceId)
	request := ecs.CreateStartInstanceRequest()
	request.InstanceId = instanceId
	response, err := client.StartInstance(request)
	assertErrorNil(t, err, "Failed to start instance "+instanceId)
	assert.Equal(t, 200, response.GetHttpStatus(), response.GetHttpContentString())
	fmt.Println("success!")
}

func stopInstance(t *testing.T, client *ecs.Client, instanceId string) {
	fmt.Printf("stopping instance(%s)...", instanceId)
	request := ecs.CreateStopInstanceRequest()
	request.InstanceId = instanceId
	response, err := client.StopInstance(request)
	assertErrorNil(t, err, "Failed to stop instance "+instanceId)
	assert.Equal(t, 200, response.GetHttpStatus(), response.GetHttpContentString())
	fmt.Println("success!")
}

func deleteInstance(t *testing.T, client *ecs.Client, instanceId string) {
	fmt.Printf("deleting instance(%s)...", instanceId)
	request := ecs.CreateDeleteInstanceRequest()
	request.InstanceId = instanceId
	response, err := client.DeleteInstance(request)
	assertErrorNil(t, err, "Failed to delete instance "+instanceId)
	assert.Equal(t, 200, response.GetHttpStatus(), response.GetHttpContentString())
	fmt.Println("success!")
}

func assertErrorNil(t *testing.T, err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, message+": %v\n", err)
	}
}

func waitForInstance(t *testing.T, client *ecs.Client, instanceId string, status string, timeout int) {
	if timeout <= 0 {
		timeout = InstanceDefaultTimeout
	}
	for {
		request := ecs.CreateDescribeInstanceAttributeRequest()
		request.InstanceId = instanceId
		response, err := client.DescribeInstanceAttribute(request)
		assertErrorNil(t, err, "Failed to create describe instance request\n")

		if status == Deleted {
			if response.GetHttpStatus() == 404 || response.Status == Deleted {
				fmt.Printf("delete instance(%s) success\n", instanceId)
				break
			}
		} else {
			if response.Status == status {
				fmt.Printf("instance(%s) status changed to %s, wait a moment\n", instanceId, status)
				time.Sleep(DefaultWaitForInterval * time.Second)
				break
			} else {
				fmt.Printf("instance(%s) status is %s, wait for changing to %s\n", instanceId, response.Status, status)
			}
		}

		timeout = timeout - DefaultWaitForInterval
		if timeout <= 0 {
			t.Errorf(fmt.Sprintf("wait for instance(%s) status to %s timeout(%d)\n", instanceId, status, timeout))
		}
		time.Sleep(DefaultWaitForInterval * time.Second)
	}
}
