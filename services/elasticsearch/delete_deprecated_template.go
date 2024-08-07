package elasticsearch

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

// DeleteDeprecatedTemplate invokes the elasticsearch.DeleteDeprecatedTemplate API synchronously
func (client *Client) DeleteDeprecatedTemplate(request *DeleteDeprecatedTemplateRequest) (response *DeleteDeprecatedTemplateResponse, err error) {
	response = CreateDeleteDeprecatedTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeprecatedTemplateWithChan invokes the elasticsearch.DeleteDeprecatedTemplate API asynchronously
func (client *Client) DeleteDeprecatedTemplateWithChan(request *DeleteDeprecatedTemplateRequest) (<-chan *DeleteDeprecatedTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteDeprecatedTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDeprecatedTemplate(request)
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

// DeleteDeprecatedTemplateWithCallback invokes the elasticsearch.DeleteDeprecatedTemplate API asynchronously
func (client *Client) DeleteDeprecatedTemplateWithCallback(request *DeleteDeprecatedTemplateRequest, callback func(response *DeleteDeprecatedTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDeprecatedTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteDeprecatedTemplate(request)
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

// DeleteDeprecatedTemplateRequest is the request struct for api DeleteDeprecatedTemplate
type DeleteDeprecatedTemplateRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Name       string `position:"Path" name:"name"`
}

// DeleteDeprecatedTemplateResponse is the response struct for api DeleteDeprecatedTemplate
type DeleteDeprecatedTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteDeprecatedTemplateRequest creates a request to invoke DeleteDeprecatedTemplate API
func CreateDeleteDeprecatedTemplateRequest() (request *DeleteDeprecatedTemplateRequest) {
	request = &DeleteDeprecatedTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteDeprecatedTemplate", "/openapi/instances/[InstanceId]/deprecated-templates/[name]", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteDeprecatedTemplateResponse creates a response to parse from DeleteDeprecatedTemplate response
func CreateDeleteDeprecatedTemplateResponse() (response *DeleteDeprecatedTemplateResponse) {
	response = &DeleteDeprecatedTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
