package live

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

// DescribeLiveRealtimeLogAuthorized invokes the live.DescribeLiveRealtimeLogAuthorized API synchronously
func (client *Client) DescribeLiveRealtimeLogAuthorized(request *DescribeLiveRealtimeLogAuthorizedRequest) (response *DescribeLiveRealtimeLogAuthorizedResponse, err error) {
	response = CreateDescribeLiveRealtimeLogAuthorizedResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveRealtimeLogAuthorizedWithChan invokes the live.DescribeLiveRealtimeLogAuthorized API asynchronously
func (client *Client) DescribeLiveRealtimeLogAuthorizedWithChan(request *DescribeLiveRealtimeLogAuthorizedRequest) (<-chan *DescribeLiveRealtimeLogAuthorizedResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveRealtimeLogAuthorizedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveRealtimeLogAuthorized(request)
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

// DescribeLiveRealtimeLogAuthorizedWithCallback invokes the live.DescribeLiveRealtimeLogAuthorized API asynchronously
func (client *Client) DescribeLiveRealtimeLogAuthorizedWithCallback(request *DescribeLiveRealtimeLogAuthorizedRequest, callback func(response *DescribeLiveRealtimeLogAuthorizedResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveRealtimeLogAuthorizedResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveRealtimeLogAuthorized(request)
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

// DescribeLiveRealtimeLogAuthorizedRequest is the request struct for api DescribeLiveRealtimeLogAuthorized
type DescribeLiveRealtimeLogAuthorizedRequest struct {
	*requests.RpcRequest
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	LiveOpenapiReserve string           `position:"Query" name:"LiveOpenapiReserve"`
}

// DescribeLiveRealtimeLogAuthorizedResponse is the response struct for api DescribeLiveRealtimeLogAuthorized
type DescribeLiveRealtimeLogAuthorizedResponse struct {
	*responses.BaseResponse
	AuthorizedStatus string `json:"AuthorizedStatus" xml:"AuthorizedStatus"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeLiveRealtimeLogAuthorizedRequest creates a request to invoke DescribeLiveRealtimeLogAuthorized API
func CreateDescribeLiveRealtimeLogAuthorizedRequest() (request *DescribeLiveRealtimeLogAuthorizedRequest) {
	request = &DescribeLiveRealtimeLogAuthorizedRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRealtimeLogAuthorized", "live", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeLiveRealtimeLogAuthorizedResponse creates a response to parse from DescribeLiveRealtimeLogAuthorized response
func CreateDescribeLiveRealtimeLogAuthorizedResponse() (response *DescribeLiveRealtimeLogAuthorizedResponse) {
	response = &DescribeLiveRealtimeLogAuthorizedResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
