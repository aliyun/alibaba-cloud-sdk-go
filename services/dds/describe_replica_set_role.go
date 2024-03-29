package dds

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

// DescribeReplicaSetRole invokes the dds.DescribeReplicaSetRole API synchronously
func (client *Client) DescribeReplicaSetRole(request *DescribeReplicaSetRoleRequest) (response *DescribeReplicaSetRoleResponse, err error) {
	response = CreateDescribeReplicaSetRoleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReplicaSetRoleWithChan invokes the dds.DescribeReplicaSetRole API asynchronously
func (client *Client) DescribeReplicaSetRoleWithChan(request *DescribeReplicaSetRoleRequest) (<-chan *DescribeReplicaSetRoleResponse, <-chan error) {
	responseChan := make(chan *DescribeReplicaSetRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReplicaSetRole(request)
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

// DescribeReplicaSetRoleWithCallback invokes the dds.DescribeReplicaSetRole API asynchronously
func (client *Client) DescribeReplicaSetRoleWithCallback(request *DescribeReplicaSetRoleRequest, callback func(response *DescribeReplicaSetRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReplicaSetRoleResponse
		var err error
		defer close(result)
		response, err = client.DescribeReplicaSetRole(request)
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

// DescribeReplicaSetRoleRequest is the request struct for api DescribeReplicaSetRole
type DescribeReplicaSetRoleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeReplicaSetRoleResponse is the response struct for api DescribeReplicaSetRole
type DescribeReplicaSetRoleResponse struct {
	*responses.BaseResponse
	RequestId    string                              `json:"RequestId" xml:"RequestId"`
	DBInstanceId string                              `json:"DBInstanceId" xml:"DBInstanceId"`
	ReplicaSets  ReplicaSetsInDescribeReplicaSetRole `json:"ReplicaSets" xml:"ReplicaSets"`
}

// CreateDescribeReplicaSetRoleRequest creates a request to invoke DescribeReplicaSetRole API
func CreateDescribeReplicaSetRoleRequest() (request *DescribeReplicaSetRoleRequest) {
	request = &DescribeReplicaSetRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeReplicaSetRole", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeReplicaSetRoleResponse creates a response to parse from DescribeReplicaSetRole response
func CreateDescribeReplicaSetRoleResponse() (response *DescribeReplicaSetRoleResponse) {
	response = &DescribeReplicaSetRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
