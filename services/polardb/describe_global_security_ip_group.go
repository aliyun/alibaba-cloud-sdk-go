package polardb

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

// DescribeGlobalSecurityIPGroup invokes the polardb.DescribeGlobalSecurityIPGroup API synchronously
func (client *Client) DescribeGlobalSecurityIPGroup(request *DescribeGlobalSecurityIPGroupRequest) (response *DescribeGlobalSecurityIPGroupResponse, err error) {
	response = CreateDescribeGlobalSecurityIPGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGlobalSecurityIPGroupWithChan invokes the polardb.DescribeGlobalSecurityIPGroup API asynchronously
func (client *Client) DescribeGlobalSecurityIPGroupWithChan(request *DescribeGlobalSecurityIPGroupRequest) (<-chan *DescribeGlobalSecurityIPGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeGlobalSecurityIPGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGlobalSecurityIPGroup(request)
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

// DescribeGlobalSecurityIPGroupWithCallback invokes the polardb.DescribeGlobalSecurityIPGroup API asynchronously
func (client *Client) DescribeGlobalSecurityIPGroupWithCallback(request *DescribeGlobalSecurityIPGroupRequest, callback func(response *DescribeGlobalSecurityIPGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGlobalSecurityIPGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeGlobalSecurityIPGroup(request)
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

// DescribeGlobalSecurityIPGroupRequest is the request struct for api DescribeGlobalSecurityIPGroup
type DescribeGlobalSecurityIPGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	GlobalSecurityGroupId string           `position:"Query" name:"GlobalSecurityGroupId"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeGlobalSecurityIPGroupResponse is the response struct for api DescribeGlobalSecurityIPGroup
type DescribeGlobalSecurityIPGroupResponse struct {
	*responses.BaseResponse
	RequestId             string                      `json:"RequestId" xml:"RequestId"`
	GlobalSecurityIPGroup []GlobalSecurityIPGroupItem `json:"GlobalSecurityIPGroup" xml:"GlobalSecurityIPGroup"`
}

// CreateDescribeGlobalSecurityIPGroupRequest creates a request to invoke DescribeGlobalSecurityIPGroup API
func CreateDescribeGlobalSecurityIPGroupRequest() (request *DescribeGlobalSecurityIPGroupRequest) {
	request = &DescribeGlobalSecurityIPGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeGlobalSecurityIPGroup", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGlobalSecurityIPGroupResponse creates a response to parse from DescribeGlobalSecurityIPGroup response
func CreateDescribeGlobalSecurityIPGroupResponse() (response *DescribeGlobalSecurityIPGroupResponse) {
	response = &DescribeGlobalSecurityIPGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
