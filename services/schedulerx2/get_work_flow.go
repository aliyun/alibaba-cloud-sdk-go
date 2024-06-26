package schedulerx2

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

// GetWorkFlow invokes the schedulerx2.GetWorkFlow API synchronously
func (client *Client) GetWorkFlow(request *GetWorkFlowRequest) (response *GetWorkFlowResponse, err error) {
	response = CreateGetWorkFlowResponse()
	err = client.DoAction(request, response)
	return
}

// GetWorkFlowWithChan invokes the schedulerx2.GetWorkFlow API asynchronously
func (client *Client) GetWorkFlowWithChan(request *GetWorkFlowRequest) (<-chan *GetWorkFlowResponse, <-chan error) {
	responseChan := make(chan *GetWorkFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWorkFlow(request)
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

// GetWorkFlowWithCallback invokes the schedulerx2.GetWorkFlow API asynchronously
func (client *Client) GetWorkFlowWithCallback(request *GetWorkFlowRequest, callback func(response *GetWorkFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWorkFlowResponse
		var err error
		defer close(result)
		response, err = client.GetWorkFlow(request)
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

// GetWorkFlowRequest is the request struct for api GetWorkFlow
type GetWorkFlowRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	Namespace       string           `position:"Query" name:"Namespace"`
	WorkflowId      requests.Integer `position:"Query" name:"WorkflowId"`
}

// GetWorkFlowResponse is the response struct for api GetWorkFlow
type GetWorkFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetWorkFlowRequest creates a request to invoke GetWorkFlow API
func CreateGetWorkFlowRequest() (request *GetWorkFlowRequest) {
	request = &GetWorkFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "GetWorkFlow", "schedulerx2", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetWorkFlowResponse creates a response to parse from GetWorkFlow response
func CreateGetWorkFlowResponse() (response *GetWorkFlowResponse) {
	response = &GetWorkFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
