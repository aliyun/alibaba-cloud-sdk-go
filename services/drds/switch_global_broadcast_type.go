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

// SwitchGlobalBroadcastType invokes the drds.SwitchGlobalBroadcastType API synchronously
func (client *Client) SwitchGlobalBroadcastType(request *SwitchGlobalBroadcastTypeRequest) (response *SwitchGlobalBroadcastTypeResponse, err error) {
	response = CreateSwitchGlobalBroadcastTypeResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchGlobalBroadcastTypeWithChan invokes the drds.SwitchGlobalBroadcastType API asynchronously
func (client *Client) SwitchGlobalBroadcastTypeWithChan(request *SwitchGlobalBroadcastTypeRequest) (<-chan *SwitchGlobalBroadcastTypeResponse, <-chan error) {
	responseChan := make(chan *SwitchGlobalBroadcastTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchGlobalBroadcastType(request)
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

// SwitchGlobalBroadcastTypeWithCallback invokes the drds.SwitchGlobalBroadcastType API asynchronously
func (client *Client) SwitchGlobalBroadcastTypeWithCallback(request *SwitchGlobalBroadcastTypeRequest, callback func(response *SwitchGlobalBroadcastTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchGlobalBroadcastTypeResponse
		var err error
		defer close(result)
		response, err = client.SwitchGlobalBroadcastType(request)
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

// SwitchGlobalBroadcastTypeRequest is the request struct for api SwitchGlobalBroadcastType
type SwitchGlobalBroadcastTypeRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// SwitchGlobalBroadcastTypeResponse is the response struct for api SwitchGlobalBroadcastType
type SwitchGlobalBroadcastTypeResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSwitchGlobalBroadcastTypeRequest creates a request to invoke SwitchGlobalBroadcastType API
func CreateSwitchGlobalBroadcastTypeRequest() (request *SwitchGlobalBroadcastTypeRequest) {
	request = &SwitchGlobalBroadcastTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SwitchGlobalBroadcastType", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchGlobalBroadcastTypeResponse creates a response to parse from SwitchGlobalBroadcastType response
func CreateSwitchGlobalBroadcastTypeResponse() (response *SwitchGlobalBroadcastTypeResponse) {
	response = &SwitchGlobalBroadcastTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
