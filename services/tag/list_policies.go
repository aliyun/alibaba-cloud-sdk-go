package tag

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

// ListPolicies invokes the tag.ListPolicies API synchronously
func (client *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
	response = CreateListPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPoliciesWithChan invokes the tag.ListPolicies API asynchronously
func (client *Client) ListPoliciesWithChan(request *ListPoliciesRequest) (<-chan *ListPoliciesResponse, <-chan error) {
	responseChan := make(chan *ListPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPolicies(request)
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

// ListPoliciesWithCallback invokes the tag.ListPolicies API asynchronously
func (client *Client) ListPoliciesWithCallback(request *ListPoliciesRequest, callback func(response *ListPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPoliciesResponse
		var err error
		defer close(result)
		response, err = client.ListPolicies(request)
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

// ListPoliciesRequest is the request struct for api ListPolicies
type ListPoliciesRequest struct {
	*requests.RpcRequest
	PolicyNames          *[]string        `position:"Query" name:"PolicyNames"  type:"Repeated"`
	PolicyIds            *[]string        `position:"Query" name:"PolicyIds"  type:"Repeated"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UserType             string           `position:"Query" name:"UserType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaxResult            requests.Integer `position:"Query" name:"MaxResult"`
}

// ListPoliciesResponse is the response struct for api ListPolicies
type ListPoliciesResponse struct {
	*responses.BaseResponse
	NextToken  string   `json:"NextToken" xml:"NextToken"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PolicyList []Policy `json:"PolicyList" xml:"PolicyList"`
}

// CreateListPoliciesRequest creates a request to invoke ListPolicies API
func CreateListPoliciesRequest() (request *ListPoliciesRequest) {
	request = &ListPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "ListPolicies", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPoliciesResponse creates a response to parse from ListPolicies response
func CreateListPoliciesResponse() (response *ListPoliciesResponse) {
	response = &ListPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
