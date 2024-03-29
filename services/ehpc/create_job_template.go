package ehpc

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

// CreateJobTemplate invokes the ehpc.CreateJobTemplate API synchronously
func (client *Client) CreateJobTemplate(request *CreateJobTemplateRequest) (response *CreateJobTemplateResponse, err error) {
	response = CreateCreateJobTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateJobTemplateWithChan invokes the ehpc.CreateJobTemplate API asynchronously
func (client *Client) CreateJobTemplateWithChan(request *CreateJobTemplateRequest) (<-chan *CreateJobTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateJobTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateJobTemplate(request)
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

// CreateJobTemplateWithCallback invokes the ehpc.CreateJobTemplate API asynchronously
func (client *Client) CreateJobTemplateWithCallback(request *CreateJobTemplateRequest, callback func(response *CreateJobTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateJobTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateJobTemplate(request)
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

// CreateJobTemplateRequest is the request struct for api CreateJobTemplate
type CreateJobTemplateRequest struct {
	*requests.RpcRequest
	StderrRedirectPath string           `position:"Query" name:"StderrRedirectPath"`
	ClockTime          string           `position:"Query" name:"ClockTime"`
	CommandLine        string           `position:"Query" name:"CommandLine"`
	ArrayRequest       string           `position:"Query" name:"ArrayRequest"`
	UnzipCmd           string           `position:"Query" name:"UnzipCmd"`
	PackagePath        string           `position:"Query" name:"PackagePath"`
	Mem                string           `position:"Query" name:"Mem"`
	StdoutRedirectPath string           `position:"Query" name:"StdoutRedirectPath"`
	Variables          string           `position:"Query" name:"Variables"`
	RunasUser          string           `position:"Query" name:"RunasUser"`
	ReRunable          requests.Boolean `position:"Query" name:"ReRunable"`
	Thread             requests.Integer `position:"Query" name:"Thread"`
	Priority           requests.Integer `position:"Query" name:"Priority"`
	Gpu                requests.Integer `position:"Query" name:"Gpu"`
	WithUnzipCmd       requests.Boolean `position:"Query" name:"WithUnzipCmd"`
	Node               requests.Integer `position:"Query" name:"Node"`
	Task               requests.Integer `position:"Query" name:"Task"`
	InputFileUrl       string           `position:"Query" name:"InputFileUrl"`
	Name               string           `position:"Query" name:"Name"`
	Queue              string           `position:"Query" name:"Queue"`
}

// CreateJobTemplateResponse is the response struct for api CreateJobTemplate
type CreateJobTemplateResponse struct {
	*responses.BaseResponse
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateJobTemplateRequest creates a request to invoke CreateJobTemplate API
func CreateCreateJobTemplateRequest() (request *CreateJobTemplateRequest) {
	request = &CreateJobTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "CreateJobTemplate", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateCreateJobTemplateResponse creates a response to parse from CreateJobTemplate response
func CreateCreateJobTemplateResponse() (response *CreateJobTemplateResponse) {
	response = &CreateJobTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
