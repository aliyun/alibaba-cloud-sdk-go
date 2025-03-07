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

// ListLoadBalancers invokes the gwlb.ListLoadBalancers API synchronously
func (client *Client) ListLoadBalancers(request *ListLoadBalancersRequest) (response *ListLoadBalancersResponse, err error) {
	response = CreateListLoadBalancersResponse()
	err = client.DoAction(request, response)
	return
}

// ListLoadBalancersWithChan invokes the gwlb.ListLoadBalancers API asynchronously
func (client *Client) ListLoadBalancersWithChan(request *ListLoadBalancersRequest) (<-chan *ListLoadBalancersResponse, <-chan error) {
	responseChan := make(chan *ListLoadBalancersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLoadBalancers(request)
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

// ListLoadBalancersWithCallback invokes the gwlb.ListLoadBalancers API asynchronously
func (client *Client) ListLoadBalancersWithCallback(request *ListLoadBalancersRequest, callback func(response *ListLoadBalancersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLoadBalancersResponse
		var err error
		defer close(result)
		response, err = client.ListLoadBalancers(request)
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

// ListLoadBalancersRequest is the request struct for api ListLoadBalancers
type ListLoadBalancersRequest struct {
	*requests.RpcRequest
	LoadBalancerNames          *[]string               `position:"Body" name:"LoadBalancerNames"  type:"Repeated"`
	LoadBalancerIds            *[]string               `position:"Body" name:"LoadBalancerIds"  type:"Repeated"`
	Skip                       requests.Integer        `position:"Body" name:"Skip"`
	AddressIpVersion           string                  `position:"Body" name:"AddressIpVersion"`
	ResourceGroupId            string                  `position:"Body" name:"ResourceGroupId"`
	ZoneIds                    *[]string               `position:"Body" name:"ZoneIds"  type:"Repeated"`
	NextToken                  string                  `position:"Body" name:"NextToken"`
	VpcIds                     *[]string               `position:"Body" name:"VpcIds"  type:"Repeated"`
	Tag                        *[]ListLoadBalancersTag `position:"Body" name:"Tag"  type:"Repeated"`
	LoadBalancerBusinessStatus string                  `position:"Body" name:"LoadBalancerBusinessStatus"`
	LoadBalancerStatus         string                  `position:"Body" name:"LoadBalancerStatus"`
	MaxResults                 requests.Integer        `position:"Body" name:"MaxResults"`
}

// ListLoadBalancersTag is a repeated param struct in ListLoadBalancersRequest
type ListLoadBalancersTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListLoadBalancersResponse is the response struct for api ListLoadBalancers
type ListLoadBalancersResponse struct {
	*responses.BaseResponse
	MaxResults    int    `json:"MaxResults" xml:"MaxResults"`
	NextToken     string `json:"NextToken" xml:"NextToken"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TotalCount    int    `json:"TotalCount" xml:"TotalCount"`
	LoadBalancers []Data `json:"LoadBalancers" xml:"LoadBalancers"`
}

// CreateListLoadBalancersRequest creates a request to invoke ListLoadBalancers API
func CreateListLoadBalancersRequest() (request *ListLoadBalancersRequest) {
	request = &ListLoadBalancersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Gwlb", "2024-04-15", "ListLoadBalancers", "gwlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLoadBalancersResponse creates a response to parse from ListLoadBalancers response
func CreateListLoadBalancersResponse() (response *ListLoadBalancersResponse) {
	response = &ListLoadBalancersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
