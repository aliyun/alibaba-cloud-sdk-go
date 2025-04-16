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

// GetOSSInfoForCardTemplate invokes the dysmsapi.GetOSSInfoForCardTemplate API synchronously
func (client *Client) GetOSSInfoForCardTemplate(request *GetOSSInfoForCardTemplateRequest) (response *GetOSSInfoForCardTemplateResponse, err error) {
	response = CreateGetOSSInfoForCardTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetOSSInfoForCardTemplateWithChan invokes the dysmsapi.GetOSSInfoForCardTemplate API asynchronously
func (client *Client) GetOSSInfoForCardTemplateWithChan(request *GetOSSInfoForCardTemplateRequest) (<-chan *GetOSSInfoForCardTemplateResponse, <-chan error) {
	responseChan := make(chan *GetOSSInfoForCardTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOSSInfoForCardTemplate(request)
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

// GetOSSInfoForCardTemplateWithCallback invokes the dysmsapi.GetOSSInfoForCardTemplate API asynchronously
func (client *Client) GetOSSInfoForCardTemplateWithCallback(request *GetOSSInfoForCardTemplateRequest, callback func(response *GetOSSInfoForCardTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOSSInfoForCardTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetOSSInfoForCardTemplate(request)
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

// GetOSSInfoForCardTemplateRequest is the request struct for api GetOSSInfoForCardTemplate
type GetOSSInfoForCardTemplateRequest struct {
	*requests.RpcRequest
	ProductCode string `position:"Query" name:"ProductCode"`
	BizType     string `position:"Query" name:"BizType"`
}

// GetOSSInfoForCardTemplateResponse is the response struct for api GetOSSInfoForCardTemplate
type GetOSSInfoForCardTemplateResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetOSSInfoForCardTemplateRequest creates a request to invoke GetOSSInfoForCardTemplate API
func CreateGetOSSInfoForCardTemplateRequest() (request *GetOSSInfoForCardTemplateRequest) {
	request = &GetOSSInfoForCardTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "GetOSSInfoForCardTemplate", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOSSInfoForCardTemplateResponse creates a response to parse from GetOSSInfoForCardTemplate response
func CreateGetOSSInfoForCardTemplateResponse() (response *GetOSSInfoForCardTemplateResponse) {
	response = &GetOSSInfoForCardTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
