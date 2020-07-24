package codeup

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

// UpdateGroupMember invokes the codeup.UpdateGroupMember API synchronously
// api document: https://help.aliyun.com/api/codeup/updategroupmember.html
func (client *Client) UpdateGroupMember(request *UpdateGroupMemberRequest) (response *UpdateGroupMemberResponse, err error) {
	response = CreateUpdateGroupMemberResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGroupMemberWithChan invokes the codeup.UpdateGroupMember API asynchronously
// api document: https://help.aliyun.com/api/codeup/updategroupmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGroupMemberWithChan(request *UpdateGroupMemberRequest) (<-chan *UpdateGroupMemberResponse, <-chan error) {
	responseChan := make(chan *UpdateGroupMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGroupMember(request)
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

// UpdateGroupMemberWithCallback invokes the codeup.UpdateGroupMember API asynchronously
// api document: https://help.aliyun.com/api/codeup/updategroupmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGroupMemberWithCallback(request *UpdateGroupMemberRequest, callback func(response *UpdateGroupMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGroupMemberResponse
		var err error
		defer close(result)
		response, err = client.UpdateGroupMember(request)
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

// UpdateGroupMemberRequest is the request struct for api UpdateGroupMember
type UpdateGroupMemberRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	GroupId        requests.Integer `position:"Path" name:"GroupId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	UserId         requests.Integer `position:"Path" name:"UserId"`
}

// UpdateGroupMemberResponse is the response struct for api UpdateGroupMember
type UpdateGroupMemberResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateUpdateGroupMemberRequest creates a request to invoke UpdateGroupMember API
func CreateUpdateGroupMemberRequest() (request *UpdateGroupMemberRequest) {
	request = &UpdateGroupMemberRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "UpdateGroupMember", "/api/v3/groups/[GroupId]/members/[UserId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateGroupMemberResponse creates a response to parse from UpdateGroupMember response
func CreateUpdateGroupMemberResponse() (response *UpdateGroupMemberResponse) {
	response = &UpdateGroupMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
