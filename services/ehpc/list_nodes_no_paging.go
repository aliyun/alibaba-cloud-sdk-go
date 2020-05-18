package ehpc

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

// ListNodesNoPaging invokes the ehpc.ListNodesNoPaging API synchronously
// api document: https://help.aliyun.com/api/ehpc/listnodesnopaging.html
func (client *Client) ListNodesNoPaging(request *ListNodesNoPagingRequest) (response *ListNodesNoPagingResponse, err error) {
	response = CreateListNodesNoPagingResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodesNoPagingWithChan invokes the ehpc.ListNodesNoPaging API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listnodesnopaging.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNodesNoPagingWithChan(request *ListNodesNoPagingRequest) (<-chan *ListNodesNoPagingResponse, <-chan error) {
	responseChan := make(chan *ListNodesNoPagingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodesNoPaging(request)
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

// ListNodesNoPagingWithCallback invokes the ehpc.ListNodesNoPaging API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listnodesnopaging.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNodesNoPagingWithCallback(request *ListNodesNoPagingRequest, callback func(response *ListNodesNoPagingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodesNoPagingResponse
		var err error
		defer close(result)
		response, err = client.ListNodesNoPaging(request)
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

// ListNodesNoPagingRequest is the request struct for api ListNodesNoPaging
type ListNodesNoPagingRequest struct {
	*requests.RpcRequest
	Role         string           `position:"Query" name:"Role"`
	ClusterId    string           `position:"Query" name:"ClusterId"`
	Sequence     string           `position:"Query" name:"Sequence"`
	HostName     string           `position:"Query" name:"HostName"`
	OnlyDetached requests.Boolean `position:"Query" name:"OnlyDetached"`
}

// ListNodesNoPagingResponse is the response struct for api ListNodesNoPaging
type ListNodesNoPagingResponse struct {
	*responses.BaseResponse
	RequestId  string                   `json:"RequestId" xml:"RequestId"`
	TotalCount int                      `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                      `json:"PageSize" xml:"PageSize"`
	Nodes      NodesInListNodesNoPaging `json:"Nodes" xml:"Nodes"`
}

// CreateListNodesNoPagingRequest creates a request to invoke ListNodesNoPaging API
func CreateListNodesNoPagingRequest() (request *ListNodesNoPagingRequest) {
	request = &ListNodesNoPagingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListNodesNoPaging", "", "")
	return
}

// CreateListNodesNoPagingResponse creates a response to parse from ListNodesNoPaging response
func CreateListNodesNoPagingResponse() (response *ListNodesNoPagingResponse) {
	response = &ListNodesNoPagingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
