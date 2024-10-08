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

// DescribeUsersPassword invokes the ecd.DescribeUsersPassword API synchronously
func (client *Client) DescribeUsersPassword(request *DescribeUsersPasswordRequest) (response *DescribeUsersPasswordResponse, err error) {
	response = CreateDescribeUsersPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUsersPasswordWithChan invokes the ecd.DescribeUsersPassword API asynchronously
func (client *Client) DescribeUsersPasswordWithChan(request *DescribeUsersPasswordRequest) (<-chan *DescribeUsersPasswordResponse, <-chan error) {
	responseChan := make(chan *DescribeUsersPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUsersPassword(request)
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

// DescribeUsersPasswordWithCallback invokes the ecd.DescribeUsersPassword API asynchronously
func (client *Client) DescribeUsersPasswordWithCallback(request *DescribeUsersPasswordRequest, callback func(response *DescribeUsersPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUsersPasswordResponse
		var err error
		defer close(result)
		response, err = client.DescribeUsersPassword(request)
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

// DescribeUsersPasswordRequest is the request struct for api DescribeUsersPassword
type DescribeUsersPasswordRequest struct {
	*requests.RpcRequest
	DesktopId string `position:"Query" name:"DesktopId"`
}

// DescribeUsersPasswordResponse is the response struct for api DescribeUsersPassword
type DescribeUsersPasswordResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	DesktopUsers []DesktopUser `json:"DesktopUsers" xml:"DesktopUsers"`
}

// CreateDescribeUsersPasswordRequest creates a request to invoke DescribeUsersPassword API
func CreateDescribeUsersPasswordRequest() (request *DescribeUsersPasswordRequest) {
	request = &DescribeUsersPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeUsersPassword", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUsersPasswordResponse creates a response to parse from DescribeUsersPassword response
func CreateDescribeUsersPasswordResponse() (response *DescribeUsersPasswordResponse) {
	response = &DescribeUsersPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
