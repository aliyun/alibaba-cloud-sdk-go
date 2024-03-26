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

// CreateTransitRouterPrefixListAssociation invokes the cbn.CreateTransitRouterPrefixListAssociation API synchronously
func (client *Client) CreateTransitRouterPrefixListAssociation(request *CreateTransitRouterPrefixListAssociationRequest) (response *CreateTransitRouterPrefixListAssociationResponse, err error) {
	response = CreateCreateTransitRouterPrefixListAssociationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTransitRouterPrefixListAssociationWithChan invokes the cbn.CreateTransitRouterPrefixListAssociation API asynchronously
func (client *Client) CreateTransitRouterPrefixListAssociationWithChan(request *CreateTransitRouterPrefixListAssociationRequest) (<-chan *CreateTransitRouterPrefixListAssociationResponse, <-chan error) {
	responseChan := make(chan *CreateTransitRouterPrefixListAssociationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTransitRouterPrefixListAssociation(request)
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

// CreateTransitRouterPrefixListAssociationWithCallback invokes the cbn.CreateTransitRouterPrefixListAssociation API asynchronously
func (client *Client) CreateTransitRouterPrefixListAssociationWithCallback(request *CreateTransitRouterPrefixListAssociationRequest, callback func(response *CreateTransitRouterPrefixListAssociationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTransitRouterPrefixListAssociationResponse
		var err error
		defer close(result)
		response, err = client.CreateTransitRouterPrefixListAssociation(request)
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

// CreateTransitRouterPrefixListAssociationRequest is the request struct for api CreateTransitRouterPrefixListAssociation
type CreateTransitRouterPrefixListAssociationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PrefixListId         string           `position:"Query" name:"PrefixListId"`
	OwnerUid             requests.Integer `position:"Query" name:"OwnerUid"`
	NextHopType          string           `position:"Query" name:"NextHopType"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TransitRouterId      string           `position:"Query" name:"TransitRouterId"`
	TransitRouterTableId string           `position:"Query" name:"TransitRouterTableId"`
	Version              string           `position:"Query" name:"Version"`
	NextHop              string           `position:"Query" name:"NextHop"`
}

// CreateTransitRouterPrefixListAssociationResponse is the response struct for api CreateTransitRouterPrefixListAssociation
type CreateTransitRouterPrefixListAssociationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTransitRouterPrefixListAssociationRequest creates a request to invoke CreateTransitRouterPrefixListAssociation API
func CreateCreateTransitRouterPrefixListAssociationRequest() (request *CreateTransitRouterPrefixListAssociationRequest) {
	request = &CreateTransitRouterPrefixListAssociationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateTransitRouterPrefixListAssociation", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTransitRouterPrefixListAssociationResponse creates a response to parse from CreateTransitRouterPrefixListAssociation response
func CreateCreateTransitRouterPrefixListAssociationResponse() (response *CreateTransitRouterPrefixListAssociationResponse) {
	response = &CreateTransitRouterPrefixListAssociationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
