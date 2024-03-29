package cloud_siem

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

// DescribeAutomateResponseConfigPlayBooks invokes the cloud_siem.DescribeAutomateResponseConfigPlayBooks API synchronously
func (client *Client) DescribeAutomateResponseConfigPlayBooks(request *DescribeAutomateResponseConfigPlayBooksRequest) (response *DescribeAutomateResponseConfigPlayBooksResponse, err error) {
	response = CreateDescribeAutomateResponseConfigPlayBooksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutomateResponseConfigPlayBooksWithChan invokes the cloud_siem.DescribeAutomateResponseConfigPlayBooks API asynchronously
func (client *Client) DescribeAutomateResponseConfigPlayBooksWithChan(request *DescribeAutomateResponseConfigPlayBooksRequest) (<-chan *DescribeAutomateResponseConfigPlayBooksResponse, <-chan error) {
	responseChan := make(chan *DescribeAutomateResponseConfigPlayBooksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutomateResponseConfigPlayBooks(request)
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

// DescribeAutomateResponseConfigPlayBooksWithCallback invokes the cloud_siem.DescribeAutomateResponseConfigPlayBooks API asynchronously
func (client *Client) DescribeAutomateResponseConfigPlayBooksWithCallback(request *DescribeAutomateResponseConfigPlayBooksRequest, callback func(response *DescribeAutomateResponseConfigPlayBooksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutomateResponseConfigPlayBooksResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutomateResponseConfigPlayBooks(request)
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

// DescribeAutomateResponseConfigPlayBooksRequest is the request struct for api DescribeAutomateResponseConfigPlayBooks
type DescribeAutomateResponseConfigPlayBooksRequest struct {
	*requests.RpcRequest
	AutoResponseType string `position:"Body" name:"AutoResponseType"`
	EntityType       string `position:"Body" name:"EntityType"`
}

// DescribeAutomateResponseConfigPlayBooksResponse is the response struct for api DescribeAutomateResponseConfigPlayBooks
type DescribeAutomateResponseConfigPlayBooksResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeAutomateResponseConfigPlayBooksRequest creates a request to invoke DescribeAutomateResponseConfigPlayBooks API
func CreateDescribeAutomateResponseConfigPlayBooksRequest() (request *DescribeAutomateResponseConfigPlayBooksRequest) {
	request = &DescribeAutomateResponseConfigPlayBooksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAutomateResponseConfigPlayBooks", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutomateResponseConfigPlayBooksResponse creates a response to parse from DescribeAutomateResponseConfigPlayBooks response
func CreateDescribeAutomateResponseConfigPlayBooksResponse() (response *DescribeAutomateResponseConfigPlayBooksResponse) {
	response = &DescribeAutomateResponseConfigPlayBooksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
