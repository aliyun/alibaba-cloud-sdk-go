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

// DescribeInvocations invokes the ecd.DescribeInvocations API synchronously
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
	response = CreateDescribeInvocationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInvocationsWithChan invokes the ecd.DescribeInvocations API asynchronously
func (client *Client) DescribeInvocationsWithChan(request *DescribeInvocationsRequest) (<-chan *DescribeInvocationsResponse, <-chan error) {
	responseChan := make(chan *DescribeInvocationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInvocations(request)
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

// DescribeInvocationsWithCallback invokes the ecd.DescribeInvocations API asynchronously
func (client *Client) DescribeInvocationsWithCallback(request *DescribeInvocationsRequest, callback func(response *DescribeInvocationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInvocationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInvocations(request)
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

// DescribeInvocationsRequest is the request struct for api DescribeInvocations
type DescribeInvocationsRequest struct {
	*requests.RpcRequest
	InvokeStatus          string           `position:"Query" name:"InvokeStatus"`
	DesktopIds            *[]string        `position:"Query" name:"DesktopIds"  type:"Repeated"`
	IncludeOutput         requests.Boolean `position:"Query" name:"IncludeOutput"`
	NextToken             string           `position:"Query" name:"NextToken"`
	ContentEncoding       string           `position:"Query" name:"ContentEncoding"`
	EndUserId             string           `position:"Query" name:"EndUserId"`
	DesktopId             string           `position:"Query" name:"DesktopId"`
	InvokeId              string           `position:"Query" name:"InvokeId"`
	CommandType           string           `position:"Query" name:"CommandType"`
	MaxResults            requests.Integer `position:"Query" name:"MaxResults"`
	IncludeInvokeDesktops requests.Boolean `position:"Query" name:"IncludeInvokeDesktops"`
}

// DescribeInvocationsResponse is the response struct for api DescribeInvocations
type DescribeInvocationsResponse struct {
	*responses.BaseResponse
	NextToken   string       `json:"NextToken" xml:"NextToken"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	Invocations []Invocation `json:"Invocations" xml:"Invocations"`
}

// CreateDescribeInvocationsRequest creates a request to invoke DescribeInvocations API
func CreateDescribeInvocationsRequest() (request *DescribeInvocationsRequest) {
	request = &DescribeInvocationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeInvocations", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInvocationsResponse creates a response to parse from DescribeInvocations response
func CreateDescribeInvocationsResponse() (response *DescribeInvocationsResponse) {
	response = &DescribeInvocationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
