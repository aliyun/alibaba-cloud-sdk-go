package quickbi_public

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

// ListPortalMenus invokes the quickbi_public.ListPortalMenus API synchronously
func (client *Client) ListPortalMenus(request *ListPortalMenusRequest) (response *ListPortalMenusResponse, err error) {
	response = CreateListPortalMenusResponse()
	err = client.DoAction(request, response)
	return
}

// ListPortalMenusWithChan invokes the quickbi_public.ListPortalMenus API asynchronously
func (client *Client) ListPortalMenusWithChan(request *ListPortalMenusRequest) (<-chan *ListPortalMenusResponse, <-chan error) {
	responseChan := make(chan *ListPortalMenusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPortalMenus(request)
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

// ListPortalMenusWithCallback invokes the quickbi_public.ListPortalMenus API asynchronously
func (client *Client) ListPortalMenusWithCallback(request *ListPortalMenusRequest, callback func(response *ListPortalMenusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPortalMenusResponse
		var err error
		defer close(result)
		response, err = client.ListPortalMenus(request)
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

// ListPortalMenusRequest is the request struct for api ListPortalMenus
type ListPortalMenusRequest struct {
	*requests.RpcRequest
	DataPortalId string `position:"Query" name:"DataPortalId"`
	AccessPoint  string `position:"Query" name:"AccessPoint"`
	UserId       string `position:"Query" name:"UserId"`
	SignType     string `position:"Query" name:"SignType"`
}

// ListPortalMenusResponse is the response struct for api ListPortalMenus
type ListPortalMenusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateListPortalMenusRequest creates a request to invoke ListPortalMenus API
func CreateListPortalMenusRequest() (request *ListPortalMenusRequest) {
	request = &ListPortalMenusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListPortalMenus", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPortalMenusResponse creates a response to parse from ListPortalMenus response
func CreateListPortalMenusResponse() (response *ListPortalMenusResponse) {
	response = &ListPortalMenusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
