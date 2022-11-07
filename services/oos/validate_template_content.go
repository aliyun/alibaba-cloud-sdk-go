package oos

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

// ValidateTemplateContent invokes the oos.ValidateTemplateContent API synchronously
func (client *Client) ValidateTemplateContent(request *ValidateTemplateContentRequest) (response *ValidateTemplateContentResponse, err error) {
	response = CreateValidateTemplateContentResponse()
	err = client.DoAction(request, response)
	return
}

// ValidateTemplateContentWithChan invokes the oos.ValidateTemplateContent API asynchronously
func (client *Client) ValidateTemplateContentWithChan(request *ValidateTemplateContentRequest) (<-chan *ValidateTemplateContentResponse, <-chan error) {
	responseChan := make(chan *ValidateTemplateContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValidateTemplateContent(request)
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

// ValidateTemplateContentWithCallback invokes the oos.ValidateTemplateContent API asynchronously
func (client *Client) ValidateTemplateContentWithCallback(request *ValidateTemplateContentRequest, callback func(response *ValidateTemplateContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValidateTemplateContentResponse
		var err error
		defer close(result)
		response, err = client.ValidateTemplateContent(request)
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

// ValidateTemplateContentRequest is the request struct for api ValidateTemplateContent
type ValidateTemplateContentRequest struct {
	*requests.RpcRequest
	Content string `position:"Query" name:"Content"`
}

// ValidateTemplateContentResponse is the response struct for api ValidateTemplateContent
type ValidateTemplateContentResponse struct {
	*responses.BaseResponse
	Outputs    string `json:"Outputs" xml:"Outputs"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Parameters string `json:"Parameters" xml:"Parameters"`
	RamRole    string `json:"RamRole" xml:"RamRole"`
	Tasks      []Task `json:"Tasks" xml:"Tasks"`
}

// CreateValidateTemplateContentRequest creates a request to invoke ValidateTemplateContent API
func CreateValidateTemplateContentRequest() (request *ValidateTemplateContentRequest) {
	request = &ValidateTemplateContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ValidateTemplateContent", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateValidateTemplateContentResponse creates a response to parse from ValidateTemplateContent response
func CreateValidateTemplateContentResponse() (response *ValidateTemplateContentResponse) {
	response = &ValidateTemplateContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
