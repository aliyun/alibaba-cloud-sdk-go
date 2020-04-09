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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeAutoSnapshotTasks invokes the nas.DescribeAutoSnapshotTasks API synchronously
// api document: https://help.aliyun.com/api/nas/describeautosnapshottasks.html
func (client *Client) DescribeAutoSnapshotTasks(request *DescribeAutoSnapshotTasksRequest) (response *DescribeAutoSnapshotTasksResponse, err error) {
	response = CreateDescribeAutoSnapshotTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoSnapshotTasksWithChan invokes the nas.DescribeAutoSnapshotTasks API asynchronously
// api document: https://help.aliyun.com/api/nas/describeautosnapshottasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAutoSnapshotTasksWithChan(request *DescribeAutoSnapshotTasksRequest) (<-chan *DescribeAutoSnapshotTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoSnapshotTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoSnapshotTasks(request)
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

// DescribeAutoSnapshotTasksWithCallback invokes the nas.DescribeAutoSnapshotTasks API asynchronously
// api document: https://help.aliyun.com/api/nas/describeautosnapshottasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAutoSnapshotTasksWithCallback(request *DescribeAutoSnapshotTasksRequest, callback func(response *DescribeAutoSnapshotTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoSnapshotTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoSnapshotTasks(request)
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

// DescribeAutoSnapshotTasksRequest is the request struct for api DescribeAutoSnapshotTasks
type DescribeAutoSnapshotTasksRequest struct {
	*requests.RpcRequest
	FileSystemIds         string           `position:"Query" name:"FileSystemIds"`
	AutoSnapshotPolicyIds string           `position:"Query" name:"AutoSnapshotPolicyIds"`
	FileSystemType        string           `position:"Query" name:"FileSystemType"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeAutoSnapshotTasksResponse is the response struct for api DescribeAutoSnapshotTasks
type DescribeAutoSnapshotTasksResponse struct {
	*responses.BaseResponse
	RequestId         string                                       `json:"RequestId" xml:"RequestId"`
	TotalCount        int                                          `json:"TotalCount" xml:"TotalCount"`
	PageSize          int                                          `json:"PageSize" xml:"PageSize"`
	PageNumber        int                                          `json:"PageNumber" xml:"PageNumber"`
	AutoSnapshotTasks []DescribeAutoSnapshotTasksAutoSnapshotTask0 `json:"AutoSnapshotTasks" xml:"AutoSnapshotTasks"`
}

type DescribeAutoSnapshotTasksAutoSnapshotTask0 struct {
	SourceFileSystemId   string `json:"SourceFileSystemId" xml:"SourceFileSystemId"`
	AutoSnapshotPolicyId string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
}

// CreateDescribeAutoSnapshotTasksRequest creates a request to invoke DescribeAutoSnapshotTasks API
func CreateDescribeAutoSnapshotTasksRequest() (request *DescribeAutoSnapshotTasksRequest) {
	request = &DescribeAutoSnapshotTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeAutoSnapshotTasks", "nas", "openAPI")
	return
}

// CreateDescribeAutoSnapshotTasksResponse creates a response to parse from DescribeAutoSnapshotTasks response
func CreateDescribeAutoSnapshotTasksResponse() (response *DescribeAutoSnapshotTasksResponse) {
	response = &DescribeAutoSnapshotTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
