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

// StopServerlessJobs invokes the ehpc.StopServerlessJobs API synchronously
func (client *Client) StopServerlessJobs(request *StopServerlessJobsRequest) (response *StopServerlessJobsResponse, err error) {
	response = CreateStopServerlessJobsResponse()
	err = client.DoAction(request, response)
	return
}

// StopServerlessJobsWithChan invokes the ehpc.StopServerlessJobs API asynchronously
func (client *Client) StopServerlessJobsWithChan(request *StopServerlessJobsRequest) (<-chan *StopServerlessJobsResponse, <-chan error) {
	responseChan := make(chan *StopServerlessJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopServerlessJobs(request)
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

// StopServerlessJobsWithCallback invokes the ehpc.StopServerlessJobs API asynchronously
func (client *Client) StopServerlessJobsWithCallback(request *StopServerlessJobsRequest, callback func(response *StopServerlessJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopServerlessJobsResponse
		var err error
		defer close(result)
		response, err = client.StopServerlessJobs(request)
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

// StopServerlessJobsRequest is the request struct for api StopServerlessJobs
type StopServerlessJobsRequest struct {
	*requests.RpcRequest
	ClusterId string    `position:"Query" name:"ClusterId"`
	JobIds    *[]string `position:"Query" name:"JobIds"  type:"Repeated"`
}

// StopServerlessJobsResponse is the response struct for api StopServerlessJobs
type StopServerlessJobsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopServerlessJobsRequest creates a request to invoke StopServerlessJobs API
func CreateStopServerlessJobsRequest() (request *StopServerlessJobsRequest) {
	request = &StopServerlessJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StopServerlessJobs", "ehs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopServerlessJobsResponse creates a response to parse from StopServerlessJobs response
func CreateStopServerlessJobsResponse() (response *StopServerlessJobsResponse) {
	response = &StopServerlessJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
