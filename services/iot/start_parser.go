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

// StartParser invokes the iot.StartParser API synchronously
func (client *Client) StartParser(request *StartParserRequest) (response *StartParserResponse, err error) {
	response = CreateStartParserResponse()
	err = client.DoAction(request, response)
	return
}

// StartParserWithChan invokes the iot.StartParser API asynchronously
func (client *Client) StartParserWithChan(request *StartParserRequest) (<-chan *StartParserResponse, <-chan error) {
	responseChan := make(chan *StartParserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartParser(request)
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

// StartParserWithCallback invokes the iot.StartParser API asynchronously
func (client *Client) StartParserWithCallback(request *StartParserRequest, callback func(response *StartParserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartParserResponse
		var err error
		defer close(result)
		response, err = client.StartParser(request)
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

// StartParserRequest is the request struct for api StartParser
type StartParserRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ParserId      requests.Integer `position:"Query" name:"ParserId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// StartParserResponse is the response struct for api StartParser
type StartParserResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateStartParserRequest creates a request to invoke StartParser API
func CreateStartParserRequest() (request *StartParserRequest) {
	request = &StartParserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "StartParser", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartParserResponse creates a response to parse from StartParser response
func CreateStartParserResponse() (response *StartParserResponse) {
	response = &StartParserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
