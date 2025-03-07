package live

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

// AddLiveDomain invokes the live.AddLiveDomain API synchronously
func (client *Client) AddLiveDomain(request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
	response = CreateAddLiveDomainResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveDomainWithChan invokes the live.AddLiveDomain API asynchronously
func (client *Client) AddLiveDomainWithChan(request *AddLiveDomainRequest) (<-chan *AddLiveDomainResponse, <-chan error) {
	responseChan := make(chan *AddLiveDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveDomain(request)
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

// AddLiveDomainWithCallback invokes the live.AddLiveDomain API asynchronously
func (client *Client) AddLiveDomainWithCallback(request *AddLiveDomainRequest, callback func(response *AddLiveDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveDomainResponse
		var err error
		defer close(result)
		response, err = client.AddLiveDomain(request)
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

// AddLiveDomainRequest is the request struct for api AddLiveDomain
type AddLiveDomainRequest struct {
	*requests.RpcRequest
	ResourceGroupId string              `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string              `position:"Query" name:"SecurityToken"`
	Scope           string              `position:"Query" name:"Scope"`
	Tag             *[]AddLiveDomainTag `position:"Query" name:"Tag"  type:"Repeated"`
	TopLevelDomain  string              `position:"Query" name:"TopLevelDomain"`
	OwnerAccount    string              `position:"Query" name:"OwnerAccount"`
	DomainName      string              `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer    `position:"Query" name:"OwnerId"`
	Region          string              `position:"Query" name:"Region"`
	CheckUrl        string              `position:"Query" name:"CheckUrl"`
	LiveDomainType  string              `position:"Query" name:"LiveDomainType"`
}

// AddLiveDomainTag is a repeated param struct in AddLiveDomainRequest
type AddLiveDomainTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// AddLiveDomainResponse is the response struct for api AddLiveDomain
type AddLiveDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveDomainRequest creates a request to invoke AddLiveDomain API
func CreateAddLiveDomainRequest() (request *AddLiveDomainRequest) {
	request = &AddLiveDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLiveDomain", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLiveDomainResponse creates a response to parse from AddLiveDomain response
func CreateAddLiveDomainResponse() (response *AddLiveDomainResponse) {
	response = &AddLiveDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
