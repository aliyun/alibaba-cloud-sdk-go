package cbn

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

// RoutePrivateZoneInCenToVpc invokes the cbn.RoutePrivateZoneInCenToVpc API synchronously
func (client *Client) RoutePrivateZoneInCenToVpc(request *RoutePrivateZoneInCenToVpcRequest) (response *RoutePrivateZoneInCenToVpcResponse, err error) {
	response = CreateRoutePrivateZoneInCenToVpcResponse()
	err = client.DoAction(request, response)
	return
}

// RoutePrivateZoneInCenToVpcWithChan invokes the cbn.RoutePrivateZoneInCenToVpc API asynchronously
func (client *Client) RoutePrivateZoneInCenToVpcWithChan(request *RoutePrivateZoneInCenToVpcRequest) (<-chan *RoutePrivateZoneInCenToVpcResponse, <-chan error) {
	responseChan := make(chan *RoutePrivateZoneInCenToVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RoutePrivateZoneInCenToVpc(request)
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

// RoutePrivateZoneInCenToVpcWithCallback invokes the cbn.RoutePrivateZoneInCenToVpc API asynchronously
func (client *Client) RoutePrivateZoneInCenToVpcWithCallback(request *RoutePrivateZoneInCenToVpcRequest, callback func(response *RoutePrivateZoneInCenToVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RoutePrivateZoneInCenToVpcResponse
		var err error
		defer close(result)
		response, err = client.RoutePrivateZoneInCenToVpc(request)
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

// RoutePrivateZoneInCenToVpcRequest is the request struct for api RoutePrivateZoneInCenToVpc
type RoutePrivateZoneInCenToVpcRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	AccessRegionId       string           `position:"Query" name:"AccessRegionId"`
	HostRegionId         string           `position:"Query" name:"HostRegionId"`
	HostVpcId            string           `position:"Query" name:"HostVpcId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
}

// RoutePrivateZoneInCenToVpcResponse is the response struct for api RoutePrivateZoneInCenToVpc
type RoutePrivateZoneInCenToVpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRoutePrivateZoneInCenToVpcRequest creates a request to invoke RoutePrivateZoneInCenToVpc API
func CreateRoutePrivateZoneInCenToVpcRequest() (request *RoutePrivateZoneInCenToVpcRequest) {
	request = &RoutePrivateZoneInCenToVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "RoutePrivateZoneInCenToVpc", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRoutePrivateZoneInCenToVpcResponse creates a response to parse from RoutePrivateZoneInCenToVpc response
func CreateRoutePrivateZoneInCenToVpcResponse() (response *RoutePrivateZoneInCenToVpcResponse) {
	response = &RoutePrivateZoneInCenToVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
