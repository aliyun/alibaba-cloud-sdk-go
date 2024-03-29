package qualitycheck

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

// GetRule invokes the qualitycheck.GetRule API synchronously
func (client *Client) GetRule(request *GetRuleRequest) (response *GetRuleResponse, err error) {
	response = CreateGetRuleResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleWithChan invokes the qualitycheck.GetRule API asynchronously
func (client *Client) GetRuleWithChan(request *GetRuleRequest) (<-chan *GetRuleResponse, <-chan error) {
	responseChan := make(chan *GetRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRule(request)
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

// GetRuleWithCallback invokes the qualitycheck.GetRule API asynchronously
func (client *Client) GetRuleWithCallback(request *GetRuleRequest, callback func(response *GetRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleResponse
		var err error
		defer close(result)
		response, err = client.GetRule(request)
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

// GetRuleRequest is the request struct for api GetRule
type GetRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// GetRuleResponse is the response struct for api GetRule
type GetRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetRuleRequest creates a request to invoke GetRule API
func CreateGetRuleRequest() (request *GetRuleRequest) {
	request = &GetRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetRule", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRuleResponse creates a response to parse from GetRule response
func CreateGetRuleResponse() (response *GetRuleResponse) {
	response = &GetRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
