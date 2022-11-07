package cloudesl

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

// DescribeEventReason invokes the cloudesl.DescribeEventReason API synchronously
func (client *Client) DescribeEventReason(request *DescribeEventReasonRequest) (response *DescribeEventReasonResponse, err error) {
	response = CreateDescribeEventReasonResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEventReasonWithChan invokes the cloudesl.DescribeEventReason API asynchronously
func (client *Client) DescribeEventReasonWithChan(request *DescribeEventReasonRequest) (<-chan *DescribeEventReasonResponse, <-chan error) {
	responseChan := make(chan *DescribeEventReasonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEventReason(request)
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

// DescribeEventReasonWithCallback invokes the cloudesl.DescribeEventReason API asynchronously
func (client *Client) DescribeEventReasonWithCallback(request *DescribeEventReasonRequest, callback func(response *DescribeEventReasonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEventReasonResponse
		var err error
		defer close(result)
		response, err = client.DescribeEventReason(request)
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

// DescribeEventReasonRequest is the request struct for api DescribeEventReason
type DescribeEventReasonRequest struct {
	*requests.RpcRequest
	ExtraParams string `position:"Body" name:"ExtraParams"`
}

// DescribeEventReasonResponse is the response struct for api DescribeEventReason
type DescribeEventReasonResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	Message        string        `json:"Message" xml:"Message"`
	ErrorCode      string        `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string        `json:"ErrorMessage" xml:"ErrorMessage"`
	Code           string        `json:"Code" xml:"Code"`
	DynamicCode    string        `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string        `json:"DynamicMessage" xml:"DynamicMessage"`
	EventReasons   []EventReason `json:"EventReasons" xml:"EventReasons"`
}

// CreateDescribeEventReasonRequest creates a request to invoke DescribeEventReason API
func CreateDescribeEventReasonRequest() (request *DescribeEventReasonRequest) {
	request = &DescribeEventReasonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeEventReason", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeEventReasonResponse creates a response to parse from DescribeEventReason response
func CreateDescribeEventReasonResponse() (response *DescribeEventReasonResponse) {
	response = &DescribeEventReasonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
