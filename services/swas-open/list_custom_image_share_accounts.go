package swas_open

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

// ListCustomImageShareAccounts invokes the swas_open.ListCustomImageShareAccounts API synchronously
func (client *Client) ListCustomImageShareAccounts(request *ListCustomImageShareAccountsRequest) (response *ListCustomImageShareAccountsResponse, err error) {
	response = CreateListCustomImageShareAccountsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCustomImageShareAccountsWithChan invokes the swas_open.ListCustomImageShareAccounts API asynchronously
func (client *Client) ListCustomImageShareAccountsWithChan(request *ListCustomImageShareAccountsRequest) (<-chan *ListCustomImageShareAccountsResponse, <-chan error) {
	responseChan := make(chan *ListCustomImageShareAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCustomImageShareAccounts(request)
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

// ListCustomImageShareAccountsWithCallback invokes the swas_open.ListCustomImageShareAccounts API asynchronously
func (client *Client) ListCustomImageShareAccountsWithCallback(request *ListCustomImageShareAccountsRequest, callback func(response *ListCustomImageShareAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCustomImageShareAccountsResponse
		var err error
		defer close(result)
		response, err = client.ListCustomImageShareAccounts(request)
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

// ListCustomImageShareAccountsRequest is the request struct for api ListCustomImageShareAccounts
type ListCustomImageShareAccountsRequest struct {
	*requests.RpcRequest
	ImageId     string           `position:"Query" name:"ImageId"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// ListCustomImageShareAccountsResponse is the response struct for api ListCustomImageShareAccounts
type ListCustomImageShareAccountsResponse struct {
	*responses.BaseResponse
	TotalCount      int              `json:"TotalCount" xml:"TotalCount"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	PageSize        int              `json:"PageSize" xml:"PageSize"`
	PageNumber      int              `json:"PageNumber" xml:"PageNumber"`
	ImageShareUsers []ImageShareUser `json:"ImageShareUsers" xml:"ImageShareUsers"`
}

// CreateListCustomImageShareAccountsRequest creates a request to invoke ListCustomImageShareAccounts API
func CreateListCustomImageShareAccountsRequest() (request *ListCustomImageShareAccountsRequest) {
	request = &ListCustomImageShareAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ListCustomImageShareAccounts", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCustomImageShareAccountsResponse creates a response to parse from ListCustomImageShareAccounts response
func CreateListCustomImageShareAccountsResponse() (response *ListCustomImageShareAccountsResponse) {
	response = &ListCustomImageShareAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
