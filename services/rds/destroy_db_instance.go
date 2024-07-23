package rds

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

// DestroyDBInstance invokes the rds.DestroyDBInstance API synchronously
func (client *Client) DestroyDBInstance(request *DestroyDBInstanceRequest) (response *DestroyDBInstanceResponse, err error) {
	response = CreateDestroyDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DestroyDBInstanceWithChan invokes the rds.DestroyDBInstance API asynchronously
func (client *Client) DestroyDBInstanceWithChan(request *DestroyDBInstanceRequest) (<-chan *DestroyDBInstanceResponse, <-chan error) {
	responseChan := make(chan *DestroyDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DestroyDBInstance(request)
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

// DestroyDBInstanceWithCallback invokes the rds.DestroyDBInstance API asynchronously
func (client *Client) DestroyDBInstanceWithCallback(request *DestroyDBInstanceRequest, callback func(response *DestroyDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DestroyDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.DestroyDBInstance(request)
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

// DestroyDBInstanceRequest is the request struct for api DestroyDBInstance
type DestroyDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DestroyDBInstanceResponse is the response struct for api DestroyDBInstance
type DestroyDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDestroyDBInstanceRequest creates a request to invoke DestroyDBInstance API
func CreateDestroyDBInstanceRequest() (request *DestroyDBInstanceRequest) {
	request = &DestroyDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DestroyDBInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDestroyDBInstanceResponse creates a response to parse from DestroyDBInstance response
func CreateDestroyDBInstanceResponse() (response *DestroyDBInstanceResponse) {
	response = &DestroyDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
