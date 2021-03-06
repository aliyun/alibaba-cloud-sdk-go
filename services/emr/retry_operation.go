package emr

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

// RetryOperation invokes the emr.RetryOperation API synchronously
func (client *Client) RetryOperation(request *RetryOperationRequest) (response *RetryOperationResponse, err error) {
	response = CreateRetryOperationResponse()
	err = client.DoAction(request, response)
	return
}

// RetryOperationWithChan invokes the emr.RetryOperation API asynchronously
func (client *Client) RetryOperationWithChan(request *RetryOperationRequest) (<-chan *RetryOperationResponse, <-chan error) {
	responseChan := make(chan *RetryOperationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetryOperation(request)
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

// RetryOperationWithCallback invokes the emr.RetryOperation API asynchronously
func (client *Client) RetryOperationWithCallback(request *RetryOperationRequest, callback func(response *RetryOperationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetryOperationResponse
		var err error
		defer close(result)
		response, err = client.RetryOperation(request)
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

// RetryOperationRequest is the request struct for api RetryOperation
type RetryOperationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OperationId     string           `position:"Query" name:"OperationId"`
}

// RetryOperationResponse is the response struct for api RetryOperation
type RetryOperationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
}

// CreateRetryOperationRequest creates a request to invoke RetryOperation API
func CreateRetryOperationRequest() (request *RetryOperationRequest) {
	request = &RetryOperationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "RetryOperation", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRetryOperationResponse creates a response to parse from RetryOperation response
func CreateRetryOperationResponse() (response *RetryOperationResponse) {
	response = &RetryOperationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
