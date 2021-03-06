package arms

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

// UpdateAlertTemplate invokes the arms.UpdateAlertTemplate API synchronously
func (client *Client) UpdateAlertTemplate(request *UpdateAlertTemplateRequest) (response *UpdateAlertTemplateResponse, err error) {
	response = CreateUpdateAlertTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAlertTemplateWithChan invokes the arms.UpdateAlertTemplate API asynchronously
func (client *Client) UpdateAlertTemplateWithChan(request *UpdateAlertTemplateRequest) (<-chan *UpdateAlertTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateAlertTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAlertTemplate(request)
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

// UpdateAlertTemplateWithCallback invokes the arms.UpdateAlertTemplate API asynchronously
func (client *Client) UpdateAlertTemplateWithCallback(request *UpdateAlertTemplateRequest, callback func(response *UpdateAlertTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAlertTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateAlertTemplate(request)
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

// UpdateAlertTemplateRequest is the request struct for api UpdateAlertTemplate
type UpdateAlertTemplateRequest struct {
	*requests.RpcRequest
	Annotations      string           `position:"Query" name:"Annotations"`
	Rule             string           `position:"Query" name:"Rule"`
	ProxyUserId      string           `position:"Query" name:"ProxyUserId"`
	Type             string           `position:"Query" name:"Type"`
	Message          string           `position:"Query" name:"Message"`
	Labels           string           `position:"Query" name:"Labels"`
	Name             string           `position:"Query" name:"Name"`
	Id               requests.Integer `position:"Query" name:"Id"`
	MatchExpressions string           `position:"Query" name:"MatchExpressions"`
	Status           requests.Boolean `position:"Query" name:"Status"`
}

// UpdateAlertTemplateResponse is the response struct for api UpdateAlertTemplate
type UpdateAlertTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateAlertTemplateRequest creates a request to invoke UpdateAlertTemplate API
func CreateUpdateAlertTemplateRequest() (request *UpdateAlertTemplateRequest) {
	request = &UpdateAlertTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateAlertTemplate", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAlertTemplateResponse creates a response to parse from UpdateAlertTemplate response
func CreateUpdateAlertTemplateResponse() (response *UpdateAlertTemplateResponse) {
	response = &UpdateAlertTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
