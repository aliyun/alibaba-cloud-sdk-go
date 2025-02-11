package swas_open

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

// ResetSystem invokes the swas_open.ResetSystem API synchronously
func (client *Client) ResetSystem(request *ResetSystemRequest) (response *ResetSystemResponse, err error) {
	response = CreateResetSystemResponse()
	err = client.DoAction(request, response)
	return
}

// ResetSystemWithChan invokes the swas_open.ResetSystem API asynchronously
func (client *Client) ResetSystemWithChan(request *ResetSystemRequest) (<-chan *ResetSystemResponse, <-chan error) {
	responseChan := make(chan *ResetSystemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetSystem(request)
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

// ResetSystemWithCallback invokes the swas_open.ResetSystem API asynchronously
func (client *Client) ResetSystemWithCallback(request *ResetSystemRequest, callback func(response *ResetSystemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetSystemResponse
		var err error
		defer close(result)
		response, err = client.ResetSystem(request)
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

// ResetSystemRequest is the request struct for api ResetSystem
type ResetSystemRequest struct {
	*requests.RpcRequest
	ImageId          string                      `position:"Query" name:"ImageId"`
	ClientToken      string                      `position:"Query" name:"ClientToken"`
	LoginCredentials ResetSystemLoginCredentials `position:"Query" name:"LoginCredentials"  type:"Struct"`
	InstanceId       string                      `position:"Query" name:"InstanceId"`
}

// ResetSystemLoginCredentials is a repeated param struct in ResetSystemRequest
type ResetSystemLoginCredentials struct {
	Password    string `name:"Password"`
	KeyPairName string `name:"KeyPairName"`
}

// ResetSystemResponse is the response struct for api ResetSystem
type ResetSystemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetSystemRequest creates a request to invoke ResetSystem API
func CreateResetSystemRequest() (request *ResetSystemRequest) {
	request = &ResetSystemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ResetSystem", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetSystemResponse creates a response to parse from ResetSystem response
func CreateResetSystemResponse() (response *ResetSystemResponse) {
	response = &ResetSystemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
