package ens

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

// UnAssociateEnsEipAddress invokes the ens.UnAssociateEnsEipAddress API synchronously
func (client *Client) UnAssociateEnsEipAddress(request *UnAssociateEnsEipAddressRequest) (response *UnAssociateEnsEipAddressResponse, err error) {
	response = CreateUnAssociateEnsEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// UnAssociateEnsEipAddressWithChan invokes the ens.UnAssociateEnsEipAddress API asynchronously
func (client *Client) UnAssociateEnsEipAddressWithChan(request *UnAssociateEnsEipAddressRequest) (<-chan *UnAssociateEnsEipAddressResponse, <-chan error) {
	responseChan := make(chan *UnAssociateEnsEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnAssociateEnsEipAddress(request)
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

// UnAssociateEnsEipAddressWithCallback invokes the ens.UnAssociateEnsEipAddress API asynchronously
func (client *Client) UnAssociateEnsEipAddressWithCallback(request *UnAssociateEnsEipAddressRequest, callback func(response *UnAssociateEnsEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnAssociateEnsEipAddressResponse
		var err error
		defer close(result)
		response, err = client.UnAssociateEnsEipAddress(request)
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

// UnAssociateEnsEipAddressRequest is the request struct for api UnAssociateEnsEipAddress
type UnAssociateEnsEipAddressRequest struct {
	*requests.RpcRequest
	AllocationId string           `position:"Query" name:"AllocationId"`
	Force        requests.Boolean `position:"Query" name:"Force"`
}

// UnAssociateEnsEipAddressResponse is the response struct for api UnAssociateEnsEipAddress
type UnAssociateEnsEipAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnAssociateEnsEipAddressRequest creates a request to invoke UnAssociateEnsEipAddress API
func CreateUnAssociateEnsEipAddressRequest() (request *UnAssociateEnsEipAddressRequest) {
	request = &UnAssociateEnsEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "UnAssociateEnsEipAddress", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnAssociateEnsEipAddressResponse creates a response to parse from UnAssociateEnsEipAddress response
func CreateUnAssociateEnsEipAddressResponse() (response *UnAssociateEnsEipAddressResponse) {
	response = &UnAssociateEnsEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
