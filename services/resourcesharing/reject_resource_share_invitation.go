package resourcesharing

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

// RejectResourceShareInvitation invokes the resourcesharing.RejectResourceShareInvitation API synchronously
func (client *Client) RejectResourceShareInvitation(request *RejectResourceShareInvitationRequest) (response *RejectResourceShareInvitationResponse, err error) {
	response = CreateRejectResourceShareInvitationResponse()
	err = client.DoAction(request, response)
	return
}

// RejectResourceShareInvitationWithChan invokes the resourcesharing.RejectResourceShareInvitation API asynchronously
func (client *Client) RejectResourceShareInvitationWithChan(request *RejectResourceShareInvitationRequest) (<-chan *RejectResourceShareInvitationResponse, <-chan error) {
	responseChan := make(chan *RejectResourceShareInvitationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RejectResourceShareInvitation(request)
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

// RejectResourceShareInvitationWithCallback invokes the resourcesharing.RejectResourceShareInvitation API asynchronously
func (client *Client) RejectResourceShareInvitationWithCallback(request *RejectResourceShareInvitationRequest, callback func(response *RejectResourceShareInvitationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RejectResourceShareInvitationResponse
		var err error
		defer close(result)
		response, err = client.RejectResourceShareInvitation(request)
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

// RejectResourceShareInvitationRequest is the request struct for api RejectResourceShareInvitation
type RejectResourceShareInvitationRequest struct {
	*requests.RpcRequest
	ResourceShareInvitationId string `position:"Query" name:"ResourceShareInvitationId"`
}

// RejectResourceShareInvitationResponse is the response struct for api RejectResourceShareInvitation
type RejectResourceShareInvitationResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	ResourceShareInvitation ResourceShareInvitation `json:"ResourceShareInvitation" xml:"ResourceShareInvitation"`
}

// CreateRejectResourceShareInvitationRequest creates a request to invoke RejectResourceShareInvitation API
func CreateRejectResourceShareInvitationRequest() (request *RejectResourceShareInvitationRequest) {
	request = &RejectResourceShareInvitationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceSharing", "2020-01-10", "RejectResourceShareInvitation", "ressharing", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRejectResourceShareInvitationResponse creates a response to parse from RejectResourceShareInvitation response
func CreateRejectResourceShareInvitationResponse() (response *RejectResourceShareInvitationResponse) {
	response = &RejectResourceShareInvitationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
