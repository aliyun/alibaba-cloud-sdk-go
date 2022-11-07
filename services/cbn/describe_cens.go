package cbn

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

// DescribeCens invokes the cbn.DescribeCens API synchronously
func (client *Client) DescribeCens(request *DescribeCensRequest) (response *DescribeCensResponse, err error) {
	response = CreateDescribeCensResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCensWithChan invokes the cbn.DescribeCens API asynchronously
func (client *Client) DescribeCensWithChan(request *DescribeCensRequest) (<-chan *DescribeCensResponse, <-chan error) {
	responseChan := make(chan *DescribeCensResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCens(request)
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

// DescribeCensWithCallback invokes the cbn.DescribeCens API asynchronously
func (client *Client) DescribeCensWithCallback(request *DescribeCensRequest, callback func(response *DescribeCensResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCensResponse
		var err error
		defer close(result)
		response, err = client.DescribeCens(request)
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

// DescribeCensRequest is the request struct for api DescribeCens
type DescribeCensRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer      `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer      `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer      `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeCensTag    `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer      `position:"Query" name:"OwnerId"`
	Filter               *[]DescribeCensFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

// DescribeCensTag is a repeated param struct in DescribeCensRequest
type DescribeCensTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeCensFilter is a repeated param struct in DescribeCensRequest
type DescribeCensFilter struct {
	Value *[]string `name:"Value" type:"Repeated"`
	Key   string    `name:"Key"`
}

// DescribeCensResponse is the response struct for api DescribeCens
type DescribeCensResponse struct {
	*responses.BaseResponse
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Cens       Cens   `json:"Cens" xml:"Cens"`
}

// CreateDescribeCensRequest creates a request to invoke DescribeCens API
func CreateDescribeCensRequest() (request *DescribeCensRequest) {
	request = &DescribeCensRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCens", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCensResponse creates a response to parse from DescribeCens response
func CreateDescribeCensResponse() (response *DescribeCensResponse) {
	response = &DescribeCensResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
