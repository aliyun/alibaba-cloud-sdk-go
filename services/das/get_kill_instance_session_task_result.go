package das

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

// GetKillInstanceSessionTaskResult invokes the das.GetKillInstanceSessionTaskResult API synchronously
func (client *Client) GetKillInstanceSessionTaskResult(request *GetKillInstanceSessionTaskResultRequest) (response *GetKillInstanceSessionTaskResultResponse, err error) {
	response = CreateGetKillInstanceSessionTaskResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetKillInstanceSessionTaskResultWithChan invokes the das.GetKillInstanceSessionTaskResult API asynchronously
func (client *Client) GetKillInstanceSessionTaskResultWithChan(request *GetKillInstanceSessionTaskResultRequest) (<-chan *GetKillInstanceSessionTaskResultResponse, <-chan error) {
	responseChan := make(chan *GetKillInstanceSessionTaskResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetKillInstanceSessionTaskResult(request)
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

// GetKillInstanceSessionTaskResultWithCallback invokes the das.GetKillInstanceSessionTaskResult API asynchronously
func (client *Client) GetKillInstanceSessionTaskResultWithCallback(request *GetKillInstanceSessionTaskResultRequest, callback func(response *GetKillInstanceSessionTaskResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetKillInstanceSessionTaskResultResponse
		var err error
		defer close(result)
		response, err = client.GetKillInstanceSessionTaskResult(request)
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

// GetKillInstanceSessionTaskResultRequest is the request struct for api GetKillInstanceSessionTaskResult
type GetKillInstanceSessionTaskResultRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	NodeId     string `position:"Query" name:"NodeId"`
	TaskId     string `position:"Query" name:"TaskId"`
}

// GetKillInstanceSessionTaskResultResponse is the response struct for api GetKillInstanceSessionTaskResult
type GetKillInstanceSessionTaskResultResponse struct {
	*responses.BaseResponse
	Code      int64                                  `json:"Code" xml:"Code"`
	Message   string                                 `json:"Message" xml:"Message"`
	RequestId string                                 `json:"RequestId" xml:"RequestId"`
	Success   bool                                   `json:"Success" xml:"Success"`
	Data      DataInGetKillInstanceSessionTaskResult `json:"Data" xml:"Data"`
}

// CreateGetKillInstanceSessionTaskResultRequest creates a request to invoke GetKillInstanceSessionTaskResult API
func CreateGetKillInstanceSessionTaskResultRequest() (request *GetKillInstanceSessionTaskResultRequest) {
	request = &GetKillInstanceSessionTaskResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetKillInstanceSessionTaskResult", "", "")
	request.Method = requests.POST
	return
}

// CreateGetKillInstanceSessionTaskResultResponse creates a response to parse from GetKillInstanceSessionTaskResult response
func CreateGetKillInstanceSessionTaskResultResponse() (response *GetKillInstanceSessionTaskResultResponse) {
	response = &GetKillInstanceSessionTaskResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
