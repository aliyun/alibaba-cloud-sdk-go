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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SwitchSynchronizationEndpoint invokes the dts.SwitchSynchronizationEndpoint API synchronously
// api document: https://help.aliyun.com/api/dts/switchsynchronizationendpoint.html
func (client *Client) SwitchSynchronizationEndpoint(request *SwitchSynchronizationEndpointRequest) (response *SwitchSynchronizationEndpointResponse, err error) {
	response = CreateSwitchSynchronizationEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchSynchronizationEndpointWithChan invokes the dts.SwitchSynchronizationEndpoint API asynchronously
// api document: https://help.aliyun.com/api/dts/switchsynchronizationendpoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchSynchronizationEndpointWithChan(request *SwitchSynchronizationEndpointRequest) (<-chan *SwitchSynchronizationEndpointResponse, <-chan error) {
	responseChan := make(chan *SwitchSynchronizationEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchSynchronizationEndpoint(request)
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

// SwitchSynchronizationEndpointWithCallback invokes the dts.SwitchSynchronizationEndpoint API asynchronously
// api document: https://help.aliyun.com/api/dts/switchsynchronizationendpoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchSynchronizationEndpointWithCallback(request *SwitchSynchronizationEndpointRequest, callback func(response *SwitchSynchronizationEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchSynchronizationEndpointResponse
		var err error
		defer close(result)
		response, err = client.SwitchSynchronizationEndpoint(request)
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

// SwitchSynchronizationEndpointRequest is the request struct for api SwitchSynchronizationEndpoint
type SwitchSynchronizationEndpointRequest struct {
	*requests.RpcRequest
	SynchronizationJobId     string                                `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection string                                `position:"Query" name:"SynchronizationDirection"`
	OwnerId                  string                                `position:"Query" name:"OwnerId"`
	Endpoint                 SwitchSynchronizationEndpointEndpoint `position:"Query" name:"Endpoint" type:"Struct"`
}

type SwitchSynchronizationEndpointEndpoint struct {
	Type         string `name:"Type"`
	InstanceType string `name:"InstanceType"`
	InstanceId   string `name:"InstanceId"`
	IP           string `name:"IP"`
	Port         string `name:"Port"`
}

// SwitchSynchronizationEndpointResponse is the response struct for api SwitchSynchronizationEndpoint
type SwitchSynchronizationEndpointResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
}

// CreateSwitchSynchronizationEndpointRequest creates a request to invoke SwitchSynchronizationEndpoint API
func CreateSwitchSynchronizationEndpointRequest() (request *SwitchSynchronizationEndpointRequest) {
	request = &SwitchSynchronizationEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "SwitchSynchronizationEndpoint", "dts", "openAPI")
	return
}

// CreateSwitchSynchronizationEndpointResponse creates a response to parse from SwitchSynchronizationEndpoint response
func CreateSwitchSynchronizationEndpointResponse() (response *SwitchSynchronizationEndpointResponse) {
	response = &SwitchSynchronizationEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
