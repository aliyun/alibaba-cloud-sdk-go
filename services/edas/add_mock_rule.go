package edas

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

// AddMockRule invokes the edas.AddMockRule API synchronously
func (client *Client) AddMockRule(request *AddMockRuleRequest) (response *AddMockRuleResponse, err error) {
	response = CreateAddMockRuleResponse()
	err = client.DoAction(request, response)
	return
}

// AddMockRuleWithChan invokes the edas.AddMockRule API asynchronously
func (client *Client) AddMockRuleWithChan(request *AddMockRuleRequest) (<-chan *AddMockRuleResponse, <-chan error) {
	responseChan := make(chan *AddMockRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddMockRule(request)
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

// AddMockRuleWithCallback invokes the edas.AddMockRule API asynchronously
func (client *Client) AddMockRuleWithCallback(request *AddMockRuleRequest, callback func(response *AddMockRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddMockRuleResponse
		var err error
		defer close(result)
		response, err = client.AddMockRule(request)
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

// AddMockRuleRequest is the request struct for api AddMockRule
type AddMockRuleRequest struct {
	*requests.RoaRequest
	ScMockItemJson    string `position:"Query" name:"ScMockItemJson"`
	DubboMockItemJson string `position:"Query" name:"DubboMockItemJson"`
	ExtraJson         string `position:"Query" name:"ExtraJson"`
	ProviderAppId     string `position:"Query" name:"ProviderAppId"`
	Source            string `position:"Query" name:"Source"`
	Enable            string `position:"Query" name:"Enable"`
	ProviderAppName   string `position:"Query" name:"ProviderAppName"`
	Name              string `position:"Query" name:"Name"`
	Namespace         string `position:"Query" name:"Namespace"`
	ConsumerAppsJson  string `position:"Query" name:"ConsumerAppsJson"`
	MockType          string `position:"Query" name:"MockType"`
	Region            string `position:"Query" name:"Region"`
}

// AddMockRuleResponse is the response struct for api AddMockRule
type AddMockRuleResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddMockRuleRequest creates a request to invoke AddMockRule API
func CreateAddMockRuleRequest() (request *AddMockRuleRequest) {
	request = &AddMockRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "AddMockRule", "/pop/sp/api/mock/addMockRule", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddMockRuleResponse creates a response to parse from AddMockRule response
func CreateAddMockRuleResponse() (response *AddMockRuleResponse) {
	response = &AddMockRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
