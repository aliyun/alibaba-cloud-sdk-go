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

// CreateSmsTemplate invokes the dysmsapi.CreateSmsTemplate API synchronously
func (client *Client) CreateSmsTemplate(request *CreateSmsTemplateRequest) (response *CreateSmsTemplateResponse, err error) {
	response = CreateCreateSmsTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSmsTemplateWithChan invokes the dysmsapi.CreateSmsTemplate API asynchronously
func (client *Client) CreateSmsTemplateWithChan(request *CreateSmsTemplateRequest) (<-chan *CreateSmsTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateSmsTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSmsTemplate(request)
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

// CreateSmsTemplateWithCallback invokes the dysmsapi.CreateSmsTemplate API asynchronously
func (client *Client) CreateSmsTemplateWithCallback(request *CreateSmsTemplateRequest, callback func(response *CreateSmsTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSmsTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateSmsTemplate(request)
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

// CreateSmsTemplateRequest is the request struct for api CreateSmsTemplate
type CreateSmsTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ApplySceneContent    string           `position:"Query" name:"ApplySceneContent"`
	MoreData             string           `position:"Query" name:"MoreData"`
	Remark               string           `position:"Query" name:"Remark"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	RelatedSignName      string           `position:"Query" name:"RelatedSignName"`
	ApplySource          string           `position:"Query" name:"ApplySource"`
	TemplateType         requests.Integer `position:"Query" name:"TemplateType"`
	TemplateRule         string           `position:"Query" name:"TemplateRule"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	RouteName            string           `position:"Query" name:"RouteName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateContent      string           `position:"Query" name:"TemplateContent"`
	IntlType             requests.Integer `position:"Query" name:"IntlType"`
}

// CreateSmsTemplateResponse is the response struct for api CreateSmsTemplate
type CreateSmsTemplateResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Code         string `json:"Code" xml:"Code"`
	Message      string `json:"Message" xml:"Message"`
	TemplateName string `json:"TemplateName" xml:"TemplateName"`
	TemplateCode string `json:"TemplateCode" xml:"TemplateCode"`
	OrderId      string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateSmsTemplateRequest creates a request to invoke CreateSmsTemplate API
func CreateCreateSmsTemplateRequest() (request *CreateSmsTemplateRequest) {
	request = &CreateSmsTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "CreateSmsTemplate", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSmsTemplateResponse creates a response to parse from CreateSmsTemplate response
func CreateCreateSmsTemplateResponse() (response *CreateSmsTemplateResponse) {
	response = &CreateSmsTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
