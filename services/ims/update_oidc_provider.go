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

// UpdateOIDCProvider invokes the ims.UpdateOIDCProvider API synchronously
func (client *Client) UpdateOIDCProvider(request *UpdateOIDCProviderRequest) (response *UpdateOIDCProviderResponse, err error) {
	response = CreateUpdateOIDCProviderResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOIDCProviderWithChan invokes the ims.UpdateOIDCProvider API asynchronously
func (client *Client) UpdateOIDCProviderWithChan(request *UpdateOIDCProviderRequest) (<-chan *UpdateOIDCProviderResponse, <-chan error) {
	responseChan := make(chan *UpdateOIDCProviderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOIDCProvider(request)
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

// UpdateOIDCProviderWithCallback invokes the ims.UpdateOIDCProvider API asynchronously
func (client *Client) UpdateOIDCProviderWithCallback(request *UpdateOIDCProviderRequest, callback func(response *UpdateOIDCProviderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOIDCProviderResponse
		var err error
		defer close(result)
		response, err = client.UpdateOIDCProvider(request)
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

// UpdateOIDCProviderRequest is the request struct for api UpdateOIDCProvider
type UpdateOIDCProviderRequest struct {
	*requests.RpcRequest
	IssuanceLimitTime requests.Integer `position:"Query" name:"IssuanceLimitTime"`
	AkProxySuffix     string           `position:"Query" name:"AkProxySuffix"`
	OIDCProviderName  string           `position:"Query" name:"OIDCProviderName"`
	ClientIds         string           `position:"Query" name:"ClientIds"`
	NewDescription    string           `position:"Query" name:"NewDescription"`
}

// UpdateOIDCProviderResponse is the response struct for api UpdateOIDCProvider
type UpdateOIDCProviderResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	OIDCProvider OIDCProvider `json:"OIDCProvider" xml:"OIDCProvider"`
}

// CreateUpdateOIDCProviderRequest creates a request to invoke UpdateOIDCProvider API
func CreateUpdateOIDCProviderRequest() (request *UpdateOIDCProviderRequest) {
	request = &UpdateOIDCProviderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "UpdateOIDCProvider", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateOIDCProviderResponse creates a response to parse from UpdateOIDCProvider response
func CreateUpdateOIDCProviderResponse() (response *UpdateOIDCProviderResponse) {
	response = &UpdateOIDCProviderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
