package mts

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

// UpdateMediaWorkflow invokes the mts.UpdateMediaWorkflow API synchronously
func (client *Client) UpdateMediaWorkflow(request *UpdateMediaWorkflowRequest) (response *UpdateMediaWorkflowResponse, err error) {
	response = CreateUpdateMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMediaWorkflowWithChan invokes the mts.UpdateMediaWorkflow API asynchronously
func (client *Client) UpdateMediaWorkflowWithChan(request *UpdateMediaWorkflowRequest) (<-chan *UpdateMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *UpdateMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMediaWorkflow(request)
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

// UpdateMediaWorkflowWithCallback invokes the mts.UpdateMediaWorkflow API asynchronously
func (client *Client) UpdateMediaWorkflowWithCallback(request *UpdateMediaWorkflowRequest, callback func(response *UpdateMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.UpdateMediaWorkflow(request)
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

// UpdateMediaWorkflowRequest is the request struct for api UpdateMediaWorkflow
type UpdateMediaWorkflowRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MediaWorkflowId      string           `position:"Query" name:"MediaWorkflowId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string           `position:"Query" name:"Topology"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TriggerMode          string           `position:"Query" name:"TriggerMode"`
	Name                 string           `position:"Query" name:"Name"`
}

// UpdateMediaWorkflowResponse is the response struct for api UpdateMediaWorkflow
type UpdateMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	MediaWorkflow MediaWorkflow `json:"MediaWorkflow" xml:"MediaWorkflow"`
}

// CreateUpdateMediaWorkflowRequest creates a request to invoke UpdateMediaWorkflow API
func CreateUpdateMediaWorkflowRequest() (request *UpdateMediaWorkflowRequest) {
	request = &UpdateMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaWorkflow", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateMediaWorkflowResponse creates a response to parse from UpdateMediaWorkflow response
func CreateUpdateMediaWorkflowResponse() (response *UpdateMediaWorkflowResponse) {
	response = &UpdateMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
