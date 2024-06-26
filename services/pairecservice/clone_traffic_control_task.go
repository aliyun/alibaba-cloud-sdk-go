package pairecservice

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

// CloneTrafficControlTask invokes the pairecservice.CloneTrafficControlTask API synchronously
func (client *Client) CloneTrafficControlTask(request *CloneTrafficControlTaskRequest) (response *CloneTrafficControlTaskResponse, err error) {
	response = CreateCloneTrafficControlTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CloneTrafficControlTaskWithChan invokes the pairecservice.CloneTrafficControlTask API asynchronously
func (client *Client) CloneTrafficControlTaskWithChan(request *CloneTrafficControlTaskRequest) (<-chan *CloneTrafficControlTaskResponse, <-chan error) {
	responseChan := make(chan *CloneTrafficControlTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloneTrafficControlTask(request)
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

// CloneTrafficControlTaskWithCallback invokes the pairecservice.CloneTrafficControlTask API asynchronously
func (client *Client) CloneTrafficControlTaskWithCallback(request *CloneTrafficControlTaskRequest, callback func(response *CloneTrafficControlTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloneTrafficControlTaskResponse
		var err error
		defer close(result)
		response, err = client.CloneTrafficControlTask(request)
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

// CloneTrafficControlTaskRequest is the request struct for api CloneTrafficControlTask
type CloneTrafficControlTaskRequest struct {
	*requests.RoaRequest
	TrafficControlTaskId string `position:"Path" name:"TrafficControlTaskId"`
	Body                 string `position:"Body" name:"body"`
}

// CloneTrafficControlTaskResponse is the response struct for api CloneTrafficControlTask
type CloneTrafficControlTaskResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	TrafficControlTaskId string `json:"TrafficControlTaskId" xml:"TrafficControlTaskId"`
}

// CreateCloneTrafficControlTaskRequest creates a request to invoke CloneTrafficControlTask API
func CreateCloneTrafficControlTaskRequest() (request *CloneTrafficControlTaskRequest) {
	request = &CloneTrafficControlTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CloneTrafficControlTask", "/api/v1/trafficcontroltasks/[TrafficControlTaskId]/action/clone", "", "")
	request.Method = requests.POST
	return
}

// CreateCloneTrafficControlTaskResponse creates a response to parse from CloneTrafficControlTask response
func CreateCloneTrafficControlTaskResponse() (response *CloneTrafficControlTaskResponse) {
	response = &CloneTrafficControlTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
