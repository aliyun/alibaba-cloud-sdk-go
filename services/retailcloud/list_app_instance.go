package retailcloud

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

// ListAppInstance invokes the retailcloud.ListAppInstance API synchronously
// api document: https://help.aliyun.com/api/retailcloud/listappinstance.html
func (client *Client) ListAppInstance(request *ListAppInstanceRequest) (response *ListAppInstanceResponse, err error) {
	response = CreateListAppInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppInstanceWithChan invokes the retailcloud.ListAppInstance API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listappinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAppInstanceWithChan(request *ListAppInstanceRequest) (<-chan *ListAppInstanceResponse, <-chan error) {
	responseChan := make(chan *ListAppInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppInstance(request)
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

// ListAppInstanceWithCallback invokes the retailcloud.ListAppInstance API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listappinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAppInstanceWithCallback(request *ListAppInstanceRequest, callback func(response *ListAppInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppInstanceResponse
		var err error
		defer close(result)
		response, err = client.ListAppInstance(request)
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

// ListAppInstanceRequest is the request struct for api ListAppInstance
type ListAppInstanceRequest struct {
	*requests.RpcRequest
	AppId      requests.Integer `position:"Body" name:"AppId"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	EnvId      requests.Integer `position:"Body" name:"EnvId"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
}

// ListAppInstanceResponse is the response struct for api ListAppInstance
type ListAppInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string            `json:"RequestId" xml:"RequestId"`
	Code       int               `json:"Code" xml:"Code"`
	PageSize   int               `json:"PageSize" xml:"PageSize"`
	PageNumber int               `json:"PageNumber" xml:"PageNumber"`
	TotalCount int64             `json:"TotalCount" xml:"TotalCount"`
	ErrMsg     string            `json:"ErrMsg" xml:"ErrMsg"`
	Data       []ListAppInstance `json:"Data" xml:"Data"`
}

// CreateListAppInstanceRequest creates a request to invoke ListAppInstance API
func CreateListAppInstanceRequest() (request *ListAppInstanceRequest) {
	request = &ListAppInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListAppInstance", "retailcloud", "openAPI")
	return
}

// CreateListAppInstanceResponse creates a response to parse from ListAppInstance response
func CreateListAppInstanceResponse() (response *ListAppInstanceResponse) {
	response = &ListAppInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
