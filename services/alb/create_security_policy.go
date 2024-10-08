package alb

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

// CreateSecurityPolicy invokes the alb.CreateSecurityPolicy API synchronously
func (client *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
	response = CreateCreateSecurityPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSecurityPolicyWithChan invokes the alb.CreateSecurityPolicy API asynchronously
func (client *Client) CreateSecurityPolicyWithChan(request *CreateSecurityPolicyRequest) (<-chan *CreateSecurityPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateSecurityPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSecurityPolicy(request)
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

// CreateSecurityPolicyWithCallback invokes the alb.CreateSecurityPolicy API asynchronously
func (client *Client) CreateSecurityPolicyWithCallback(request *CreateSecurityPolicyRequest, callback func(response *CreateSecurityPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSecurityPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateSecurityPolicy(request)
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

// CreateSecurityPolicyRequest is the request struct for api CreateSecurityPolicy
type CreateSecurityPolicyRequest struct {
	*requests.RpcRequest
	ClientToken        string                     `position:"Query" name:"ClientToken"`
	ResourceGroupId    string                     `position:"Query" name:"ResourceGroupId"`
	Ciphers            *[]string                  `position:"Query" name:"Ciphers"  type:"Repeated"`
	Tag                *[]CreateSecurityPolicyTag `position:"Query" name:"Tag"  type:"Repeated"`
	TLSVersions        *[]string                  `position:"Query" name:"TLSVersions"  type:"Repeated"`
	SecurityPolicyName string                     `position:"Query" name:"SecurityPolicyName"`
	DryRun             requests.Boolean           `position:"Query" name:"DryRun"`
}

// CreateSecurityPolicyTag is a repeated param struct in CreateSecurityPolicyRequest
type CreateSecurityPolicyTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateSecurityPolicyResponse is the response struct for api CreateSecurityPolicy
type CreateSecurityPolicyResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	SecurityPolicyId string `json:"SecurityPolicyId" xml:"SecurityPolicyId"`
}

// CreateCreateSecurityPolicyRequest creates a request to invoke CreateSecurityPolicy API
func CreateCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
	request = &CreateSecurityPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "CreateSecurityPolicy", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSecurityPolicyResponse creates a response to parse from CreateSecurityPolicy response
func CreateCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
	response = &CreateSecurityPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
