package mse

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

// DeleteGatewayFlowRule invokes the mse.DeleteGatewayFlowRule API synchronously
func (client *Client) DeleteGatewayFlowRule(request *DeleteGatewayFlowRuleRequest) (response *DeleteGatewayFlowRuleResponse, err error) {
	response = CreateDeleteGatewayFlowRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayFlowRuleWithChan invokes the mse.DeleteGatewayFlowRule API asynchronously
func (client *Client) DeleteGatewayFlowRuleWithChan(request *DeleteGatewayFlowRuleRequest) (<-chan *DeleteGatewayFlowRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayFlowRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGatewayFlowRule(request)
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

// DeleteGatewayFlowRuleWithCallback invokes the mse.DeleteGatewayFlowRule API asynchronously
func (client *Client) DeleteGatewayFlowRuleWithCallback(request *DeleteGatewayFlowRuleRequest, callback func(response *DeleteGatewayFlowRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayFlowRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteGatewayFlowRule(request)
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

// DeleteGatewayFlowRuleRequest is the request struct for api DeleteGatewayFlowRule
type DeleteGatewayFlowRuleRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	RouteId         requests.Integer `position:"Query" name:"RouteId"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	RuleId          requests.Integer `position:"Query" name:"RuleId"`
}

// DeleteGatewayFlowRuleResponse is the response struct for api DeleteGatewayFlowRule
type DeleteGatewayFlowRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteGatewayFlowRuleRequest creates a request to invoke DeleteGatewayFlowRule API
func CreateDeleteGatewayFlowRuleRequest() (request *DeleteGatewayFlowRuleRequest) {
	request = &DeleteGatewayFlowRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteGatewayFlowRule", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayFlowRuleResponse creates a response to parse from DeleteGatewayFlowRule response
func CreateDeleteGatewayFlowRuleResponse() (response *DeleteGatewayFlowRuleResponse) {
	response = &DeleteGatewayFlowRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
