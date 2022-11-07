package resourcemanager

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

// SetMemberDeletionPermission invokes the resourcemanager.SetMemberDeletionPermission API synchronously
func (client *Client) SetMemberDeletionPermission(request *SetMemberDeletionPermissionRequest) (response *SetMemberDeletionPermissionResponse, err error) {
	response = CreateSetMemberDeletionPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// SetMemberDeletionPermissionWithChan invokes the resourcemanager.SetMemberDeletionPermission API asynchronously
func (client *Client) SetMemberDeletionPermissionWithChan(request *SetMemberDeletionPermissionRequest) (<-chan *SetMemberDeletionPermissionResponse, <-chan error) {
	responseChan := make(chan *SetMemberDeletionPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetMemberDeletionPermission(request)
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

// SetMemberDeletionPermissionWithCallback invokes the resourcemanager.SetMemberDeletionPermission API asynchronously
func (client *Client) SetMemberDeletionPermissionWithCallback(request *SetMemberDeletionPermissionRequest, callback func(response *SetMemberDeletionPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetMemberDeletionPermissionResponse
		var err error
		defer close(result)
		response, err = client.SetMemberDeletionPermission(request)
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

// SetMemberDeletionPermissionRequest is the request struct for api SetMemberDeletionPermission
type SetMemberDeletionPermissionRequest struct {
	*requests.RpcRequest
	Status string `position:"Query" name:"Status"`
}

// SetMemberDeletionPermissionResponse is the response struct for api SetMemberDeletionPermission
type SetMemberDeletionPermissionResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	MemberDeletionStatus string `json:"MemberDeletionStatus" xml:"MemberDeletionStatus"`
	ManagementAccountId  string `json:"ManagementAccountId" xml:"ManagementAccountId"`
	ResourceDirectoryId  string `json:"ResourceDirectoryId" xml:"ResourceDirectoryId"`
}

// CreateSetMemberDeletionPermissionRequest creates a request to invoke SetMemberDeletionPermission API
func CreateSetMemberDeletionPermissionRequest() (request *SetMemberDeletionPermissionRequest) {
	request = &SetMemberDeletionPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "SetMemberDeletionPermission", "", "")
	request.Method = requests.POST
	return
}

// CreateSetMemberDeletionPermissionResponse creates a response to parse from SetMemberDeletionPermission response
func CreateSetMemberDeletionPermissionResponse() (response *SetMemberDeletionPermissionResponse) {
	response = &SetMemberDeletionPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
