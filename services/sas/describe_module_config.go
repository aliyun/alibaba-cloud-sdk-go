package sas

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

// DescribeModuleConfig invokes the sas.DescribeModuleConfig API synchronously
func (client *Client) DescribeModuleConfig(request *DescribeModuleConfigRequest) (response *DescribeModuleConfigResponse, err error) {
	response = CreateDescribeModuleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeModuleConfigWithChan invokes the sas.DescribeModuleConfig API asynchronously
func (client *Client) DescribeModuleConfigWithChan(request *DescribeModuleConfigRequest) (<-chan *DescribeModuleConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeModuleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeModuleConfig(request)
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

// DescribeModuleConfigWithCallback invokes the sas.DescribeModuleConfig API asynchronously
func (client *Client) DescribeModuleConfigWithCallback(request *DescribeModuleConfigRequest, callback func(response *DescribeModuleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeModuleConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeModuleConfig(request)
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

// DescribeModuleConfigRequest is the request struct for api DescribeModuleConfig
type DescribeModuleConfigRequest struct {
	*requests.RpcRequest
	SourceIp                   string `position:"Query" name:"SourceIp"`
	ResourceDirectoryAccountId string `position:"Query" name:"ResourceDirectoryAccountId"`
}

// DescribeModuleConfigResponse is the response struct for api DescribeModuleConfig
type DescribeModuleConfigResponse struct {
	*responses.BaseResponse
	HttpStatusCode   int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	Success          bool           `json:"Success" xml:"Success"`
	Count            int            `json:"Count" xml:"Count"`
	ModuleConfigList []ModuleConfig `json:"ModuleConfigList" xml:"ModuleConfigList"`
}

// CreateDescribeModuleConfigRequest creates a request to invoke DescribeModuleConfig API
func CreateDescribeModuleConfigRequest() (request *DescribeModuleConfigRequest) {
	request = &DescribeModuleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeModuleConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeModuleConfigResponse creates a response to parse from DescribeModuleConfig response
func CreateDescribeModuleConfigResponse() (response *DescribeModuleConfigResponse) {
	response = &DescribeModuleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
