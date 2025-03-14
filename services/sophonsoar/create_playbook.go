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

// CreatePlaybook invokes the sophonsoar.CreatePlaybook API synchronously
func (client *Client) CreatePlaybook(request *CreatePlaybookRequest) (response *CreatePlaybookResponse, err error) {
	response = CreateCreatePlaybookResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePlaybookWithChan invokes the sophonsoar.CreatePlaybook API asynchronously
func (client *Client) CreatePlaybookWithChan(request *CreatePlaybookRequest) (<-chan *CreatePlaybookResponse, <-chan error) {
	responseChan := make(chan *CreatePlaybookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePlaybook(request)
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

// CreatePlaybookWithCallback invokes the sophonsoar.CreatePlaybook API asynchronously
func (client *Client) CreatePlaybookWithCallback(request *CreatePlaybookRequest, callback func(response *CreatePlaybookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePlaybookResponse
		var err error
		defer close(result)
		response, err = client.CreatePlaybook(request)
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

// CreatePlaybookRequest is the request struct for api CreatePlaybook
type CreatePlaybookRequest struct {
	*requests.RpcRequest
	RoleFor      string `position:"Query" name:"RoleFor"`
	TaskflowType string `position:"Body" name:"TaskflowType"`
	Description  string `position:"Body" name:"Description"`
	DisplayName  string `position:"Body" name:"DisplayName"`
	RoleType     string `position:"Query" name:"RoleType"`
	Lang         string `position:"Body" name:"Lang"`
}

// CreatePlaybookResponse is the response struct for api CreatePlaybook
type CreatePlaybookResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreatePlaybookRequest creates a request to invoke CreatePlaybook API
func CreateCreatePlaybookRequest() (request *CreatePlaybookRequest) {
	request = &CreatePlaybookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "CreatePlaybook", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePlaybookResponse creates a response to parse from CreatePlaybook response
func CreateCreatePlaybookResponse() (response *CreatePlaybookResponse) {
	response = &CreatePlaybookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
