package ecd

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

// ModifyOfficeSiteMfaEnabled invokes the ecd.ModifyOfficeSiteMfaEnabled API synchronously
func (client *Client) ModifyOfficeSiteMfaEnabled(request *ModifyOfficeSiteMfaEnabledRequest) (response *ModifyOfficeSiteMfaEnabledResponse, err error) {
	response = CreateModifyOfficeSiteMfaEnabledResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyOfficeSiteMfaEnabledWithChan invokes the ecd.ModifyOfficeSiteMfaEnabled API asynchronously
func (client *Client) ModifyOfficeSiteMfaEnabledWithChan(request *ModifyOfficeSiteMfaEnabledRequest) (<-chan *ModifyOfficeSiteMfaEnabledResponse, <-chan error) {
	responseChan := make(chan *ModifyOfficeSiteMfaEnabledResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyOfficeSiteMfaEnabled(request)
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

// ModifyOfficeSiteMfaEnabledWithCallback invokes the ecd.ModifyOfficeSiteMfaEnabled API asynchronously
func (client *Client) ModifyOfficeSiteMfaEnabledWithCallback(request *ModifyOfficeSiteMfaEnabledRequest, callback func(response *ModifyOfficeSiteMfaEnabledResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyOfficeSiteMfaEnabledResponse
		var err error
		defer close(result)
		response, err = client.ModifyOfficeSiteMfaEnabled(request)
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

// ModifyOfficeSiteMfaEnabledRequest is the request struct for api ModifyOfficeSiteMfaEnabled
type ModifyOfficeSiteMfaEnabledRequest struct {
	*requests.RpcRequest
	OfficeSiteId string           `position:"Query" name:"OfficeSiteId"`
	MfaEnabled   requests.Boolean `position:"Query" name:"MfaEnabled"`
}

// ModifyOfficeSiteMfaEnabledResponse is the response struct for api ModifyOfficeSiteMfaEnabled
type ModifyOfficeSiteMfaEnabledResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyOfficeSiteMfaEnabledRequest creates a request to invoke ModifyOfficeSiteMfaEnabled API
func CreateModifyOfficeSiteMfaEnabledRequest() (request *ModifyOfficeSiteMfaEnabledRequest) {
	request = &ModifyOfficeSiteMfaEnabledRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyOfficeSiteMfaEnabled", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyOfficeSiteMfaEnabledResponse creates a response to parse from ModifyOfficeSiteMfaEnabled response
func CreateModifyOfficeSiteMfaEnabledResponse() (response *ModifyOfficeSiteMfaEnabledResponse) {
	response = &ModifyOfficeSiteMfaEnabledResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
