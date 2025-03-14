package sophonsoar

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

// TriggerProcessTask invokes the sophonsoar.TriggerProcessTask API synchronously
func (client *Client) TriggerProcessTask(request *TriggerProcessTaskRequest) (response *TriggerProcessTaskResponse, err error) {
	response = CreateTriggerProcessTaskResponse()
	err = client.DoAction(request, response)
	return
}

// TriggerProcessTaskWithChan invokes the sophonsoar.TriggerProcessTask API asynchronously
func (client *Client) TriggerProcessTaskWithChan(request *TriggerProcessTaskRequest) (<-chan *TriggerProcessTaskResponse, <-chan error) {
	responseChan := make(chan *TriggerProcessTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TriggerProcessTask(request)
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

// TriggerProcessTaskWithCallback invokes the sophonsoar.TriggerProcessTask API asynchronously
func (client *Client) TriggerProcessTaskWithCallback(request *TriggerProcessTaskRequest, callback func(response *TriggerProcessTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TriggerProcessTaskResponse
		var err error
		defer close(result)
		response, err = client.TriggerProcessTask(request)
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

// TriggerProcessTaskRequest is the request struct for api TriggerProcessTask
type TriggerProcessTaskRequest struct {
	*requests.RpcRequest
	ActionType string `position:"Query" name:"ActionType"`
	RoleFor    string `position:"Query" name:"RoleFor"`
	RoleType   string `position:"Query" name:"RoleType"`
	Lang       string `position:"Query" name:"Lang"`
	TaskId     string `position:"Body" name:"TaskId"`
}

// TriggerProcessTaskResponse is the response struct for api TriggerProcessTask
type TriggerProcessTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTriggerProcessTaskRequest creates a request to invoke TriggerProcessTask API
func CreateTriggerProcessTaskRequest() (request *TriggerProcessTaskRequest) {
	request = &TriggerProcessTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "TriggerProcessTask", "", "")
	request.Method = requests.POST
	return
}

// CreateTriggerProcessTaskResponse creates a response to parse from TriggerProcessTask response
func CreateTriggerProcessTaskResponse() (response *TriggerProcessTaskResponse) {
	response = &TriggerProcessTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
