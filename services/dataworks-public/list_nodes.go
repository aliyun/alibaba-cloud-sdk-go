package dataworks_public

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

// ListNodes invokes the dataworks_public.ListNodes API synchronously
func (client *Client) ListNodes(request *ListNodesRequest) (response *ListNodesResponse, err error) {
	response = CreateListNodesResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodesWithChan invokes the dataworks_public.ListNodes API asynchronously
func (client *Client) ListNodesWithChan(request *ListNodesRequest) (<-chan *ListNodesResponse, <-chan error) {
	responseChan := make(chan *ListNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodes(request)
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

// ListNodesWithCallback invokes the dataworks_public.ListNodes API asynchronously
func (client *Client) ListNodesWithCallback(request *ListNodesRequest, callback func(response *ListNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodesResponse
		var err error
		defer close(result)
		response, err = client.ListNodes(request)
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

// ListNodesRequest is the request struct for api ListNodes
type ListNodesRequest struct {
	*requests.RpcRequest
	Owner         string           `position:"Body" name:"Owner"`
	ProjectEnv    string           `position:"Body" name:"ProjectEnv"`
	BizName       string           `position:"Body" name:"BizName"`
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	NodeName      string           `position:"Body" name:"NodeName"`
	ProgramType   string           `position:"Body" name:"ProgramType"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	SchedulerType string           `position:"Body" name:"SchedulerType"`
	ProjectId     requests.Integer `position:"Body" name:"ProjectId"`
}

// ListNodesResponse is the response struct for api ListNodes
type ListNodesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string          `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string          `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool            `json:"Success" xml:"Success"`
	Data           DataInListNodes `json:"Data" xml:"Data"`
}

// CreateListNodesRequest creates a request to invoke ListNodes API
func CreateListNodesRequest() (request *ListNodesRequest) {
	request = &ListNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListNodes", "", "")
	request.Method = requests.POST
	return
}

// CreateListNodesResponse creates a response to parse from ListNodes response
func CreateListNodesResponse() (response *ListNodesResponse) {
	response = &ListNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
