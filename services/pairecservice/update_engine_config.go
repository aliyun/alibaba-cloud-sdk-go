package pairecservice

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

// UpdateEngineConfig invokes the pairecservice.UpdateEngineConfig API synchronously
func (client *Client) UpdateEngineConfig(request *UpdateEngineConfigRequest) (response *UpdateEngineConfigResponse, err error) {
	response = CreateUpdateEngineConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEngineConfigWithChan invokes the pairecservice.UpdateEngineConfig API asynchronously
func (client *Client) UpdateEngineConfigWithChan(request *UpdateEngineConfigRequest) (<-chan *UpdateEngineConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateEngineConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEngineConfig(request)
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

// UpdateEngineConfigWithCallback invokes the pairecservice.UpdateEngineConfig API asynchronously
func (client *Client) UpdateEngineConfigWithCallback(request *UpdateEngineConfigRequest, callback func(response *UpdateEngineConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEngineConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateEngineConfig(request)
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

// UpdateEngineConfigRequest is the request struct for api UpdateEngineConfig
type UpdateEngineConfigRequest struct {
	*requests.RoaRequest
	EngineConfigId string `position:"Path" name:"EngineConfigId"`
	Body           string `position:"Body" name:"body"`
}

// UpdateEngineConfigResponse is the response struct for api UpdateEngineConfig
type UpdateEngineConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateEngineConfigRequest creates a request to invoke UpdateEngineConfig API
func CreateUpdateEngineConfigRequest() (request *UpdateEngineConfigRequest) {
	request = &UpdateEngineConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateEngineConfig", "/api/v1/engineconfigs/[EngineConfigId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateEngineConfigResponse creates a response to parse from UpdateEngineConfig response
func CreateUpdateEngineConfigResponse() (response *UpdateEngineConfigResponse) {
	response = &UpdateEngineConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
