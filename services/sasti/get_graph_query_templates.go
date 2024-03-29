package sasti

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

// GetGraphQueryTemplates invokes the sasti.GetGraphQueryTemplates API synchronously
func (client *Client) GetGraphQueryTemplates(request *GetGraphQueryTemplatesRequest) (response *GetGraphQueryTemplatesResponse, err error) {
	response = CreateGetGraphQueryTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// GetGraphQueryTemplatesWithChan invokes the sasti.GetGraphQueryTemplates API asynchronously
func (client *Client) GetGraphQueryTemplatesWithChan(request *GetGraphQueryTemplatesRequest) (<-chan *GetGraphQueryTemplatesResponse, <-chan error) {
	responseChan := make(chan *GetGraphQueryTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGraphQueryTemplates(request)
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

// GetGraphQueryTemplatesWithCallback invokes the sasti.GetGraphQueryTemplates API asynchronously
func (client *Client) GetGraphQueryTemplatesWithCallback(request *GetGraphQueryTemplatesRequest, callback func(response *GetGraphQueryTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGraphQueryTemplatesResponse
		var err error
		defer close(result)
		response, err = client.GetGraphQueryTemplates(request)
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

// GetGraphQueryTemplatesRequest is the request struct for api GetGraphQueryTemplates
type GetGraphQueryTemplatesRequest struct {
	*requests.RpcRequest
	ServiceUnit string `position:"Body" name:"ServiceUnit"`
	Env         string `position:"Body" name:"Env"`
}

// GetGraphQueryTemplatesResponse is the response struct for api GetGraphQueryTemplates
type GetGraphQueryTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetGraphQueryTemplatesRequest creates a request to invoke GetGraphQueryTemplates API
func CreateGetGraphQueryTemplatesRequest() (request *GetGraphQueryTemplatesRequest) {
	request = &GetGraphQueryTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sasti", "2020-05-12", "GetGraphQueryTemplates", "", "")
	request.Method = requests.POST
	return
}

// CreateGetGraphQueryTemplatesResponse creates a response to parse from GetGraphQueryTemplates response
func CreateGetGraphQueryTemplatesResponse() (response *GetGraphQueryTemplatesResponse) {
	response = &GetGraphQueryTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
