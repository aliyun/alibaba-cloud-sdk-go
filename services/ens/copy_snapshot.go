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

// CopySnapshot invokes the ens.CopySnapshot API synchronously
func (client *Client) CopySnapshot(request *CopySnapshotRequest) (response *CopySnapshotResponse, err error) {
	response = CreateCopySnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// CopySnapshotWithChan invokes the ens.CopySnapshot API asynchronously
func (client *Client) CopySnapshotWithChan(request *CopySnapshotRequest) (<-chan *CopySnapshotResponse, <-chan error) {
	responseChan := make(chan *CopySnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopySnapshot(request)
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

// CopySnapshotWithCallback invokes the ens.CopySnapshot API asynchronously
func (client *Client) CopySnapshotWithCallback(request *CopySnapshotRequest, callback func(response *CopySnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopySnapshotResponse
		var err error
		defer close(result)
		response, err = client.CopySnapshot(request)
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

// CopySnapshotRequest is the request struct for api CopySnapshot
type CopySnapshotRequest struct {
	*requests.RpcRequest
	SnapshotId                     string    `position:"Query" name:"SnapshotId"`
	DestinationRegionIds           *[]string `position:"Query" name:"DestinationRegionIds"  type:"Json"`
	DestinationSnapshotName        string    `position:"Query" name:"DestinationSnapshotName"`
	DestinationSnapshotDescription string    `position:"Query" name:"DestinationSnapshotDescription"`
}

// CopySnapshotResponse is the response struct for api CopySnapshot
type CopySnapshotResponse struct {
	*responses.BaseResponse
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	BizStatusCode  string               `json:"BizStatusCode" xml:"BizStatusCode"`
	AllocationId   []AllocationIdItem   `json:"AllocationId" xml:"AllocationId"`
	UnAllocationId []UnAllocationIdItem `json:"UnAllocationId" xml:"UnAllocationId"`
}

// CreateCopySnapshotRequest creates a request to invoke CopySnapshot API
func CreateCopySnapshotRequest() (request *CopySnapshotRequest) {
	request = &CopySnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CopySnapshot", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCopySnapshotResponse creates a response to parse from CopySnapshot response
func CreateCopySnapshotResponse() (response *CopySnapshotResponse) {
	response = &CopySnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
