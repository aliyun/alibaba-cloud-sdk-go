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

// CheckIntlFixPriceDomainStatus invokes the domain.CheckIntlFixPriceDomainStatus API synchronously
func (client *Client) CheckIntlFixPriceDomainStatus(request *CheckIntlFixPriceDomainStatusRequest) (response *CheckIntlFixPriceDomainStatusResponse, err error) {
	response = CreateCheckIntlFixPriceDomainStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckIntlFixPriceDomainStatusWithChan invokes the domain.CheckIntlFixPriceDomainStatus API asynchronously
func (client *Client) CheckIntlFixPriceDomainStatusWithChan(request *CheckIntlFixPriceDomainStatusRequest) (<-chan *CheckIntlFixPriceDomainStatusResponse, <-chan error) {
	responseChan := make(chan *CheckIntlFixPriceDomainStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckIntlFixPriceDomainStatus(request)
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

// CheckIntlFixPriceDomainStatusWithCallback invokes the domain.CheckIntlFixPriceDomainStatus API asynchronously
func (client *Client) CheckIntlFixPriceDomainStatusWithCallback(request *CheckIntlFixPriceDomainStatusRequest, callback func(response *CheckIntlFixPriceDomainStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckIntlFixPriceDomainStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckIntlFixPriceDomainStatus(request)
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

// CheckIntlFixPriceDomainStatusRequest is the request struct for api CheckIntlFixPriceDomainStatus
type CheckIntlFixPriceDomainStatusRequest struct {
	*requests.RpcRequest
	Domain string `position:"Query" name:"Domain"`
}

// CheckIntlFixPriceDomainStatusResponse is the response struct for api CheckIntlFixPriceDomainStatus
type CheckIntlFixPriceDomainStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    Module `json:"Module" xml:"Module"`
}

// CreateCheckIntlFixPriceDomainStatusRequest creates a request to invoke CheckIntlFixPriceDomainStatus API
func CreateCheckIntlFixPriceDomainStatusRequest() (request *CheckIntlFixPriceDomainStatusRequest) {
	request = &CheckIntlFixPriceDomainStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "CheckIntlFixPriceDomainStatus", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckIntlFixPriceDomainStatusResponse creates a response to parse from CheckIntlFixPriceDomainStatus response
func CreateCheckIntlFixPriceDomainStatusResponse() (response *CheckIntlFixPriceDomainStatusResponse) {
	response = &CheckIntlFixPriceDomainStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
