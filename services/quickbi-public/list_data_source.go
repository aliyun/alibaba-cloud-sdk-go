package quickbi_public

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

// ListDataSource invokes the quickbi_public.ListDataSource API synchronously
func (client *Client) ListDataSource(request *ListDataSourceRequest) (response *ListDataSourceResponse, err error) {
	response = CreateListDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataSourceWithChan invokes the quickbi_public.ListDataSource API asynchronously
func (client *Client) ListDataSourceWithChan(request *ListDataSourceRequest) (<-chan *ListDataSourceResponse, <-chan error) {
	responseChan := make(chan *ListDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataSource(request)
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

// ListDataSourceWithCallback invokes the quickbi_public.ListDataSource API asynchronously
func (client *Client) ListDataSourceWithCallback(request *ListDataSourceRequest, callback func(response *ListDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataSourceResponse
		var err error
		defer close(result)
		response, err = client.ListDataSource(request)
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

// ListDataSourceRequest is the request struct for api ListDataSource
type ListDataSourceRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	DsType      string `position:"Query" name:"DsType"`
	SignType    string `position:"Query" name:"SignType"`
	WorkspaceId string `position:"Query" name:"WorkspaceId"`
}

// ListDataSourceResponse is the response struct for api ListDataSource
type ListDataSourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateListDataSourceRequest creates a request to invoke ListDataSource API
func CreateListDataSourceRequest() (request *ListDataSourceRequest) {
	request = &ListDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListDataSource", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataSourceResponse creates a response to parse from ListDataSource response
func CreateListDataSourceResponse() (response *ListDataSourceResponse) {
	response = &ListDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
