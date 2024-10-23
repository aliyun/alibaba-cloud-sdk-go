package ims

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

// ListPredefinedScopes invokes the ims.ListPredefinedScopes API synchronously
func (client *Client) ListPredefinedScopes(request *ListPredefinedScopesRequest) (response *ListPredefinedScopesResponse, err error) {
	response = CreateListPredefinedScopesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPredefinedScopesWithChan invokes the ims.ListPredefinedScopes API asynchronously
func (client *Client) ListPredefinedScopesWithChan(request *ListPredefinedScopesRequest) (<-chan *ListPredefinedScopesResponse, <-chan error) {
	responseChan := make(chan *ListPredefinedScopesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPredefinedScopes(request)
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

// ListPredefinedScopesWithCallback invokes the ims.ListPredefinedScopes API asynchronously
func (client *Client) ListPredefinedScopesWithCallback(request *ListPredefinedScopesRequest, callback func(response *ListPredefinedScopesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPredefinedScopesResponse
		var err error
		defer close(result)
		response, err = client.ListPredefinedScopes(request)
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

// ListPredefinedScopesRequest is the request struct for api ListPredefinedScopes
type ListPredefinedScopesRequest struct {
	*requests.RpcRequest
	AkProxySuffix string `position:"Query" name:"AkProxySuffix"`
	AppType       string `position:"Query" name:"AppType"`
}

// ListPredefinedScopesResponse is the response struct for api ListPredefinedScopes
type ListPredefinedScopesResponse struct {
	*responses.BaseResponse
	RequestId        string                                 `json:"RequestId" xml:"RequestId"`
	PredefinedScopes PredefinedScopesInListPredefinedScopes `json:"PredefinedScopes" xml:"PredefinedScopes"`
}

// CreateListPredefinedScopesRequest creates a request to invoke ListPredefinedScopes API
func CreateListPredefinedScopesRequest() (request *ListPredefinedScopesRequest) {
	request = &ListPredefinedScopesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "ListPredefinedScopes", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPredefinedScopesResponse creates a response to parse from ListPredefinedScopes response
func CreateListPredefinedScopesResponse() (response *ListPredefinedScopesResponse) {
	response = &ListPredefinedScopesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
