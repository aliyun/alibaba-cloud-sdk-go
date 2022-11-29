package das

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

// DisableAutoResourceOptimizeRules invokes the das.DisableAutoResourceOptimizeRules API synchronously
func (client *Client) DisableAutoResourceOptimizeRules(request *DisableAutoResourceOptimizeRulesRequest) (response *DisableAutoResourceOptimizeRulesResponse, err error) {
	response = CreateDisableAutoResourceOptimizeRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DisableAutoResourceOptimizeRulesWithChan invokes the das.DisableAutoResourceOptimizeRules API asynchronously
func (client *Client) DisableAutoResourceOptimizeRulesWithChan(request *DisableAutoResourceOptimizeRulesRequest) (<-chan *DisableAutoResourceOptimizeRulesResponse, <-chan error) {
	responseChan := make(chan *DisableAutoResourceOptimizeRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableAutoResourceOptimizeRules(request)
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

// DisableAutoResourceOptimizeRulesWithCallback invokes the das.DisableAutoResourceOptimizeRules API asynchronously
func (client *Client) DisableAutoResourceOptimizeRulesWithCallback(request *DisableAutoResourceOptimizeRulesRequest, callback func(response *DisableAutoResourceOptimizeRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableAutoResourceOptimizeRulesResponse
		var err error
		defer close(result)
		response, err = client.DisableAutoResourceOptimizeRules(request)
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

// DisableAutoResourceOptimizeRulesRequest is the request struct for api DisableAutoResourceOptimizeRules
type DisableAutoResourceOptimizeRulesRequest struct {
	*requests.RpcRequest
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	InstanceIds    string `position:"Query" name:"InstanceIds"`
}

// DisableAutoResourceOptimizeRulesResponse is the response struct for api DisableAutoResourceOptimizeRules
type DisableAutoResourceOptimizeRulesResponse struct {
	*responses.BaseResponse
	Code      int64  `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDisableAutoResourceOptimizeRulesRequest creates a request to invoke DisableAutoResourceOptimizeRules API
func CreateDisableAutoResourceOptimizeRulesRequest() (request *DisableAutoResourceOptimizeRulesRequest) {
	request = &DisableAutoResourceOptimizeRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "DisableAutoResourceOptimizeRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableAutoResourceOptimizeRulesResponse creates a response to parse from DisableAutoResourceOptimizeRules response
func CreateDisableAutoResourceOptimizeRulesResponse() (response *DisableAutoResourceOptimizeRulesResponse) {
	response = &DisableAutoResourceOptimizeRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
