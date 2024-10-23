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

// GetListenerAttribute invokes the gwlb.GetListenerAttribute API synchronously
func (client *Client) GetListenerAttribute(request *GetListenerAttributeRequest) (response *GetListenerAttributeResponse, err error) {
	response = CreateGetListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// GetListenerAttributeWithChan invokes the gwlb.GetListenerAttribute API asynchronously
func (client *Client) GetListenerAttributeWithChan(request *GetListenerAttributeRequest) (<-chan *GetListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *GetListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetListenerAttribute(request)
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

// GetListenerAttributeWithCallback invokes the gwlb.GetListenerAttribute API asynchronously
func (client *Client) GetListenerAttributeWithCallback(request *GetListenerAttributeRequest, callback func(response *GetListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.GetListenerAttribute(request)
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

// GetListenerAttributeRequest is the request struct for api GetListenerAttribute
type GetListenerAttributeRequest struct {
	*requests.RpcRequest
	ListenerId string `position:"Body" name:"ListenerId"`
}

// GetListenerAttributeResponse is the response struct for api GetListenerAttribute
type GetListenerAttributeResponse struct {
	*responses.BaseResponse
	ListenerDescription string     `json:"ListenerDescription" xml:"ListenerDescription"`
	ListenerId          string     `json:"ListenerId" xml:"ListenerId"`
	ListenerStatus      string     `json:"ListenerStatus" xml:"ListenerStatus"`
	LoadBalancerId      string     `json:"LoadBalancerId" xml:"LoadBalancerId"`
	RegionId            string     `json:"RegionId" xml:"RegionId"`
	RequestId           string     `json:"RequestId" xml:"RequestId"`
	ServerGroupId       string     `json:"ServerGroupId" xml:"ServerGroupId"`
	Tags                []TagModel `json:"Tags" xml:"Tags"`
}

// CreateGetListenerAttributeRequest creates a request to invoke GetListenerAttribute API
func CreateGetListenerAttributeRequest() (request *GetListenerAttributeRequest) {
	request = &GetListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Gwlb", "2024-04-15", "GetListenerAttribute", "gwlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetListenerAttributeResponse creates a response to parse from GetListenerAttribute response
func CreateGetListenerAttributeResponse() (response *GetListenerAttributeResponse) {
	response = &GetListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
