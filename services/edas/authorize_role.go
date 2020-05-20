package edas

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

// AuthorizeRole invokes the edas.AuthorizeRole API synchronously
// api document: https://help.aliyun.com/api/edas/authorizerole.html
func (client *Client) AuthorizeRole(request *AuthorizeRoleRequest) (response *AuthorizeRoleResponse, err error) {
	response = CreateAuthorizeRoleResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeRoleWithChan invokes the edas.AuthorizeRole API asynchronously
// api document: https://help.aliyun.com/api/edas/authorizerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizeRoleWithChan(request *AuthorizeRoleRequest) (<-chan *AuthorizeRoleResponse, <-chan error) {
	responseChan := make(chan *AuthorizeRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeRole(request)
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

// AuthorizeRoleWithCallback invokes the edas.AuthorizeRole API asynchronously
// api document: https://help.aliyun.com/api/edas/authorizerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizeRoleWithCallback(request *AuthorizeRoleRequest, callback func(response *AuthorizeRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeRoleResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeRole(request)
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

// AuthorizeRoleRequest is the request struct for api AuthorizeRole
type AuthorizeRoleRequest struct {
	*requests.RoaRequest
	RoleIds      string `position:"Query" name:"RoleIds"`
	TargetUserId string `position:"Query" name:"TargetUserId"`
}

// AuthorizeRoleResponse is the response struct for api AuthorizeRole
type AuthorizeRoleResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAuthorizeRoleRequest creates a request to invoke AuthorizeRole API
func CreateAuthorizeRoleRequest() (request *AuthorizeRoleRequest) {
	request = &AuthorizeRoleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "AuthorizeRole", "/pop/v5/account/authorize_role", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAuthorizeRoleResponse creates a response to parse from AuthorizeRole response
func CreateAuthorizeRoleResponse() (response *AuthorizeRoleResponse) {
	response = &AuthorizeRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
