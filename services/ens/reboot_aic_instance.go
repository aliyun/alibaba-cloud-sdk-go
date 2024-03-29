package ens

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

// RebootAICInstance invokes the ens.RebootAICInstance API synchronously
func (client *Client) RebootAICInstance(request *RebootAICInstanceRequest) (response *RebootAICInstanceResponse, err error) {
	response = CreateRebootAICInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RebootAICInstanceWithChan invokes the ens.RebootAICInstance API asynchronously
func (client *Client) RebootAICInstanceWithChan(request *RebootAICInstanceRequest) (<-chan *RebootAICInstanceResponse, <-chan error) {
	responseChan := make(chan *RebootAICInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebootAICInstance(request)
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

// RebootAICInstanceWithCallback invokes the ens.RebootAICInstance API asynchronously
func (client *Client) RebootAICInstanceWithCallback(request *RebootAICInstanceRequest, callback func(response *RebootAICInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebootAICInstanceResponse
		var err error
		defer close(result)
		response, err = client.RebootAICInstance(request)
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

// RebootAICInstanceRequest is the request struct for api RebootAICInstance
type RebootAICInstanceRequest struct {
	*requests.RpcRequest
	ServerId    string    `position:"Query" name:"ServerId"`
	InstanceId  string    `position:"Query" name:"InstanceId"`
	InstanceIds *[]string `position:"Query" name:"InstanceIds"  type:"Json"`
}

// RebootAICInstanceResponse is the response struct for api RebootAICInstance
type RebootAICInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRebootAICInstanceRequest creates a request to invoke RebootAICInstance API
func CreateRebootAICInstanceRequest() (request *RebootAICInstanceRequest) {
	request = &RebootAICInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "RebootAICInstance", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateRebootAICInstanceResponse creates a response to parse from RebootAICInstance response
func CreateRebootAICInstanceResponse() (response *RebootAICInstanceResponse) {
	response = &RebootAICInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
