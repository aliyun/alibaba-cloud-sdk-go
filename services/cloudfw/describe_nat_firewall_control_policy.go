package cloudfw

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

// DescribeNatFirewallControlPolicy invokes the cloudfw.DescribeNatFirewallControlPolicy API synchronously
func (client *Client) DescribeNatFirewallControlPolicy(request *DescribeNatFirewallControlPolicyRequest) (response *DescribeNatFirewallControlPolicyResponse, err error) {
	response = CreateDescribeNatFirewallControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNatFirewallControlPolicyWithChan invokes the cloudfw.DescribeNatFirewallControlPolicy API asynchronously
func (client *Client) DescribeNatFirewallControlPolicyWithChan(request *DescribeNatFirewallControlPolicyRequest) (<-chan *DescribeNatFirewallControlPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeNatFirewallControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNatFirewallControlPolicy(request)
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

// DescribeNatFirewallControlPolicyWithCallback invokes the cloudfw.DescribeNatFirewallControlPolicy API asynchronously
func (client *Client) DescribeNatFirewallControlPolicyWithCallback(request *DescribeNatFirewallControlPolicyRequest, callback func(response *DescribeNatFirewallControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNatFirewallControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeNatFirewallControlPolicy(request)
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

// DescribeNatFirewallControlPolicyRequest is the request struct for api DescribeNatFirewallControlPolicy
type DescribeNatFirewallControlPolicyRequest struct {
	*requests.RpcRequest
	Release      string `position:"Query" name:"Release"`
	Destination  string `position:"Query" name:"Destination"`
	Description  string `position:"Query" name:"Description"`
	Source       string `position:"Query" name:"Source"`
	AclUuid      string `position:"Query" name:"AclUuid"`
	AclAction    string `position:"Query" name:"AclAction"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	PageSize     string `position:"Query" name:"PageSize"`
	NatGatewayId string `position:"Query" name:"NatGatewayId"`
	Lang         string `position:"Query" name:"Lang"`
	Direction    string `position:"Query" name:"Direction"`
	CurrentPage  string `position:"Query" name:"CurrentPage"`
	Proto        string `position:"Query" name:"Proto"`
}

// DescribeNatFirewallControlPolicyResponse is the response struct for api DescribeNatFirewallControlPolicy
type DescribeNatFirewallControlPolicyResponse struct {
	*responses.BaseResponse
	TotalCount string                                       `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                                       `json:"RequestId" xml:"RequestId"`
	Policys    []DataItemInDescribeNatFirewallControlPolicy `json:"Policys" xml:"Policys"`
}

// CreateDescribeNatFirewallControlPolicyRequest creates a request to invoke DescribeNatFirewallControlPolicy API
func CreateDescribeNatFirewallControlPolicyRequest() (request *DescribeNatFirewallControlPolicyRequest) {
	request = &DescribeNatFirewallControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeNatFirewallControlPolicy", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNatFirewallControlPolicyResponse creates a response to parse from DescribeNatFirewallControlPolicy response
func CreateDescribeNatFirewallControlPolicyResponse() (response *DescribeNatFirewallControlPolicyResponse) {
	response = &DescribeNatFirewallControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
