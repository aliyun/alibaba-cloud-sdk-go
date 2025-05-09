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

// SubmitHotExpandPreCheckTask invokes the drds.SubmitHotExpandPreCheckTask API synchronously
func (client *Client) SubmitHotExpandPreCheckTask(request *SubmitHotExpandPreCheckTaskRequest) (response *SubmitHotExpandPreCheckTaskResponse, err error) {
	response = CreateSubmitHotExpandPreCheckTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitHotExpandPreCheckTaskWithChan invokes the drds.SubmitHotExpandPreCheckTask API asynchronously
func (client *Client) SubmitHotExpandPreCheckTaskWithChan(request *SubmitHotExpandPreCheckTaskRequest) (<-chan *SubmitHotExpandPreCheckTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitHotExpandPreCheckTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitHotExpandPreCheckTask(request)
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

// SubmitHotExpandPreCheckTaskWithCallback invokes the drds.SubmitHotExpandPreCheckTask API asynchronously
func (client *Client) SubmitHotExpandPreCheckTaskWithCallback(request *SubmitHotExpandPreCheckTaskRequest, callback func(response *SubmitHotExpandPreCheckTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitHotExpandPreCheckTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitHotExpandPreCheckTask(request)
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

// SubmitHotExpandPreCheckTaskRequest is the request struct for api SubmitHotExpandPreCheckTask
type SubmitHotExpandPreCheckTaskRequest struct {
	*requests.RpcRequest
	TableList      *[]string `position:"Query" name:"TableList"  type:"Repeated"`
	DrdsInstanceId string    `position:"Query" name:"DrdsInstanceId"`
	DbName         string    `position:"Query" name:"DbName"`
	DbInstType     string    `position:"Query" name:"DbInstType"`
}

// SubmitHotExpandPreCheckTaskResponse is the response struct for api SubmitHotExpandPreCheckTask
type SubmitHotExpandPreCheckTaskResponse struct {
	*responses.BaseResponse
	Msg       string `json:"Msg" xml:"Msg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSubmitHotExpandPreCheckTaskRequest creates a request to invoke SubmitHotExpandPreCheckTask API
func CreateSubmitHotExpandPreCheckTaskRequest() (request *SubmitHotExpandPreCheckTaskRequest) {
	request = &SubmitHotExpandPreCheckTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SubmitHotExpandPreCheckTask", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitHotExpandPreCheckTaskResponse creates a response to parse from SubmitHotExpandPreCheckTask response
func CreateSubmitHotExpandPreCheckTaskResponse() (response *SubmitHotExpandPreCheckTaskResponse) {
	response = &SubmitHotExpandPreCheckTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
