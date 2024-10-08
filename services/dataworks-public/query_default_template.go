package dataworks_public

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

// QueryDefaultTemplate invokes the dataworks_public.QueryDefaultTemplate API synchronously
func (client *Client) QueryDefaultTemplate(request *QueryDefaultTemplateRequest) (response *QueryDefaultTemplateResponse, err error) {
	response = CreateQueryDefaultTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDefaultTemplateWithChan invokes the dataworks_public.QueryDefaultTemplate API asynchronously
func (client *Client) QueryDefaultTemplateWithChan(request *QueryDefaultTemplateRequest) (<-chan *QueryDefaultTemplateResponse, <-chan error) {
	responseChan := make(chan *QueryDefaultTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDefaultTemplate(request)
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

// QueryDefaultTemplateWithCallback invokes the dataworks_public.QueryDefaultTemplate API asynchronously
func (client *Client) QueryDefaultTemplateWithCallback(request *QueryDefaultTemplateRequest, callback func(response *QueryDefaultTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDefaultTemplateResponse
		var err error
		defer close(result)
		response, err = client.QueryDefaultTemplate(request)
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

// QueryDefaultTemplateRequest is the request struct for api QueryDefaultTemplate
type QueryDefaultTemplateRequest struct {
	*requests.RpcRequest
	TenantId string `position:"Body" name:"TenantId"`
}

// QueryDefaultTemplateResponse is the response struct for api QueryDefaultTemplate
type QueryDefaultTemplateResponse struct {
	*responses.BaseResponse
}

// CreateQueryDefaultTemplateRequest creates a request to invoke QueryDefaultTemplate API
func CreateQueryDefaultTemplateRequest() (request *QueryDefaultTemplateRequest) {
	request = &QueryDefaultTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "QueryDefaultTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDefaultTemplateResponse creates a response to parse from QueryDefaultTemplate response
func CreateQueryDefaultTemplateResponse() (response *QueryDefaultTemplateResponse) {
	response = &QueryDefaultTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
