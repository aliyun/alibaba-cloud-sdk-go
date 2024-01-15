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

// GetRuleCategory invokes the qualitycheck.GetRuleCategory API synchronously
func (client *Client) GetRuleCategory(request *GetRuleCategoryRequest) (response *GetRuleCategoryResponse, err error) {
	response = CreateGetRuleCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleCategoryWithChan invokes the qualitycheck.GetRuleCategory API asynchronously
func (client *Client) GetRuleCategoryWithChan(request *GetRuleCategoryRequest) (<-chan *GetRuleCategoryResponse, <-chan error) {
	responseChan := make(chan *GetRuleCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleCategory(request)
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

// GetRuleCategoryWithCallback invokes the qualitycheck.GetRuleCategory API asynchronously
func (client *Client) GetRuleCategoryWithCallback(request *GetRuleCategoryRequest, callback func(response *GetRuleCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleCategoryResponse
		var err error
		defer close(result)
		response, err = client.GetRuleCategory(request)
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

// GetRuleCategoryRequest is the request struct for api GetRuleCategory
type GetRuleCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// GetRuleCategoryResponse is the response struct for api GetRuleCategory
type GetRuleCategoryResponse struct {
	*responses.BaseResponse
	Code      string                `json:"Code" xml:"Code"`
	Message   string                `json:"Message" xml:"Message"`
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Success   bool                  `json:"Success" xml:"Success"`
	Data      DataInGetRuleCategory `json:"Data" xml:"Data"`
}

// CreateGetRuleCategoryRequest creates a request to invoke GetRuleCategory API
func CreateGetRuleCategoryRequest() (request *GetRuleCategoryRequest) {
	request = &GetRuleCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetRuleCategory", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRuleCategoryResponse creates a response to parse from GetRuleCategory response
func CreateGetRuleCategoryResponse() (response *GetRuleCategoryResponse) {
	response = &GetRuleCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
