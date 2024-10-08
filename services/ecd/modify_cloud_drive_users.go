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

// ModifyCloudDriveUsers invokes the ecd.ModifyCloudDriveUsers API synchronously
func (client *Client) ModifyCloudDriveUsers(request *ModifyCloudDriveUsersRequest) (response *ModifyCloudDriveUsersResponse, err error) {
	response = CreateModifyCloudDriveUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCloudDriveUsersWithChan invokes the ecd.ModifyCloudDriveUsers API asynchronously
func (client *Client) ModifyCloudDriveUsersWithChan(request *ModifyCloudDriveUsersRequest) (<-chan *ModifyCloudDriveUsersResponse, <-chan error) {
	responseChan := make(chan *ModifyCloudDriveUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCloudDriveUsers(request)
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

// ModifyCloudDriveUsersWithCallback invokes the ecd.ModifyCloudDriveUsers API asynchronously
func (client *Client) ModifyCloudDriveUsersWithCallback(request *ModifyCloudDriveUsersRequest, callback func(response *ModifyCloudDriveUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCloudDriveUsersResponse
		var err error
		defer close(result)
		response, err = client.ModifyCloudDriveUsers(request)
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

// ModifyCloudDriveUsersRequest is the request struct for api ModifyCloudDriveUsers
type ModifyCloudDriveUsersRequest struct {
	*requests.RpcRequest
	CdsId       string           `position:"Query" name:"CdsId"`
	EndUserId   *[]string        `position:"Query" name:"EndUserId"  type:"Repeated"`
	UserMaxSize requests.Integer `position:"Query" name:"UserMaxSize"`
	Status      string           `position:"Query" name:"Status"`
}

// ModifyCloudDriveUsersResponse is the response struct for api ModifyCloudDriveUsers
type ModifyCloudDriveUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCloudDriveUsersRequest creates a request to invoke ModifyCloudDriveUsers API
func CreateModifyCloudDriveUsersRequest() (request *ModifyCloudDriveUsersRequest) {
	request = &ModifyCloudDriveUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyCloudDriveUsers", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCloudDriveUsersResponse creates a response to parse from ModifyCloudDriveUsers response
func CreateModifyCloudDriveUsersResponse() (response *ModifyCloudDriveUsersResponse) {
	response = &ModifyCloudDriveUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
