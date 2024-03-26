package cbn

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

// UpdateTransitRouterVbrAttachmentAttribute invokes the cbn.UpdateTransitRouterVbrAttachmentAttribute API synchronously
func (client *Client) UpdateTransitRouterVbrAttachmentAttribute(request *UpdateTransitRouterVbrAttachmentAttributeRequest) (response *UpdateTransitRouterVbrAttachmentAttributeResponse, err error) {
	response = CreateUpdateTransitRouterVbrAttachmentAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTransitRouterVbrAttachmentAttributeWithChan invokes the cbn.UpdateTransitRouterVbrAttachmentAttribute API asynchronously
func (client *Client) UpdateTransitRouterVbrAttachmentAttributeWithChan(request *UpdateTransitRouterVbrAttachmentAttributeRequest) (<-chan *UpdateTransitRouterVbrAttachmentAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateTransitRouterVbrAttachmentAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTransitRouterVbrAttachmentAttribute(request)
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

// UpdateTransitRouterVbrAttachmentAttributeWithCallback invokes the cbn.UpdateTransitRouterVbrAttachmentAttribute API asynchronously
func (client *Client) UpdateTransitRouterVbrAttachmentAttributeWithCallback(request *UpdateTransitRouterVbrAttachmentAttributeRequest, callback func(response *UpdateTransitRouterVbrAttachmentAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTransitRouterVbrAttachmentAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateTransitRouterVbrAttachmentAttribute(request)
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

// UpdateTransitRouterVbrAttachmentAttributeRequest is the request struct for api UpdateTransitRouterVbrAttachmentAttribute
type UpdateTransitRouterVbrAttachmentAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                    requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                        string           `position:"Query" name:"ClientToken"`
	TransitRouterAttachmentName        string           `position:"Query" name:"TransitRouterAttachmentName"`
	AutoPublishRouteEnabled            requests.Boolean `position:"Query" name:"AutoPublishRouteEnabled"`
	DryRun                             requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount               string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                       string           `position:"Query" name:"OwnerAccount"`
	OwnerId                            requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType                       string           `position:"Query" name:"ResourceType"`
	Version                            string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId          string           `position:"Query" name:"TransitRouterAttachmentId"`
	TransitRouterAttachmentDescription string           `position:"Query" name:"TransitRouterAttachmentDescription"`
}

// UpdateTransitRouterVbrAttachmentAttributeResponse is the response struct for api UpdateTransitRouterVbrAttachmentAttribute
type UpdateTransitRouterVbrAttachmentAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateTransitRouterVbrAttachmentAttributeRequest creates a request to invoke UpdateTransitRouterVbrAttachmentAttribute API
func CreateUpdateTransitRouterVbrAttachmentAttributeRequest() (request *UpdateTransitRouterVbrAttachmentAttributeRequest) {
	request = &UpdateTransitRouterVbrAttachmentAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "UpdateTransitRouterVbrAttachmentAttribute", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTransitRouterVbrAttachmentAttributeResponse creates a response to parse from UpdateTransitRouterVbrAttachmentAttribute response
func CreateUpdateTransitRouterVbrAttachmentAttributeResponse() (response *UpdateTransitRouterVbrAttachmentAttributeResponse) {
	response = &UpdateTransitRouterVbrAttachmentAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
