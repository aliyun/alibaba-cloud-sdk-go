package dcdn

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

// DescribeDcdnWafDomain invokes the dcdn.DescribeDcdnWafDomain API synchronously
func (client *Client) DescribeDcdnWafDomain(request *DescribeDcdnWafDomainRequest) (response *DescribeDcdnWafDomainResponse, err error) {
	response = CreateDescribeDcdnWafDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnWafDomainWithChan invokes the dcdn.DescribeDcdnWafDomain API asynchronously
func (client *Client) DescribeDcdnWafDomainWithChan(request *DescribeDcdnWafDomainRequest) (<-chan *DescribeDcdnWafDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnWafDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnWafDomain(request)
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

// DescribeDcdnWafDomainWithCallback invokes the dcdn.DescribeDcdnWafDomain API asynchronously
func (client *Client) DescribeDcdnWafDomainWithCallback(request *DescribeDcdnWafDomainRequest, callback func(response *DescribeDcdnWafDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnWafDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnWafDomain(request)
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

// DescribeDcdnWafDomainRequest is the request struct for api DescribeDcdnWafDomain
type DescribeDcdnWafDomainRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnWafDomainResponse is the response struct for api DescribeDcdnWafDomain
type DescribeDcdnWafDomainResponse struct {
	*responses.BaseResponse
	TotalCount    int            `json:"TotalCount" xml:"TotalCount"`
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	OutPutDomains []OutPutDomain `json:"OutPutDomains" xml:"OutPutDomains"`
}

// CreateDescribeDcdnWafDomainRequest creates a request to invoke DescribeDcdnWafDomain API
func CreateDescribeDcdnWafDomainRequest() (request *DescribeDcdnWafDomainRequest) {
	request = &DescribeDcdnWafDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnWafDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnWafDomainResponse creates a response to parse from DescribeDcdnWafDomain response
func CreateDescribeDcdnWafDomainResponse() (response *DescribeDcdnWafDomainResponse) {
	response = &DescribeDcdnWafDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
