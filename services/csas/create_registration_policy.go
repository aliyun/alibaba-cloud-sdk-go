package csas

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

// CreateRegistrationPolicy invokes the csas.CreateRegistrationPolicy API synchronously
func (client *Client) CreateRegistrationPolicy(request *CreateRegistrationPolicyRequest) (response *CreateRegistrationPolicyResponse, err error) {
	response = CreateCreateRegistrationPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRegistrationPolicyWithChan invokes the csas.CreateRegistrationPolicy API asynchronously
func (client *Client) CreateRegistrationPolicyWithChan(request *CreateRegistrationPolicyRequest) (<-chan *CreateRegistrationPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateRegistrationPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRegistrationPolicy(request)
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

// CreateRegistrationPolicyWithCallback invokes the csas.CreateRegistrationPolicy API asynchronously
func (client *Client) CreateRegistrationPolicyWithCallback(request *CreateRegistrationPolicyRequest, callback func(response *CreateRegistrationPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRegistrationPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateRegistrationPolicy(request)
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

// CreateRegistrationPolicyRequest is the request struct for api CreateRegistrationPolicy
type CreateRegistrationPolicyRequest struct {
	*requests.RpcRequest
	Description        string                                     `position:"Body" name:"Description"`
	MatchMode          string                                     `position:"Body" name:"MatchMode"`
	SourceIp           string                                     `position:"Query" name:"SourceIp"`
	CompanyLimitCount  CreateRegistrationPolicyCompanyLimitCount  `position:"Body" name:"CompanyLimitCount"  type:"Struct"`
	PersonalLimitCount CreateRegistrationPolicyPersonalLimitCount `position:"Body" name:"PersonalLimitCount"  type:"Struct"`
	UserGroupIds       *[]string                                  `position:"Body" name:"UserGroupIds"  type:"Repeated"`
	Whitelist          *[]string                                  `position:"Body" name:"Whitelist"  type:"Repeated"`
	Priority           requests.Integer                           `position:"Body" name:"Priority"`
	PersonalLimitType  string                                     `position:"Body" name:"PersonalLimitType"`
	Name               string                                     `position:"Body" name:"Name"`
	CompanyLimitType   string                                     `position:"Body" name:"CompanyLimitType"`
	Status             string                                     `position:"Body" name:"Status"`
}

// CreateRegistrationPolicyCompanyLimitCount is a repeated param struct in CreateRegistrationPolicyRequest
type CreateRegistrationPolicyCompanyLimitCount struct {
	All    string `name:"All"`
	PC     string `name:"PC"`
	Mobile string `name:"Mobile"`
}

// CreateRegistrationPolicyPersonalLimitCount is a repeated param struct in CreateRegistrationPolicyRequest
type CreateRegistrationPolicyPersonalLimitCount struct {
	All    string `name:"All"`
	PC     string `name:"PC"`
	Mobile string `name:"Mobile"`
}

// CreateRegistrationPolicyResponse is the response struct for api CreateRegistrationPolicy
type CreateRegistrationPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Policy    Policy `json:"Policy" xml:"Policy"`
}

// CreateCreateRegistrationPolicyRequest creates a request to invoke CreateRegistrationPolicy API
func CreateCreateRegistrationPolicyRequest() (request *CreateRegistrationPolicyRequest) {
	request = &CreateRegistrationPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "CreateRegistrationPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRegistrationPolicyResponse creates a response to parse from CreateRegistrationPolicy response
func CreateCreateRegistrationPolicyResponse() (response *CreateRegistrationPolicyResponse) {
	response = &CreateRegistrationPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
