package mts

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

// SubmitImageCopyright invokes the mts.SubmitImageCopyright API synchronously
func (client *Client) SubmitImageCopyright(request *SubmitImageCopyrightRequest) (response *SubmitImageCopyrightResponse, err error) {
	response = CreateSubmitImageCopyrightResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitImageCopyrightWithChan invokes the mts.SubmitImageCopyright API asynchronously
func (client *Client) SubmitImageCopyrightWithChan(request *SubmitImageCopyrightRequest) (<-chan *SubmitImageCopyrightResponse, <-chan error) {
	responseChan := make(chan *SubmitImageCopyrightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitImageCopyright(request)
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

// SubmitImageCopyrightWithCallback invokes the mts.SubmitImageCopyright API asynchronously
func (client *Client) SubmitImageCopyrightWithCallback(request *SubmitImageCopyrightRequest, callback func(response *SubmitImageCopyrightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitImageCopyrightResponse
		var err error
		defer close(result)
		response, err = client.SubmitImageCopyright(request)
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

// SubmitImageCopyrightRequest is the request struct for api SubmitImageCopyright
type SubmitImageCopyrightRequest struct {
	*requests.RpcRequest
	Level   requests.Integer `position:"Query" name:"Level"`
	Message string           `position:"Query" name:"Message"`
	Params  string           `position:"Query" name:"Params"`
	Url     string           `position:"Query" name:"Url"`
	Output  string           `position:"Query" name:"Output"`
	Input   string           `position:"Query" name:"Input"`
}

// SubmitImageCopyrightResponse is the response struct for api SubmitImageCopyright
type SubmitImageCopyrightResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Message    string `json:"Message" xml:"Message"`
	StatusCode int64  `json:"StatusCode" xml:"StatusCode"`
	Data       Data   `json:"Data" xml:"Data"`
}

// CreateSubmitImageCopyrightRequest creates a request to invoke SubmitImageCopyright API
func CreateSubmitImageCopyrightRequest() (request *SubmitImageCopyrightRequest) {
	request = &SubmitImageCopyrightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitImageCopyright", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitImageCopyrightResponse creates a response to parse from SubmitImageCopyright response
func CreateSubmitImageCopyrightResponse() (response *SubmitImageCopyrightResponse) {
	response = &SubmitImageCopyrightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
