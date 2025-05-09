package drds

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

// SubmitSmoothExpandPreCheck invokes the drds.SubmitSmoothExpandPreCheck API synchronously
func (client *Client) SubmitSmoothExpandPreCheck(request *SubmitSmoothExpandPreCheckRequest) (response *SubmitSmoothExpandPreCheckResponse, err error) {
	response = CreateSubmitSmoothExpandPreCheckResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitSmoothExpandPreCheckWithChan invokes the drds.SubmitSmoothExpandPreCheck API asynchronously
func (client *Client) SubmitSmoothExpandPreCheckWithChan(request *SubmitSmoothExpandPreCheckRequest) (<-chan *SubmitSmoothExpandPreCheckResponse, <-chan error) {
	responseChan := make(chan *SubmitSmoothExpandPreCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSmoothExpandPreCheck(request)
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

// SubmitSmoothExpandPreCheckWithCallback invokes the drds.SubmitSmoothExpandPreCheck API asynchronously
func (client *Client) SubmitSmoothExpandPreCheckWithCallback(request *SubmitSmoothExpandPreCheckRequest, callback func(response *SubmitSmoothExpandPreCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSmoothExpandPreCheckResponse
		var err error
		defer close(result)
		response, err = client.SubmitSmoothExpandPreCheck(request)
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

// SubmitSmoothExpandPreCheckRequest is the request struct for api SubmitSmoothExpandPreCheck
type SubmitSmoothExpandPreCheckRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	DbInstType     string `position:"Query" name:"DbInstType"`
}

// SubmitSmoothExpandPreCheckResponse is the response struct for api SubmitSmoothExpandPreCheck
type SubmitSmoothExpandPreCheckResponse struct {
	*responses.BaseResponse
	Msg       string `json:"Msg" xml:"Msg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSubmitSmoothExpandPreCheckRequest creates a request to invoke SubmitSmoothExpandPreCheck API
func CreateSubmitSmoothExpandPreCheckRequest() (request *SubmitSmoothExpandPreCheckRequest) {
	request = &SubmitSmoothExpandPreCheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SubmitSmoothExpandPreCheck", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitSmoothExpandPreCheckResponse creates a response to parse from SubmitSmoothExpandPreCheck response
func CreateSubmitSmoothExpandPreCheckResponse() (response *SubmitSmoothExpandPreCheckResponse) {
	response = &SubmitSmoothExpandPreCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
