package swas_open

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

// DescribeFirewallTemplateApplyResults invokes the swas_open.DescribeFirewallTemplateApplyResults API synchronously
func (client *Client) DescribeFirewallTemplateApplyResults(request *DescribeFirewallTemplateApplyResultsRequest) (response *DescribeFirewallTemplateApplyResultsResponse, err error) {
	response = CreateDescribeFirewallTemplateApplyResultsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFirewallTemplateApplyResultsWithChan invokes the swas_open.DescribeFirewallTemplateApplyResults API asynchronously
func (client *Client) DescribeFirewallTemplateApplyResultsWithChan(request *DescribeFirewallTemplateApplyResultsRequest) (<-chan *DescribeFirewallTemplateApplyResultsResponse, <-chan error) {
	responseChan := make(chan *DescribeFirewallTemplateApplyResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFirewallTemplateApplyResults(request)
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

// DescribeFirewallTemplateApplyResultsWithCallback invokes the swas_open.DescribeFirewallTemplateApplyResults API asynchronously
func (client *Client) DescribeFirewallTemplateApplyResultsWithCallback(request *DescribeFirewallTemplateApplyResultsRequest, callback func(response *DescribeFirewallTemplateApplyResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFirewallTemplateApplyResultsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFirewallTemplateApplyResults(request)
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

// DescribeFirewallTemplateApplyResultsRequest is the request struct for api DescribeFirewallTemplateApplyResults
type DescribeFirewallTemplateApplyResultsRequest struct {
	*requests.RpcRequest
	FirewallTemplateId string           `position:"Query" name:"FirewallTemplateId"`
	ClientToken        string           `position:"Query" name:"ClientToken"`
	PageNumber         requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	TaskId             *[]string        `position:"Query" name:"TaskId"  type:"Repeated"`
}

// DescribeFirewallTemplateApplyResultsResponse is the response struct for api DescribeFirewallTemplateApplyResults
type DescribeFirewallTemplateApplyResultsResponse struct {
	*responses.BaseResponse
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	PageNumber string                        `json:"PageNumber" xml:"PageNumber"`
	PageSize   string                        `json:"PageSize" xml:"PageSize"`
	TotalCount string                        `json:"TotalCount" xml:"TotalCount"`
	Data       []ApplyFirewallTemplateResult `json:"data" xml:"data"`
}

// CreateDescribeFirewallTemplateApplyResultsRequest creates a request to invoke DescribeFirewallTemplateApplyResults API
func CreateDescribeFirewallTemplateApplyResultsRequest() (request *DescribeFirewallTemplateApplyResultsRequest) {
	request = &DescribeFirewallTemplateApplyResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeFirewallTemplateApplyResults", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFirewallTemplateApplyResultsResponse creates a response to parse from DescribeFirewallTemplateApplyResults response
func CreateDescribeFirewallTemplateApplyResultsResponse() (response *DescribeFirewallTemplateApplyResultsResponse) {
	response = &DescribeFirewallTemplateApplyResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
