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

// DescribeVSwitchAttributes invokes the ens.DescribeVSwitchAttributes API synchronously
func (client *Client) DescribeVSwitchAttributes(request *DescribeVSwitchAttributesRequest) (response *DescribeVSwitchAttributesResponse, err error) {
	response = CreateDescribeVSwitchAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVSwitchAttributesWithChan invokes the ens.DescribeVSwitchAttributes API asynchronously
func (client *Client) DescribeVSwitchAttributesWithChan(request *DescribeVSwitchAttributesRequest) (<-chan *DescribeVSwitchAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeVSwitchAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVSwitchAttributes(request)
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

// DescribeVSwitchAttributesWithCallback invokes the ens.DescribeVSwitchAttributes API asynchronously
func (client *Client) DescribeVSwitchAttributesWithCallback(request *DescribeVSwitchAttributesRequest, callback func(response *DescribeVSwitchAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVSwitchAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVSwitchAttributes(request)
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

// DescribeVSwitchAttributesRequest is the request struct for api DescribeVSwitchAttributes
type DescribeVSwitchAttributesRequest struct {
	*requests.RpcRequest
	VSwitchId string `position:"Query" name:"VSwitchId"`
}

// DescribeVSwitchAttributesResponse is the response struct for api DescribeVSwitchAttributes
type DescribeVSwitchAttributesResponse struct {
	*responses.BaseResponse
	RequestId               string                                         `json:"RequestId" xml:"RequestId"`
	EnsRegionId             string                                         `json:"EnsRegionId" xml:"EnsRegionId"`
	NetworkId               string                                         `json:"NetworkId" xml:"NetworkId"`
	VSwitchName             string                                         `json:"VSwitchName" xml:"VSwitchName"`
	CidrBlock               string                                         `json:"CidrBlock" xml:"CidrBlock"`
	Status                  string                                         `json:"Status" xml:"Status"`
	Description             string                                         `json:"Description" xml:"Description"`
	CreatedTime             string                                         `json:"CreatedTime" xml:"CreatedTime"`
	VSwitchId               string                                         `json:"VSwitchId" xml:"VSwitchId"`
	AvailableIpAddressCount int64                                          `json:"AvailableIpAddressCount" xml:"AvailableIpAddressCount"`
	InstanceIds             InstanceIdsInDescribeVSwitchAttributes         `json:"InstanceIds" xml:"InstanceIds"`
	NetworkInterfaceIds     NetworkInterfaceIdsInDescribeVSwitchAttributes `json:"NetworkInterfaceIds" xml:"NetworkInterfaceIds"`
	LoadBalancerIds         LoadBalancerIdsInDescribeVSwitchAttributes     `json:"LoadBalancerIds" xml:"LoadBalancerIds"`
	NatGatewayIds           NatGatewayIdsInDescribeVSwitchAttributes       `json:"NatGatewayIds" xml:"NatGatewayIds"`
	HaVipIds                HaVipIdsInDescribeVSwitchAttributes            `json:"HaVipIds" xml:"HaVipIds"`
}

// CreateDescribeVSwitchAttributesRequest creates a request to invoke DescribeVSwitchAttributes API
func CreateDescribeVSwitchAttributesRequest() (request *DescribeVSwitchAttributesRequest) {
	request = &DescribeVSwitchAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeVSwitchAttributes", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVSwitchAttributesResponse creates a response to parse from DescribeVSwitchAttributes response
func CreateDescribeVSwitchAttributesResponse() (response *DescribeVSwitchAttributesResponse) {
	response = &DescribeVSwitchAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
