package dataworks_public

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

// DeleteLineageRelation invokes the dataworks_public.DeleteLineageRelation API synchronously
func (client *Client) DeleteLineageRelation(request *DeleteLineageRelationRequest) (response *DeleteLineageRelationResponse, err error) {
	response = CreateDeleteLineageRelationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLineageRelationWithChan invokes the dataworks_public.DeleteLineageRelation API asynchronously
func (client *Client) DeleteLineageRelationWithChan(request *DeleteLineageRelationRequest) (<-chan *DeleteLineageRelationResponse, <-chan error) {
	responseChan := make(chan *DeleteLineageRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLineageRelation(request)
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

// DeleteLineageRelationWithCallback invokes the dataworks_public.DeleteLineageRelation API asynchronously
func (client *Client) DeleteLineageRelationWithCallback(request *DeleteLineageRelationRequest, callback func(response *DeleteLineageRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLineageRelationResponse
		var err error
		defer close(result)
		response, err = client.DeleteLineageRelation(request)
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

// DeleteLineageRelationRequest is the request struct for api DeleteLineageRelation
type DeleteLineageRelationRequest struct {
	*requests.RpcRequest
	SrcEntityQualifiedName  string `position:"Query" name:"SrcEntityQualifiedName"`
	DestEntityQualifiedName string `position:"Query" name:"DestEntityQualifiedName"`
	RelationshipType        string `position:"Query" name:"RelationshipType"`
	RelationshipGuid        string `position:"Query" name:"RelationshipGuid"`
}

// DeleteLineageRelationResponse is the response struct for api DeleteLineageRelation
type DeleteLineageRelationResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Status         bool   `json:"Status" xml:"Status"`
}

// CreateDeleteLineageRelationRequest creates a request to invoke DeleteLineageRelation API
func CreateDeleteLineageRelationRequest() (request *DeleteLineageRelationRequest) {
	request = &DeleteLineageRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteLineageRelation", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteLineageRelationResponse creates a response to parse from DeleteLineageRelation response
func CreateDeleteLineageRelationResponse() (response *DeleteLineageRelationResponse) {
	response = &DeleteLineageRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
