package ehpc

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

// StopJobs invokes the ehpc.StopJobs API synchronously
func (client *Client) StopJobs(request *StopJobsRequest) (response *StopJobsResponse, err error) {
	response = CreateStopJobsResponse()
	err = client.DoAction(request, response)
	return
}

// StopJobsWithChan invokes the ehpc.StopJobs API asynchronously
func (client *Client) StopJobsWithChan(request *StopJobsRequest) (<-chan *StopJobsResponse, <-chan error) {
	responseChan := make(chan *StopJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopJobs(request)
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

// StopJobsWithCallback invokes the ehpc.StopJobs API asynchronously
func (client *Client) StopJobsWithCallback(request *StopJobsRequest, callback func(response *StopJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopJobsResponse
		var err error
		defer close(result)
		response, err = client.StopJobs(request)
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

// StopJobsRequest is the request struct for api StopJobs
type StopJobsRequest struct {
	*requests.RpcRequest
	Jobs      string           `position:"Query" name:"Jobs"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	Async     requests.Boolean `position:"Query" name:"Async"`
}

// StopJobsResponse is the response struct for api StopJobs
type StopJobsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopJobsRequest creates a request to invoke StopJobs API
func CreateStopJobsRequest() (request *StopJobsRequest) {
	request = &StopJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StopJobs", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateStopJobsResponse creates a response to parse from StopJobs response
func CreateStopJobsResponse() (response *StopJobsResponse) {
	response = &StopJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
