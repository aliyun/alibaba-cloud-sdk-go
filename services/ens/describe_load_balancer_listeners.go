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

// DescribeLoadBalancerListeners invokes the ens.DescribeLoadBalancerListeners API synchronously
func (client *Client) DescribeLoadBalancerListeners(request *DescribeLoadBalancerListenersRequest) (response *DescribeLoadBalancerListenersResponse, err error) {
	response = CreateDescribeLoadBalancerListenersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerListenersWithChan invokes the ens.DescribeLoadBalancerListeners API asynchronously
func (client *Client) DescribeLoadBalancerListenersWithChan(request *DescribeLoadBalancerListenersRequest) (<-chan *DescribeLoadBalancerListenersResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerListenersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerListeners(request)
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

// DescribeLoadBalancerListenersWithCallback invokes the ens.DescribeLoadBalancerListeners API asynchronously
func (client *Client) DescribeLoadBalancerListenersWithCallback(request *DescribeLoadBalancerListenersRequest, callback func(response *DescribeLoadBalancerListenersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerListenersResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerListeners(request)
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

// DescribeLoadBalancerListenersRequest is the request struct for api DescribeLoadBalancerListeners
type DescribeLoadBalancerListenersRequest struct {
	*requests.RpcRequest
	Description    string           `position:"Query" name:"Description"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ListenerPort   requests.Integer `position:"Query" name:"ListenerPort"`
	LoadBalancerId string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerListenersResponse is the response struct for api DescribeLoadBalancerListeners
type DescribeLoadBalancerListenersResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	Listeners  Listeners `json:"Listeners" xml:"Listeners"`
}

// CreateDescribeLoadBalancerListenersRequest creates a request to invoke DescribeLoadBalancerListeners API
func CreateDescribeLoadBalancerListenersRequest() (request *DescribeLoadBalancerListenersRequest) {
	request = &DescribeLoadBalancerListenersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeLoadBalancerListeners", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerListenersResponse creates a response to parse from DescribeLoadBalancerListeners response
func CreateDescribeLoadBalancerListenersResponse() (response *DescribeLoadBalancerListenersResponse) {
	response = &DescribeLoadBalancerListenersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
