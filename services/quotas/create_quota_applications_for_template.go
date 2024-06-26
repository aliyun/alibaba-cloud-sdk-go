package quotas

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

// CreateQuotaApplicationsForTemplate invokes the quotas.CreateQuotaApplicationsForTemplate API synchronously
func (client *Client) CreateQuotaApplicationsForTemplate(request *CreateQuotaApplicationsForTemplateRequest) (response *CreateQuotaApplicationsForTemplateResponse, err error) {
	response = CreateCreateQuotaApplicationsForTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQuotaApplicationsForTemplateWithChan invokes the quotas.CreateQuotaApplicationsForTemplate API asynchronously
func (client *Client) CreateQuotaApplicationsForTemplateWithChan(request *CreateQuotaApplicationsForTemplateRequest) (<-chan *CreateQuotaApplicationsForTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateQuotaApplicationsForTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQuotaApplicationsForTemplate(request)
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

// CreateQuotaApplicationsForTemplateWithCallback invokes the quotas.CreateQuotaApplicationsForTemplate API asynchronously
func (client *Client) CreateQuotaApplicationsForTemplateWithCallback(request *CreateQuotaApplicationsForTemplateRequest, callback func(response *CreateQuotaApplicationsForTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQuotaApplicationsForTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateQuotaApplicationsForTemplate(request)
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

// CreateQuotaApplicationsForTemplateRequest is the request struct for api CreateQuotaApplicationsForTemplate
type CreateQuotaApplicationsForTemplateRequest struct {
	*requests.RpcRequest
	Reason          string                                          `position:"Body" name:"Reason"`
	ProductCode     string                                          `position:"Body" name:"ProductCode"`
	QuotaActionCode string                                          `position:"Body" name:"QuotaActionCode"`
	DesireValue     string                                          `position:"Body" name:"DesireValue"`
	EffectiveTime   string                                          `position:"Body" name:"EffectiveTime"`
	AliyunUids      *[]string                                       `position:"Body" name:"AliyunUids"  type:"Repeated"`
	QuotaCategory   string                                          `position:"Body" name:"QuotaCategory"`
	OriginalContext string                                          `position:"Body" name:"OriginalContext"`
	ExpireTime      string                                          `position:"Body" name:"ExpireTime"`
	EnvLanguage     string                                          `position:"Body" name:"EnvLanguage"`
	NoticeType      requests.Integer                                `position:"Body" name:"NoticeType"`
	Dimensions      *[]CreateQuotaApplicationsForTemplateDimensions `position:"Body" name:"Dimensions"  type:"Repeated"`
}

// CreateQuotaApplicationsForTemplateDimensions is a repeated param struct in CreateQuotaApplicationsForTemplateRequest
type CreateQuotaApplicationsForTemplateDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateQuotaApplicationsForTemplateResponse is the response struct for api CreateQuotaApplicationsForTemplate
type CreateQuotaApplicationsForTemplateResponse struct {
	*responses.BaseResponse
	RequestId               string            `json:"RequestId" xml:"RequestId"`
	BatchQuotaApplicationId string            `json:"BatchQuotaApplicationId" xml:"BatchQuotaApplicationId"`
	AliyunUids              []string          `json:"AliyunUids" xml:"AliyunUids"`
	FailResults             []FailResultsItem `json:"FailResults" xml:"FailResults"`
}

// CreateCreateQuotaApplicationsForTemplateRequest creates a request to invoke CreateQuotaApplicationsForTemplate API
func CreateCreateQuotaApplicationsForTemplateRequest() (request *CreateQuotaApplicationsForTemplateRequest) {
	request = &CreateQuotaApplicationsForTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "CreateQuotaApplicationsForTemplate", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateQuotaApplicationsForTemplateResponse creates a response to parse from CreateQuotaApplicationsForTemplate response
func CreateCreateQuotaApplicationsForTemplateResponse() (response *CreateQuotaApplicationsForTemplateResponse) {
	response = &CreateQuotaApplicationsForTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
