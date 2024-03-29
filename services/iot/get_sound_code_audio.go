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

// GetSoundCodeAudio invokes the iot.GetSoundCodeAudio API synchronously
func (client *Client) GetSoundCodeAudio(request *GetSoundCodeAudioRequest) (response *GetSoundCodeAudioResponse, err error) {
	response = CreateGetSoundCodeAudioResponse()
	err = client.DoAction(request, response)
	return
}

// GetSoundCodeAudioWithChan invokes the iot.GetSoundCodeAudio API asynchronously
func (client *Client) GetSoundCodeAudioWithChan(request *GetSoundCodeAudioRequest) (<-chan *GetSoundCodeAudioResponse, <-chan error) {
	responseChan := make(chan *GetSoundCodeAudioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSoundCodeAudio(request)
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

// GetSoundCodeAudioWithCallback invokes the iot.GetSoundCodeAudio API asynchronously
func (client *Client) GetSoundCodeAudioWithCallback(request *GetSoundCodeAudioRequest, callback func(response *GetSoundCodeAudioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSoundCodeAudioResponse
		var err error
		defer close(result)
		response, err = client.GetSoundCodeAudio(request)
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

// GetSoundCodeAudioRequest is the request struct for api GetSoundCodeAudio
type GetSoundCodeAudioRequest struct {
	*requests.RpcRequest
	IotInstanceId string    `position:"Body" name:"IotInstanceId"`
	SoundCodeList *[]string `position:"Body" name:"SoundCodeList"  type:"Repeated"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// GetSoundCodeAudioResponse is the response struct for api GetSoundCodeAudio
type GetSoundCodeAudioResponse struct {
	*responses.BaseResponse
	RequestId    string                  `json:"RequestId" xml:"RequestId"`
	Success      bool                    `json:"Success" xml:"Success"`
	Code         string                  `json:"Code" xml:"Code"`
	ErrorMessage string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetSoundCodeAudio `json:"Data" xml:"Data"`
}

// CreateGetSoundCodeAudioRequest creates a request to invoke GetSoundCodeAudio API
func CreateGetSoundCodeAudioRequest() (request *GetSoundCodeAudioRequest) {
	request = &GetSoundCodeAudioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetSoundCodeAudio", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSoundCodeAudioResponse creates a response to parse from GetSoundCodeAudio response
func CreateGetSoundCodeAudioResponse() (response *GetSoundCodeAudioResponse) {
	response = &GetSoundCodeAudioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
