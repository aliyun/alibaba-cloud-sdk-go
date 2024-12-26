package schedulerx3

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

// ListAppNames invokes the schedulerx3.ListAppNames API synchronously
func (client *Client) ListAppNames(request *ListAppNamesRequest) (response *ListAppNamesResponse, err error) {
	response = CreateListAppNamesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppNamesWithChan invokes the schedulerx3.ListAppNames API asynchronously
func (client *Client) ListAppNamesWithChan(request *ListAppNamesRequest) (<-chan *ListAppNamesResponse, <-chan error) {
	responseChan := make(chan *ListAppNamesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppNames(request)
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

// ListAppNamesWithCallback invokes the schedulerx3.ListAppNames API asynchronously
func (client *Client) ListAppNamesWithCallback(request *ListAppNamesRequest, callback func(response *ListAppNamesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppNamesResponse
		var err error
		defer close(result)
		response, err = client.ListAppNames(request)
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

// ListAppNamesRequest is the request struct for api ListAppNames
type ListAppNamesRequest struct {
	*requests.RpcRequest
	MseSessionId string `position:"Query" name:"MseSessionId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
	AppName      string `position:"Query" name:"AppName"`
}

// ListAppNamesResponse is the response struct for api ListAppNames
type ListAppNamesResponse struct {
	*responses.BaseResponse
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListAppNamesRequest creates a request to invoke ListAppNames API
func CreateListAppNamesRequest() (request *ListAppNamesRequest) {
	request = &ListAppNamesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SchedulerX3", "2024-06-24", "ListAppNames", "", "")
	request.Method = requests.GET
	return
}

// CreateListAppNamesResponse creates a response to parse from ListAppNames response
func CreateListAppNamesResponse() (response *ListAppNamesResponse) {
	response = &ListAppNamesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
