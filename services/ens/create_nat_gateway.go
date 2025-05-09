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

// CreateNatGateway invokes the ens.CreateNatGateway API synchronously
func (client *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
	response = CreateCreateNatGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNatGatewayWithChan invokes the ens.CreateNatGateway API asynchronously
func (client *Client) CreateNatGatewayWithChan(request *CreateNatGatewayRequest) (<-chan *CreateNatGatewayResponse, <-chan error) {
	responseChan := make(chan *CreateNatGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNatGateway(request)
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

// CreateNatGatewayWithCallback invokes the ens.CreateNatGateway API asynchronously
func (client *Client) CreateNatGatewayWithCallback(request *CreateNatGatewayRequest, callback func(response *CreateNatGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNatGatewayResponse
		var err error
		defer close(result)
		response, err = client.CreateNatGateway(request)
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

// CreateNatGatewayRequest is the request struct for api CreateNatGateway
type CreateNatGatewayRequest struct {
	*requests.RpcRequest
	EnsRegionId  string                 `position:"Query" name:"EnsRegionId"`
	InstanceType string                 `position:"Query" name:"InstanceType"`
	Tag          *[]CreateNatGatewayTag `position:"Query" name:"Tag"  type:"Repeated"`
	VSwitchId    string                 `position:"Query" name:"VSwitchId"`
	Name         string                 `position:"Query" name:"Name"`
	NetworkId    string                 `position:"Query" name:"NetworkId"`
}

// CreateNatGatewayTag is a repeated param struct in CreateNatGatewayRequest
type CreateNatGatewayTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateNatGatewayResponse is the response struct for api CreateNatGateway
type CreateNatGatewayResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	NatGatewayId string `json:"NatGatewayId" xml:"NatGatewayId"`
}

// CreateCreateNatGatewayRequest creates a request to invoke CreateNatGateway API
func CreateCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
	request = &CreateNatGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateNatGateway", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNatGatewayResponse creates a response to parse from CreateNatGateway response
func CreateCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
	response = &CreateNatGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
