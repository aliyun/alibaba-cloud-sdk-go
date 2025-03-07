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

// ListAuthorizedInstancesForUser invokes the dms_enterprise.ListAuthorizedInstancesForUser API synchronously
func (client *Client) ListAuthorizedInstancesForUser(request *ListAuthorizedInstancesForUserRequest) (response *ListAuthorizedInstancesForUserResponse, err error) {
	response = CreateListAuthorizedInstancesForUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListAuthorizedInstancesForUserWithChan invokes the dms_enterprise.ListAuthorizedInstancesForUser API asynchronously
func (client *Client) ListAuthorizedInstancesForUserWithChan(request *ListAuthorizedInstancesForUserRequest) (<-chan *ListAuthorizedInstancesForUserResponse, <-chan error) {
	responseChan := make(chan *ListAuthorizedInstancesForUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAuthorizedInstancesForUser(request)
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

// ListAuthorizedInstancesForUserWithCallback invokes the dms_enterprise.ListAuthorizedInstancesForUser API asynchronously
func (client *Client) ListAuthorizedInstancesForUserWithCallback(request *ListAuthorizedInstancesForUserRequest, callback func(response *ListAuthorizedInstancesForUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAuthorizedInstancesForUserResponse
		var err error
		defer close(result)
		response, err = client.ListAuthorizedInstancesForUser(request)
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

// ListAuthorizedInstancesForUserRequest is the request struct for api ListAuthorizedInstancesForUser
type ListAuthorizedInstancesForUserRequest struct {
	*requests.RpcRequest
	SearchKey  string           `position:"Query" name:"SearchKey"`
	UserId     string           `position:"Query" name:"UserId"`
	PageNumber string           `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	EnvType    string           `position:"Query" name:"EnvType"`
	PageSize   string           `position:"Query" name:"PageSize"`
	DbType     string           `position:"Query" name:"DbType"`
}

// ListAuthorizedInstancesForUserResponse is the response struct for api ListAuthorizedInstancesForUser
type ListAuthorizedInstancesForUserResponse struct {
	*responses.BaseResponse
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Instances []InstancesItem `json:"Instances" xml:"Instances"`
}

// CreateListAuthorizedInstancesForUserRequest creates a request to invoke ListAuthorizedInstancesForUser API
func CreateListAuthorizedInstancesForUserRequest() (request *ListAuthorizedInstancesForUserRequest) {
	request = &ListAuthorizedInstancesForUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListAuthorizedInstancesForUser", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAuthorizedInstancesForUserResponse creates a response to parse from ListAuthorizedInstancesForUser response
func CreateListAuthorizedInstancesForUserResponse() (response *ListAuthorizedInstancesForUserResponse) {
	response = &ListAuthorizedInstancesForUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
