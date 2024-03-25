package fnf

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

// ListFlows invokes the fnf.ListFlows API synchronously
func (client *Client) ListFlows(request *ListFlowsRequest) (response *ListFlowsResponse, err error) {
	response = CreateListFlowsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowsWithChan invokes the fnf.ListFlows API asynchronously
func (client *Client) ListFlowsWithChan(request *ListFlowsRequest) (<-chan *ListFlowsResponse, <-chan error) {
	responseChan := make(chan *ListFlowsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlows(request)
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

// ListFlowsWithCallback invokes the fnf.ListFlows API asynchronously
func (client *Client) ListFlowsWithCallback(request *ListFlowsRequest, callback func(response *ListFlowsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowsResponse
		var err error
		defer close(result)
		response, err = client.ListFlows(request)
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

// ListFlowsRequest is the request struct for api ListFlows
type ListFlowsRequest struct {
	*requests.RpcRequest
	NextToken string           `position:"Query" name:"NextToken"`
	Limit     requests.Integer `position:"Query" name:"Limit"`
}

// ListFlowsResponse is the response struct for api ListFlows
type ListFlowsResponse struct {
	*responses.BaseResponse
	NextToken string      `json:"NextToken" xml:"NextToken"`
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Flows     []FlowsItem `json:"Flows" xml:"Flows"`
}

// CreateListFlowsRequest creates a request to invoke ListFlows API
func CreateListFlowsRequest() (request *ListFlowsRequest) {
	request = &ListFlowsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "ListFlows", "fnf", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListFlowsResponse creates a response to parse from ListFlows response
func CreateListFlowsResponse() (response *ListFlowsResponse) {
	response = &ListFlowsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
