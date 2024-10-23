package gwlb

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

// GetListenerHealthStatus invokes the gwlb.GetListenerHealthStatus API synchronously
func (client *Client) GetListenerHealthStatus(request *GetListenerHealthStatusRequest) (response *GetListenerHealthStatusResponse, err error) {
	response = CreateGetListenerHealthStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetListenerHealthStatusWithChan invokes the gwlb.GetListenerHealthStatus API asynchronously
func (client *Client) GetListenerHealthStatusWithChan(request *GetListenerHealthStatusRequest) (<-chan *GetListenerHealthStatusResponse, <-chan error) {
	responseChan := make(chan *GetListenerHealthStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetListenerHealthStatus(request)
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

// GetListenerHealthStatusWithCallback invokes the gwlb.GetListenerHealthStatus API asynchronously
func (client *Client) GetListenerHealthStatusWithCallback(request *GetListenerHealthStatusRequest, callback func(response *GetListenerHealthStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetListenerHealthStatusResponse
		var err error
		defer close(result)
		response, err = client.GetListenerHealthStatus(request)
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

// GetListenerHealthStatusRequest is the request struct for api GetListenerHealthStatus
type GetListenerHealthStatusRequest struct {
	*requests.RpcRequest
	Skip       requests.Integer                 `position:"Body" name:"Skip"`
	ListenerId string                           `position:"Body" name:"ListenerId"`
	NextToken  string                           `position:"Body" name:"NextToken"`
	Filter     *[]GetListenerHealthStatusFilter `position:"Body" name:"Filter"  type:"Repeated"`
	MaxResults requests.Integer                 `position:"Body" name:"MaxResults"`
}

// GetListenerHealthStatusFilter is a repeated param struct in GetListenerHealthStatusRequest
type GetListenerHealthStatusFilter struct {
	Name   string    `name:"Name"`
	Values *[]string `name:"Values" type:"Repeated"`
}

// GetListenerHealthStatusResponse is the response struct for api GetListenerHealthStatus
type GetListenerHealthStatusResponse struct {
	*responses.BaseResponse
	MaxResults           int                        `json:"MaxResults" xml:"MaxResults"`
	NextToken            string                     `json:"NextToken" xml:"NextToken"`
	RequestId            string                     `json:"RequestId" xml:"RequestId"`
	TotalCount           int                        `json:"TotalCount" xml:"TotalCount"`
	ListenerHealthStatus []ListenerHealthStatusItem `json:"ListenerHealthStatus" xml:"ListenerHealthStatus"`
}

// CreateGetListenerHealthStatusRequest creates a request to invoke GetListenerHealthStatus API
func CreateGetListenerHealthStatusRequest() (request *GetListenerHealthStatusRequest) {
	request = &GetListenerHealthStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Gwlb", "2024-04-15", "GetListenerHealthStatus", "gwlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetListenerHealthStatusResponse creates a response to parse from GetListenerHealthStatus response
func CreateGetListenerHealthStatusResponse() (response *GetListenerHealthStatusResponse) {
	response = &GetListenerHealthStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
