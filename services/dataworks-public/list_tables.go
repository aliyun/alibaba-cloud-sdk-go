package dataworks_public

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

// ListTables invokes the dataworks_public.ListTables API synchronously
func (client *Client) ListTables(request *ListTablesRequest) (response *ListTablesResponse, err error) {
	response = CreateListTablesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTablesWithChan invokes the dataworks_public.ListTables API asynchronously
func (client *Client) ListTablesWithChan(request *ListTablesRequest) (<-chan *ListTablesResponse, <-chan error) {
	responseChan := make(chan *ListTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTables(request)
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

// ListTablesWithCallback invokes the dataworks_public.ListTables API asynchronously
func (client *Client) ListTablesWithCallback(request *ListTablesRequest, callback func(response *ListTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTablesResponse
		var err error
		defer close(result)
		response, err = client.ListTables(request)
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

// ListTablesRequest is the request struct for api ListTables
type ListTablesRequest struct {
	*requests.RpcRequest
	DataSourceType string           `position:"Query" name:"DataSourceType"`
	NextToken      string           `position:"Query" name:"NextToken"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// ListTablesResponse is the response struct for api ListTables
type ListTablesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListTablesRequest creates a request to invoke ListTables API
func CreateListTablesRequest() (request *ListTablesRequest) {
	request = &ListTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListTables", "", "")
	request.Method = requests.POST
	return
}

// CreateListTablesResponse creates a response to parse from ListTables response
func CreateListTablesResponse() (response *ListTablesResponse) {
	response = &ListTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
