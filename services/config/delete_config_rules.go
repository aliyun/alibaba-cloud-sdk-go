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

// DeleteConfigRules invokes the config.DeleteConfigRules API synchronously
func (client *Client) DeleteConfigRules(request *DeleteConfigRulesRequest) (response *DeleteConfigRulesResponse, err error) {
	response = CreateDeleteConfigRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteConfigRulesWithChan invokes the config.DeleteConfigRules API asynchronously
func (client *Client) DeleteConfigRulesWithChan(request *DeleteConfigRulesRequest) (<-chan *DeleteConfigRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteConfigRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteConfigRules(request)
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

// DeleteConfigRulesWithCallback invokes the config.DeleteConfigRules API asynchronously
func (client *Client) DeleteConfigRulesWithCallback(request *DeleteConfigRulesRequest, callback func(response *DeleteConfigRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteConfigRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteConfigRules(request)
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

// DeleteConfigRulesRequest is the request struct for api DeleteConfigRules
type DeleteConfigRulesRequest struct {
	*requests.RpcRequest
	ConfigRuleIds string `position:"Query" name:"ConfigRuleIds"`
}

// DeleteConfigRulesResponse is the response struct for api DeleteConfigRules
type DeleteConfigRulesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateDeleteConfigRulesRequest creates a request to invoke DeleteConfigRules API
func CreateDeleteConfigRulesRequest() (request *DeleteConfigRulesRequest) {
	request = &DeleteConfigRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "DeleteConfigRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteConfigRulesResponse creates a response to parse from DeleteConfigRules response
func CreateDeleteConfigRulesResponse() (response *DeleteConfigRulesResponse) {
	response = &DeleteConfigRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
