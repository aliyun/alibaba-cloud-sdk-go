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

// ReleaseLindormInstance invokes the hitsdb.ReleaseLindormInstance API synchronously
func (client *Client) ReleaseLindormInstance(request *ReleaseLindormInstanceRequest) (response *ReleaseLindormInstanceResponse, err error) {
	response = CreateReleaseLindormInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseLindormInstanceWithChan invokes the hitsdb.ReleaseLindormInstance API asynchronously
func (client *Client) ReleaseLindormInstanceWithChan(request *ReleaseLindormInstanceRequest) (<-chan *ReleaseLindormInstanceResponse, <-chan error) {
	responseChan := make(chan *ReleaseLindormInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseLindormInstance(request)
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

// ReleaseLindormInstanceWithCallback invokes the hitsdb.ReleaseLindormInstance API asynchronously
func (client *Client) ReleaseLindormInstanceWithCallback(request *ReleaseLindormInstanceRequest, callback func(response *ReleaseLindormInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseLindormInstanceResponse
		var err error
		defer close(result)
		response, err = client.ReleaseLindormInstance(request)
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

// ReleaseLindormInstanceRequest is the request struct for api ReleaseLindormInstance
type ReleaseLindormInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Immediately          requests.Boolean `position:"Query" name:"Immediately"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ReleaseLindormInstanceResponse is the response struct for api ReleaseLindormInstance
type ReleaseLindormInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseLindormInstanceRequest creates a request to invoke ReleaseLindormInstance API
func CreateReleaseLindormInstanceRequest() (request *ReleaseLindormInstanceRequest) {
	request = &ReleaseLindormInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "ReleaseLindormInstance", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleaseLindormInstanceResponse creates a response to parse from ReleaseLindormInstance response
func CreateReleaseLindormInstanceResponse() (response *ReleaseLindormInstanceResponse) {
	response = &ReleaseLindormInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
