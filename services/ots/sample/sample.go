package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ots"
	"log"
	"os"
)

var otsDomainTemplate = "ots.%s.aliyuncs.com"
var triggerArnTemplate = "acs:fc:%s:%s:services/%s/functions/%s/triggers/%s"

func main() {
	akId := os.Getenv("ACCESSKEY_ID")
	akSec := os.Getenv("ACCESSKEY_SECRET")

	testInstance := "existInstance"
	testTable := "StreamTable"
	region := "cn-beijing"
	uid := os.Getenv("USER_ID")

	adapter, err := NewOTSAdapter(region, akId, akSec)
	if err != nil {
		log.Fatal(err)
	}

	// ERROR_ACCESS_DENIED

	// ... ignore, return when ram checking permission failed for sub account

	// ERROR_TRIGGER_NOT_EXIST, arn is empty
	info, err := adapter.ReadTrigger(testInstance, testTable, "not-exist-trigger", "")
	if err == nil {
		log.Fatal("Unexpected trigger info:", info)
	}
	if ok, popErr := parsePopError(err); ok {
		log.Println("http code:", popErr.HttpStatus())
		log.Println("error code:", popErr.ErrorCode())
		log.Println("error message:", popErr.Message())
	} else {
		log.Println("Unexpected error:", err)
	}

	// ERROR_TRIGGER_NOT_EXIST, arn is not empty
	arn := fmt.Sprintf(triggerArnTemplate, region, uid, "fcService", "fcFunction", "not-exist-trigger")
	err = adapter.DeleteTrigger(testInstance, testTable, "not-exist-trigger", arn)
	if err == nil {
		log.Fatal("Unexpected delete succeed")
	}
	if ok, popErr := parsePopError(err); ok {
		log.Println("http code:", popErr.HttpStatus())
		log.Println("error code:", popErr.ErrorCode())
		log.Println("error message:", popErr.Message())
	} else {
		log.Println("Unexpected error:", err)
	}

	// ERROR_INVALID_ARGUMENT
	info, err = adapter.ReadTrigger("notExistInstance", testTable, "not-exist-trigger", "")
	if err == nil {
		log.Fatal("Unexpected trigger info:", info)
	}
	if ok, popErr := parsePopError(err); ok {
		log.Println("http code:", popErr.HttpStatus())
		log.Println("error code:", popErr.ErrorCode())
		log.Println("error message:", popErr.Message())
	} else {
		log.Println("Unexpected error:", err)
	}

	// CRUD
	arn = fmt.Sprintf(triggerArnTemplate, region, uid, "fcService", "fcFunction", "testTrigger")
	triggerMeta := &ots.CreateTriggerRequestBody{
		TriggerName: "testTrigger",
		TriggerArn:  arn,
		RoleArn:     fmt.Sprintf("acs:ram::%s:role/aliyuntablestorestreamnotificationrole", uid),
		UdfInfo: &ots.UdfInfo{
			ServiceName:  "fcService",
			FunctionName: "fcFunction",
		},
	}
	// CREATE
	newTrigger, err := adapter.CreateTrigger(testInstance, testTable, triggerMeta)
	if err != nil {
		log.Fatal("Create trigger failed:", err)
	}
	log.Println("new trigger id:", newTrigger.Etag)

	// READ
	info, err = adapter.ReadTrigger(testInstance, testTable, triggerMeta.TriggerName, arn)
	if err != nil {
		log.Fatal("Read trigger failed:", err)
	}
	log.Println("read trigger info:", info.Trigger)

	// UPDATE
	// not support update

	// ERROR_RESOURCE_CONFLICT
	conflictMeta := &ots.CreateTriggerRequestBody{
		TriggerName: "testTrigger",
		TriggerArn:  arn,
		RoleArn:     fmt.Sprintf("acs:ram::%s:role/aliyuntablestorestreamnotificationrole", uid),
		UdfInfo: &ots.UdfInfo{
			ServiceName:  "fcService",
			FunctionName: "fcFunction",
		},
	}
	conflictTrigger, err := adapter.CreateTrigger(testInstance, testTable, conflictMeta)
	if err == nil {
		log.Fatal("Uexpected trigger:", conflictTrigger)
	}
	if ok, popErr := parsePopError(err); ok {
		log.Println("http code:", popErr.HttpStatus())
		log.Println("error code:", popErr.ErrorCode())
		log.Println("error message:", popErr.Message())
	} else {
		log.Println("Unexpected error:", err)
	}

	// DELETE
	err = adapter.DeleteTrigger(testInstance, testTable, triggerMeta.TriggerName, arn)
	if err != nil {
		log.Fatal("Delete trigger failed:", err)
	}
	log.Println("CRUD done")
}

type OTSAdapter struct {
	client *ots.Client
	domain string
}

func NewOTSAdapter(regionId string, accessKeyId, accessKeySecret string) (*OTSAdapter, error) {
	client, err := ots.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &OTSAdapter{client: client, domain: fmt.Sprintf(otsDomainTemplate, regionId)}, nil
}

func (o *OTSAdapter) CreateTrigger(instanceName, tableName string, triggerInfo *ots.CreateTriggerRequestBody) (resp *ots.CreateTriggerResponseBody, err error) {
	req, err := ots.NewCreateTriggerRequest(o.domain, instanceName, tableName, triggerInfo)
	if err != nil {
		return
	}
	popResp, err := o.client.CreateTrigger(req)
	if err != nil {
		return
	}
	resp, err = popResp.GetBody()
	return
}

func (o *OTSAdapter) ReadTrigger(instanceName, tableName, triggerName, triggerArn string) (triggerInfo *ots.GetTriggerResponseBody, err error) {
	popResp, err := o.client.GetTrigger(ots.NewGetTriggerRequest(o.domain, instanceName, tableName, triggerName, triggerArn))
	if err != nil {
		return
	}
	triggerInfo, err = popResp.GetBody()
	return
}

func (o *OTSAdapter) DeleteTrigger(instanceName, tableName, triggerName, triggerArn string) error {
	_, err := o.client.DeleteTrigger(ots.NewDeleteTriggerRequest(o.domain, instanceName, tableName, triggerName, triggerArn))
	return err
}

// fc应该用不到这个接口
func (o *OTSAdapter) ListTriggers(instanceName, tableName string) (triggers *ots.ListTriggerResponseBody, err error) {
	popResp, err := o.client.ListTrigger(ots.NewListTriggerRequest(o.domain, instanceName, tableName))
	if err != nil {
		return
	}
	triggers, err = popResp.GetBody()
	return
}

func parsePopError(err error) (succeed bool, error *errors.ServerError) {
	if sdkErr, ok := err.(*errors.ServerError); ok {
		return true, sdkErr
	}
	return false, nil
}
