package adb

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

// BindDBResourceGroupWithUser invokes the adb.BindDBResourceGroupWithUser API synchronously
func (client *Client) BindDBResourceGroupWithUser(request *BindDBResourceGroupWithUserRequest) (response *BindDBResourceGroupWithUserResponse, err error) {
	response = CreateBindDBResourceGroupWithUserResponse()
	err = client.DoAction(request, response)
	return
}

// BindDBResourceGroupWithUserWithChan invokes the adb.BindDBResourceGroupWithUser API asynchronously
func (client *Client) BindDBResourceGroupWithUserWithChan(request *BindDBResourceGroupWithUserRequest) (<-chan *BindDBResourceGroupWithUserResponse, <-chan error) {
	responseChan := make(chan *BindDBResourceGroupWithUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindDBResourceGroupWithUser(request)
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

// BindDBResourceGroupWithUserWithCallback invokes the adb.BindDBResourceGroupWithUser API asynchronously
func (client *Client) BindDBResourceGroupWithUserWithCallback(request *BindDBResourceGroupWithUserRequest, callback func(response *BindDBResourceGroupWithUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindDBResourceGroupWithUserResponse
		var err error
		defer close(result)
		response, err = client.BindDBResourceGroupWithUser(request)
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

// BindDBResourceGroupWithUserRequest is the request struct for api BindDBResourceGroupWithUser
type BindDBResourceGroupWithUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	GroupUser            string           `position:"Query" name:"GroupUser"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
}

// BindDBResourceGroupWithUserResponse is the response struct for api BindDBResourceGroupWithUser
type BindDBResourceGroupWithUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindDBResourceGroupWithUserRequest creates a request to invoke BindDBResourceGroupWithUser API
func CreateBindDBResourceGroupWithUserRequest() (request *BindDBResourceGroupWithUserRequest) {
	request = &BindDBResourceGroupWithUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "BindDBResourceGroupWithUser", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindDBResourceGroupWithUserResponse creates a response to parse from BindDBResourceGroupWithUser response
func CreateBindDBResourceGroupWithUserResponse() (response *BindDBResourceGroupWithUserResponse) {
	response = &BindDBResourceGroupWithUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
