package sophonsoar

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

// VerifyPlaybook invokes the sophonsoar.VerifyPlaybook API synchronously
func (client *Client) VerifyPlaybook(request *VerifyPlaybookRequest) (response *VerifyPlaybookResponse, err error) {
	response = CreateVerifyPlaybookResponse()
	err = client.DoAction(request, response)
	return
}

// VerifyPlaybookWithChan invokes the sophonsoar.VerifyPlaybook API asynchronously
func (client *Client) VerifyPlaybookWithChan(request *VerifyPlaybookRequest) (<-chan *VerifyPlaybookResponse, <-chan error) {
	responseChan := make(chan *VerifyPlaybookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyPlaybook(request)
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

// VerifyPlaybookWithCallback invokes the sophonsoar.VerifyPlaybook API asynchronously
func (client *Client) VerifyPlaybookWithCallback(request *VerifyPlaybookRequest, callback func(response *VerifyPlaybookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyPlaybookResponse
		var err error
		defer close(result)
		response, err = client.VerifyPlaybook(request)
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

// VerifyPlaybookRequest is the request struct for api VerifyPlaybook
type VerifyPlaybookRequest struct {
	*requests.RpcRequest
	RoleFor      string `position:"Query" name:"RoleFor"`
	TaskFlow     string `position:"Body" name:"TaskFlow"`
	PlaybookUuid string `position:"Body" name:"PlaybookUuid"`
	RoleType     string `position:"Query" name:"RoleType"`
	Lang         string `position:"Body" name:"Lang"`
}

// VerifyPlaybookResponse is the response struct for api VerifyPlaybook
type VerifyPlaybookResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	CheckTaskInfos []Data `json:"CheckTaskInfos" xml:"CheckTaskInfos"`
}

// CreateVerifyPlaybookRequest creates a request to invoke VerifyPlaybook API
func CreateVerifyPlaybookRequest() (request *VerifyPlaybookRequest) {
	request = &VerifyPlaybookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "VerifyPlaybook", "", "")
	request.Method = requests.POST
	return
}

// CreateVerifyPlaybookResponse creates a response to parse from VerifyPlaybook response
func CreateVerifyPlaybookResponse() (response *VerifyPlaybookResponse) {
	response = &VerifyPlaybookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
