package lmztest

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

// GrayReleaseDemo invokes the lmztest.GrayReleaseDemo API synchronously
func (client *Client) GrayReleaseDemo(request *GrayReleaseDemoRequest) (response *GrayReleaseDemoResponse, err error) {
	response = CreateGrayReleaseDemoResponse()
	err = client.DoAction(request, response)
	return
}

// GrayReleaseDemoWithChan invokes the lmztest.GrayReleaseDemo API asynchronously
func (client *Client) GrayReleaseDemoWithChan(request *GrayReleaseDemoRequest) (<-chan *GrayReleaseDemoResponse, <-chan error) {
	responseChan := make(chan *GrayReleaseDemoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrayReleaseDemo(request)
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

// GrayReleaseDemoWithCallback invokes the lmztest.GrayReleaseDemo API asynchronously
func (client *Client) GrayReleaseDemoWithCallback(request *GrayReleaseDemoRequest, callback func(response *GrayReleaseDemoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrayReleaseDemoResponse
		var err error
		defer close(result)
		response, err = client.GrayReleaseDemo(request)
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

// GrayReleaseDemoRequest is the request struct for api GrayReleaseDemo
type GrayReleaseDemoRequest struct {
	*requests.RpcRequest
	Id string `position:"Query" name:"Id"`
}

// GrayReleaseDemoResponse is the response struct for api GrayReleaseDemo
type GrayReleaseDemoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Year      string `json:"Year" xml:"Year"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateGrayReleaseDemoRequest creates a request to invoke GrayReleaseDemo API
func CreateGrayReleaseDemoRequest() (request *GrayReleaseDemoRequest) {
	request = &GrayReleaseDemoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LmzTest", "2010-10-11", "GrayReleaseDemo", "", "")
	request.Method = requests.POST
	return
}

// CreateGrayReleaseDemoResponse creates a response to parse from GrayReleaseDemo response
func CreateGrayReleaseDemoResponse() (response *GrayReleaseDemoResponse) {
	response = &GrayReleaseDemoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
