package dms_enterprise

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

// ListSLARules invokes the dms_enterprise.ListSLARules API synchronously
func (client *Client) ListSLARules(request *ListSLARulesRequest) (response *ListSLARulesResponse, err error) {
	response = CreateListSLARulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListSLARulesWithChan invokes the dms_enterprise.ListSLARules API asynchronously
func (client *Client) ListSLARulesWithChan(request *ListSLARulesRequest) (<-chan *ListSLARulesResponse, <-chan error) {
	responseChan := make(chan *ListSLARulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSLARules(request)
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

// ListSLARulesWithCallback invokes the dms_enterprise.ListSLARules API asynchronously
func (client *Client) ListSLARulesWithCallback(request *ListSLARulesRequest, callback func(response *ListSLARulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSLARulesResponse
		var err error
		defer close(result)
		response, err = client.ListSLARules(request)
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

// ListSLARulesRequest is the request struct for api ListSLARules
type ListSLARulesRequest struct {
	*requests.RpcRequest
	DagId requests.Integer `position:"Query" name:"DagId"`
	Tid   requests.Integer `position:"Query" name:"Tid"`
}

// ListSLARulesResponse is the response struct for api ListSLARules
type ListSLARulesResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	ErrorCode    string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool                      `json:"Success" xml:"Success"`
	SLARuleList  SLARuleListInListSLARules `json:"SLARuleList" xml:"SLARuleList"`
}

// CreateListSLARulesRequest creates a request to invoke ListSLARules API
func CreateListSLARulesRequest() (request *ListSLARulesRequest) {
	request = &ListSLARulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListSLARules", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSLARulesResponse creates a response to parse from ListSLARules response
func CreateListSLARulesResponse() (response *ListSLARulesResponse) {
	response = &ListSLARulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
