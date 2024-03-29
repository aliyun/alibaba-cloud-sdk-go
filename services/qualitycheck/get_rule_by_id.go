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

// GetRuleById invokes the qualitycheck.GetRuleById API synchronously
func (client *Client) GetRuleById(request *GetRuleByIdRequest) (response *GetRuleByIdResponse, err error) {
	response = CreateGetRuleByIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleByIdWithChan invokes the qualitycheck.GetRuleById API asynchronously
func (client *Client) GetRuleByIdWithChan(request *GetRuleByIdRequest) (<-chan *GetRuleByIdResponse, <-chan error) {
	responseChan := make(chan *GetRuleByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleById(request)
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

// GetRuleByIdWithCallback invokes the qualitycheck.GetRuleById API asynchronously
func (client *Client) GetRuleByIdWithCallback(request *GetRuleByIdRequest, callback func(response *GetRuleByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleByIdResponse
		var err error
		defer close(result)
		response, err = client.GetRuleById(request)
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

// GetRuleByIdRequest is the request struct for api GetRuleById
type GetRuleByIdRequest struct {
	*requests.RpcRequest
	IsSchemeData  requests.Integer `position:"Body" name:"IsSchemeData"`
	RuleId        requests.Integer `position:"Body" name:"RuleId"`
	BaseMeAgentId requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// GetRuleByIdResponse is the response struct for api GetRuleById
type GetRuleByIdResponse struct {
	*responses.BaseResponse
}

// CreateGetRuleByIdRequest creates a request to invoke GetRuleById API
func CreateGetRuleByIdRequest() (request *GetRuleByIdRequest) {
	request = &GetRuleByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetRuleById", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRuleByIdResponse creates a response to parse from GetRuleById response
func CreateGetRuleByIdResponse() (response *GetRuleByIdResponse) {
	response = &GetRuleByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
