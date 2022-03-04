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

// ListUserGroupsByUserId invokes the quickbi_public.ListUserGroupsByUserId API synchronously
func (client *Client) ListUserGroupsByUserId(request *ListUserGroupsByUserIdRequest) (response *ListUserGroupsByUserIdResponse, err error) {
	response = CreateListUserGroupsByUserIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserGroupsByUserIdWithChan invokes the quickbi_public.ListUserGroupsByUserId API asynchronously
func (client *Client) ListUserGroupsByUserIdWithChan(request *ListUserGroupsByUserIdRequest) (<-chan *ListUserGroupsByUserIdResponse, <-chan error) {
	responseChan := make(chan *ListUserGroupsByUserIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserGroupsByUserId(request)
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

// ListUserGroupsByUserIdWithCallback invokes the quickbi_public.ListUserGroupsByUserId API asynchronously
func (client *Client) ListUserGroupsByUserIdWithCallback(request *ListUserGroupsByUserIdRequest, callback func(response *ListUserGroupsByUserIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserGroupsByUserIdResponse
		var err error
		defer close(result)
		response, err = client.ListUserGroupsByUserId(request)
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

// ListUserGroupsByUserIdRequest is the request struct for api ListUserGroupsByUserId
type ListUserGroupsByUserIdRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
}

// ListUserGroupsByUserIdResponse is the response struct for api ListUserGroupsByUserId
type ListUserGroupsByUserIdResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateListUserGroupsByUserIdRequest creates a request to invoke ListUserGroupsByUserId API
func CreateListUserGroupsByUserIdRequest() (request *ListUserGroupsByUserIdRequest) {
	request = &ListUserGroupsByUserIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListUserGroupsByUserId", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUserGroupsByUserIdResponse creates a response to parse from ListUserGroupsByUserId response
func CreateListUserGroupsByUserIdResponse() (response *ListUserGroupsByUserIdResponse) {
	response = &ListUserGroupsByUserIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
