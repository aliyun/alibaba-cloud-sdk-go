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

// QuerySmsTemplate invokes the dysmsapi.QuerySmsTemplate API synchronously
func (client *Client) QuerySmsTemplate(request *QuerySmsTemplateRequest) (response *QuerySmsTemplateResponse, err error) {
	response = CreateQuerySmsTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySmsTemplateWithChan invokes the dysmsapi.QuerySmsTemplate API asynchronously
func (client *Client) QuerySmsTemplateWithChan(request *QuerySmsTemplateRequest) (<-chan *QuerySmsTemplateResponse, <-chan error) {
	responseChan := make(chan *QuerySmsTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySmsTemplate(request)
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

// QuerySmsTemplateWithCallback invokes the dysmsapi.QuerySmsTemplate API asynchronously
func (client *Client) QuerySmsTemplateWithCallback(request *QuerySmsTemplateRequest, callback func(response *QuerySmsTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySmsTemplateResponse
		var err error
		defer close(result)
		response, err = client.QuerySmsTemplate(request)
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

// QuerySmsTemplateRequest is the request struct for api QuerySmsTemplate
type QuerySmsTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateCode         string           `position:"Query" name:"TemplateCode"`
}

// QuerySmsTemplateResponse is the response struct for api QuerySmsTemplate
type QuerySmsTemplateResponse struct {
	*responses.BaseResponse
	TemplateContent string `json:"TemplateContent" xml:"TemplateContent"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	TemplateCode    string `json:"TemplateCode" xml:"TemplateCode"`
	TemplateStatus  int    `json:"TemplateStatus" xml:"TemplateStatus"`
	Code            string `json:"Code" xml:"Code"`
	TemplateType    int    `json:"TemplateType" xml:"TemplateType"`
	Message         string `json:"Message" xml:"Message"`
	TemplateName    string `json:"TemplateName" xml:"TemplateName"`
	CreateDate      string `json:"CreateDate" xml:"CreateDate"`
	Reason          string `json:"Reason" xml:"Reason"`
}

// CreateQuerySmsTemplateRequest creates a request to invoke QuerySmsTemplate API
func CreateQuerySmsTemplateRequest() (request *QuerySmsTemplateRequest) {
	request = &QuerySmsTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "QuerySmsTemplate", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySmsTemplateResponse creates a response to parse from QuerySmsTemplate response
func CreateQuerySmsTemplateResponse() (response *QuerySmsTemplateResponse) {
	response = &QuerySmsTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
