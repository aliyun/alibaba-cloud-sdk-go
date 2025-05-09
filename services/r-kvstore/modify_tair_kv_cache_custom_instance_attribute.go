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

// ModifyTairKVCacheCustomInstanceAttribute invokes the r_kvstore.ModifyTairKVCacheCustomInstanceAttribute API synchronously
func (client *Client) ModifyTairKVCacheCustomInstanceAttribute(request *ModifyTairKVCacheCustomInstanceAttributeRequest) (response *ModifyTairKVCacheCustomInstanceAttributeResponse, err error) {
	response = CreateModifyTairKVCacheCustomInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTairKVCacheCustomInstanceAttributeWithChan invokes the r_kvstore.ModifyTairKVCacheCustomInstanceAttribute API asynchronously
func (client *Client) ModifyTairKVCacheCustomInstanceAttributeWithChan(request *ModifyTairKVCacheCustomInstanceAttributeRequest) (<-chan *ModifyTairKVCacheCustomInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyTairKVCacheCustomInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTairKVCacheCustomInstanceAttribute(request)
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

// ModifyTairKVCacheCustomInstanceAttributeWithCallback invokes the r_kvstore.ModifyTairKVCacheCustomInstanceAttribute API asynchronously
func (client *Client) ModifyTairKVCacheCustomInstanceAttributeWithCallback(request *ModifyTairKVCacheCustomInstanceAttributeRequest, callback func(response *ModifyTairKVCacheCustomInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTairKVCacheCustomInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyTairKVCacheCustomInstanceAttribute(request)
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

// ModifyTairKVCacheCustomInstanceAttributeRequest is the request struct for api ModifyTairKVCacheCustomInstanceAttribute
type ModifyTairKVCacheCustomInstanceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SourceBiz            string           `position:"Query" name:"SourceBiz"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	InstanceName         string           `position:"Query" name:"InstanceName"`
}

// ModifyTairKVCacheCustomInstanceAttributeResponse is the response struct for api ModifyTairKVCacheCustomInstanceAttribute
type ModifyTairKVCacheCustomInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyTairKVCacheCustomInstanceAttributeRequest creates a request to invoke ModifyTairKVCacheCustomInstanceAttribute API
func CreateModifyTairKVCacheCustomInstanceAttributeRequest() (request *ModifyTairKVCacheCustomInstanceAttributeRequest) {
	request = &ModifyTairKVCacheCustomInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyTairKVCacheCustomInstanceAttribute", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyTairKVCacheCustomInstanceAttributeResponse creates a response to parse from ModifyTairKVCacheCustomInstanceAttribute response
func CreateModifyTairKVCacheCustomInstanceAttributeResponse() (response *ModifyTairKVCacheCustomInstanceAttributeResponse) {
	response = &ModifyTairKVCacheCustomInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
