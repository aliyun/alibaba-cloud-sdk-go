package dbfs

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

// ListConstants invokes the dbfs.ListConstants API synchronously
func (client *Client) ListConstants(request *ListConstantsRequest) (response *ListConstantsResponse, err error) {
	response = CreateListConstantsResponse()
	err = client.DoAction(request, response)
	return
}

// ListConstantsWithChan invokes the dbfs.ListConstants API asynchronously
func (client *Client) ListConstantsWithChan(request *ListConstantsRequest) (<-chan *ListConstantsResponse, <-chan error) {
	responseChan := make(chan *ListConstantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConstants(request)
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

// ListConstantsWithCallback invokes the dbfs.ListConstants API asynchronously
func (client *Client) ListConstantsWithCallback(request *ListConstantsRequest, callback func(response *ListConstantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConstantsResponse
		var err error
		defer close(result)
		response, err = client.ListConstants(request)
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

// ListConstantsRequest is the request struct for api ListConstants
type ListConstantsRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	ConstantsData string           `position:"Query" name:"ConstantsData"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListConstantsResponse is the response struct for api ListConstants
type ListConstantsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Data       string `json:"Data" xml:"Data"`
	PageSize   int64  `json:"PageSize" xml:"PageSize"`
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber int64  `json:"PageNumber" xml:"PageNumber"`
}

// CreateListConstantsRequest creates a request to invoke ListConstants API
func CreateListConstantsRequest() (request *ListConstantsRequest) {
	request = &ListConstantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "ListConstants", "", "")
	request.Method = requests.POST
	return
}

// CreateListConstantsResponse creates a response to parse from ListConstants response
func CreateListConstantsResponse() (response *ListConstantsResponse) {
	response = &ListConstantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
