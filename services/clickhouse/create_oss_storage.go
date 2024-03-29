package clickhouse

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

// CreateOSSStorage invokes the clickhouse.CreateOSSStorage API synchronously
func (client *Client) CreateOSSStorage(request *CreateOSSStorageRequest) (response *CreateOSSStorageResponse, err error) {
	response = CreateCreateOSSStorageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOSSStorageWithChan invokes the clickhouse.CreateOSSStorage API asynchronously
func (client *Client) CreateOSSStorageWithChan(request *CreateOSSStorageRequest) (<-chan *CreateOSSStorageResponse, <-chan error) {
	responseChan := make(chan *CreateOSSStorageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOSSStorage(request)
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

// CreateOSSStorageWithCallback invokes the clickhouse.CreateOSSStorage API asynchronously
func (client *Client) CreateOSSStorageWithCallback(request *CreateOSSStorageRequest, callback func(response *CreateOSSStorageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOSSStorageResponse
		var err error
		defer close(result)
		response, err = client.CreateOSSStorage(request)
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

// CreateOSSStorageRequest is the request struct for api CreateOSSStorage
type CreateOSSStorageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateOSSStorageResponse is the response struct for api CreateOSSStorage
type CreateOSSStorageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOSSStorageRequest creates a request to invoke CreateOSSStorage API
func CreateCreateOSSStorageRequest() (request *CreateOSSStorageRequest) {
	request = &CreateOSSStorageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "CreateOSSStorage", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateOSSStorageResponse creates a response to parse from CreateOSSStorage response
func CreateCreateOSSStorageResponse() (response *CreateOSSStorageResponse) {
	response = &CreateOSSStorageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
