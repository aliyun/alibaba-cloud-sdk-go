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

// ListTransitRouterRouteTablePropagations invokes the cbn.ListTransitRouterRouteTablePropagations API synchronously
func (client *Client) ListTransitRouterRouteTablePropagations(request *ListTransitRouterRouteTablePropagationsRequest) (response *ListTransitRouterRouteTablePropagationsResponse, err error) {
	response = CreateListTransitRouterRouteTablePropagationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterRouteTablePropagationsWithChan invokes the cbn.ListTransitRouterRouteTablePropagations API asynchronously
func (client *Client) ListTransitRouterRouteTablePropagationsWithChan(request *ListTransitRouterRouteTablePropagationsRequest) (<-chan *ListTransitRouterRouteTablePropagationsResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterRouteTablePropagationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterRouteTablePropagations(request)
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

// ListTransitRouterRouteTablePropagationsWithCallback invokes the cbn.ListTransitRouterRouteTablePropagations API asynchronously
func (client *Client) ListTransitRouterRouteTablePropagationsWithCallback(request *ListTransitRouterRouteTablePropagationsRequest, callback func(response *ListTransitRouterRouteTablePropagationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterRouteTablePropagationsResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterRouteTablePropagations(request)
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

// ListTransitRouterRouteTablePropagationsRequest is the request struct for api ListTransitRouterRouteTablePropagations
type ListTransitRouterRouteTablePropagationsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TransitRouterRouteTableId  string           `position:"Query" name:"TransitRouterRouteTableId"`
	NextToken                  string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	TransitRouterAttachmentId  string           `position:"Query" name:"TransitRouterAttachmentId"`
	MaxResults                 requests.Integer `position:"Query" name:"MaxResults"`
	TransitRouterAttachmentIds *[]string        `position:"Query" name:"TransitRouterAttachmentIds"  type:"Repeated"`
}

// ListTransitRouterRouteTablePropagationsResponse is the response struct for api ListTransitRouterRouteTablePropagations
type ListTransitRouterRouteTablePropagationsResponse struct {
	*responses.BaseResponse
	NextToken                 string                     `json:"NextToken" xml:"NextToken"`
	RequestId                 string                     `json:"RequestId" xml:"RequestId"`
	TotalCount                int                        `json:"TotalCount" xml:"TotalCount"`
	MaxResults                int                        `json:"MaxResults" xml:"MaxResults"`
	TransitRouterPropagations []TransitRouterPropagation `json:"TransitRouterPropagations" xml:"TransitRouterPropagations"`
}

// CreateListTransitRouterRouteTablePropagationsRequest creates a request to invoke ListTransitRouterRouteTablePropagations API
func CreateListTransitRouterRouteTablePropagationsRequest() (request *ListTransitRouterRouteTablePropagationsRequest) {
	request = &ListTransitRouterRouteTablePropagationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterRouteTablePropagations", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterRouteTablePropagationsResponse creates a response to parse from ListTransitRouterRouteTablePropagations response
func CreateListTransitRouterRouteTablePropagationsResponse() (response *ListTransitRouterRouteTablePropagationsResponse) {
	response = &ListTransitRouterRouteTablePropagationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
