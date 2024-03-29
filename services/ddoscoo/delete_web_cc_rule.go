package ddoscoo

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

// DeleteWebCCRule invokes the ddoscoo.DeleteWebCCRule API synchronously
func (client *Client) DeleteWebCCRule(request *DeleteWebCCRuleRequest) (response *DeleteWebCCRuleResponse, err error) {
	response = CreateDeleteWebCCRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWebCCRuleWithChan invokes the ddoscoo.DeleteWebCCRule API asynchronously
func (client *Client) DeleteWebCCRuleWithChan(request *DeleteWebCCRuleRequest) (<-chan *DeleteWebCCRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteWebCCRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWebCCRule(request)
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

// DeleteWebCCRuleWithCallback invokes the ddoscoo.DeleteWebCCRule API asynchronously
func (client *Client) DeleteWebCCRuleWithCallback(request *DeleteWebCCRuleRequest, callback func(response *DeleteWebCCRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWebCCRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteWebCCRule(request)
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

// DeleteWebCCRuleRequest is the request struct for api DeleteWebCCRule
type DeleteWebCCRuleRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
	Name            string `position:"Query" name:"Name"`
}

// DeleteWebCCRuleResponse is the response struct for api DeleteWebCCRule
type DeleteWebCCRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteWebCCRuleRequest creates a request to invoke DeleteWebCCRule API
func CreateDeleteWebCCRuleRequest() (request *DeleteWebCCRuleRequest) {
	request = &DeleteWebCCRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DeleteWebCCRule", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteWebCCRuleResponse creates a response to parse from DeleteWebCCRule response
func CreateDeleteWebCCRuleResponse() (response *DeleteWebCCRuleResponse) {
	response = &DeleteWebCCRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
