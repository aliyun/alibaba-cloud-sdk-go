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

// CreateNetworkPath invokes the nis.CreateNetworkPath API synchronously
func (client *Client) CreateNetworkPath(request *CreateNetworkPathRequest) (response *CreateNetworkPathResponse, err error) {
	response = CreateCreateNetworkPathResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNetworkPathWithChan invokes the nis.CreateNetworkPath API asynchronously
func (client *Client) CreateNetworkPathWithChan(request *CreateNetworkPathRequest) (<-chan *CreateNetworkPathResponse, <-chan error) {
	responseChan := make(chan *CreateNetworkPathResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNetworkPath(request)
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

// CreateNetworkPathWithCallback invokes the nis.CreateNetworkPath API asynchronously
func (client *Client) CreateNetworkPathWithCallback(request *CreateNetworkPathRequest, callback func(response *CreateNetworkPathResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNetworkPathResponse
		var err error
		defer close(result)
		response, err = client.CreateNetworkPath(request)
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

// CreateNetworkPathRequest is the request struct for api CreateNetworkPath
type CreateNetworkPathRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	TargetId               string                        `position:"Query" name:"TargetId"`
	TargetType             string                        `position:"Query" name:"TargetType"`
	ClientToken            string                        `position:"Query" name:"ClientToken"`
	SystemTag              *[]CreateNetworkPathSystemTag `position:"Query" name:"SystemTag"  type:"Repeated"`
	TargetIpAddress        string                        `position:"Query" name:"TargetIpAddress"`
	NetworkPathName        string                        `position:"Query" name:"NetworkPathName"`
	SourcePort             requests.Integer              `position:"Query" name:"SourcePort"`
	ResourceGroupId        string                        `position:"Query" name:"ResourceGroupId"`
	Protocol               string                        `position:"Query" name:"Protocol"`
	SourceAliUid           requests.Integer              `position:"Query" name:"SourceAliUid"`
	SourceType             string                        `position:"Query" name:"SourceType"`
	Tag                    *[]CreateNetworkPathTag       `position:"Query" name:"Tag"  type:"Repeated"`
	TargetPort             requests.Integer              `position:"Query" name:"TargetPort"`
	SourceId               string                        `position:"Query" name:"SourceId"`
	SourceIpAddress        string                        `position:"Query" name:"SourceIpAddress"`
	DryRun                 requests.Boolean              `position:"Query" name:"DryRun"`
	UseMultiAccount        requests.Boolean              `position:"Query" name:"UseMultiAccount"`
	NetworkPathDescription string                        `position:"Query" name:"NetworkPathDescription"`
	TargetAliUid           requests.Integer              `position:"Query" name:"TargetAliUid"`
}

// CreateNetworkPathSystemTag is a repeated param struct in CreateNetworkPathRequest
type CreateNetworkPathSystemTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Scope string `name:"Scope"`
}

// CreateNetworkPathTag is a repeated param struct in CreateNetworkPathRequest
type CreateNetworkPathTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateNetworkPathResponse is the response struct for api CreateNetworkPath
type CreateNetworkPathResponse struct {
	*responses.BaseResponse
	NetworkPathId string `json:"NetworkPathId" xml:"NetworkPathId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateNetworkPathRequest creates a request to invoke CreateNetworkPath API
func CreateCreateNetworkPathRequest() (request *CreateNetworkPathRequest) {
	request = &CreateNetworkPathRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nis", "2021-12-16", "CreateNetworkPath", "networkana", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNetworkPathResponse creates a response to parse from CreateNetworkPath response
func CreateCreateNetworkPathResponse() (response *CreateNetworkPathResponse) {
	response = &CreateNetworkPathResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
