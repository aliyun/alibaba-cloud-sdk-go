package selectdb

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

// EnDisableScalingRules invokes the selectdb.EnDisableScalingRules API synchronously
func (client *Client) EnDisableScalingRules(request *EnDisableScalingRulesRequest) (response *EnDisableScalingRulesResponse, err error) {
	response = CreateEnDisableScalingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// EnDisableScalingRulesWithChan invokes the selectdb.EnDisableScalingRules API asynchronously
func (client *Client) EnDisableScalingRulesWithChan(request *EnDisableScalingRulesRequest) (<-chan *EnDisableScalingRulesResponse, <-chan error) {
	responseChan := make(chan *EnDisableScalingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnDisableScalingRules(request)
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

// EnDisableScalingRulesWithCallback invokes the selectdb.EnDisableScalingRules API asynchronously
func (client *Client) EnDisableScalingRulesWithCallback(request *EnDisableScalingRulesRequest, callback func(response *EnDisableScalingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnDisableScalingRulesResponse
		var err error
		defer close(result)
		response, err = client.EnDisableScalingRules(request)
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

// EnDisableScalingRulesRequest is the request struct for api EnDisableScalingRules
type EnDisableScalingRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId    requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Product            string           `position:"Query" name:"Product"`
	ClusterId          string           `position:"Query" name:"ClusterId"`
	DbInstanceId       string           `position:"Query" name:"DbInstanceId"`
	ScalingRulesEnable requests.Boolean `position:"Query" name:"ScalingRulesEnable"`
}

// EnDisableScalingRulesResponse is the response struct for api EnDisableScalingRules
type EnDisableScalingRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnDisableScalingRulesRequest creates a request to invoke EnDisableScalingRules API
func CreateEnDisableScalingRulesRequest() (request *EnDisableScalingRulesRequest) {
	request = &EnDisableScalingRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("selectdb", "2023-05-22", "EnDisableScalingRules", "", "")
	request.Method = requests.POST
	return
}

// CreateEnDisableScalingRulesResponse creates a response to parse from EnDisableScalingRules response
func CreateEnDisableScalingRulesResponse() (response *EnDisableScalingRulesResponse) {
	response = &EnDisableScalingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
