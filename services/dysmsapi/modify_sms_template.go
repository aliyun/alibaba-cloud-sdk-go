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

// ModifySmsTemplate invokes the dysmsapi.ModifySmsTemplate API synchronously
func (client *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
	response = CreateModifySmsTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySmsTemplateWithChan invokes the dysmsapi.ModifySmsTemplate API asynchronously
func (client *Client) ModifySmsTemplateWithChan(request *ModifySmsTemplateRequest) (<-chan *ModifySmsTemplateResponse, <-chan error) {
	responseChan := make(chan *ModifySmsTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySmsTemplate(request)
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

// ModifySmsTemplateWithCallback invokes the dysmsapi.ModifySmsTemplate API asynchronously
func (client *Client) ModifySmsTemplateWithCallback(request *ModifySmsTemplateRequest, callback func(response *ModifySmsTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySmsTemplateResponse
		var err error
		defer close(result)
		response, err = client.ModifySmsTemplate(request)
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

// ModifySmsTemplateRequest is the request struct for api ModifySmsTemplate
type ModifySmsTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Remark               string           `position:"Query" name:"Remark"`
	TemplateType         requests.Integer `position:"Query" name:"TemplateType"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateContent      string           `position:"Query" name:"TemplateContent"`
	TemplateCode         string           `position:"Query" name:"TemplateCode"`
}

// ModifySmsTemplateResponse is the response struct for api ModifySmsTemplate
type ModifySmsTemplateResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	Message      string `json:"Message" xml:"Message"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TemplateCode string `json:"TemplateCode" xml:"TemplateCode"`
}

// CreateModifySmsTemplateRequest creates a request to invoke ModifySmsTemplate API
func CreateModifySmsTemplateRequest() (request *ModifySmsTemplateRequest) {
	request = &ModifySmsTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "ModifySmsTemplate", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySmsTemplateResponse creates a response to parse from ModifySmsTemplate response
func CreateModifySmsTemplateResponse() (response *ModifySmsTemplateResponse) {
	response = &ModifySmsTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
