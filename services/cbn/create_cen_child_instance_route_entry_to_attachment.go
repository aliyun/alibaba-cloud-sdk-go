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

// CreateCenChildInstanceRouteEntryToAttachment invokes the cbn.CreateCenChildInstanceRouteEntryToAttachment API synchronously
func (client *Client) CreateCenChildInstanceRouteEntryToAttachment(request *CreateCenChildInstanceRouteEntryToAttachmentRequest) (response *CreateCenChildInstanceRouteEntryToAttachmentResponse, err error) {
	response = CreateCreateCenChildInstanceRouteEntryToAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCenChildInstanceRouteEntryToAttachmentWithChan invokes the cbn.CreateCenChildInstanceRouteEntryToAttachment API asynchronously
func (client *Client) CreateCenChildInstanceRouteEntryToAttachmentWithChan(request *CreateCenChildInstanceRouteEntryToAttachmentRequest) (<-chan *CreateCenChildInstanceRouteEntryToAttachmentResponse, <-chan error) {
	responseChan := make(chan *CreateCenChildInstanceRouteEntryToAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCenChildInstanceRouteEntryToAttachment(request)
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

// CreateCenChildInstanceRouteEntryToAttachmentWithCallback invokes the cbn.CreateCenChildInstanceRouteEntryToAttachment API asynchronously
func (client *Client) CreateCenChildInstanceRouteEntryToAttachmentWithCallback(request *CreateCenChildInstanceRouteEntryToAttachmentRequest, callback func(response *CreateCenChildInstanceRouteEntryToAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCenChildInstanceRouteEntryToAttachmentResponse
		var err error
		defer close(result)
		response, err = client.CreateCenChildInstanceRouteEntryToAttachment(request)
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

// CreateCenChildInstanceRouteEntryToAttachmentRequest is the request struct for api CreateCenChildInstanceRouteEntryToAttachment
type CreateCenChildInstanceRouteEntryToAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	CenId                     string           `position:"Query" name:"CenId"`
	RouteTableId              string           `position:"Query" name:"RouteTableId"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	DestinationCidrBlock      string           `position:"Query" name:"DestinationCidrBlock"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	Version                   string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId string           `position:"Query" name:"TransitRouterAttachmentId"`
}

// CreateCenChildInstanceRouteEntryToAttachmentResponse is the response struct for api CreateCenChildInstanceRouteEntryToAttachment
type CreateCenChildInstanceRouteEntryToAttachmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateCenChildInstanceRouteEntryToAttachmentRequest creates a request to invoke CreateCenChildInstanceRouteEntryToAttachment API
func CreateCreateCenChildInstanceRouteEntryToAttachmentRequest() (request *CreateCenChildInstanceRouteEntryToAttachmentRequest) {
	request = &CreateCenChildInstanceRouteEntryToAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateCenChildInstanceRouteEntryToAttachment", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCenChildInstanceRouteEntryToAttachmentResponse creates a response to parse from CreateCenChildInstanceRouteEntryToAttachment response
func CreateCreateCenChildInstanceRouteEntryToAttachmentResponse() (response *CreateCenChildInstanceRouteEntryToAttachmentResponse) {
	response = &CreateCenChildInstanceRouteEntryToAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
