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

// DisassociateResourceShare invokes the resourcesharing.DisassociateResourceShare API synchronously
func (client *Client) DisassociateResourceShare(request *DisassociateResourceShareRequest) (response *DisassociateResourceShareResponse, err error) {
	response = CreateDisassociateResourceShareResponse()
	err = client.DoAction(request, response)
	return
}

// DisassociateResourceShareWithChan invokes the resourcesharing.DisassociateResourceShare API asynchronously
func (client *Client) DisassociateResourceShareWithChan(request *DisassociateResourceShareRequest) (<-chan *DisassociateResourceShareResponse, <-chan error) {
	responseChan := make(chan *DisassociateResourceShareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisassociateResourceShare(request)
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

// DisassociateResourceShareWithCallback invokes the resourcesharing.DisassociateResourceShare API asynchronously
func (client *Client) DisassociateResourceShareWithCallback(request *DisassociateResourceShareRequest, callback func(response *DisassociateResourceShareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisassociateResourceShareResponse
		var err error
		defer close(result)
		response, err = client.DisassociateResourceShare(request)
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

// DisassociateResourceShareRequest is the request struct for api DisassociateResourceShare
type DisassociateResourceShareRequest struct {
	*requests.RpcRequest
	ResourceOwner   string                                `position:"Query" name:"ResourceOwner"`
	Resources       *[]DisassociateResourceShareResources `position:"Query" name:"Resources"  type:"Repeated"`
	Targets         *[]string                             `position:"Query" name:"Targets"  type:"Repeated"`
	ResourceShareId string                                `position:"Query" name:"ResourceShareId"`
}

// DisassociateResourceShareResources is a repeated param struct in DisassociateResourceShareRequest
type DisassociateResourceShareResources struct {
	ResourceId   string `name:"ResourceId"`
	ResourceType string `name:"ResourceType"`
}

// DisassociateResourceShareResponse is the response struct for api DisassociateResourceShare
type DisassociateResourceShareResponse struct {
	*responses.BaseResponse
	RequestId                 string                     `json:"RequestId" xml:"RequestId"`
	ResourceShareAssociations []ResourceShareAssociation `json:"ResourceShareAssociations" xml:"ResourceShareAssociations"`
}

// CreateDisassociateResourceShareRequest creates a request to invoke DisassociateResourceShare API
func CreateDisassociateResourceShareRequest() (request *DisassociateResourceShareRequest) {
	request = &DisassociateResourceShareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceSharing", "2020-01-10", "DisassociateResourceShare", "ressharing", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisassociateResourceShareResponse creates a response to parse from DisassociateResourceShare response
func CreateDisassociateResourceShareResponse() (response *DisassociateResourceShareResponse) {
	response = &DisassociateResourceShareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
