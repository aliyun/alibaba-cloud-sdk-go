package alb

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

// ListAcls invokes the alb.ListAcls API synchronously
func (client *Client) ListAcls(request *ListAclsRequest) (response *ListAclsResponse, err error) {
	response = CreateListAclsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAclsWithChan invokes the alb.ListAcls API asynchronously
func (client *Client) ListAclsWithChan(request *ListAclsRequest) (<-chan *ListAclsResponse, <-chan error) {
	responseChan := make(chan *ListAclsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAcls(request)
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

// ListAclsWithCallback invokes the alb.ListAcls API asynchronously
func (client *Client) ListAclsWithCallback(request *ListAclsRequest, callback func(response *ListAclsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAclsResponse
		var err error
		defer close(result)
		response, err = client.ListAcls(request)
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

// ListAclsRequest is the request struct for api ListAcls
type ListAclsRequest struct {
	*requests.RpcRequest
	AclIds           *[]string        `position:"Query" name:"AclIds"  type:"Repeated"`
	AddressIPVersion string           `position:"Query" name:"AddressIPVersion"`
	ResourceGroupId  string           `position:"Query" name:"ResourceGroupId"`
	NextToken        string           `position:"Query" name:"NextToken"`
	Tag              *[]ListAclsTag   `position:"Query" name:"Tag"  type:"Repeated"`
	AclNames         *[]string        `position:"Query" name:"AclNames"  type:"Repeated"`
	MaxResults       requests.Integer `position:"Query" name:"MaxResults"`
}

// ListAclsTag is a repeated param struct in ListAclsRequest
type ListAclsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListAclsResponse is the response struct for api ListAcls
type ListAclsResponse struct {
	*responses.BaseResponse
	MaxResults int    `json:"MaxResults" xml:"MaxResults"`
	NextToken  string `json:"NextToken" xml:"NextToken"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Acls       []Acl  `json:"Acls" xml:"Acls"`
}

// CreateListAclsRequest creates a request to invoke ListAcls API
func CreateListAclsRequest() (request *ListAclsRequest) {
	request = &ListAclsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "ListAcls", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAclsResponse creates a response to parse from ListAcls response
func CreateListAclsResponse() (response *ListAclsResponse) {
	response = &ListAclsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
