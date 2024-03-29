package kms

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

// ListKmsInstances invokes the kms.ListKmsInstances API synchronously
func (client *Client) ListKmsInstances(request *ListKmsInstancesRequest) (response *ListKmsInstancesResponse, err error) {
	response = CreateListKmsInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListKmsInstancesWithChan invokes the kms.ListKmsInstances API asynchronously
func (client *Client) ListKmsInstancesWithChan(request *ListKmsInstancesRequest) (<-chan *ListKmsInstancesResponse, <-chan error) {
	responseChan := make(chan *ListKmsInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListKmsInstances(request)
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

// ListKmsInstancesWithCallback invokes the kms.ListKmsInstances API asynchronously
func (client *Client) ListKmsInstancesWithCallback(request *ListKmsInstancesRequest, callback func(response *ListKmsInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListKmsInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListKmsInstances(request)
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

// ListKmsInstancesRequest is the request struct for api ListKmsInstances
type ListKmsInstancesRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListKmsInstancesResponse is the response struct for api ListKmsInstances
type ListKmsInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TotalCount   int64        `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int64        `json:"PageNumber" xml:"PageNumber"`
	PageSize     int64        `json:"PageSize" xml:"PageSize"`
	KmsInstances KmsInstances `json:"KmsInstances" xml:"KmsInstances"`
}

// CreateListKmsInstancesRequest creates a request to invoke ListKmsInstances API
func CreateListKmsInstancesRequest() (request *ListKmsInstancesRequest) {
	request = &ListKmsInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListKmsInstances", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListKmsInstancesResponse creates a response to parse from ListKmsInstances response
func CreateListKmsInstancesResponse() (response *ListKmsInstancesResponse) {
	response = &ListKmsInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
