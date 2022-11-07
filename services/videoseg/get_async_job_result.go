package videoseg

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

// GetAsyncJobResult invokes the videoseg.GetAsyncJobResult API synchronously
func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (response *GetAsyncJobResultResponse, err error) {
	response = CreateGetAsyncJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetAsyncJobResultWithChan invokes the videoseg.GetAsyncJobResult API asynchronously
func (client *Client) GetAsyncJobResultWithChan(request *GetAsyncJobResultRequest) (<-chan *GetAsyncJobResultResponse, <-chan error) {
	responseChan := make(chan *GetAsyncJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAsyncJobResult(request)
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

// GetAsyncJobResultWithCallback invokes the videoseg.GetAsyncJobResult API asynchronously
func (client *Client) GetAsyncJobResultWithCallback(request *GetAsyncJobResultRequest, callback func(response *GetAsyncJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAsyncJobResultResponse
		var err error
		defer close(result)
		response, err = client.GetAsyncJobResult(request)
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

// GetAsyncJobResultRequest is the request struct for api GetAsyncJobResult
type GetAsyncJobResultRequest struct {
	*requests.RpcRequest
	JobId string           `position:"Body" name:"JobId"`
	Async requests.Boolean `position:"Body" name:"Async"`
}

// GetAsyncJobResultResponse is the response struct for api GetAsyncJobResult
type GetAsyncJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAsyncJobResultRequest creates a request to invoke GetAsyncJobResult API
func CreateGetAsyncJobResultRequest() (request *GetAsyncJobResultRequest) {
	request = &GetAsyncJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoseg", "2020-03-20", "GetAsyncJobResult", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAsyncJobResultResponse creates a response to parse from GetAsyncJobResult response
func CreateGetAsyncJobResultResponse() (response *GetAsyncJobResultResponse) {
	response = &GetAsyncJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
