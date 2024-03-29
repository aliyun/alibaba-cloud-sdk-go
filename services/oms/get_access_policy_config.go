package oms

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

// GetAccessPolicyConfig invokes the oms.GetAccessPolicyConfig API synchronously
func (client *Client) GetAccessPolicyConfig(request *GetAccessPolicyConfigRequest) (response *GetAccessPolicyConfigResponse, err error) {
	response = CreateGetAccessPolicyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccessPolicyConfigWithChan invokes the oms.GetAccessPolicyConfig API asynchronously
func (client *Client) GetAccessPolicyConfigWithChan(request *GetAccessPolicyConfigRequest) (<-chan *GetAccessPolicyConfigResponse, <-chan error) {
	responseChan := make(chan *GetAccessPolicyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccessPolicyConfig(request)
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

// GetAccessPolicyConfigWithCallback invokes the oms.GetAccessPolicyConfig API asynchronously
func (client *Client) GetAccessPolicyConfigWithCallback(request *GetAccessPolicyConfigRequest, callback func(response *GetAccessPolicyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccessPolicyConfigResponse
		var err error
		defer close(result)
		response, err = client.GetAccessPolicyConfig(request)
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

// GetAccessPolicyConfigRequest is the request struct for api GetAccessPolicyConfig
type GetAccessPolicyConfigRequest struct {
	*requests.RpcRequest
	CompressEnable requests.Boolean `position:"Query" name:"CompressEnable"`
	AliUid         string           `position:"Query" name:"AliUid"`
}

// GetAccessPolicyConfigResponse is the response struct for api GetAccessPolicyConfig
type GetAccessPolicyConfigResponse struct {
	*responses.BaseResponse
	Compressed bool   `json:"Compressed" xml:"Compressed"`
	AliUid     string `json:"AliUid" xml:"AliUid"`
	Data       string `json:"Data" xml:"Data"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateGetAccessPolicyConfigRequest creates a request to invoke GetAccessPolicyConfig API
func CreateGetAccessPolicyConfigRequest() (request *GetAccessPolicyConfigRequest) {
	request = &GetAccessPolicyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "GetAccessPolicyConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAccessPolicyConfigResponse creates a response to parse from GetAccessPolicyConfig response
func CreateGetAccessPolicyConfigResponse() (response *GetAccessPolicyConfigResponse) {
	response = &GetAccessPolicyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
