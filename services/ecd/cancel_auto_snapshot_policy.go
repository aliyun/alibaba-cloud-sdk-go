package ecd

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

// CancelAutoSnapshotPolicy invokes the ecd.CancelAutoSnapshotPolicy API synchronously
func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (response *CancelAutoSnapshotPolicyResponse, err error) {
	response = CreateCancelAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CancelAutoSnapshotPolicyWithChan invokes the ecd.CancelAutoSnapshotPolicy API asynchronously
func (client *Client) CancelAutoSnapshotPolicyWithChan(request *CancelAutoSnapshotPolicyRequest) (<-chan *CancelAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *CancelAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelAutoSnapshotPolicy(request)
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

// CancelAutoSnapshotPolicyWithCallback invokes the ecd.CancelAutoSnapshotPolicy API asynchronously
func (client *Client) CancelAutoSnapshotPolicyWithCallback(request *CancelAutoSnapshotPolicyRequest, callback func(response *CancelAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.CancelAutoSnapshotPolicy(request)
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

// CancelAutoSnapshotPolicyRequest is the request struct for api CancelAutoSnapshotPolicy
type CancelAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	PolicyId  string    `position:"Query" name:"PolicyId"`
	DesktopId *[]string `position:"Query" name:"DesktopId"  type:"Repeated"`
}

// CancelAutoSnapshotPolicyResponse is the response struct for api CancelAutoSnapshotPolicy
type CancelAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelAutoSnapshotPolicyRequest creates a request to invoke CancelAutoSnapshotPolicy API
func CreateCancelAutoSnapshotPolicyRequest() (request *CancelAutoSnapshotPolicyRequest) {
	request = &CancelAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CancelAutoSnapshotPolicy", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelAutoSnapshotPolicyResponse creates a response to parse from CancelAutoSnapshotPolicy response
func CreateCancelAutoSnapshotPolicyResponse() (response *CancelAutoSnapshotPolicyResponse) {
	response = &CancelAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
