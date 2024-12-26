package schedulerx3

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

// OperateDesignateExecutors invokes the schedulerx3.OperateDesignateExecutors API synchronously
func (client *Client) OperateDesignateExecutors(request *OperateDesignateExecutorsRequest) (response *OperateDesignateExecutorsResponse, err error) {
	response = CreateOperateDesignateExecutorsResponse()
	err = client.DoAction(request, response)
	return
}

// OperateDesignateExecutorsWithChan invokes the schedulerx3.OperateDesignateExecutors API asynchronously
func (client *Client) OperateDesignateExecutorsWithChan(request *OperateDesignateExecutorsRequest) (<-chan *OperateDesignateExecutorsResponse, <-chan error) {
	responseChan := make(chan *OperateDesignateExecutorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateDesignateExecutors(request)
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

// OperateDesignateExecutorsWithCallback invokes the schedulerx3.OperateDesignateExecutors API asynchronously
func (client *Client) OperateDesignateExecutorsWithCallback(request *OperateDesignateExecutorsRequest, callback func(response *OperateDesignateExecutorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateDesignateExecutorsResponse
		var err error
		defer close(result)
		response, err = client.OperateDesignateExecutors(request)
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

// OperateDesignateExecutorsRequest is the request struct for api OperateDesignateExecutors
type OperateDesignateExecutorsRequest struct {
	*requests.RpcRequest
	MseSessionId  string           `position:"Query" name:"MseSessionId"`
	Transferable  requests.Boolean `position:"Body" name:"Transferable"`
	DesignateType requests.Integer `position:"Body" name:"DesignateType"`
	JobId         requests.Integer `position:"Body" name:"JobId"`
	AppName       string           `position:"Body" name:"AppName"`
	AddressList   *[]string        `position:"Body" name:"AddressList"  type:"Json"`
	ClusterId     string           `position:"Body" name:"ClusterId"`
}

// OperateDesignateExecutorsResponse is the response struct for api OperateDesignateExecutors
type OperateDesignateExecutorsResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateOperateDesignateExecutorsRequest creates a request to invoke OperateDesignateExecutors API
func CreateOperateDesignateExecutorsRequest() (request *OperateDesignateExecutorsRequest) {
	request = &OperateDesignateExecutorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SchedulerX3", "2024-06-24", "OperateDesignateExecutors", "", "")
	request.Method = requests.POST
	return
}

// CreateOperateDesignateExecutorsResponse creates a response to parse from OperateDesignateExecutors response
func CreateOperateDesignateExecutorsResponse() (response *OperateDesignateExecutorsResponse) {
	response = &OperateDesignateExecutorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
