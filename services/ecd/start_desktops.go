package ecd

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

// StartDesktops invokes the ecd.StartDesktops API synchronously
func (client *Client) StartDesktops(request *StartDesktopsRequest) (response *StartDesktopsResponse, err error) {
	response = CreateStartDesktopsResponse()
	err = client.DoAction(request, response)
	return
}

// StartDesktopsWithChan invokes the ecd.StartDesktops API asynchronously
func (client *Client) StartDesktopsWithChan(request *StartDesktopsRequest) (<-chan *StartDesktopsResponse, <-chan error) {
	responseChan := make(chan *StartDesktopsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartDesktops(request)
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

// StartDesktopsWithCallback invokes the ecd.StartDesktops API asynchronously
func (client *Client) StartDesktopsWithCallback(request *StartDesktopsRequest, callback func(response *StartDesktopsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartDesktopsResponse
		var err error
		defer close(result)
		response, err = client.StartDesktops(request)
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

// StartDesktopsRequest is the request struct for api StartDesktops
type StartDesktopsRequest struct {
	*requests.RpcRequest
	DesktopId *[]string `position:"Query" name:"DesktopId"  type:"Repeated"`
}

// StartDesktopsResponse is the response struct for api StartDesktops
type StartDesktopsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartDesktopsRequest creates a request to invoke StartDesktops API
func CreateStartDesktopsRequest() (request *StartDesktopsRequest) {
	request = &StartDesktopsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "StartDesktops", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartDesktopsResponse creates a response to parse from StartDesktops response
func CreateStartDesktopsResponse() (response *StartDesktopsResponse) {
	response = &StartDesktopsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
