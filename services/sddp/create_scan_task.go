package sddp

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

// CreateScanTask invokes the sddp.CreateScanTask API synchronously
func (client *Client) CreateScanTask(request *CreateScanTaskRequest) (response *CreateScanTaskResponse, err error) {
	response = CreateCreateScanTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScanTaskWithChan invokes the sddp.CreateScanTask API asynchronously
func (client *Client) CreateScanTaskWithChan(request *CreateScanTaskRequest) (<-chan *CreateScanTaskResponse, <-chan error) {
	responseChan := make(chan *CreateScanTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScanTask(request)
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

// CreateScanTaskWithCallback invokes the sddp.CreateScanTask API asynchronously
func (client *Client) CreateScanTaskWithCallback(request *CreateScanTaskRequest, callback func(response *CreateScanTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScanTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateScanTask(request)
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

// CreateScanTaskRequest is the request struct for api CreateScanTask
type CreateScanTaskRequest struct {
	*requests.RpcRequest
	RunHour          requests.Integer `position:"Query" name:"RunHour"`
	ScanRangeContent string           `position:"Query" name:"ScanRangeContent"`
	TaskName         string           `position:"Query" name:"TaskName"`
	DataLimitId      requests.Integer `position:"Query" name:"DataLimitId"`
	RunMinute        requests.Integer `position:"Query" name:"RunMinute"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	IntervalDay      requests.Integer `position:"Query" name:"IntervalDay"`
	ScanRange        requests.Integer `position:"Query" name:"ScanRange"`
	Lang             string           `position:"Query" name:"Lang"`
	FeatureType      requests.Integer `position:"Query" name:"FeatureType"`
	OssScanPath      string           `position:"Query" name:"OssScanPath"`
	ResourceType     requests.Integer `position:"Query" name:"ResourceType"`
	TaskUserName     string           `position:"Query" name:"TaskUserName"`
}

// CreateScanTaskResponse is the response struct for api CreateScanTask
type CreateScanTaskResponse struct {
	*responses.BaseResponse
	Id        int    `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateScanTaskRequest creates a request to invoke CreateScanTask API
func CreateCreateScanTaskRequest() (request *CreateScanTaskRequest) {
	request = &CreateScanTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "CreateScanTask", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateScanTaskResponse creates a response to parse from CreateScanTask response
func CreateCreateScanTaskResponse() (response *CreateScanTaskResponse) {
	response = &CreateScanTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
