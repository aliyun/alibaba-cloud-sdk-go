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

// ListRepositoryMember invokes the codeup.ListRepositoryMember API synchronously
// api document: https://help.aliyun.com/api/codeup/listrepositorymember.html
func (client *Client) ListRepositoryMember(request *ListRepositoryMemberRequest) (response *ListRepositoryMemberResponse, err error) {
	response = CreateListRepositoryMemberResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepositoryMemberWithChan invokes the codeup.ListRepositoryMember API asynchronously
// api document: https://help.aliyun.com/api/codeup/listrepositorymember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepositoryMemberWithChan(request *ListRepositoryMemberRequest) (<-chan *ListRepositoryMemberResponse, <-chan error) {
	responseChan := make(chan *ListRepositoryMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepositoryMember(request)
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

// ListRepositoryMemberWithCallback invokes the codeup.ListRepositoryMember API asynchronously
// api document: https://help.aliyun.com/api/codeup/listrepositorymember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepositoryMemberWithCallback(request *ListRepositoryMemberRequest, callback func(response *ListRepositoryMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepositoryMemberResponse
		var err error
		defer close(result)
		response, err = client.ListRepositoryMember(request)
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

// ListRepositoryMemberRequest is the request struct for api ListRepositoryMember
type ListRepositoryMemberRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	Query          string           `position:"Query" name:"Query"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	Page           requests.Integer `position:"Query" name:"Page"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// ListRepositoryMemberResponse is the response struct for api ListRepositoryMember
type ListRepositoryMemberResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool         `json:"Success" xml:"Success"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       []ResultItem `json:"Result" xml:"Result"`
}

// CreateListRepositoryMemberRequest creates a request to invoke ListRepositoryMember API
func CreateListRepositoryMemberRequest() (request *ListRepositoryMemberRequest) {
	request = &ListRepositoryMemberRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListRepositoryMember", "/api/v3/projects/[ProjectId]/members", "", "")
	request.Method = requests.GET
	return
}

// CreateListRepositoryMemberResponse creates a response to parse from ListRepositoryMember response
func CreateListRepositoryMemberResponse() (response *ListRepositoryMemberResponse) {
	response = &ListRepositoryMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
