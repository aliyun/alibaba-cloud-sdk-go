package iot

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

// AsyncRRpc invokes the iot.AsyncRRpc API synchronously
func (client *Client) AsyncRRpc(request *AsyncRRpcRequest) (response *AsyncRRpcResponse, err error) {
	response = CreateAsyncRRpcResponse()
	err = client.DoAction(request, response)
	return
}

// AsyncRRpcWithChan invokes the iot.AsyncRRpc API asynchronously
func (client *Client) AsyncRRpcWithChan(request *AsyncRRpcRequest) (<-chan *AsyncRRpcResponse, <-chan error) {
	responseChan := make(chan *AsyncRRpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AsyncRRpc(request)
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

// AsyncRRpcWithCallback invokes the iot.AsyncRRpc API asynchronously
func (client *Client) AsyncRRpcWithCallback(request *AsyncRRpcRequest, callback func(response *AsyncRRpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AsyncRRpcResponse
		var err error
		defer close(result)
		response, err = client.AsyncRRpc(request)
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

// AsyncRRpcRequest is the request struct for api AsyncRRpc
type AsyncRRpcRequest struct {
	*requests.RpcRequest
	MessageContent string `position:"Body" name:"MessageContent"`
	IotInstanceId  string `position:"Query" name:"IotInstanceId"`
	ExtInfo        string `position:"Query" name:"ExtInfo"`
	TopicFullName  string `position:"Query" name:"TopicFullName"`
	ProductKey     string `position:"Query" name:"ProductKey"`
	ApiProduct     string `position:"Body" name:"ApiProduct"`
	ApiRevision    string `position:"Body" name:"ApiRevision"`
	DeviceName     string `position:"Query" name:"DeviceName"`
}

// AsyncRRpcResponse is the response struct for api AsyncRRpc
type AsyncRRpcResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	MessageId    int64  `json:"MessageId" xml:"MessageId"`
}

// CreateAsyncRRpcRequest creates a request to invoke AsyncRRpc API
func CreateAsyncRRpcRequest() (request *AsyncRRpcRequest) {
	request = &AsyncRRpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "AsyncRRpc", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAsyncRRpcResponse creates a response to parse from AsyncRRpc response
func CreateAsyncRRpcResponse() (response *AsyncRRpcResponse) {
	response = &AsyncRRpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
