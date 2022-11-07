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

// DescribeUsersInGroup invokes the ecd.DescribeUsersInGroup API synchronously
func (client *Client) DescribeUsersInGroup(request *DescribeUsersInGroupRequest) (response *DescribeUsersInGroupResponse, err error) {
	response = CreateDescribeUsersInGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUsersInGroupWithChan invokes the ecd.DescribeUsersInGroup API asynchronously
func (client *Client) DescribeUsersInGroupWithChan(request *DescribeUsersInGroupRequest) (<-chan *DescribeUsersInGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeUsersInGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUsersInGroup(request)
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

// DescribeUsersInGroupWithCallback invokes the ecd.DescribeUsersInGroup API asynchronously
func (client *Client) DescribeUsersInGroupWithCallback(request *DescribeUsersInGroupRequest, callback func(response *DescribeUsersInGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUsersInGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeUsersInGroup(request)
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

// DescribeUsersInGroupRequest is the request struct for api DescribeUsersInGroup
type DescribeUsersInGroupRequest struct {
	*requests.RpcRequest
	ConnectState    requests.Integer `position:"Query" name:"ConnectState"`
	Filter          string           `position:"Query" name:"Filter"`
	DesktopGroupId  string           `position:"Query" name:"DesktopGroupId"`
	NextToken       string           `position:"Query" name:"NextToken"`
	QueryUserDetail requests.Boolean `position:"Query" name:"QueryUserDetail"`
	MaxResults      requests.Integer `position:"Query" name:"MaxResults"`
	EndUserId       string           `position:"Query" name:"EndUserId"`
}

// DescribeUsersInGroupResponse is the response struct for api DescribeUsersInGroup
type DescribeUsersInGroupResponse struct {
	*responses.BaseResponse
	NextToken        string    `json:"NextToken" xml:"NextToken"`
	RequestId        string    `json:"RequestId" xml:"RequestId"`
	UsersCount       int       `json:"UsersCount" xml:"UsersCount"`
	OnlineUsersCount int       `json:"OnlineUsersCount" xml:"OnlineUsersCount"`
	EndUsers         []EndUser `json:"EndUsers" xml:"EndUsers"`
}

// CreateDescribeUsersInGroupRequest creates a request to invoke DescribeUsersInGroup API
func CreateDescribeUsersInGroupRequest() (request *DescribeUsersInGroupRequest) {
	request = &DescribeUsersInGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeUsersInGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUsersInGroupResponse creates a response to parse from DescribeUsersInGroup response
func CreateDescribeUsersInGroupResponse() (response *DescribeUsersInGroupResponse) {
	response = &DescribeUsersInGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
