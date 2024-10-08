package biz_ipdb

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

// AddIpPolicy invokes the biz_ipdb.AddIpPolicy API synchronously
func (client *Client) AddIpPolicy(request *AddIpPolicyRequest) (response *AddIpPolicyResponse, err error) {
	response = CreateAddIpPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// AddIpPolicyWithChan invokes the biz_ipdb.AddIpPolicy API asynchronously
func (client *Client) AddIpPolicyWithChan(request *AddIpPolicyRequest) (<-chan *AddIpPolicyResponse, <-chan error) {
	responseChan := make(chan *AddIpPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddIpPolicy(request)
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

// AddIpPolicyWithCallback invokes the biz_ipdb.AddIpPolicy API asynchronously
func (client *Client) AddIpPolicyWithCallback(request *AddIpPolicyRequest, callback func(response *AddIpPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddIpPolicyResponse
		var err error
		defer close(result)
		response, err = client.AddIpPolicy(request)
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

// AddIpPolicyRequest is the request struct for api AddIpPolicy
type AddIpPolicyRequest struct {
	*requests.RpcRequest
	PairValue   string           `position:"Query" name:"PairValue"`
	ActionTag   requests.Integer `position:"Query" name:"ActionTag"`
	Description string           `position:"Query" name:"Description"`
	Value       string           `position:"Query" name:"Value"`
}

// AddIpPolicyResponse is the response struct for api AddIpPolicy
type AddIpPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddIpPolicyRequest creates a request to invoke AddIpPolicy API
func CreateAddIpPolicyRequest() (request *AddIpPolicyRequest) {
	request = &AddIpPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Biz-ipdb", "2016-08-08", "AddIpPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateAddIpPolicyResponse creates a response to parse from AddIpPolicy response
func CreateAddIpPolicyResponse() (response *AddIpPolicyResponse) {
	response = &AddIpPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
