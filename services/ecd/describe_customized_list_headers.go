package ecd

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

// DescribeCustomizedListHeaders invokes the ecd.DescribeCustomizedListHeaders API synchronously
func (client *Client) DescribeCustomizedListHeaders(request *DescribeCustomizedListHeadersRequest) (response *DescribeCustomizedListHeadersResponse, err error) {
	response = CreateDescribeCustomizedListHeadersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomizedListHeadersWithChan invokes the ecd.DescribeCustomizedListHeaders API asynchronously
func (client *Client) DescribeCustomizedListHeadersWithChan(request *DescribeCustomizedListHeadersRequest) (<-chan *DescribeCustomizedListHeadersResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomizedListHeadersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomizedListHeaders(request)
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

// DescribeCustomizedListHeadersWithCallback invokes the ecd.DescribeCustomizedListHeaders API asynchronously
func (client *Client) DescribeCustomizedListHeadersWithCallback(request *DescribeCustomizedListHeadersRequest, callback func(response *DescribeCustomizedListHeadersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomizedListHeadersResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomizedListHeaders(request)
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

// DescribeCustomizedListHeadersRequest is the request struct for api DescribeCustomizedListHeaders
type DescribeCustomizedListHeadersRequest struct {
	*requests.RpcRequest
	LangType string `position:"Query" name:"LangType"`
	ListType string `position:"Query" name:"ListType"`
}

// DescribeCustomizedListHeadersResponse is the response struct for api DescribeCustomizedListHeaders
type DescribeCustomizedListHeadersResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Headers   []Header `json:"Headers" xml:"Headers"`
}

// CreateDescribeCustomizedListHeadersRequest creates a request to invoke DescribeCustomizedListHeaders API
func CreateDescribeCustomizedListHeadersRequest() (request *DescribeCustomizedListHeadersRequest) {
	request = &DescribeCustomizedListHeadersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeCustomizedListHeaders", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCustomizedListHeadersResponse creates a response to parse from DescribeCustomizedListHeaders response
func CreateDescribeCustomizedListHeadersResponse() (response *DescribeCustomizedListHeadersResponse) {
	response = &DescribeCustomizedListHeadersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
