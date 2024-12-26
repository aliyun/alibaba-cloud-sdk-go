package schedulerx3

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

// ImportJobs invokes the schedulerx3.ImportJobs API synchronously
func (client *Client) ImportJobs(request *ImportJobsRequest) (response *ImportJobsResponse, err error) {
	response = CreateImportJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ImportJobsWithChan invokes the schedulerx3.ImportJobs API asynchronously
func (client *Client) ImportJobsWithChan(request *ImportJobsRequest) (<-chan *ImportJobsResponse, <-chan error) {
	responseChan := make(chan *ImportJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportJobs(request)
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

// ImportJobsWithCallback invokes the schedulerx3.ImportJobs API asynchronously
func (client *Client) ImportJobsWithCallback(request *ImportJobsRequest, callback func(response *ImportJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportJobsResponse
		var err error
		defer close(result)
		response, err = client.ImportJobs(request)
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

// ImportJobsRequest is the request struct for api ImportJobs
type ImportJobsRequest struct {
	*requests.RpcRequest
	MseSessionId  string           `position:"Query" name:"MseSessionId"`
	Content       string           `position:"Body" name:"Content"`
	Overwrite     requests.Boolean `position:"Body" name:"Overwrite"`
	ClusterId     string           `position:"Body" name:"ClusterId"`
	AutoCreateApp requests.Boolean `position:"Body" name:"AutoCreateApp"`
}

// ImportJobsResponse is the response struct for api ImportJobs
type ImportJobsResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateImportJobsRequest creates a request to invoke ImportJobs API
func CreateImportJobsRequest() (request *ImportJobsRequest) {
	request = &ImportJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SchedulerX3", "2024-06-24", "ImportJobs", "", "")
	request.Method = requests.POST
	return
}

// CreateImportJobsResponse creates a response to parse from ImportJobs response
func CreateImportJobsResponse() (response *ImportJobsResponse) {
	response = &ImportJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
