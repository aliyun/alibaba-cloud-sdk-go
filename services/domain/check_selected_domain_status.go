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

// CheckSelectedDomainStatus invokes the domain.CheckSelectedDomainStatus API synchronously
func (client *Client) CheckSelectedDomainStatus(request *CheckSelectedDomainStatusRequest) (response *CheckSelectedDomainStatusResponse, err error) {
	response = CreateCheckSelectedDomainStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSelectedDomainStatusWithChan invokes the domain.CheckSelectedDomainStatus API asynchronously
func (client *Client) CheckSelectedDomainStatusWithChan(request *CheckSelectedDomainStatusRequest) (<-chan *CheckSelectedDomainStatusResponse, <-chan error) {
	responseChan := make(chan *CheckSelectedDomainStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSelectedDomainStatus(request)
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

// CheckSelectedDomainStatusWithCallback invokes the domain.CheckSelectedDomainStatus API asynchronously
func (client *Client) CheckSelectedDomainStatusWithCallback(request *CheckSelectedDomainStatusRequest, callback func(response *CheckSelectedDomainStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSelectedDomainStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckSelectedDomainStatus(request)
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

// CheckSelectedDomainStatusRequest is the request struct for api CheckSelectedDomainStatus
type CheckSelectedDomainStatusRequest struct {
	*requests.RpcRequest
	Domain string `position:"Query" name:"Domain"`
}

// CheckSelectedDomainStatusResponse is the response struct for api CheckSelectedDomainStatus
type CheckSelectedDomainStatusResponse struct {
	*responses.BaseResponse
	RequestId      string                            `json:"RequestId" xml:"RequestId"`
	ErrorCode      string                            `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool                              `json:"Success" xml:"Success"`
	HttpStatusCode int                               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Module         ModuleInCheckSelectedDomainStatus `json:"Module" xml:"Module"`
}

// CreateCheckSelectedDomainStatusRequest creates a request to invoke CheckSelectedDomainStatus API
func CreateCheckSelectedDomainStatusRequest() (request *CheckSelectedDomainStatusRequest) {
	request = &CheckSelectedDomainStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "CheckSelectedDomainStatus", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckSelectedDomainStatusResponse creates a response to parse from CheckSelectedDomainStatus response
func CreateCheckSelectedDomainStatusResponse() (response *CheckSelectedDomainStatusResponse) {
	response = &CheckSelectedDomainStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
