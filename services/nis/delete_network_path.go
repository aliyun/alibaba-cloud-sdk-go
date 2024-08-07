package nis

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

// DeleteNetworkPath invokes the nis.DeleteNetworkPath API synchronously
func (client *Client) DeleteNetworkPath(request *DeleteNetworkPathRequest) (response *DeleteNetworkPathResponse, err error) {
	response = CreateDeleteNetworkPathResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNetworkPathWithChan invokes the nis.DeleteNetworkPath API asynchronously
func (client *Client) DeleteNetworkPathWithChan(request *DeleteNetworkPathRequest) (<-chan *DeleteNetworkPathResponse, <-chan error) {
	responseChan := make(chan *DeleteNetworkPathResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNetworkPath(request)
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

// DeleteNetworkPathWithCallback invokes the nis.DeleteNetworkPath API asynchronously
func (client *Client) DeleteNetworkPathWithCallback(request *DeleteNetworkPathRequest, callback func(response *DeleteNetworkPathResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNetworkPathResponse
		var err error
		defer close(result)
		response, err = client.DeleteNetworkPath(request)
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

// DeleteNetworkPathRequest is the request struct for api DeleteNetworkPath
type DeleteNetworkPathRequest struct {
	*requests.RpcRequest
	DryRun         requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken    string           `position:"Query" name:"ClientToken"`
	NetworkPathIds *[]string        `position:"Query" name:"NetworkPathIds"  type:"Json"`
}

// DeleteNetworkPathResponse is the response struct for api DeleteNetworkPath
type DeleteNetworkPathResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteNetworkPathRequest creates a request to invoke DeleteNetworkPath API
func CreateDeleteNetworkPathRequest() (request *DeleteNetworkPathRequest) {
	request = &DeleteNetworkPathRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nis", "2021-12-16", "DeleteNetworkPath", "networkana", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNetworkPathResponse creates a response to parse from DeleteNetworkPath response
func CreateDeleteNetworkPathResponse() (response *DeleteNetworkPathResponse) {
	response = &DeleteNetworkPathResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
