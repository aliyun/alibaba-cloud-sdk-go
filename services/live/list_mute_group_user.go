package live

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

// ListMuteGroupUser invokes the live.ListMuteGroupUser API synchronously
func (client *Client) ListMuteGroupUser(request *ListMuteGroupUserRequest) (response *ListMuteGroupUserResponse, err error) {
	response = CreateListMuteGroupUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListMuteGroupUserWithChan invokes the live.ListMuteGroupUser API asynchronously
func (client *Client) ListMuteGroupUserWithChan(request *ListMuteGroupUserRequest) (<-chan *ListMuteGroupUserResponse, <-chan error) {
	responseChan := make(chan *ListMuteGroupUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMuteGroupUser(request)
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

// ListMuteGroupUserWithCallback invokes the live.ListMuteGroupUser API asynchronously
func (client *Client) ListMuteGroupUserWithCallback(request *ListMuteGroupUserRequest, callback func(response *ListMuteGroupUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMuteGroupUserResponse
		var err error
		defer close(result)
		response, err = client.ListMuteGroupUser(request)
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

// ListMuteGroupUserRequest is the request struct for api ListMuteGroupUser
type ListMuteGroupUserRequest struct {
	*requests.RpcRequest
	PageNum        requests.Integer `position:"Body" name:"PageNum"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	OperatorUserId string           `position:"Body" name:"OperatorUserId"`
	GroupId        string           `position:"Body" name:"GroupId"`
	AppId          string           `position:"Body" name:"AppId"`
}

// ListMuteGroupUserResponse is the response struct for api ListMuteGroupUser
type ListMuteGroupUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateListMuteGroupUserRequest creates a request to invoke ListMuteGroupUser API
func CreateListMuteGroupUserRequest() (request *ListMuteGroupUserRequest) {
	request = &ListMuteGroupUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ListMuteGroupUser", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListMuteGroupUserResponse creates a response to parse from ListMuteGroupUser response
func CreateListMuteGroupUserResponse() (response *ListMuteGroupUserResponse) {
	response = &ListMuteGroupUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
