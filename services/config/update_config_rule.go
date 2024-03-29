package config

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

// UpdateConfigRule invokes the config.UpdateConfigRule API synchronously
func (client *Client) UpdateConfigRule(request *UpdateConfigRuleRequest) (response *UpdateConfigRuleResponse, err error) {
	response = CreateUpdateConfigRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConfigRuleWithChan invokes the config.UpdateConfigRule API asynchronously
func (client *Client) UpdateConfigRuleWithChan(request *UpdateConfigRuleRequest) (<-chan *UpdateConfigRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateConfigRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConfigRule(request)
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

// UpdateConfigRuleWithCallback invokes the config.UpdateConfigRule API asynchronously
func (client *Client) UpdateConfigRuleWithCallback(request *UpdateConfigRuleRequest, callback func(response *UpdateConfigRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConfigRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateConfigRule(request)
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

// UpdateConfigRuleRequest is the request struct for api UpdateConfigRule
type UpdateConfigRuleRequest struct {
	*requests.RpcRequest
	ConfigRuleId              string           `position:"Body" name:"ConfigRuleId"`
	TagKeyScope               string           `position:"Body" name:"TagKeyScope"`
	ClientToken               string           `position:"Body" name:"ClientToken"`
	ResourceTypesScope        *[]string        `position:"Body" name:"ResourceTypesScope"  type:"Repeated"`
	Description               string           `position:"Body" name:"Description"`
	ConfigRuleTriggerTypes    string           `position:"Body" name:"ConfigRuleTriggerTypes"`
	TagValueScope             string           `position:"Body" name:"TagValueScope"`
	RegionIdsScope            string           `position:"Body" name:"RegionIdsScope"`
	RiskLevel                 requests.Integer `position:"Body" name:"RiskLevel"`
	ResourceGroupIdsScope     string           `position:"Body" name:"ResourceGroupIdsScope"`
	InputParameters           string           `position:"Body" name:"InputParameters"`
	ConfigRuleName            string           `position:"Body" name:"ConfigRuleName"`
	TagKeyLogicScope          string           `position:"Body" name:"TagKeyLogicScope"`
	MaximumExecutionFrequency string           `position:"Body" name:"MaximumExecutionFrequency"`
	ExcludeResourceIdsScope   string           `position:"Body" name:"ExcludeResourceIdsScope"`
	Conditions                string           `position:"Body" name:"Conditions"`
}

// UpdateConfigRuleResponse is the response struct for api UpdateConfigRule
type UpdateConfigRuleResponse struct {
	*responses.BaseResponse
	ConfigRuleId string `json:"ConfigRuleId" xml:"ConfigRuleId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateConfigRuleRequest creates a request to invoke UpdateConfigRule API
func CreateUpdateConfigRuleRequest() (request *UpdateConfigRuleRequest) {
	request = &UpdateConfigRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "UpdateConfigRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateConfigRuleResponse creates a response to parse from UpdateConfigRule response
func CreateUpdateConfigRuleResponse() (response *UpdateConfigRuleResponse) {
	response = &UpdateConfigRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
