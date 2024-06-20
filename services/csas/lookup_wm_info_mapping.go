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

// LookupWmInfoMapping invokes the csas.LookupWmInfoMapping API synchronously
func (client *Client) LookupWmInfoMapping(request *LookupWmInfoMappingRequest) (response *LookupWmInfoMappingResponse, err error) {
	response = CreateLookupWmInfoMappingResponse()
	err = client.DoAction(request, response)
	return
}

// LookupWmInfoMappingWithChan invokes the csas.LookupWmInfoMapping API asynchronously
func (client *Client) LookupWmInfoMappingWithChan(request *LookupWmInfoMappingRequest) (<-chan *LookupWmInfoMappingResponse, <-chan error) {
	responseChan := make(chan *LookupWmInfoMappingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LookupWmInfoMapping(request)
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

// LookupWmInfoMappingWithCallback invokes the csas.LookupWmInfoMapping API asynchronously
func (client *Client) LookupWmInfoMappingWithCallback(request *LookupWmInfoMappingRequest, callback func(response *LookupWmInfoMappingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LookupWmInfoMappingResponse
		var err error
		defer close(result)
		response, err = client.LookupWmInfoMapping(request)
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

// LookupWmInfoMappingRequest is the request struct for api LookupWmInfoMapping
type LookupWmInfoMappingRequest struct {
	*requests.RpcRequest
	WmInfoUint string           `position:"Query" name:"WmInfoUint"`
	WmInfoSize requests.Integer `position:"Query" name:"WmInfoSize"`
	WmType     string           `position:"Query" name:"WmType"`
}

// LookupWmInfoMappingResponse is the response struct for api LookupWmInfoMapping
type LookupWmInfoMappingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateLookupWmInfoMappingRequest creates a request to invoke LookupWmInfoMapping API
func CreateLookupWmInfoMappingRequest() (request *LookupWmInfoMappingRequest) {
	request = &LookupWmInfoMappingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "LookupWmInfoMapping", "", "")
	request.Method = requests.GET
	return
}

// CreateLookupWmInfoMappingResponse creates a response to parse from LookupWmInfoMapping response
func CreateLookupWmInfoMappingResponse() (response *LookupWmInfoMappingResponse) {
	response = &LookupWmInfoMappingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
