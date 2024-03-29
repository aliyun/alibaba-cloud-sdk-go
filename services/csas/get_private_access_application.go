package csas

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

// GetPrivateAccessApplication invokes the csas.GetPrivateAccessApplication API synchronously
func (client *Client) GetPrivateAccessApplication(request *GetPrivateAccessApplicationRequest) (response *GetPrivateAccessApplicationResponse, err error) {
	response = CreateGetPrivateAccessApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// GetPrivateAccessApplicationWithChan invokes the csas.GetPrivateAccessApplication API asynchronously
func (client *Client) GetPrivateAccessApplicationWithChan(request *GetPrivateAccessApplicationRequest) (<-chan *GetPrivateAccessApplicationResponse, <-chan error) {
	responseChan := make(chan *GetPrivateAccessApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPrivateAccessApplication(request)
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

// GetPrivateAccessApplicationWithCallback invokes the csas.GetPrivateAccessApplication API asynchronously
func (client *Client) GetPrivateAccessApplicationWithCallback(request *GetPrivateAccessApplicationRequest, callback func(response *GetPrivateAccessApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPrivateAccessApplicationResponse
		var err error
		defer close(result)
		response, err = client.GetPrivateAccessApplication(request)
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

// GetPrivateAccessApplicationRequest is the request struct for api GetPrivateAccessApplication
type GetPrivateAccessApplicationRequest struct {
	*requests.RpcRequest
	SourceIp      string `position:"Query" name:"SourceIp"`
	ApplicationId string `position:"Query" name:"ApplicationId"`
}

// GetPrivateAccessApplicationResponse is the response struct for api GetPrivateAccessApplication
type GetPrivateAccessApplicationResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Application Application `json:"Application" xml:"Application"`
}

// CreateGetPrivateAccessApplicationRequest creates a request to invoke GetPrivateAccessApplication API
func CreateGetPrivateAccessApplicationRequest() (request *GetPrivateAccessApplicationRequest) {
	request = &GetPrivateAccessApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "GetPrivateAccessApplication", "", "")
	request.Method = requests.GET
	return
}

// CreateGetPrivateAccessApplicationResponse creates a response to parse from GetPrivateAccessApplication response
func CreateGetPrivateAccessApplicationResponse() (response *GetPrivateAccessApplicationResponse) {
	response = &GetPrivateAccessApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
