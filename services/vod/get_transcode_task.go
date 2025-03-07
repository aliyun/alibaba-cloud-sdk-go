package vod

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

// GetTranscodeTask invokes the vod.GetTranscodeTask API synchronously
func (client *Client) GetTranscodeTask(request *GetTranscodeTaskRequest) (response *GetTranscodeTaskResponse, err error) {
	response = CreateGetTranscodeTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetTranscodeTaskWithChan invokes the vod.GetTranscodeTask API asynchronously
func (client *Client) GetTranscodeTaskWithChan(request *GetTranscodeTaskRequest) (<-chan *GetTranscodeTaskResponse, <-chan error) {
	responseChan := make(chan *GetTranscodeTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTranscodeTask(request)
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

// GetTranscodeTaskWithCallback invokes the vod.GetTranscodeTask API asynchronously
func (client *Client) GetTranscodeTaskWithCallback(request *GetTranscodeTaskRequest, callback func(response *GetTranscodeTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTranscodeTaskResponse
		var err error
		defer close(result)
		response, err = client.GetTranscodeTask(request)
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

// GetTranscodeTaskRequest is the request struct for api GetTranscodeTask
type GetTranscodeTaskRequest struct {
	*requests.RpcRequest
	TranscodeTaskId string `position:"Query" name:"TranscodeTaskId"`
	JobIds          string `position:"Query" name:"JobIds"`
}

// GetTranscodeTaskResponse is the response struct for api GetTranscodeTask
type GetTranscodeTaskResponse struct {
	*responses.BaseResponse
	RequestId            string             `json:"RequestId" xml:"RequestId"`
	NonExistJobIds       []string           `json:"NonExistJobIds" xml:"NonExistJobIds"`
	TranscodeTask        TranscodeTask      `json:"TranscodeTask" xml:"TranscodeTask"`
	TranscodeJobInfoList []TranscodeJobInfo `json:"TranscodeJobInfoList" xml:"TranscodeJobInfoList"`
}

// CreateGetTranscodeTaskRequest creates a request to invoke GetTranscodeTask API
func CreateGetTranscodeTaskRequest() (request *GetTranscodeTaskRequest) {
	request = &GetTranscodeTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetTranscodeTask", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTranscodeTaskResponse creates a response to parse from GetTranscodeTask response
func CreateGetTranscodeTaskResponse() (response *GetTranscodeTaskResponse) {
	response = &GetTranscodeTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
