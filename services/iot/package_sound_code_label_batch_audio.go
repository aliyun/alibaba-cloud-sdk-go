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

// PackageSoundCodeLabelBatchAudio invokes the iot.PackageSoundCodeLabelBatchAudio API synchronously
func (client *Client) PackageSoundCodeLabelBatchAudio(request *PackageSoundCodeLabelBatchAudioRequest) (response *PackageSoundCodeLabelBatchAudioResponse, err error) {
	response = CreatePackageSoundCodeLabelBatchAudioResponse()
	err = client.DoAction(request, response)
	return
}

// PackageSoundCodeLabelBatchAudioWithChan invokes the iot.PackageSoundCodeLabelBatchAudio API asynchronously
func (client *Client) PackageSoundCodeLabelBatchAudioWithChan(request *PackageSoundCodeLabelBatchAudioRequest) (<-chan *PackageSoundCodeLabelBatchAudioResponse, <-chan error) {
	responseChan := make(chan *PackageSoundCodeLabelBatchAudioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PackageSoundCodeLabelBatchAudio(request)
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

// PackageSoundCodeLabelBatchAudioWithCallback invokes the iot.PackageSoundCodeLabelBatchAudio API asynchronously
func (client *Client) PackageSoundCodeLabelBatchAudioWithCallback(request *PackageSoundCodeLabelBatchAudioRequest, callback func(response *PackageSoundCodeLabelBatchAudioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PackageSoundCodeLabelBatchAudioResponse
		var err error
		defer close(result)
		response, err = client.PackageSoundCodeLabelBatchAudio(request)
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

// PackageSoundCodeLabelBatchAudioRequest is the request struct for api PackageSoundCodeLabelBatchAudio
type PackageSoundCodeLabelBatchAudioRequest struct {
	*requests.RpcRequest
	BatchCode     string `position:"Body" name:"BatchCode"`
	IotInstanceId string `position:"Body" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// PackageSoundCodeLabelBatchAudioResponse is the response struct for api PackageSoundCodeLabelBatchAudio
type PackageSoundCodeLabelBatchAudioResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         string `json:"Data" xml:"Data"`
}

// CreatePackageSoundCodeLabelBatchAudioRequest creates a request to invoke PackageSoundCodeLabelBatchAudio API
func CreatePackageSoundCodeLabelBatchAudioRequest() (request *PackageSoundCodeLabelBatchAudioRequest) {
	request = &PackageSoundCodeLabelBatchAudioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "PackageSoundCodeLabelBatchAudio", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePackageSoundCodeLabelBatchAudioResponse creates a response to parse from PackageSoundCodeLabelBatchAudio response
func CreatePackageSoundCodeLabelBatchAudioResponse() (response *PackageSoundCodeLabelBatchAudioResponse) {
	response = &PackageSoundCodeLabelBatchAudioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
