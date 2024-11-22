package qianzhou

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

// AICreateSessionMessage invokes the qianzhou.AICreateSessionMessage API synchronously
func (client *Client) AICreateSessionMessage(request *AICreateSessionMessageRequest) (response *AICreateSessionMessageResponse, err error) {
	response = CreateAICreateSessionMessageResponse()
	err = client.DoAction(request, response)
	return
}

// AICreateSessionMessageWithChan invokes the qianzhou.AICreateSessionMessage API asynchronously
func (client *Client) AICreateSessionMessageWithChan(request *AICreateSessionMessageRequest) (<-chan *AICreateSessionMessageResponse, <-chan error) {
	responseChan := make(chan *AICreateSessionMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AICreateSessionMessage(request)
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

// AICreateSessionMessageWithCallback invokes the qianzhou.AICreateSessionMessage API asynchronously
func (client *Client) AICreateSessionMessageWithCallback(request *AICreateSessionMessageRequest, callback func(response *AICreateSessionMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AICreateSessionMessageResponse
		var err error
		defer close(result)
		response, err = client.AICreateSessionMessage(request)
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

// AICreateSessionMessageRequest is the request struct for api AICreateSessionMessage
type AICreateSessionMessageRequest struct {
	*requests.RoaRequest
	EmployeeId string                     `position:"Query" name:"employee_id"`
	Body       AICreateSessionMessageBody `position:"Body" name:"body"  type:"Struct"`
}

// AICreateSessionMessageBody is a repeated param struct in AICreateSessionMessageRequest
type AICreateSessionMessageBody struct {
	Context   AICreateSessionMessageBodyContext `name:"context" type:"Struct"`
	SessionId string                            `name:"session_id"`
	Language  string                            `name:"language"`
	Type      string                            `name:"type"`
	Message   string                            `name:"message"`
}

// AICreateSessionMessageBodyContext is a repeated param struct in AICreateSessionMessageRequest
type AICreateSessionMessageBodyContext struct {
	ClusterId string `name:"cluster_id"`
	Kind      string `name:"kind"`
	Namespace string `name:"namespace"`
	Name      string `name:"name"`
	Error     string `name:"error"`
	Uuid      string `name:"uuid"`
}

// AICreateSessionMessageResponse is the response struct for api AICreateSessionMessage
type AICreateSessionMessageResponse struct {
	*responses.BaseResponse
	Answer    string `json:"answer" xml:"answer"`
	SessionId string `json:"session_id" xml:"session_id"`
	RequestId string `json:"requestId" xml:"requestId"`
	Code      int64  `json:"code" xml:"code"`
	Data      string `json:"data" xml:"data"`
	Msg       string `json:"msg" xml:"msg"`
}

// CreateAICreateSessionMessageRequest creates a request to invoke AICreateSessionMessage API
func CreateAICreateSessionMessageRequest() (request *AICreateSessionMessageRequest) {
	request = &AICreateSessionMessageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("qianzhou", "2021-11-11", "AICreateSessionMessage", "/popapi/AICreateSessionMessage", "", "")
	request.Method = requests.POST
	return
}

// CreateAICreateSessionMessageResponse creates a response to parse from AICreateSessionMessage response
func CreateAICreateSessionMessageResponse() (response *AICreateSessionMessageResponse) {
	response = &AICreateSessionMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
