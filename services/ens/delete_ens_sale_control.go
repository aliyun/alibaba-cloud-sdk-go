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

// DeleteEnsSaleControl invokes the ens.DeleteEnsSaleControl API synchronously
func (client *Client) DeleteEnsSaleControl(request *DeleteEnsSaleControlRequest) (response *DeleteEnsSaleControlResponse, err error) {
	response = CreateDeleteEnsSaleControlResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEnsSaleControlWithChan invokes the ens.DeleteEnsSaleControl API asynchronously
func (client *Client) DeleteEnsSaleControlWithChan(request *DeleteEnsSaleControlRequest) (<-chan *DeleteEnsSaleControlResponse, <-chan error) {
	responseChan := make(chan *DeleteEnsSaleControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEnsSaleControl(request)
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

// DeleteEnsSaleControlWithCallback invokes the ens.DeleteEnsSaleControl API asynchronously
func (client *Client) DeleteEnsSaleControlWithCallback(request *DeleteEnsSaleControlRequest, callback func(response *DeleteEnsSaleControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEnsSaleControlResponse
		var err error
		defer close(result)
		response, err = client.DeleteEnsSaleControl(request)
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

// DeleteEnsSaleControlRequest is the request struct for api DeleteEnsSaleControl
type DeleteEnsSaleControlRequest struct {
	*requests.RpcRequest
	SaleControls  *[]DeleteEnsSaleControlSaleControls `position:"Query" name:"SaleControls"  type:"Json"`
	CustomAccount string                              `position:"Query" name:"CustomAccount"`
	CommodityCode string                              `position:"Query" name:"CommodityCode"`
	AliUidAccount string                              `position:"Query" name:"AliUidAccount"`
}

// DeleteEnsSaleControlSaleControls is a repeated param struct in DeleteEnsSaleControlRequest
type DeleteEnsSaleControlSaleControls struct {
	ModuleCode string `name:"ModuleCode"`
	OrderType  string `name:"OrderType"`
}

// DeleteEnsSaleControlResponse is the response struct for api DeleteEnsSaleControl
type DeleteEnsSaleControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEnsSaleControlRequest creates a request to invoke DeleteEnsSaleControl API
func CreateDeleteEnsSaleControlRequest() (request *DeleteEnsSaleControlRequest) {
	request = &DeleteEnsSaleControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DeleteEnsSaleControl", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEnsSaleControlResponse creates a response to parse from DeleteEnsSaleControl response
func CreateDeleteEnsSaleControlResponse() (response *DeleteEnsSaleControlResponse) {
	response = &DeleteEnsSaleControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
