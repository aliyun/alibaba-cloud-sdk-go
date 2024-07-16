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

// ModifyFirewallTemplate invokes the swas_open.ModifyFirewallTemplate API synchronously
func (client *Client) ModifyFirewallTemplate(request *ModifyFirewallTemplateRequest) (response *ModifyFirewallTemplateResponse, err error) {
	response = CreateModifyFirewallTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFirewallTemplateWithChan invokes the swas_open.ModifyFirewallTemplate API asynchronously
func (client *Client) ModifyFirewallTemplateWithChan(request *ModifyFirewallTemplateRequest) (<-chan *ModifyFirewallTemplateResponse, <-chan error) {
	responseChan := make(chan *ModifyFirewallTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFirewallTemplate(request)
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

// ModifyFirewallTemplateWithCallback invokes the swas_open.ModifyFirewallTemplate API asynchronously
func (client *Client) ModifyFirewallTemplateWithCallback(request *ModifyFirewallTemplateRequest, callback func(response *ModifyFirewallTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFirewallTemplateResponse
		var err error
		defer close(result)
		response, err = client.ModifyFirewallTemplate(request)
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

// ModifyFirewallTemplateRequest is the request struct for api ModifyFirewallTemplate
type ModifyFirewallTemplateRequest struct {
	*requests.RpcRequest
	FirewallTemplateId   string                                        `position:"Query" name:"FirewallTemplateId"`
	ClientToken          string                                        `position:"Query" name:"ClientToken"`
	Description          string                                        `position:"Query" name:"Description"`
	InstanceId           string                                        `position:"Query" name:"InstanceId"`
	FirewallTemplateRule *[]ModifyFirewallTemplateFirewallTemplateRule `position:"Query" name:"FirewallTemplateRule"  type:"Repeated"`
	Name                 string                                        `position:"Query" name:"Name"`
}

// ModifyFirewallTemplateFirewallTemplateRule is a repeated param struct in ModifyFirewallTemplateRequest
type ModifyFirewallTemplateFirewallTemplateRule struct {
	FirewallTemplateRuleId string `name:"FirewallTemplateRuleId"`
	RuleProtocol           string `name:"RuleProtocol"`
	Port                   string `name:"Port"`
	SourceCidrIp           string `name:"SourceCidrIp"`
	Remark                 string `name:"Remark"`
}

// ModifyFirewallTemplateResponse is the response struct for api ModifyFirewallTemplate
type ModifyFirewallTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyFirewallTemplateRequest creates a request to invoke ModifyFirewallTemplate API
func CreateModifyFirewallTemplateRequest() (request *ModifyFirewallTemplateRequest) {
	request = &ModifyFirewallTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ModifyFirewallTemplate", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyFirewallTemplateResponse creates a response to parse from ModifyFirewallTemplate response
func CreateModifyFirewallTemplateResponse() (response *ModifyFirewallTemplateResponse) {
	response = &ModifyFirewallTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
