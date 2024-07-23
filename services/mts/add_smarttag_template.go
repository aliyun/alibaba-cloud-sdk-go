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

// AddSmarttagTemplate invokes the mts.AddSmarttagTemplate API synchronously
func (client *Client) AddSmarttagTemplate(request *AddSmarttagTemplateRequest) (response *AddSmarttagTemplateResponse, err error) {
	response = CreateAddSmarttagTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// AddSmarttagTemplateWithChan invokes the mts.AddSmarttagTemplate API asynchronously
func (client *Client) AddSmarttagTemplateWithChan(request *AddSmarttagTemplateRequest) (<-chan *AddSmarttagTemplateResponse, <-chan error) {
	responseChan := make(chan *AddSmarttagTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSmarttagTemplate(request)
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

// AddSmarttagTemplateWithCallback invokes the mts.AddSmarttagTemplate API asynchronously
func (client *Client) AddSmarttagTemplateWithCallback(request *AddSmarttagTemplateRequest, callback func(response *AddSmarttagTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSmarttagTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddSmarttagTemplate(request)
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

// AddSmarttagTemplateRequest is the request struct for api AddSmarttagTemplate
type AddSmarttagTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	KnowledgeConfig        string           `position:"Query" name:"KnowledgeConfig"`
	Industry               string           `position:"Query" name:"Industry"`
	LabelVersion           string           `position:"Query" name:"LabelVersion"`
	Scene                  string           `position:"Query" name:"Scene"`
	FaceCustomParamsConfig string           `position:"Query" name:"FaceCustomParamsConfig"`
	TemplateName           string           `position:"Query" name:"TemplateName"`
	IsDefault              requests.Boolean `position:"Query" name:"IsDefault"`
	FaceCategoryIds        string           `position:"Query" name:"FaceCategoryIds"`
	KeywordConfig          string           `position:"Query" name:"KeywordConfig"`
	LandmarkGroupIds       string           `position:"Query" name:"LandmarkGroupIds"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	ObjectGroupIds         string           `position:"Query" name:"ObjectGroupIds"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	AnalyseTypes           string           `position:"Query" name:"AnalyseTypes"`
	LabelType              string           `position:"Query" name:"LabelType"`
}

// AddSmarttagTemplateResponse is the response struct for api AddSmarttagTemplate
type AddSmarttagTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateAddSmarttagTemplateRequest creates a request to invoke AddSmarttagTemplate API
func CreateAddSmarttagTemplateRequest() (request *AddSmarttagTemplateRequest) {
	request = &AddSmarttagTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddSmarttagTemplate", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddSmarttagTemplateResponse creates a response to parse from AddSmarttagTemplate response
func CreateAddSmarttagTemplateResponse() (response *AddSmarttagTemplateResponse) {
	response = &AddSmarttagTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
