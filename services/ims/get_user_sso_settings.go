package ims

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

// GetUserSsoSettings invokes the ims.GetUserSsoSettings API synchronously
func (client *Client) GetUserSsoSettings(request *GetUserSsoSettingsRequest) (response *GetUserSsoSettingsResponse, err error) {
	response = CreateGetUserSsoSettingsResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserSsoSettingsWithChan invokes the ims.GetUserSsoSettings API asynchronously
func (client *Client) GetUserSsoSettingsWithChan(request *GetUserSsoSettingsRequest) (<-chan *GetUserSsoSettingsResponse, <-chan error) {
	responseChan := make(chan *GetUserSsoSettingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserSsoSettings(request)
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

// GetUserSsoSettingsWithCallback invokes the ims.GetUserSsoSettings API asynchronously
func (client *Client) GetUserSsoSettingsWithCallback(request *GetUserSsoSettingsRequest, callback func(response *GetUserSsoSettingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserSsoSettingsResponse
		var err error
		defer close(result)
		response, err = client.GetUserSsoSettings(request)
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

// GetUserSsoSettingsRequest is the request struct for api GetUserSsoSettings
type GetUserSsoSettingsRequest struct {
	*requests.RpcRequest
	AkProxySuffix string `position:"Query" name:"AkProxySuffix"`
}

// GetUserSsoSettingsResponse is the response struct for api GetUserSsoSettings
type GetUserSsoSettingsResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	UserSsoSettings UserSsoSettings `json:"UserSsoSettings" xml:"UserSsoSettings"`
}

// CreateGetUserSsoSettingsRequest creates a request to invoke GetUserSsoSettings API
func CreateGetUserSsoSettingsRequest() (request *GetUserSsoSettingsRequest) {
	request = &GetUserSsoSettingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "GetUserSsoSettings", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserSsoSettingsResponse creates a response to parse from GetUserSsoSettings response
func CreateGetUserSsoSettingsResponse() (response *GetUserSsoSettingsResponse) {
	response = &GetUserSsoSettingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
