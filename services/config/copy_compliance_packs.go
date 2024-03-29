package config

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

// CopyCompliancePacks invokes the config.CopyCompliancePacks API synchronously
func (client *Client) CopyCompliancePacks(request *CopyCompliancePacksRequest) (response *CopyCompliancePacksResponse, err error) {
	response = CreateCopyCompliancePacksResponse()
	err = client.DoAction(request, response)
	return
}

// CopyCompliancePacksWithChan invokes the config.CopyCompliancePacks API asynchronously
func (client *Client) CopyCompliancePacksWithChan(request *CopyCompliancePacksRequest) (<-chan *CopyCompliancePacksResponse, <-chan error) {
	responseChan := make(chan *CopyCompliancePacksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyCompliancePacks(request)
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

// CopyCompliancePacksWithCallback invokes the config.CopyCompliancePacks API asynchronously
func (client *Client) CopyCompliancePacksWithCallback(request *CopyCompliancePacksRequest, callback func(response *CopyCompliancePacksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyCompliancePacksResponse
		var err error
		defer close(result)
		response, err = client.CopyCompliancePacks(request)
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

// CopyCompliancePacksRequest is the request struct for api CopyCompliancePacks
type CopyCompliancePacksRequest struct {
	*requests.RpcRequest
	DesAggregatorIds     string `position:"Query" name:"DesAggregatorIds"`
	SrcAggregatorId      string `position:"Query" name:"SrcAggregatorId"`
	SrcCompliancePackIds string `position:"Query" name:"SrcCompliancePackIds"`
}

// CopyCompliancePacksResponse is the response struct for api CopyCompliancePacks
type CopyCompliancePacksResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	CopyRulesResult bool   `json:"CopyRulesResult" xml:"CopyRulesResult"`
}

// CreateCopyCompliancePacksRequest creates a request to invoke CopyCompliancePacks API
func CreateCopyCompliancePacksRequest() (request *CopyCompliancePacksRequest) {
	request = &CopyCompliancePacksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CopyCompliancePacks", "", "")
	request.Method = requests.POST
	return
}

// CreateCopyCompliancePacksResponse creates a response to parse from CopyCompliancePacks response
func CreateCopyCompliancePacksResponse() (response *CopyCompliancePacksResponse) {
	response = &CopyCompliancePacksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
