package config

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

// StopConfigurationRecorder invokes the config.StopConfigurationRecorder API synchronously
func (client *Client) StopConfigurationRecorder(request *StopConfigurationRecorderRequest) (response *StopConfigurationRecorderResponse, err error) {
	response = CreateStopConfigurationRecorderResponse()
	err = client.DoAction(request, response)
	return
}

// StopConfigurationRecorderWithChan invokes the config.StopConfigurationRecorder API asynchronously
func (client *Client) StopConfigurationRecorderWithChan(request *StopConfigurationRecorderRequest) (<-chan *StopConfigurationRecorderResponse, <-chan error) {
	responseChan := make(chan *StopConfigurationRecorderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopConfigurationRecorder(request)
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

// StopConfigurationRecorderWithCallback invokes the config.StopConfigurationRecorder API asynchronously
func (client *Client) StopConfigurationRecorderWithCallback(request *StopConfigurationRecorderRequest, callback func(response *StopConfigurationRecorderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopConfigurationRecorderResponse
		var err error
		defer close(result)
		response, err = client.StopConfigurationRecorder(request)
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

// StopConfigurationRecorderRequest is the request struct for api StopConfigurationRecorder
type StopConfigurationRecorderRequest struct {
	*requests.RpcRequest
}

// StopConfigurationRecorderResponse is the response struct for api StopConfigurationRecorder
type StopConfigurationRecorderResponse struct {
	*responses.BaseResponse
	RequestId                       string `json:"RequestId" xml:"RequestId"`
	StopConfigurationRecorderResult bool   `json:"StopConfigurationRecorderResult" xml:"StopConfigurationRecorderResult"`
}

// CreateStopConfigurationRecorderRequest creates a request to invoke StopConfigurationRecorder API
func CreateStopConfigurationRecorderRequest() (request *StopConfigurationRecorderRequest) {
	request = &StopConfigurationRecorderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "StopConfigurationRecorder", "", "")
	request.Method = requests.POST
	return
}

// CreateStopConfigurationRecorderResponse creates a response to parse from StopConfigurationRecorder response
func CreateStopConfigurationRecorderResponse() (response *StopConfigurationRecorderResponse) {
	response = &StopConfigurationRecorderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
