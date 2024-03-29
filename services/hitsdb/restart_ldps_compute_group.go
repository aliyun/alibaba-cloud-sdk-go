package hitsdb

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

// RestartLdpsComputeGroup invokes the hitsdb.RestartLdpsComputeGroup API synchronously
func (client *Client) RestartLdpsComputeGroup(request *RestartLdpsComputeGroupRequest) (response *RestartLdpsComputeGroupResponse, err error) {
	response = CreateRestartLdpsComputeGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RestartLdpsComputeGroupWithChan invokes the hitsdb.RestartLdpsComputeGroup API asynchronously
func (client *Client) RestartLdpsComputeGroupWithChan(request *RestartLdpsComputeGroupRequest) (<-chan *RestartLdpsComputeGroupResponse, <-chan error) {
	responseChan := make(chan *RestartLdpsComputeGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartLdpsComputeGroup(request)
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

// RestartLdpsComputeGroupWithCallback invokes the hitsdb.RestartLdpsComputeGroup API asynchronously
func (client *Client) RestartLdpsComputeGroupWithCallback(request *RestartLdpsComputeGroupRequest, callback func(response *RestartLdpsComputeGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartLdpsComputeGroupResponse
		var err error
		defer close(result)
		response, err = client.RestartLdpsComputeGroup(request)
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

// RestartLdpsComputeGroupRequest is the request struct for api RestartLdpsComputeGroup
type RestartLdpsComputeGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// RestartLdpsComputeGroupResponse is the response struct for api RestartLdpsComputeGroup
type RestartLdpsComputeGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestartLdpsComputeGroupRequest creates a request to invoke RestartLdpsComputeGroup API
func CreateRestartLdpsComputeGroupRequest() (request *RestartLdpsComputeGroupRequest) {
	request = &RestartLdpsComputeGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "RestartLdpsComputeGroup", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestartLdpsComputeGroupResponse creates a response to parse from RestartLdpsComputeGroup response
func CreateRestartLdpsComputeGroupResponse() (response *RestartLdpsComputeGroupResponse) {
	response = &RestartLdpsComputeGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
