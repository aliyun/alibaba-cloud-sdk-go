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

// GetSecurityPreference invokes the ims.GetSecurityPreference API synchronously
func (client *Client) GetSecurityPreference(request *GetSecurityPreferenceRequest) (response *GetSecurityPreferenceResponse, err error) {
	response = CreateGetSecurityPreferenceResponse()
	err = client.DoAction(request, response)
	return
}

// GetSecurityPreferenceWithChan invokes the ims.GetSecurityPreference API asynchronously
func (client *Client) GetSecurityPreferenceWithChan(request *GetSecurityPreferenceRequest) (<-chan *GetSecurityPreferenceResponse, <-chan error) {
	responseChan := make(chan *GetSecurityPreferenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSecurityPreference(request)
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

// GetSecurityPreferenceWithCallback invokes the ims.GetSecurityPreference API asynchronously
func (client *Client) GetSecurityPreferenceWithCallback(request *GetSecurityPreferenceRequest, callback func(response *GetSecurityPreferenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSecurityPreferenceResponse
		var err error
		defer close(result)
		response, err = client.GetSecurityPreference(request)
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

// GetSecurityPreferenceRequest is the request struct for api GetSecurityPreference
type GetSecurityPreferenceRequest struct {
	*requests.RpcRequest
	AkProxySuffix string `position:"Query" name:"AkProxySuffix"`
}

// GetSecurityPreferenceResponse is the response struct for api GetSecurityPreference
type GetSecurityPreferenceResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	SecurityPreference SecurityPreference `json:"SecurityPreference" xml:"SecurityPreference"`
}

// CreateGetSecurityPreferenceRequest creates a request to invoke GetSecurityPreference API
func CreateGetSecurityPreferenceRequest() (request *GetSecurityPreferenceRequest) {
	request = &GetSecurityPreferenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "GetSecurityPreference", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSecurityPreferenceResponse creates a response to parse from GetSecurityPreference response
func CreateGetSecurityPreferenceResponse() (response *GetSecurityPreferenceResponse) {
	response = &GetSecurityPreferenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
