package iot

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

// GetThingScript invokes the iot.GetThingScript API synchronously
func (client *Client) GetThingScript(request *GetThingScriptRequest) (response *GetThingScriptResponse, err error) {
	response = CreateGetThingScriptResponse()
	err = client.DoAction(request, response)
	return
}

// GetThingScriptWithChan invokes the iot.GetThingScript API asynchronously
func (client *Client) GetThingScriptWithChan(request *GetThingScriptRequest) (<-chan *GetThingScriptResponse, <-chan error) {
	responseChan := make(chan *GetThingScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetThingScript(request)
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

// GetThingScriptWithCallback invokes the iot.GetThingScript API asynchronously
func (client *Client) GetThingScriptWithCallback(request *GetThingScriptRequest, callback func(response *GetThingScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetThingScriptResponse
		var err error
		defer close(result)
		response, err = client.GetThingScript(request)
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

// GetThingScriptRequest is the request struct for api GetThingScript
type GetThingScriptRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// GetThingScriptResponse is the response struct for api GetThingScript
type GetThingScriptResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGetThingScriptRequest creates a request to invoke GetThingScript API
func CreateGetThingScriptRequest() (request *GetThingScriptRequest) {
	request = &GetThingScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetThingScript", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetThingScriptResponse creates a response to parse from GetThingScript response
func CreateGetThingScriptResponse() (response *GetThingScriptResponse) {
	response = &GetThingScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
