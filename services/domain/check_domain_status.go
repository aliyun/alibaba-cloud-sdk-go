package domain

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

// CheckDomainStatus invokes the domain.CheckDomainStatus API synchronously
func (client *Client) CheckDomainStatus(request *CheckDomainStatusRequest) (response *CheckDomainStatusResponse, err error) {
	response = CreateCheckDomainStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDomainStatusWithChan invokes the domain.CheckDomainStatus API asynchronously
func (client *Client) CheckDomainStatusWithChan(request *CheckDomainStatusRequest) (<-chan *CheckDomainStatusResponse, <-chan error) {
	responseChan := make(chan *CheckDomainStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDomainStatus(request)
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

// CheckDomainStatusWithCallback invokes the domain.CheckDomainStatus API asynchronously
func (client *Client) CheckDomainStatusWithCallback(request *CheckDomainStatusRequest, callback func(response *CheckDomainStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDomainStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckDomainStatus(request)
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

// CheckDomainStatusRequest is the request struct for api CheckDomainStatus
type CheckDomainStatusRequest struct {
	*requests.RpcRequest
	Domain string `position:"Query" name:"Domain"`
}

// CheckDomainStatusResponse is the response struct for api CheckDomainStatus
type CheckDomainStatusResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Module         Module `json:"Module" xml:"Module"`
}

// CreateCheckDomainStatusRequest creates a request to invoke CheckDomainStatus API
func CreateCheckDomainStatusRequest() (request *CheckDomainStatusRequest) {
	request = &CheckDomainStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "CheckDomainStatus", "domain", "openAPI")
	request.Method = requests.GET
	return
}

// CreateCheckDomainStatusResponse creates a response to parse from CheckDomainStatus response
func CreateCheckDomainStatusResponse() (response *CheckDomainStatusResponse) {
	response = &CheckDomainStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
