package oceanbasepro

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

// UpdateProjectConfig invokes the oceanbasepro.UpdateProjectConfig API synchronously
func (client *Client) UpdateProjectConfig(request *UpdateProjectConfigRequest) (response *UpdateProjectConfigResponse, err error) {
	response = CreateUpdateProjectConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProjectConfigWithChan invokes the oceanbasepro.UpdateProjectConfig API asynchronously
func (client *Client) UpdateProjectConfigWithChan(request *UpdateProjectConfigRequest) (<-chan *UpdateProjectConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateProjectConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProjectConfig(request)
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

// UpdateProjectConfigWithCallback invokes the oceanbasepro.UpdateProjectConfig API asynchronously
func (client *Client) UpdateProjectConfigWithCallback(request *UpdateProjectConfigRequest, callback func(response *UpdateProjectConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProjectConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateProjectConfig(request)
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

// UpdateProjectConfigRequest is the request struct for api UpdateProjectConfig
type UpdateProjectConfigRequest struct {
	*requests.RpcRequest
	CommonTransferConfig      UpdateProjectConfigCommonTransferConfig      `position:"Body" name:"CommonTransferConfig"  type:"Struct"`
	ReverseIncrTransferConfig UpdateProjectConfigReverseIncrTransferConfig `position:"Body" name:"ReverseIncrTransferConfig"  type:"Struct"`
	FullTransferConfig        UpdateProjectConfigFullTransferConfig        `position:"Body" name:"FullTransferConfig"  type:"Struct"`
	Id                        string                                       `position:"Body" name:"Id"`
	IncrTransferConfig        UpdateProjectConfigIncrTransferConfig        `position:"Body" name:"IncrTransferConfig"  type:"Struct"`
}

// UpdateProjectConfigCommonTransferConfig is a repeated param struct in UpdateProjectConfigRequest
type UpdateProjectConfigCommonTransferConfig struct {
	SinkStoreFormat   string `name:"SinkStoreFormat"`
	SourceStoreFormat string `name:"SourceStoreFormat"`
}

// UpdateProjectConfigReverseIncrTransferConfig is a repeated param struct in UpdateProjectConfigRequest
type UpdateProjectConfigReverseIncrTransferConfig struct {
	ThrottleRps         string    `name:"ThrottleRps"`
	SupportDDLTypes     *[]string `name:"SupportDDLTypes" type:"Repeated"`
	ThrottleIOPS        string    `name:"ThrottleIOPS"`
	RecordTypeWhiteList *[]string `name:"RecordTypeWhiteList" type:"Repeated"`
	IncrSyncThreadCount string    `name:"IncrSyncThreadCount"`
}

// UpdateProjectConfigFullTransferConfig is a repeated param struct in UpdateProjectConfigRequest
type UpdateProjectConfigFullTransferConfig struct {
	ThrottleRps    string `name:"ThrottleRps"`
	WriteWorkerNum string `name:"WriteWorkerNum"`
	ReadWorkerNum  string `name:"ReadWorkerNum"`
	ThrottleIOPS   string `name:"ThrottleIOPS"`
}

// UpdateProjectConfigIncrTransferConfig is a repeated param struct in UpdateProjectConfigRequest
type UpdateProjectConfigIncrTransferConfig struct {
	ThrottleRps         string    `name:"ThrottleRps"`
	SupportDDLTypes     *[]string `name:"SupportDDLTypes" type:"Repeated"`
	ThrottleIOPS        string    `name:"ThrottleIOPS"`
	RecordTypeWhiteList *[]string `name:"RecordTypeWhiteList" type:"Repeated"`
	IncrSyncThreadCount string    `name:"IncrSyncThreadCount"`
}

// UpdateProjectConfigResponse is the response struct for api UpdateProjectConfig
type UpdateProjectConfigResponse struct {
	*responses.BaseResponse
}

// CreateUpdateProjectConfigRequest creates a request to invoke UpdateProjectConfig API
func CreateUpdateProjectConfigRequest() (request *UpdateProjectConfigRequest) {
	request = &UpdateProjectConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "UpdateProjectConfig", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateProjectConfigResponse creates a response to parse from UpdateProjectConfig response
func CreateUpdateProjectConfigResponse() (response *UpdateProjectConfigResponse) {
	response = &UpdateProjectConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
