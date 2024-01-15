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

// UploadRule invokes the qualitycheck.UploadRule API synchronously
func (client *Client) UploadRule(request *UploadRuleRequest) (response *UploadRuleResponse, err error) {
	response = CreateUploadRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UploadRuleWithChan invokes the qualitycheck.UploadRule API asynchronously
func (client *Client) UploadRuleWithChan(request *UploadRuleRequest) (<-chan *UploadRuleResponse, <-chan error) {
	responseChan := make(chan *UploadRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadRule(request)
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

// UploadRuleWithCallback invokes the qualitycheck.UploadRule API asynchronously
func (client *Client) UploadRuleWithCallback(request *UploadRuleRequest, callback func(response *UploadRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadRuleResponse
		var err error
		defer close(result)
		response, err = client.UploadRule(request)
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

// UploadRuleRequest is the request struct for api UploadRule
type UploadRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// UploadRuleResponse is the response struct for api UploadRule
type UploadRuleResponse struct {
	*responses.BaseResponse
	Code      string           `json:"Code" xml:"Code"`
	Message   string           `json:"Message" xml:"Message"`
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Success   bool             `json:"Success" xml:"Success"`
	Data      DataInUploadRule `json:"Data" xml:"Data"`
}

// CreateUploadRuleRequest creates a request to invoke UploadRule API
func CreateUploadRuleRequest() (request *UploadRuleRequest) {
	request = &UploadRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UploadRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadRuleResponse creates a response to parse from UploadRule response
func CreateUploadRuleResponse() (response *UploadRuleResponse) {
	response = &UploadRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
