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

// DeleteTransitRouterPrefixListAssociation invokes the cbn.DeleteTransitRouterPrefixListAssociation API synchronously
func (client *Client) DeleteTransitRouterPrefixListAssociation(request *DeleteTransitRouterPrefixListAssociationRequest) (response *DeleteTransitRouterPrefixListAssociationResponse, err error) {
	response = CreateDeleteTransitRouterPrefixListAssociationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTransitRouterPrefixListAssociationWithChan invokes the cbn.DeleteTransitRouterPrefixListAssociation API asynchronously
func (client *Client) DeleteTransitRouterPrefixListAssociationWithChan(request *DeleteTransitRouterPrefixListAssociationRequest) (<-chan *DeleteTransitRouterPrefixListAssociationResponse, <-chan error) {
	responseChan := make(chan *DeleteTransitRouterPrefixListAssociationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTransitRouterPrefixListAssociation(request)
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

// DeleteTransitRouterPrefixListAssociationWithCallback invokes the cbn.DeleteTransitRouterPrefixListAssociation API asynchronously
func (client *Client) DeleteTransitRouterPrefixListAssociationWithCallback(request *DeleteTransitRouterPrefixListAssociationRequest, callback func(response *DeleteTransitRouterPrefixListAssociationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTransitRouterPrefixListAssociationResponse
		var err error
		defer close(result)
		response, err = client.DeleteTransitRouterPrefixListAssociation(request)
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

// DeleteTransitRouterPrefixListAssociationRequest is the request struct for api DeleteTransitRouterPrefixListAssociation
type DeleteTransitRouterPrefixListAssociationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PrefixListId         string           `position:"Query" name:"PrefixListId"`
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

// DeleteTransitRouterPrefixListAssociationResponse is the response struct for api DeleteTransitRouterPrefixListAssociation
type DeleteTransitRouterPrefixListAssociationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTransitRouterPrefixListAssociationRequest creates a request to invoke DeleteTransitRouterPrefixListAssociation API
func CreateDeleteTransitRouterPrefixListAssociationRequest() (request *DeleteTransitRouterPrefixListAssociationRequest) {
	request = &DeleteTransitRouterPrefixListAssociationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteTransitRouterPrefixListAssociation", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTransitRouterPrefixListAssociationResponse creates a response to parse from DeleteTransitRouterPrefixListAssociation response
func CreateDeleteTransitRouterPrefixListAssociationResponse() (response *DeleteTransitRouterPrefixListAssociationResponse) {
	response = &DeleteTransitRouterPrefixListAssociationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
