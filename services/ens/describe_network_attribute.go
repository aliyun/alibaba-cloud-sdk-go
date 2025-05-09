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

// DescribeNetworkAttribute invokes the ens.DescribeNetworkAttribute API synchronously
func (client *Client) DescribeNetworkAttribute(request *DescribeNetworkAttributeRequest) (response *DescribeNetworkAttributeResponse, err error) {
	response = CreateDescribeNetworkAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkAttributeWithChan invokes the ens.DescribeNetworkAttribute API asynchronously
func (client *Client) DescribeNetworkAttributeWithChan(request *DescribeNetworkAttributeRequest) (<-chan *DescribeNetworkAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkAttribute(request)
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

// DescribeNetworkAttributeWithCallback invokes the ens.DescribeNetworkAttribute API asynchronously
func (client *Client) DescribeNetworkAttributeWithCallback(request *DescribeNetworkAttributeRequest, callback func(response *DescribeNetworkAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkAttribute(request)
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

// DescribeNetworkAttributeRequest is the request struct for api DescribeNetworkAttribute
type DescribeNetworkAttributeRequest struct {
	*requests.RpcRequest
	NetworkId string `position:"Query" name:"NetworkId"`
}

// DescribeNetworkAttributeResponse is the response struct for api DescribeNetworkAttribute
type DescribeNetworkAttributeResponse struct {
	*responses.BaseResponse
	RequestId           string                                        `json:"RequestId" xml:"RequestId"`
	EnsRegionId         string                                        `json:"EnsRegionId" xml:"EnsRegionId"`
	NetworkId           string                                        `json:"NetworkId" xml:"NetworkId"`
	NetworkName         string                                        `json:"NetworkName" xml:"NetworkName"`
	CidrBlock           string                                        `json:"CidrBlock" xml:"CidrBlock"`
	Status              string                                        `json:"Status" xml:"Status"`
	Description         string                                        `json:"Description" xml:"Description"`
	CreatedTime         string                                        `json:"CreatedTime" xml:"CreatedTime"`
	RouterTableId       string                                        `json:"RouterTableId" xml:"RouterTableId"`
	NetworkAclId        string                                        `json:"NetworkAclId" xml:"NetworkAclId"`
	RouteTableId        string                                        `json:"RouteTableId" xml:"RouteTableId"`
	GatewayRouteTableId string                                        `json:"GatewayRouteTableId" xml:"GatewayRouteTableId"`
	VSwitchIds          VSwitchIdsInDescribeNetworkAttribute          `json:"VSwitchIds" xml:"VSwitchIds"`
	InstanceIds         InstanceIdsInDescribeNetworkAttribute         `json:"InstanceIds" xml:"InstanceIds"`
	RouteTableIds       RouteTableIdsInDescribeNetworkAttribute       `json:"RouteTableIds" xml:"RouteTableIds"`
	NetworkInterfaceIds NetworkInterfaceIdsInDescribeNetworkAttribute `json:"NetworkInterfaceIds" xml:"NetworkInterfaceIds"`
	LoadBalancerIds     LoadBalancerIdsInDescribeNetworkAttribute     `json:"LoadBalancerIds" xml:"LoadBalancerIds"`
	NatGatewayIds       NatGatewayIdsInDescribeNetworkAttribute       `json:"NatGatewayIds" xml:"NatGatewayIds"`
	HaVipIds            HaVipIdsInDescribeNetworkAttribute            `json:"HaVipIds" xml:"HaVipIds"`
	CloudResources      CloudResources                                `json:"CloudResources" xml:"CloudResources"`
}

// CreateDescribeNetworkAttributeRequest creates a request to invoke DescribeNetworkAttribute API
func CreateDescribeNetworkAttributeRequest() (request *DescribeNetworkAttributeRequest) {
	request = &DescribeNetworkAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeNetworkAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkAttributeResponse creates a response to parse from DescribeNetworkAttribute response
func CreateDescribeNetworkAttributeResponse() (response *DescribeNetworkAttributeResponse) {
	response = &DescribeNetworkAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
