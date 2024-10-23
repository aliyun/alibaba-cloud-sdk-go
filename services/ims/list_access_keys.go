package ims

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

// ListAccessKeys invokes the ims.ListAccessKeys API synchronously
func (client *Client) ListAccessKeys(request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
	response = CreateListAccessKeysResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccessKeysWithChan invokes the ims.ListAccessKeys API asynchronously
func (client *Client) ListAccessKeysWithChan(request *ListAccessKeysRequest) (<-chan *ListAccessKeysResponse, <-chan error) {
	responseChan := make(chan *ListAccessKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccessKeys(request)
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

// ListAccessKeysWithCallback invokes the ims.ListAccessKeys API asynchronously
func (client *Client) ListAccessKeysWithCallback(request *ListAccessKeysRequest, callback func(response *ListAccessKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccessKeysResponse
		var err error
		defer close(result)
		response, err = client.ListAccessKeys(request)
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

// ListAccessKeysRequest is the request struct for api ListAccessKeys
type ListAccessKeysRequest struct {
	*requests.RpcRequest
	AkProxySuffix     string `position:"Query" name:"AkProxySuffix"`
	UserPrincipalName string `position:"Query" name:"UserPrincipalName"`
}

// ListAccessKeysResponse is the response struct for api ListAccessKeys
type ListAccessKeysResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	AccessKeys AccessKeys `json:"AccessKeys" xml:"AccessKeys"`
}

// CreateListAccessKeysRequest creates a request to invoke ListAccessKeys API
func CreateListAccessKeysRequest() (request *ListAccessKeysRequest) {
	request = &ListAccessKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "ListAccessKeys", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccessKeysResponse creates a response to parse from ListAccessKeys response
func CreateListAccessKeysResponse() (response *ListAccessKeysResponse) {
	response = &ListAccessKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
