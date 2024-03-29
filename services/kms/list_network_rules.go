package kms

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

// ListNetworkRules invokes the kms.ListNetworkRules API synchronously
func (client *Client) ListNetworkRules(request *ListNetworkRulesRequest) (response *ListNetworkRulesResponse, err error) {
	response = CreateListNetworkRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListNetworkRulesWithChan invokes the kms.ListNetworkRules API asynchronously
func (client *Client) ListNetworkRulesWithChan(request *ListNetworkRulesRequest) (<-chan *ListNetworkRulesResponse, <-chan error) {
	responseChan := make(chan *ListNetworkRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNetworkRules(request)
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

// ListNetworkRulesWithCallback invokes the kms.ListNetworkRules API asynchronously
func (client *Client) ListNetworkRulesWithCallback(request *ListNetworkRulesRequest, callback func(response *ListNetworkRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNetworkRulesResponse
		var err error
		defer close(result)
		response, err = client.ListNetworkRules(request)
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

// ListNetworkRulesRequest is the request struct for api ListNetworkRules
type ListNetworkRulesRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListNetworkRulesResponse is the response struct for api ListNetworkRules
type ListNetworkRulesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	NetworkRules NetworkRules `json:"NetworkRules" xml:"NetworkRules"`
}

// CreateListNetworkRulesRequest creates a request to invoke ListNetworkRules API
func CreateListNetworkRulesRequest() (request *ListNetworkRulesRequest) {
	request = &ListNetworkRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListNetworkRules", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNetworkRulesResponse creates a response to parse from ListNetworkRules response
func CreateListNetworkRulesResponse() (response *ListNetworkRulesResponse) {
	response = &ListNetworkRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
