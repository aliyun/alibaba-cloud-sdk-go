package ecd

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

// RenewDesktops invokes the ecd.RenewDesktops API synchronously
func (client *Client) RenewDesktops(request *RenewDesktopsRequest) (response *RenewDesktopsResponse, err error) {
	response = CreateRenewDesktopsResponse()
	err = client.DoAction(request, response)
	return
}

// RenewDesktopsWithChan invokes the ecd.RenewDesktops API asynchronously
func (client *Client) RenewDesktopsWithChan(request *RenewDesktopsRequest) (<-chan *RenewDesktopsResponse, <-chan error) {
	responseChan := make(chan *RenewDesktopsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewDesktops(request)
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

// RenewDesktopsWithCallback invokes the ecd.RenewDesktops API asynchronously
func (client *Client) RenewDesktopsWithCallback(request *RenewDesktopsRequest, callback func(response *RenewDesktopsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewDesktopsResponse
		var err error
		defer close(result)
		response, err = client.RenewDesktops(request)
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

// RenewDesktopsRequest is the request struct for api RenewDesktops
type RenewDesktopsRequest struct {
	*requests.RpcRequest
	Period       requests.Integer `position:"Query" name:"Period"`
	AutoPay      requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	PromotionId  string           `position:"Query" name:"PromotionId"`
	PeriodUnit   string           `position:"Query" name:"PeriodUnit"`
	AutoRenew    requests.Boolean `position:"Query" name:"AutoRenew"`
	DesktopId    *[]string        `position:"Query" name:"DesktopId"  type:"Repeated"`
}

// RenewDesktopsResponse is the response struct for api RenewDesktops
type RenewDesktopsResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRenewDesktopsRequest creates a request to invoke RenewDesktops API
func CreateRenewDesktopsRequest() (request *RenewDesktopsRequest) {
	request = &RenewDesktopsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "RenewDesktops", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewDesktopsResponse creates a response to parse from RenewDesktops response
func CreateRenewDesktopsResponse() (response *RenewDesktopsResponse) {
	response = &RenewDesktopsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
