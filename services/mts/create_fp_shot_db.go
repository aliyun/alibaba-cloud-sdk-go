package mts

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

// CreateFpShotDB invokes the mts.CreateFpShotDB API synchronously
func (client *Client) CreateFpShotDB(request *CreateFpShotDBRequest) (response *CreateFpShotDBResponse, err error) {
	response = CreateCreateFpShotDBResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFpShotDBWithChan invokes the mts.CreateFpShotDB API asynchronously
func (client *Client) CreateFpShotDBWithChan(request *CreateFpShotDBRequest) (<-chan *CreateFpShotDBResponse, <-chan error) {
	responseChan := make(chan *CreateFpShotDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFpShotDB(request)
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

// CreateFpShotDBWithCallback invokes the mts.CreateFpShotDB API asynchronously
func (client *Client) CreateFpShotDBWithCallback(request *CreateFpShotDBRequest, callback func(response *CreateFpShotDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFpShotDBResponse
		var err error
		defer close(result)
		response, err = client.CreateFpShotDB(request)
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

// CreateFpShotDBRequest is the request struct for api CreateFpShotDB
type CreateFpShotDBRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ModelId              requests.Integer `position:"Query" name:"ModelId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Name                 string           `position:"Query" name:"Name"`
	Config               string           `position:"Query" name:"Config"`
}

// CreateFpShotDBResponse is the response struct for api CreateFpShotDB
type CreateFpShotDBResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	FpShotDB  FpShotDB `json:"FpShotDB" xml:"FpShotDB"`
}

// CreateCreateFpShotDBRequest creates a request to invoke CreateFpShotDB API
func CreateCreateFpShotDBRequest() (request *CreateFpShotDBRequest) {
	request = &CreateFpShotDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "CreateFpShotDB", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFpShotDBResponse creates a response to parse from CreateFpShotDB response
func CreateCreateFpShotDBResponse() (response *CreateFpShotDBResponse) {
	response = &CreateFpShotDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
