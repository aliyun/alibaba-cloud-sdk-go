package iot

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

// ShareSpeechByCombination invokes the iot.ShareSpeechByCombination API synchronously
func (client *Client) ShareSpeechByCombination(request *ShareSpeechByCombinationRequest) (response *ShareSpeechByCombinationResponse, err error) {
	response = CreateShareSpeechByCombinationResponse()
	err = client.DoAction(request, response)
	return
}

// ShareSpeechByCombinationWithChan invokes the iot.ShareSpeechByCombination API asynchronously
func (client *Client) ShareSpeechByCombinationWithChan(request *ShareSpeechByCombinationRequest) (<-chan *ShareSpeechByCombinationResponse, <-chan error) {
	responseChan := make(chan *ShareSpeechByCombinationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShareSpeechByCombination(request)
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

// ShareSpeechByCombinationWithCallback invokes the iot.ShareSpeechByCombination API asynchronously
func (client *Client) ShareSpeechByCombinationWithCallback(request *ShareSpeechByCombinationRequest, callback func(response *ShareSpeechByCombinationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShareSpeechByCombinationResponse
		var err error
		defer close(result)
		response, err = client.ShareSpeechByCombination(request)
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

// ShareSpeechByCombinationRequest is the request struct for api ShareSpeechByCombination
type ShareSpeechByCombinationRequest struct {
	*requests.RpcRequest
	SpeechId        string    `position:"Body" name:"SpeechId"`
	AudioFormat     string    `position:"Body" name:"AudioFormat"`
	IotId           string    `position:"Body" name:"IotId"`
	CombinationList *[]string `position:"Body" name:"CombinationList"  type:"Repeated"`
	IotInstanceId   string    `position:"Body" name:"IotInstanceId"`
	ProductKey      string    `position:"Body" name:"ProductKey"`
	ApiProduct      string    `position:"Body" name:"ApiProduct"`
	ApiRevision     string    `position:"Body" name:"ApiRevision"`
	DeviceName      string    `position:"Body" name:"DeviceName"`
}

// ShareSpeechByCombinationResponse is the response struct for api ShareSpeechByCombination
type ShareSpeechByCombinationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateShareSpeechByCombinationRequest creates a request to invoke ShareSpeechByCombination API
func CreateShareSpeechByCombinationRequest() (request *ShareSpeechByCombinationRequest) {
	request = &ShareSpeechByCombinationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ShareSpeechByCombination", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateShareSpeechByCombinationResponse creates a response to parse from ShareSpeechByCombination response
func CreateShareSpeechByCombinationResponse() (response *ShareSpeechByCombinationResponse) {
	response = &ShareSpeechByCombinationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
