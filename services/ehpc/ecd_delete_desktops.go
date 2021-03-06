package ehpc

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

// EcdDeleteDesktops invokes the ehpc.EcdDeleteDesktops API synchronously
func (client *Client) EcdDeleteDesktops(request *EcdDeleteDesktopsRequest) (response *EcdDeleteDesktopsResponse, err error) {
	response = CreateEcdDeleteDesktopsResponse()
	err = client.DoAction(request, response)
	return
}

// EcdDeleteDesktopsWithChan invokes the ehpc.EcdDeleteDesktops API asynchronously
func (client *Client) EcdDeleteDesktopsWithChan(request *EcdDeleteDesktopsRequest) (<-chan *EcdDeleteDesktopsResponse, <-chan error) {
	responseChan := make(chan *EcdDeleteDesktopsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EcdDeleteDesktops(request)
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

// EcdDeleteDesktopsWithCallback invokes the ehpc.EcdDeleteDesktops API asynchronously
func (client *Client) EcdDeleteDesktopsWithCallback(request *EcdDeleteDesktopsRequest, callback func(response *EcdDeleteDesktopsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EcdDeleteDesktopsResponse
		var err error
		defer close(result)
		response, err = client.EcdDeleteDesktops(request)
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

// EcdDeleteDesktopsRequest is the request struct for api EcdDeleteDesktops
type EcdDeleteDesktopsRequest struct {
	*requests.RpcRequest
	DesktopId *[]string `position:"Query" name:"DesktopId"  type:"Repeated"`
}

// EcdDeleteDesktopsResponse is the response struct for api EcdDeleteDesktops
type EcdDeleteDesktopsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEcdDeleteDesktopsRequest creates a request to invoke EcdDeleteDesktops API
func CreateEcdDeleteDesktopsRequest() (request *EcdDeleteDesktopsRequest) {
	request = &EcdDeleteDesktopsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "EcdDeleteDesktops", "", "")
	request.Method = requests.GET
	return
}

// CreateEcdDeleteDesktopsResponse creates a response to parse from EcdDeleteDesktops response
func CreateEcdDeleteDesktopsResponse() (response *EcdDeleteDesktopsResponse) {
	response = &EcdDeleteDesktopsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
