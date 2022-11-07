package dms_enterprise

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

// UpdateTaskOutput invokes the dms_enterprise.UpdateTaskOutput API synchronously
func (client *Client) UpdateTaskOutput(request *UpdateTaskOutputRequest) (response *UpdateTaskOutputResponse, err error) {
	response = CreateUpdateTaskOutputResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTaskOutputWithChan invokes the dms_enterprise.UpdateTaskOutput API asynchronously
func (client *Client) UpdateTaskOutputWithChan(request *UpdateTaskOutputRequest) (<-chan *UpdateTaskOutputResponse, <-chan error) {
	responseChan := make(chan *UpdateTaskOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTaskOutput(request)
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

// UpdateTaskOutputWithCallback invokes the dms_enterprise.UpdateTaskOutput API asynchronously
func (client *Client) UpdateTaskOutputWithCallback(request *UpdateTaskOutputRequest, callback func(response *UpdateTaskOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTaskOutputResponse
		var err error
		defer close(result)
		response, err = client.UpdateTaskOutput(request)
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

// UpdateTaskOutputRequest is the request struct for api UpdateTaskOutput
type UpdateTaskOutputRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	NodeOutput string           `position:"Query" name:"NodeOutput"`
	NodeId     string           `position:"Query" name:"NodeId"`
}

// UpdateTaskOutputResponse is the response struct for api UpdateTaskOutput
type UpdateTaskOutputResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUpdateTaskOutputRequest creates a request to invoke UpdateTaskOutput API
func CreateUpdateTaskOutputRequest() (request *UpdateTaskOutputRequest) {
	request = &UpdateTaskOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "UpdateTaskOutput", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTaskOutputResponse creates a response to parse from UpdateTaskOutput response
func CreateUpdateTaskOutputResponse() (response *UpdateTaskOutputResponse) {
	response = &UpdateTaskOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
