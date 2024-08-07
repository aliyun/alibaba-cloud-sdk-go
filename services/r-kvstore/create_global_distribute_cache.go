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

// CreateGlobalDistributeCache invokes the r_kvstore.CreateGlobalDistributeCache API synchronously
func (client *Client) CreateGlobalDistributeCache(request *CreateGlobalDistributeCacheRequest) (response *CreateGlobalDistributeCacheResponse, err error) {
	response = CreateCreateGlobalDistributeCacheResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGlobalDistributeCacheWithChan invokes the r_kvstore.CreateGlobalDistributeCache API asynchronously
func (client *Client) CreateGlobalDistributeCacheWithChan(request *CreateGlobalDistributeCacheRequest) (<-chan *CreateGlobalDistributeCacheResponse, <-chan error) {
	responseChan := make(chan *CreateGlobalDistributeCacheResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGlobalDistributeCache(request)
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

// CreateGlobalDistributeCacheWithCallback invokes the r_kvstore.CreateGlobalDistributeCache API asynchronously
func (client *Client) CreateGlobalDistributeCacheWithCallback(request *CreateGlobalDistributeCacheRequest, callback func(response *CreateGlobalDistributeCacheResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGlobalDistributeCacheResponse
		var err error
		defer close(result)
		response, err = client.CreateGlobalDistributeCache(request)
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

// CreateGlobalDistributeCacheRequest is the request struct for api CreateGlobalDistributeCache
type CreateGlobalDistributeCacheRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SeedSubInstanceId    string           `position:"Query" name:"SeedSubInstanceId"`
}

// CreateGlobalDistributeCacheResponse is the response struct for api CreateGlobalDistributeCache
type CreateGlobalDistributeCacheResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateGlobalDistributeCacheRequest creates a request to invoke CreateGlobalDistributeCache API
func CreateCreateGlobalDistributeCacheRequest() (request *CreateGlobalDistributeCacheRequest) {
	request = &CreateGlobalDistributeCacheRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateGlobalDistributeCache", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGlobalDistributeCacheResponse creates a response to parse from CreateGlobalDistributeCache response
func CreateCreateGlobalDistributeCacheResponse() (response *CreateGlobalDistributeCacheResponse) {
	response = &CreateGlobalDistributeCacheResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
