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

// ResetTairKVCacheCustomInstancePassword invokes the r_kvstore.ResetTairKVCacheCustomInstancePassword API synchronously
func (client *Client) ResetTairKVCacheCustomInstancePassword(request *ResetTairKVCacheCustomInstancePasswordRequest) (response *ResetTairKVCacheCustomInstancePasswordResponse, err error) {
	response = CreateResetTairKVCacheCustomInstancePasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ResetTairKVCacheCustomInstancePasswordWithChan invokes the r_kvstore.ResetTairKVCacheCustomInstancePassword API asynchronously
func (client *Client) ResetTairKVCacheCustomInstancePasswordWithChan(request *ResetTairKVCacheCustomInstancePasswordRequest) (<-chan *ResetTairKVCacheCustomInstancePasswordResponse, <-chan error) {
	responseChan := make(chan *ResetTairKVCacheCustomInstancePasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetTairKVCacheCustomInstancePassword(request)
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

// ResetTairKVCacheCustomInstancePasswordWithCallback invokes the r_kvstore.ResetTairKVCacheCustomInstancePassword API asynchronously
func (client *Client) ResetTairKVCacheCustomInstancePasswordWithCallback(request *ResetTairKVCacheCustomInstancePasswordRequest, callback func(response *ResetTairKVCacheCustomInstancePasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetTairKVCacheCustomInstancePasswordResponse
		var err error
		defer close(result)
		response, err = client.ResetTairKVCacheCustomInstancePassword(request)
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

// ResetTairKVCacheCustomInstancePasswordRequest is the request struct for api ResetTairKVCacheCustomInstancePassword
type ResetTairKVCacheCustomInstancePasswordRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Password             string           `position:"Query" name:"Password"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SourceBiz            string           `position:"Query" name:"SourceBiz"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ResetTairKVCacheCustomInstancePasswordResponse is the response struct for api ResetTairKVCacheCustomInstancePassword
type ResetTairKVCacheCustomInstancePasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetTairKVCacheCustomInstancePasswordRequest creates a request to invoke ResetTairKVCacheCustomInstancePassword API
func CreateResetTairKVCacheCustomInstancePasswordRequest() (request *ResetTairKVCacheCustomInstancePasswordRequest) {
	request = &ResetTairKVCacheCustomInstancePasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ResetTairKVCacheCustomInstancePassword", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetTairKVCacheCustomInstancePasswordResponse creates a response to parse from ResetTairKVCacheCustomInstancePassword response
func CreateResetTairKVCacheCustomInstancePasswordResponse() (response *ResetTairKVCacheCustomInstancePasswordResponse) {
	response = &ResetTairKVCacheCustomInstancePasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
