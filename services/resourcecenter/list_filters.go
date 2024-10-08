package resourcecenter

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

// ListFilters invokes the resourcecenter.ListFilters API synchronously
func (client *Client) ListFilters(request *ListFiltersRequest) (response *ListFiltersResponse, err error) {
	response = CreateListFiltersResponse()
	err = client.DoAction(request, response)
	return
}

// ListFiltersWithChan invokes the resourcecenter.ListFilters API asynchronously
func (client *Client) ListFiltersWithChan(request *ListFiltersRequest) (<-chan *ListFiltersResponse, <-chan error) {
	responseChan := make(chan *ListFiltersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFilters(request)
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

// ListFiltersWithCallback invokes the resourcecenter.ListFilters API asynchronously
func (client *Client) ListFiltersWithCallback(request *ListFiltersRequest, callback func(response *ListFiltersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFiltersResponse
		var err error
		defer close(result)
		response, err = client.ListFilters(request)
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

// ListFiltersRequest is the request struct for api ListFilters
type ListFiltersRequest struct {
	*requests.RpcRequest
}

// ListFiltersResponse is the response struct for api ListFilters
type ListFiltersResponse struct {
	*responses.BaseResponse
	RequestId         string   `json:"RequestId" xml:"RequestId"`
	DefaultFilterName string   `json:"DefaultFilterName" xml:"DefaultFilterName"`
	Filters           []Filter `json:"Filters" xml:"Filters"`
}

// CreateListFiltersRequest creates a request to invoke ListFilters API
func CreateListFiltersRequest() (request *ListFiltersRequest) {
	request = &ListFiltersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "ListFilters", "", "")
	request.Method = requests.POST
	return
}

// CreateListFiltersResponse creates a response to parse from ListFilters response
func CreateListFiltersResponse() (response *ListFiltersResponse) {
	response = &ListFiltersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
