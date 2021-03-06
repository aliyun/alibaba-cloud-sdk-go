package dts

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

// ReplaceInstance invokes the dts.ReplaceInstance API synchronously
func (client *Client) ReplaceInstance(request *ReplaceInstanceRequest) (response *ReplaceInstanceResponse, err error) {
	response = CreateReplaceInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ReplaceInstanceWithChan invokes the dts.ReplaceInstance API asynchronously
func (client *Client) ReplaceInstanceWithChan(request *ReplaceInstanceRequest) (<-chan *ReplaceInstanceResponse, <-chan error) {
	responseChan := make(chan *ReplaceInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplaceInstance(request)
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

// ReplaceInstanceWithCallback invokes the dts.ReplaceInstance API asynchronously
func (client *Client) ReplaceInstanceWithCallback(request *ReplaceInstanceRequest, callback func(response *ReplaceInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplaceInstanceResponse
		var err error
		defer close(result)
		response, err = client.ReplaceInstance(request)
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

// ReplaceInstanceRequest is the request struct for api ReplaceInstance
type ReplaceInstanceRequest struct {
	*requests.RpcRequest
	NewInstanceId string `position:"Query" name:"NewInstanceId"`
	DtsJobId      string `position:"Query" name:"DtsJobId"`
	ChargeType    string `position:"Query" name:"ChargeType"`
}

// ReplaceInstanceResponse is the response struct for api ReplaceInstance
type ReplaceInstanceResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	DtsJobId       string `json:"DtsJobId" xml:"DtsJobId"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	ChargeType     string `json:"ChargeType" xml:"ChargeType"`
	EndTime        string `json:"EndTime" xml:"EndTime"`
}

// CreateReplaceInstanceRequest creates a request to invoke ReplaceInstance API
func CreateReplaceInstanceRequest() (request *ReplaceInstanceRequest) {
	request = &ReplaceInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ReplaceInstance", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReplaceInstanceResponse creates a response to parse from ReplaceInstance response
func CreateReplaceInstanceResponse() (response *ReplaceInstanceResponse) {
	response = &ReplaceInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
