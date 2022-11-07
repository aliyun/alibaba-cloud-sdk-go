package oos

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

// SearchInventory invokes the oos.SearchInventory API synchronously
func (client *Client) SearchInventory(request *SearchInventoryRequest) (response *SearchInventoryResponse, err error) {
	response = CreateSearchInventoryResponse()
	err = client.DoAction(request, response)
	return
}

// SearchInventoryWithChan invokes the oos.SearchInventory API asynchronously
func (client *Client) SearchInventoryWithChan(request *SearchInventoryRequest) (<-chan *SearchInventoryResponse, <-chan error) {
	responseChan := make(chan *SearchInventoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchInventory(request)
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

// SearchInventoryWithCallback invokes the oos.SearchInventory API asynchronously
func (client *Client) SearchInventoryWithCallback(request *SearchInventoryRequest, callback func(response *SearchInventoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchInventoryResponse
		var err error
		defer close(result)
		response, err = client.SearchInventory(request)
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

// SearchInventoryRequest is the request struct for api SearchInventory
type SearchInventoryRequest struct {
	*requests.RpcRequest
	Aggregator *[]string                `position:"Query" name:"Aggregator"  type:"Repeated"`
	Filter     *[]SearchInventoryFilter `position:"Query" name:"Filter"  type:"Repeated"`
	NextToken  string                   `position:"Query" name:"NextToken"`
	MaxResults requests.Integer         `position:"Query" name:"MaxResults"`
}

// SearchInventoryFilter is a repeated param struct in SearchInventoryRequest
type SearchInventoryFilter struct {
	Name     string    `name:"Name"`
	Value    *[]string `name:"Value" type:"Repeated"`
	Operator string    `name:"Operator"`
}

// SearchInventoryResponse is the response struct for api SearchInventory
type SearchInventoryResponse struct {
	*responses.BaseResponse
	NextToken  string                   `json:"NextToken" xml:"NextToken"`
	RequestId  string                   `json:"RequestId" xml:"RequestId"`
	MaxResults int                      `json:"MaxResults" xml:"MaxResults"`
	Entities   []map[string]interface{} `json:"Entities" xml:"Entities"`
}

// CreateSearchInventoryRequest creates a request to invoke SearchInventory API
func CreateSearchInventoryRequest() (request *SearchInventoryRequest) {
	request = &SearchInventoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "SearchInventory", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchInventoryResponse creates a response to parse from SearchInventory response
func CreateSearchInventoryResponse() (response *SearchInventoryResponse) {
	response = &SearchInventoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
