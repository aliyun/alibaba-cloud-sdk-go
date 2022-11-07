package ocr

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

// RecognizeIndonesiaIdentityCard invokes the ocr.RecognizeIndonesiaIdentityCard API synchronously
func (client *Client) RecognizeIndonesiaIdentityCard(request *RecognizeIndonesiaIdentityCardRequest) (response *RecognizeIndonesiaIdentityCardResponse, err error) {
	response = CreateRecognizeIndonesiaIdentityCardResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeIndonesiaIdentityCardWithChan invokes the ocr.RecognizeIndonesiaIdentityCard API asynchronously
func (client *Client) RecognizeIndonesiaIdentityCardWithChan(request *RecognizeIndonesiaIdentityCardRequest) (<-chan *RecognizeIndonesiaIdentityCardResponse, <-chan error) {
	responseChan := make(chan *RecognizeIndonesiaIdentityCardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeIndonesiaIdentityCard(request)
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

// RecognizeIndonesiaIdentityCardWithCallback invokes the ocr.RecognizeIndonesiaIdentityCard API asynchronously
func (client *Client) RecognizeIndonesiaIdentityCardWithCallback(request *RecognizeIndonesiaIdentityCardRequest, callback func(response *RecognizeIndonesiaIdentityCardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeIndonesiaIdentityCardResponse
		var err error
		defer close(result)
		response, err = client.RecognizeIndonesiaIdentityCard(request)
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

// RecognizeIndonesiaIdentityCardRequest is the request struct for api RecognizeIndonesiaIdentityCard
type RecognizeIndonesiaIdentityCardRequest struct {
	*requests.RpcRequest
	ImageUrl string `position:"Body" name:"ImageUrl"`
}

// RecognizeIndonesiaIdentityCardResponse is the response struct for api RecognizeIndonesiaIdentityCard
type RecognizeIndonesiaIdentityCardResponse struct {
	*responses.BaseResponse
	RequestId string                               `json:"RequestId" xml:"RequestId"`
	Code      string                               `json:"Code" xml:"Code"`
	Message   string                               `json:"Message" xml:"Message"`
	Data      DataInRecognizeIndonesiaIdentityCard `json:"Data" xml:"Data"`
}

// CreateRecognizeIndonesiaIdentityCardRequest creates a request to invoke RecognizeIndonesiaIdentityCard API
func CreateRecognizeIndonesiaIdentityCardRequest() (request *RecognizeIndonesiaIdentityCardRequest) {
	request = &RecognizeIndonesiaIdentityCardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeIndonesiaIdentityCard", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeIndonesiaIdentityCardResponse creates a response to parse from RecognizeIndonesiaIdentityCard response
func CreateRecognizeIndonesiaIdentityCardResponse() (response *RecognizeIndonesiaIdentityCardResponse) {
	response = &RecognizeIndonesiaIdentityCardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
