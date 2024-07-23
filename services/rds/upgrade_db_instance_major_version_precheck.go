package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpgradeDBInstanceMajorVersionPrecheck invokes the rds.UpgradeDBInstanceMajorVersionPrecheck API synchronously
func (client *Client) UpgradeDBInstanceMajorVersionPrecheck(request *UpgradeDBInstanceMajorVersionPrecheckRequest) (response *UpgradeDBInstanceMajorVersionPrecheckResponse, err error) {
	response = CreateUpgradeDBInstanceMajorVersionPrecheckResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeDBInstanceMajorVersionPrecheckWithChan invokes the rds.UpgradeDBInstanceMajorVersionPrecheck API asynchronously
func (client *Client) UpgradeDBInstanceMajorVersionPrecheckWithChan(request *UpgradeDBInstanceMajorVersionPrecheckRequest) (<-chan *UpgradeDBInstanceMajorVersionPrecheckResponse, <-chan error) {
	responseChan := make(chan *UpgradeDBInstanceMajorVersionPrecheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeDBInstanceMajorVersionPrecheck(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpgradeDBInstanceMajorVersionPrecheckWithCallback invokes the rds.UpgradeDBInstanceMajorVersionPrecheck API asynchronously
func (client *Client) UpgradeDBInstanceMajorVersionPrecheckWithCallback(request *UpgradeDBInstanceMajorVersionPrecheckRequest, callback func(response *UpgradeDBInstanceMajorVersionPrecheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeDBInstanceMajorVersionPrecheckResponse
		var err error
		defer close(result)
		response, err = client.UpgradeDBInstanceMajorVersionPrecheck(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpgradeDBInstanceMajorVersionPrecheckRequest is the request struct for api UpgradeDBInstanceMajorVersionPrecheck
type UpgradeDBInstanceMajorVersionPrecheckRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TargetMajorVersion   string           `position:"Query" name:"TargetMajorVersion"`
}

// UpgradeDBInstanceMajorVersionPrecheckResponse is the response struct for api UpgradeDBInstanceMajorVersionPrecheck
type UpgradeDBInstanceMajorVersionPrecheckResponse struct {
	*responses.BaseResponse
	DBInstanceName     string `json:"DBInstanceName" xml:"DBInstanceName"`
	TargetMajorVersion string `json:"TargetMajorVersion" xml:"TargetMajorVersion"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	TaskId             string `json:"TaskId" xml:"TaskId"`
}

// CreateUpgradeDBInstanceMajorVersionPrecheckRequest creates a request to invoke UpgradeDBInstanceMajorVersionPrecheck API
func CreateUpgradeDBInstanceMajorVersionPrecheckRequest() (request *UpgradeDBInstanceMajorVersionPrecheckRequest) {
	request = &UpgradeDBInstanceMajorVersionPrecheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceMajorVersionPrecheck", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeDBInstanceMajorVersionPrecheckResponse creates a response to parse from UpgradeDBInstanceMajorVersionPrecheck response
func CreateUpgradeDBInstanceMajorVersionPrecheckResponse() (response *UpgradeDBInstanceMajorVersionPrecheckResponse) {
	response = &UpgradeDBInstanceMajorVersionPrecheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
