package tingwu

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

// GetTaskInfo invokes the tingwu.GetTaskInfo API synchronously
func (client *Client) GetTaskInfo(request *GetTaskInfoRequest) (response *GetTaskInfoResponse, err error) {
	response = CreateGetTaskInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskInfoWithChan invokes the tingwu.GetTaskInfo API asynchronously
func (client *Client) GetTaskInfoWithChan(request *GetTaskInfoRequest) (<-chan *GetTaskInfoResponse, <-chan error) {
	responseChan := make(chan *GetTaskInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaskInfo(request)
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

// GetTaskInfoWithCallback invokes the tingwu.GetTaskInfo API asynchronously
func (client *Client) GetTaskInfoWithCallback(request *GetTaskInfoRequest, callback func(response *GetTaskInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskInfoResponse
		var err error
		defer close(result)
		response, err = client.GetTaskInfo(request)
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

// GetTaskInfoRequest is the request struct for api GetTaskInfo
type GetTaskInfoRequest struct {
	*requests.RoaRequest
	TaskId string `position:"Path" name:"TaskId"`
}

// GetTaskInfoResponse is the response struct for api GetTaskInfo
type GetTaskInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetTaskInfoRequest creates a request to invoke GetTaskInfo API
func CreateGetTaskInfoRequest() (request *GetTaskInfoRequest) {
	request = &GetTaskInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("tingwu", "2023-09-30", "GetTaskInfo", "/openapi/tingwu/v2/tasks/[taskId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetTaskInfoResponse creates a response to parse from GetTaskInfo response
func CreateGetTaskInfoResponse() (response *GetTaskInfoResponse) {
	response = &GetTaskInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
