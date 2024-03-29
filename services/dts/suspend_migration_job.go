package dts

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

// SuspendMigrationJob invokes the dts.SuspendMigrationJob API synchronously
func (client *Client) SuspendMigrationJob(request *SuspendMigrationJobRequest) (response *SuspendMigrationJobResponse, err error) {
	response = CreateSuspendMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendMigrationJobWithChan invokes the dts.SuspendMigrationJob API asynchronously
func (client *Client) SuspendMigrationJobWithChan(request *SuspendMigrationJobRequest) (<-chan *SuspendMigrationJobResponse, <-chan error) {
	responseChan := make(chan *SuspendMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendMigrationJob(request)
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

// SuspendMigrationJobWithCallback invokes the dts.SuspendMigrationJob API asynchronously
func (client *Client) SuspendMigrationJobWithCallback(request *SuspendMigrationJobRequest, callback func(response *SuspendMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.SuspendMigrationJob(request)
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

// SuspendMigrationJobRequest is the request struct for api SuspendMigrationJob
type SuspendMigrationJobRequest struct {
	*requests.RpcRequest
	ClientToken    string `position:"Query" name:"ClientToken"`
	MigrationJobId string `position:"Query" name:"MigrationJobId"`
}

// SuspendMigrationJobResponse is the response struct for api SuspendMigrationJob
type SuspendMigrationJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateSuspendMigrationJobRequest creates a request to invoke SuspendMigrationJob API
func CreateSuspendMigrationJobRequest() (request *SuspendMigrationJobRequest) {
	request = &SuspendMigrationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2017-06-01", "SuspendMigrationJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendMigrationJobResponse creates a response to parse from SuspendMigrationJob response
func CreateSuspendMigrationJobResponse() (response *SuspendMigrationJobResponse) {
	response = &SuspendMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
