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

// AddRuleV4 invokes the qualitycheck.AddRuleV4 API synchronously
func (client *Client) AddRuleV4(request *AddRuleV4Request) (response *AddRuleV4Response, err error) {
	response = CreateAddRuleV4Response()
	err = client.DoAction(request, response)
	return
}

// AddRuleV4WithChan invokes the qualitycheck.AddRuleV4 API asynchronously
func (client *Client) AddRuleV4WithChan(request *AddRuleV4Request) (<-chan *AddRuleV4Response, <-chan error) {
	responseChan := make(chan *AddRuleV4Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddRuleV4(request)
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

// AddRuleV4WithCallback invokes the qualitycheck.AddRuleV4 API asynchronously
func (client *Client) AddRuleV4WithCallback(request *AddRuleV4Request, callback func(response *AddRuleV4Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddRuleV4Response
		var err error
		defer close(result)
		response, err = client.AddRuleV4(request)
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

// AddRuleV4Request is the request struct for api AddRuleV4
type AddRuleV4Request struct {
	*requests.RpcRequest
	IsSchemeData         requests.Integer `position:"Body" name:"IsSchemeData"`
	ReturnRelatedSchemes requests.Boolean `position:"Body" name:"ReturnRelatedSchemes"`
	IsCopy               requests.Boolean `position:"Body" name:"IsCopy"`
	JsonStrForRule       string           `position:"Body" name:"JsonStrForRule"`
	RuleId               requests.Integer `position:"Body" name:"RuleId"`
	BaseMeAgentId        requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// AddRuleV4Response is the response struct for api AddRuleV4
type AddRuleV4Response struct {
	*responses.BaseResponse
	RequestId      string              `json:"RequestId" xml:"RequestId"`
	Success        bool                `json:"Success" xml:"Success"`
	Code           string              `json:"Code" xml:"Code"`
	Message        string              `json:"Message" xml:"Message"`
	HttpStatusCode int                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int64               `json:"Data" xml:"Data"`
	Messages       MessagesInAddRuleV4 `json:"Messages" xml:"Messages"`
}

// CreateAddRuleV4Request creates a request to invoke AddRuleV4 API
func CreateAddRuleV4Request() (request *AddRuleV4Request) {
	request = &AddRuleV4Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "AddRuleV4", "", "")
	request.Method = requests.POST
	return
}

// CreateAddRuleV4Response creates a response to parse from AddRuleV4 response
func CreateAddRuleV4Response() (response *AddRuleV4Response) {
	response = &AddRuleV4Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
