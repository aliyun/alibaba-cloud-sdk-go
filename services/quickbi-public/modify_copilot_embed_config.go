package quickbi_public

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

// ModifyCopilotEmbedConfig invokes the quickbi_public.ModifyCopilotEmbedConfig API synchronously
func (client *Client) ModifyCopilotEmbedConfig(request *ModifyCopilotEmbedConfigRequest) (response *ModifyCopilotEmbedConfigResponse, err error) {
	response = CreateModifyCopilotEmbedConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCopilotEmbedConfigWithChan invokes the quickbi_public.ModifyCopilotEmbedConfig API asynchronously
func (client *Client) ModifyCopilotEmbedConfigWithChan(request *ModifyCopilotEmbedConfigRequest) (<-chan *ModifyCopilotEmbedConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyCopilotEmbedConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCopilotEmbedConfig(request)
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

// ModifyCopilotEmbedConfigWithCallback invokes the quickbi_public.ModifyCopilotEmbedConfig API asynchronously
func (client *Client) ModifyCopilotEmbedConfigWithCallback(request *ModifyCopilotEmbedConfigRequest, callback func(response *ModifyCopilotEmbedConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCopilotEmbedConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyCopilotEmbedConfig(request)
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

// ModifyCopilotEmbedConfigRequest is the request struct for api ModifyCopilotEmbedConfig
type ModifyCopilotEmbedConfigRequest struct {
	*requests.RpcRequest
	CopilotId   string `position:"Query" name:"CopilotId"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	DataRange   string `position:"Query" name:"DataRange"`
	ModuleName  string `position:"Query" name:"ModuleName"`
	SignType    string `position:"Query" name:"SignType"`
	AgentName   string `position:"Query" name:"AgentName"`
}

// ModifyCopilotEmbedConfigResponse is the response struct for api ModifyCopilotEmbedConfig
type ModifyCopilotEmbedConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateModifyCopilotEmbedConfigRequest creates a request to invoke ModifyCopilotEmbedConfig API
func CreateModifyCopilotEmbedConfigRequest() (request *ModifyCopilotEmbedConfigRequest) {
	request = &ModifyCopilotEmbedConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ModifyCopilotEmbedConfig", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCopilotEmbedConfigResponse creates a response to parse from ModifyCopilotEmbedConfig response
func CreateModifyCopilotEmbedConfigResponse() (response *ModifyCopilotEmbedConfigResponse) {
	response = &ModifyCopilotEmbedConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
