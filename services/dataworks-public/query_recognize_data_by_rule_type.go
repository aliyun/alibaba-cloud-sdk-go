package dataworks_public

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

// QueryRecognizeDataByRuleType invokes the dataworks_public.QueryRecognizeDataByRuleType API synchronously
func (client *Client) QueryRecognizeDataByRuleType(request *QueryRecognizeDataByRuleTypeRequest) (response *QueryRecognizeDataByRuleTypeResponse, err error) {
	response = CreateQueryRecognizeDataByRuleTypeResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRecognizeDataByRuleTypeWithChan invokes the dataworks_public.QueryRecognizeDataByRuleType API asynchronously
func (client *Client) QueryRecognizeDataByRuleTypeWithChan(request *QueryRecognizeDataByRuleTypeRequest) (<-chan *QueryRecognizeDataByRuleTypeResponse, <-chan error) {
	responseChan := make(chan *QueryRecognizeDataByRuleTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRecognizeDataByRuleType(request)
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

// QueryRecognizeDataByRuleTypeWithCallback invokes the dataworks_public.QueryRecognizeDataByRuleType API asynchronously
func (client *Client) QueryRecognizeDataByRuleTypeWithCallback(request *QueryRecognizeDataByRuleTypeRequest, callback func(response *QueryRecognizeDataByRuleTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRecognizeDataByRuleTypeResponse
		var err error
		defer close(result)
		response, err = client.QueryRecognizeDataByRuleType(request)
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

// QueryRecognizeDataByRuleTypeRequest is the request struct for api QueryRecognizeDataByRuleType
type QueryRecognizeDataByRuleTypeRequest struct {
	*requests.RpcRequest
	TenantId           string `position:"Body" name:"TenantId"`
	RecognizeRulesType string `position:"Body" name:"RecognizeRulesType"`
}

// QueryRecognizeDataByRuleTypeResponse is the response struct for api QueryRecognizeDataByRuleType
type QueryRecognizeDataByRuleTypeResponse struct {
	*responses.BaseResponse
}

// CreateQueryRecognizeDataByRuleTypeRequest creates a request to invoke QueryRecognizeDataByRuleType API
func CreateQueryRecognizeDataByRuleTypeRequest() (request *QueryRecognizeDataByRuleTypeRequest) {
	request = &QueryRecognizeDataByRuleTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "QueryRecognizeDataByRuleType", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryRecognizeDataByRuleTypeResponse creates a response to parse from QueryRecognizeDataByRuleType response
func CreateQueryRecognizeDataByRuleTypeResponse() (response *QueryRecognizeDataByRuleTypeResponse) {
	response = &QueryRecognizeDataByRuleTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
