package aimath

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

// UpdateAnalysis invokes the aimath.UpdateAnalysis API synchronously
func (client *Client) UpdateAnalysis(request *UpdateAnalysisRequest) (response *UpdateAnalysisResponse, err error) {
	response = CreateUpdateAnalysisResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAnalysisWithChan invokes the aimath.UpdateAnalysis API asynchronously
func (client *Client) UpdateAnalysisWithChan(request *UpdateAnalysisRequest) (<-chan *UpdateAnalysisResponse, <-chan error) {
	responseChan := make(chan *UpdateAnalysisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAnalysis(request)
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

// UpdateAnalysisWithCallback invokes the aimath.UpdateAnalysis API asynchronously
func (client *Client) UpdateAnalysisWithCallback(request *UpdateAnalysisRequest, callback func(response *UpdateAnalysisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAnalysisResponse
		var err error
		defer close(result)
		response, err = client.UpdateAnalysis(request)
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

// UpdateAnalysisRequest is the request struct for api UpdateAnalysis
type UpdateAnalysisRequest struct {
	*requests.RpcRequest
	Content      string `position:"Body" name:"Content"`
	ContentCode  string `position:"Body" name:"ContentCode"`
	ExerciseCode string `position:"Body" name:"ExerciseCode"`
}

// UpdateAnalysisResponse is the response struct for api UpdateAnalysis
type UpdateAnalysisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
}

// CreateUpdateAnalysisRequest creates a request to invoke UpdateAnalysis API
func CreateUpdateAnalysisRequest() (request *UpdateAnalysisRequest) {
	request = &UpdateAnalysisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AIMath", "2024-11-14", "UpdateAnalysis", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAnalysisResponse creates a response to parse from UpdateAnalysis response
func CreateUpdateAnalysisResponse() (response *UpdateAnalysisResponse) {
	response = &UpdateAnalysisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
