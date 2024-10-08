package governance

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

// RunEvaluation invokes the governance.RunEvaluation API synchronously
func (client *Client) RunEvaluation(request *RunEvaluationRequest) (response *RunEvaluationResponse, err error) {
	response = CreateRunEvaluationResponse()
	err = client.DoAction(request, response)
	return
}

// RunEvaluationWithChan invokes the governance.RunEvaluation API asynchronously
func (client *Client) RunEvaluationWithChan(request *RunEvaluationRequest) (<-chan *RunEvaluationResponse, <-chan error) {
	responseChan := make(chan *RunEvaluationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunEvaluation(request)
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

// RunEvaluationWithCallback invokes the governance.RunEvaluation API asynchronously
func (client *Client) RunEvaluationWithCallback(request *RunEvaluationRequest, callback func(response *RunEvaluationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunEvaluationResponse
		var err error
		defer close(result)
		response, err = client.RunEvaluation(request)
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

// RunEvaluationRequest is the request struct for api RunEvaluation
type RunEvaluationRequest struct {
	*requests.RpcRequest
	AccountId   requests.Integer `position:"Query" name:"AccountId"`
	Scope       string           `position:"Query" name:"Scope"`
	PartnerCode string           `position:"Query" name:"PartnerCode"`
}

// RunEvaluationResponse is the response struct for api RunEvaluation
type RunEvaluationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRunEvaluationRequest creates a request to invoke RunEvaluation API
func CreateRunEvaluationRequest() (request *RunEvaluationRequest) {
	request = &RunEvaluationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("governance", "2021-01-20", "RunEvaluation", "governance", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunEvaluationResponse creates a response to parse from RunEvaluation response
func CreateRunEvaluationResponse() (response *RunEvaluationResponse) {
	response = &RunEvaluationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
