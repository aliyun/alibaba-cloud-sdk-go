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

// SubmitOrderApproval invokes the dms_enterprise.SubmitOrderApproval API synchronously
func (client *Client) SubmitOrderApproval(request *SubmitOrderApprovalRequest) (response *SubmitOrderApprovalResponse, err error) {
	response = CreateSubmitOrderApprovalResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitOrderApprovalWithChan invokes the dms_enterprise.SubmitOrderApproval API asynchronously
func (client *Client) SubmitOrderApprovalWithChan(request *SubmitOrderApprovalRequest) (<-chan *SubmitOrderApprovalResponse, <-chan error) {
	responseChan := make(chan *SubmitOrderApprovalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitOrderApproval(request)
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

// SubmitOrderApprovalWithCallback invokes the dms_enterprise.SubmitOrderApproval API asynchronously
func (client *Client) SubmitOrderApprovalWithCallback(request *SubmitOrderApprovalRequest, callback func(response *SubmitOrderApprovalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitOrderApprovalResponse
		var err error
		defer close(result)
		response, err = client.SubmitOrderApproval(request)
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

// SubmitOrderApprovalRequest is the request struct for api SubmitOrderApproval
type SubmitOrderApprovalRequest struct {
	*requests.RpcRequest
	Tid              requests.Integer `position:"Query" name:"Tid"`
	OrderId          requests.Integer `position:"Query" name:"OrderId"`
	RealLoginUserUid string           `position:"Query" name:"RealLoginUserUid"`
}

// SubmitOrderApprovalResponse is the response struct for api SubmitOrderApproval
type SubmitOrderApprovalResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateSubmitOrderApprovalRequest creates a request to invoke SubmitOrderApproval API
func CreateSubmitOrderApprovalRequest() (request *SubmitOrderApprovalRequest) {
	request = &SubmitOrderApprovalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SubmitOrderApproval", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitOrderApprovalResponse creates a response to parse from SubmitOrderApproval response
func CreateSubmitOrderApprovalResponse() (response *SubmitOrderApprovalResponse) {
	response = &SubmitOrderApprovalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
