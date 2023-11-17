package ens

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

// ModifySnapshotAttribute invokes the ens.ModifySnapshotAttribute API synchronously
func (client *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
	response = CreateModifySnapshotAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySnapshotAttributeWithChan invokes the ens.ModifySnapshotAttribute API asynchronously
func (client *Client) ModifySnapshotAttributeWithChan(request *ModifySnapshotAttributeRequest) (<-chan *ModifySnapshotAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifySnapshotAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySnapshotAttribute(request)
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

// ModifySnapshotAttributeWithCallback invokes the ens.ModifySnapshotAttribute API asynchronously
func (client *Client) ModifySnapshotAttributeWithCallback(request *ModifySnapshotAttributeRequest, callback func(response *ModifySnapshotAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySnapshotAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifySnapshotAttribute(request)
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

// ModifySnapshotAttributeRequest is the request struct for api ModifySnapshotAttribute
type ModifySnapshotAttributeRequest struct {
	*requests.RpcRequest
	SnapshotId   string `position:"Query" name:"SnapshotId"`
	Description  string `position:"Query" name:"Description"`
	SnapshotName string `position:"Query" name:"SnapshotName"`
}

// ModifySnapshotAttributeResponse is the response struct for api ModifySnapshotAttribute
type ModifySnapshotAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySnapshotAttributeRequest creates a request to invoke ModifySnapshotAttribute API
func CreateModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
	request = &ModifySnapshotAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ModifySnapshotAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySnapshotAttributeResponse creates a response to parse from ModifySnapshotAttribute response
func CreateModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
	response = &ModifySnapshotAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
