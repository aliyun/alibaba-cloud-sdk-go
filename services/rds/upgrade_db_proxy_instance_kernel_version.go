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

// UpgradeDBProxyInstanceKernelVersion invokes the rds.UpgradeDBProxyInstanceKernelVersion API synchronously
func (client *Client) UpgradeDBProxyInstanceKernelVersion(request *UpgradeDBProxyInstanceKernelVersionRequest) (response *UpgradeDBProxyInstanceKernelVersionResponse, err error) {
	response = CreateUpgradeDBProxyInstanceKernelVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeDBProxyInstanceKernelVersionWithChan invokes the rds.UpgradeDBProxyInstanceKernelVersion API asynchronously
func (client *Client) UpgradeDBProxyInstanceKernelVersionWithChan(request *UpgradeDBProxyInstanceKernelVersionRequest) (<-chan *UpgradeDBProxyInstanceKernelVersionResponse, <-chan error) {
	responseChan := make(chan *UpgradeDBProxyInstanceKernelVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeDBProxyInstanceKernelVersion(request)
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

// UpgradeDBProxyInstanceKernelVersionWithCallback invokes the rds.UpgradeDBProxyInstanceKernelVersion API asynchronously
func (client *Client) UpgradeDBProxyInstanceKernelVersionWithCallback(request *UpgradeDBProxyInstanceKernelVersionRequest, callback func(response *UpgradeDBProxyInstanceKernelVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeDBProxyInstanceKernelVersionResponse
		var err error
		defer close(result)
		response, err = client.UpgradeDBProxyInstanceKernelVersion(request)
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

// UpgradeDBProxyInstanceKernelVersionRequest is the request struct for api UpgradeDBProxyInstanceKernelVersion
type UpgradeDBProxyInstanceKernelVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	SwitchTime           string           `position:"Query" name:"SwitchTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBProxyEngineType    string           `position:"Query" name:"DBProxyEngineType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UpgradeTime          string           `position:"Query" name:"UpgradeTime"`
}

// UpgradeDBProxyInstanceKernelVersionResponse is the response struct for api UpgradeDBProxyInstanceKernelVersion
type UpgradeDBProxyInstanceKernelVersionResponse struct {
	*responses.BaseResponse
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
}

// CreateUpgradeDBProxyInstanceKernelVersionRequest creates a request to invoke UpgradeDBProxyInstanceKernelVersion API
func CreateUpgradeDBProxyInstanceKernelVersionRequest() (request *UpgradeDBProxyInstanceKernelVersionRequest) {
	request = &UpgradeDBProxyInstanceKernelVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBProxyInstanceKernelVersion", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeDBProxyInstanceKernelVersionResponse creates a response to parse from UpgradeDBProxyInstanceKernelVersion response
func CreateUpgradeDBProxyInstanceKernelVersionResponse() (response *UpgradeDBProxyInstanceKernelVersionResponse) {
	response = &UpgradeDBProxyInstanceKernelVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
