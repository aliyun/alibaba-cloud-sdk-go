package safe

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

// StartApprove invokes the safe.StartApprove API synchronously
func (client *Client) StartApprove(request *StartApproveRequest) (response *StartApproveResponse, err error) {
	response = CreateStartApproveResponse()
	err = client.DoAction(request, response)
	return
}

// StartApproveWithChan invokes the safe.StartApprove API asynchronously
func (client *Client) StartApproveWithChan(request *StartApproveRequest) (<-chan *StartApproveResponse, <-chan error) {
	responseChan := make(chan *StartApproveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartApprove(request)
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

// StartApproveWithCallback invokes the safe.StartApprove API asynchronously
func (client *Client) StartApproveWithCallback(request *StartApproveRequest, callback func(response *StartApproveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartApproveResponse
		var err error
		defer close(result)
		response, err = client.StartApprove(request)
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

// StartApproveRequest is the request struct for api StartApprove
type StartApproveRequest struct {
	*requests.RpcRequest
	AuthKey       string           `position:"Query" name:"AuthKey"`
	ReqTimestamp  requests.Integer `position:"Query" name:"ReqTimestamp"`
	SourceOrderId string           `position:"Query" name:"SourceOrderId"`
	AuthSign      string           `position:"Query" name:"AuthSign"`
	CreatorEmpId  string           `position:"Query" name:"CreatorEmpId"`
	ExtraInfo     string           `position:"Query" name:"ExtraInfo"`
}

// StartApproveResponse is the response struct for api StartApprove
type StartApproveResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateStartApproveRequest creates a request to invoke StartApprove API
func CreateStartApproveRequest() (request *StartApproveRequest) {
	request = &StartApproveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Safe", "2022-01-17", "StartApprove", "", "")
	request.Method = requests.POST
	return
}

// CreateStartApproveResponse creates a response to parse from StartApprove response
func CreateStartApproveResponse() (response *StartApproveResponse) {
	response = &StartApproveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
