package dataworks_public

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

// CallbackExtension invokes the dataworks_public.CallbackExtension API synchronously
func (client *Client) CallbackExtension(request *CallbackExtensionRequest) (response *CallbackExtensionResponse, err error) {
	response = CreateCallbackExtensionResponse()
	err = client.DoAction(request, response)
	return
}

// CallbackExtensionWithChan invokes the dataworks_public.CallbackExtension API asynchronously
func (client *Client) CallbackExtensionWithChan(request *CallbackExtensionRequest) (<-chan *CallbackExtensionResponse, <-chan error) {
	responseChan := make(chan *CallbackExtensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CallbackExtension(request)
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

// CallbackExtensionWithCallback invokes the dataworks_public.CallbackExtension API asynchronously
func (client *Client) CallbackExtensionWithCallback(request *CallbackExtensionRequest, callback func(response *CallbackExtensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CallbackExtensionResponse
		var err error
		defer close(result)
		response, err = client.CallbackExtension(request)
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

// CallbackExtensionRequest is the request struct for api CallbackExtension
type CallbackExtensionRequest struct {
	*requests.RpcRequest
	CheckResult   string `position:"Body" name:"CheckResult"`
	MessageId     string `position:"Body" name:"MessageId"`
	CheckMessage  string `position:"Body" name:"CheckMessage"`
	ExtensionCode string `position:"Body" name:"ExtensionCode"`
}

// CallbackExtensionResponse is the response struct for api CallbackExtension
type CallbackExtensionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateCallbackExtensionRequest creates a request to invoke CallbackExtension API
func CreateCallbackExtensionRequest() (request *CallbackExtensionRequest) {
	request = &CallbackExtensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CallbackExtension", "", "")
	request.Method = requests.POST
	return
}

// CreateCallbackExtensionResponse creates a response to parse from CallbackExtension response
func CreateCallbackExtensionResponse() (response *CallbackExtensionResponse) {
	response = &CallbackExtensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
