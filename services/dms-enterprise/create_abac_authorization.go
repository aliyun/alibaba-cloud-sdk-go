package dms_enterprise

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

// CreateAbacAuthorization invokes the dms_enterprise.CreateAbacAuthorization API synchronously
func (client *Client) CreateAbacAuthorization(request *CreateAbacAuthorizationRequest) (response *CreateAbacAuthorizationResponse, err error) {
	response = CreateCreateAbacAuthorizationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAbacAuthorizationWithChan invokes the dms_enterprise.CreateAbacAuthorization API asynchronously
func (client *Client) CreateAbacAuthorizationWithChan(request *CreateAbacAuthorizationRequest) (<-chan *CreateAbacAuthorizationResponse, <-chan error) {
	responseChan := make(chan *CreateAbacAuthorizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAbacAuthorization(request)
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

// CreateAbacAuthorizationWithCallback invokes the dms_enterprise.CreateAbacAuthorization API asynchronously
func (client *Client) CreateAbacAuthorizationWithCallback(request *CreateAbacAuthorizationRequest, callback func(response *CreateAbacAuthorizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAbacAuthorizationResponse
		var err error
		defer close(result)
		response, err = client.CreateAbacAuthorization(request)
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

// CreateAbacAuthorizationRequest is the request struct for api CreateAbacAuthorization
type CreateAbacAuthorizationRequest struct {
	*requests.RpcRequest
	RoleId       requests.Integer `position:"Query" name:"RoleId"`
	UserId       requests.Integer `position:"Query" name:"UserId"`
	Tid          requests.Integer `position:"Query" name:"Tid"`
	PolicyId     requests.Integer `position:"Query" name:"PolicyId"`
	IdentityType string           `position:"Query" name:"IdentityType"`
}

// CreateAbacAuthorizationResponse is the response struct for api CreateAbacAuthorization
type CreateAbacAuthorizationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       string `json:"Result" xml:"Result"`
}

// CreateCreateAbacAuthorizationRequest creates a request to invoke CreateAbacAuthorization API
func CreateCreateAbacAuthorizationRequest() (request *CreateAbacAuthorizationRequest) {
	request = &CreateAbacAuthorizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateAbacAuthorization", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAbacAuthorizationResponse creates a response to parse from CreateAbacAuthorization response
func CreateCreateAbacAuthorizationResponse() (response *CreateAbacAuthorizationResponse) {
	response = &CreateAbacAuthorizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
