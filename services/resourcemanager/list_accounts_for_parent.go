package resourcemanager

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

// ListAccountsForParent invokes the resourcemanager.ListAccountsForParent API synchronously
func (client *Client) ListAccountsForParent(request *ListAccountsForParentRequest) (response *ListAccountsForParentResponse, err error) {
	response = CreateListAccountsForParentResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccountsForParentWithChan invokes the resourcemanager.ListAccountsForParent API asynchronously
func (client *Client) ListAccountsForParentWithChan(request *ListAccountsForParentRequest) (<-chan *ListAccountsForParentResponse, <-chan error) {
	responseChan := make(chan *ListAccountsForParentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccountsForParent(request)
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

// ListAccountsForParentWithCallback invokes the resourcemanager.ListAccountsForParent API asynchronously
func (client *Client) ListAccountsForParentWithCallback(request *ListAccountsForParentRequest, callback func(response *ListAccountsForParentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccountsForParentResponse
		var err error
		defer close(result)
		response, err = client.ListAccountsForParent(request)
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

// ListAccountsForParentRequest is the request struct for api ListAccountsForParent
type ListAccountsForParentRequest struct {
	*requests.RpcRequest
	PageNumber     requests.Integer            `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer            `position:"Query" name:"PageSize"`
	Tag            *[]ListAccountsForParentTag `position:"Query" name:"Tag"  type:"Repeated"`
	QueryKeyword   string                      `position:"Query" name:"QueryKeyword"`
	ParentFolderId string                      `position:"Query" name:"ParentFolderId"`
	IncludeTags    requests.Boolean            `position:"Query" name:"IncludeTags"`
}

// ListAccountsForParentTag is a repeated param struct in ListAccountsForParentRequest
type ListAccountsForParentTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListAccountsForParentResponse is the response struct for api ListAccountsForParent
type ListAccountsForParentResponse struct {
	*responses.BaseResponse
	TotalCount int                             `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                          `json:"RequestId" xml:"RequestId"`
	PageSize   int                             `json:"PageSize" xml:"PageSize"`
	PageNumber int                             `json:"PageNumber" xml:"PageNumber"`
	Accounts   AccountsInListAccountsForParent `json:"Accounts" xml:"Accounts"`
}

// CreateListAccountsForParentRequest creates a request to invoke ListAccountsForParent API
func CreateListAccountsForParentRequest() (request *ListAccountsForParentRequest) {
	request = &ListAccountsForParentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListAccountsForParent", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccountsForParentResponse creates a response to parse from ListAccountsForParent response
func CreateListAccountsForParentResponse() (response *ListAccountsForParentResponse) {
	response = &ListAccountsForParentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
