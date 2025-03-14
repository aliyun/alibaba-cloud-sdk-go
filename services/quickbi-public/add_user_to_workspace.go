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

// AddUserToWorkspace invokes the quickbi_public.AddUserToWorkspace API synchronously
func (client *Client) AddUserToWorkspace(request *AddUserToWorkspaceRequest) (response *AddUserToWorkspaceResponse, err error) {
	response = CreateAddUserToWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserToWorkspaceWithChan invokes the quickbi_public.AddUserToWorkspace API asynchronously
func (client *Client) AddUserToWorkspaceWithChan(request *AddUserToWorkspaceRequest) (<-chan *AddUserToWorkspaceResponse, <-chan error) {
	responseChan := make(chan *AddUserToWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUserToWorkspace(request)
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

// AddUserToWorkspaceWithCallback invokes the quickbi_public.AddUserToWorkspace API asynchronously
func (client *Client) AddUserToWorkspaceWithCallback(request *AddUserToWorkspaceRequest, callback func(response *AddUserToWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserToWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.AddUserToWorkspace(request)
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

// AddUserToWorkspaceRequest is the request struct for api AddUserToWorkspace
type AddUserToWorkspaceRequest struct {
	*requests.RpcRequest
	RoleId      requests.Integer `position:"Query" name:"RoleId"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	UserId      string           `position:"Query" name:"UserId"`
	SignType    string           `position:"Query" name:"SignType"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// AddUserToWorkspaceResponse is the response struct for api AddUserToWorkspace
type AddUserToWorkspaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAddUserToWorkspaceRequest creates a request to invoke AddUserToWorkspace API
func CreateAddUserToWorkspaceRequest() (request *AddUserToWorkspaceRequest) {
	request = &AddUserToWorkspaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "AddUserToWorkspace", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddUserToWorkspaceResponse creates a response to parse from AddUserToWorkspace response
func CreateAddUserToWorkspaceResponse() (response *AddUserToWorkspaceResponse) {
	response = &AddUserToWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
