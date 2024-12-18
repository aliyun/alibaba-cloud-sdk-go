package selectdb

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

// CreateServiceLinkedRoleForSelectDB invokes the selectdb.CreateServiceLinkedRoleForSelectDB API synchronously
func (client *Client) CreateServiceLinkedRoleForSelectDB(request *CreateServiceLinkedRoleForSelectDBRequest) (response *CreateServiceLinkedRoleForSelectDBResponse, err error) {
	response = CreateCreateServiceLinkedRoleForSelectDBResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServiceLinkedRoleForSelectDBWithChan invokes the selectdb.CreateServiceLinkedRoleForSelectDB API asynchronously
func (client *Client) CreateServiceLinkedRoleForSelectDBWithChan(request *CreateServiceLinkedRoleForSelectDBRequest) (<-chan *CreateServiceLinkedRoleForSelectDBResponse, <-chan error) {
	responseChan := make(chan *CreateServiceLinkedRoleForSelectDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateServiceLinkedRoleForSelectDB(request)
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

// CreateServiceLinkedRoleForSelectDBWithCallback invokes the selectdb.CreateServiceLinkedRoleForSelectDB API asynchronously
func (client *Client) CreateServiceLinkedRoleForSelectDBWithCallback(request *CreateServiceLinkedRoleForSelectDBRequest, callback func(response *CreateServiceLinkedRoleForSelectDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServiceLinkedRoleForSelectDBResponse
		var err error
		defer close(result)
		response, err = client.CreateServiceLinkedRoleForSelectDB(request)
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

// CreateServiceLinkedRoleForSelectDBRequest is the request struct for api CreateServiceLinkedRoleForSelectDB
type CreateServiceLinkedRoleForSelectDBRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
}

// CreateServiceLinkedRoleForSelectDBResponse is the response struct for api CreateServiceLinkedRoleForSelectDB
type CreateServiceLinkedRoleForSelectDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateServiceLinkedRoleForSelectDBRequest creates a request to invoke CreateServiceLinkedRoleForSelectDB API
func CreateCreateServiceLinkedRoleForSelectDBRequest() (request *CreateServiceLinkedRoleForSelectDBRequest) {
	request = &CreateServiceLinkedRoleForSelectDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("selectdb", "2023-05-22", "CreateServiceLinkedRoleForSelectDB", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateServiceLinkedRoleForSelectDBResponse creates a response to parse from CreateServiceLinkedRoleForSelectDB response
func CreateCreateServiceLinkedRoleForSelectDBResponse() (response *CreateServiceLinkedRoleForSelectDBResponse) {
	response = &CreateServiceLinkedRoleForSelectDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
