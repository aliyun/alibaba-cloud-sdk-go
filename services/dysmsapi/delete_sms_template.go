package dysmsapi

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

// DeleteSmsTemplate invokes the dysmsapi.DeleteSmsTemplate API synchronously
func (client *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
	response = CreateDeleteSmsTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSmsTemplateWithChan invokes the dysmsapi.DeleteSmsTemplate API asynchronously
func (client *Client) DeleteSmsTemplateWithChan(request *DeleteSmsTemplateRequest) (<-chan *DeleteSmsTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteSmsTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSmsTemplate(request)
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

// DeleteSmsTemplateWithCallback invokes the dysmsapi.DeleteSmsTemplate API asynchronously
func (client *Client) DeleteSmsTemplateWithCallback(request *DeleteSmsTemplateRequest, callback func(response *DeleteSmsTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSmsTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteSmsTemplate(request)
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

// DeleteSmsTemplateRequest is the request struct for api DeleteSmsTemplate
type DeleteSmsTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateCode         string           `position:"Query" name:"TemplateCode"`
}

// DeleteSmsTemplateResponse is the response struct for api DeleteSmsTemplate
type DeleteSmsTemplateResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	Message      string `json:"Message" xml:"Message"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TemplateCode string `json:"TemplateCode" xml:"TemplateCode"`
}

// CreateDeleteSmsTemplateRequest creates a request to invoke DeleteSmsTemplate API
func CreateDeleteSmsTemplateRequest() (request *DeleteSmsTemplateRequest) {
	request = &DeleteSmsTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "DeleteSmsTemplate", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSmsTemplateResponse creates a response to parse from DeleteSmsTemplate response
func CreateDeleteSmsTemplateResponse() (response *DeleteSmsTemplateResponse) {
	response = &DeleteSmsTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
