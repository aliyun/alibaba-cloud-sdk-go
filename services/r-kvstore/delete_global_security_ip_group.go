package r_kvstore

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

// DeleteGlobalSecurityIPGroup invokes the r_kvstore.DeleteGlobalSecurityIPGroup API synchronously
func (client *Client) DeleteGlobalSecurityIPGroup(request *DeleteGlobalSecurityIPGroupRequest) (response *DeleteGlobalSecurityIPGroupResponse, err error) {
	response = CreateDeleteGlobalSecurityIPGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGlobalSecurityIPGroupWithChan invokes the r_kvstore.DeleteGlobalSecurityIPGroup API asynchronously
func (client *Client) DeleteGlobalSecurityIPGroupWithChan(request *DeleteGlobalSecurityIPGroupRequest) (<-chan *DeleteGlobalSecurityIPGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteGlobalSecurityIPGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGlobalSecurityIPGroup(request)
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

// DeleteGlobalSecurityIPGroupWithCallback invokes the r_kvstore.DeleteGlobalSecurityIPGroup API asynchronously
func (client *Client) DeleteGlobalSecurityIPGroupWithCallback(request *DeleteGlobalSecurityIPGroupRequest, callback func(response *DeleteGlobalSecurityIPGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGlobalSecurityIPGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteGlobalSecurityIPGroup(request)
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

// DeleteGlobalSecurityIPGroupRequest is the request struct for api DeleteGlobalSecurityIPGroup
type DeleteGlobalSecurityIPGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	GlobalSecurityGroupId string           `position:"Query" name:"GlobalSecurityGroupId"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	GlobalIgName          string           `position:"Query" name:"GlobalIgName"`
}

// DeleteGlobalSecurityIPGroupResponse is the response struct for api DeleteGlobalSecurityIPGroup
type DeleteGlobalSecurityIPGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteGlobalSecurityIPGroupRequest creates a request to invoke DeleteGlobalSecurityIPGroup API
func CreateDeleteGlobalSecurityIPGroupRequest() (request *DeleteGlobalSecurityIPGroupRequest) {
	request = &DeleteGlobalSecurityIPGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteGlobalSecurityIPGroup", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGlobalSecurityIPGroupResponse creates a response to parse from DeleteGlobalSecurityIPGroup response
func CreateDeleteGlobalSecurityIPGroupResponse() (response *DeleteGlobalSecurityIPGroupResponse) {
	response = &DeleteGlobalSecurityIPGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
