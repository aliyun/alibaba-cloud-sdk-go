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

// ModifyPlaybookInputOutput invokes the sophonsoar.ModifyPlaybookInputOutput API synchronously
func (client *Client) ModifyPlaybookInputOutput(request *ModifyPlaybookInputOutputRequest) (response *ModifyPlaybookInputOutputResponse, err error) {
	response = CreateModifyPlaybookInputOutputResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPlaybookInputOutputWithChan invokes the sophonsoar.ModifyPlaybookInputOutput API asynchronously
func (client *Client) ModifyPlaybookInputOutputWithChan(request *ModifyPlaybookInputOutputRequest) (<-chan *ModifyPlaybookInputOutputResponse, <-chan error) {
	responseChan := make(chan *ModifyPlaybookInputOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPlaybookInputOutput(request)
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

// ModifyPlaybookInputOutputWithCallback invokes the sophonsoar.ModifyPlaybookInputOutput API asynchronously
func (client *Client) ModifyPlaybookInputOutputWithCallback(request *ModifyPlaybookInputOutputRequest, callback func(response *ModifyPlaybookInputOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPlaybookInputOutputResponse
		var err error
		defer close(result)
		response, err = client.ModifyPlaybookInputOutput(request)
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

// ModifyPlaybookInputOutputRequest is the request struct for api ModifyPlaybookInputOutput
type ModifyPlaybookInputOutputRequest struct {
	*requests.RpcRequest
	ParamType    string `position:"Body" name:"ParamType"`
	RoleFor      string `position:"Query" name:"RoleFor"`
	InputParams  string `position:"Body" name:"InputParams"`
	OutputParams string `position:"Body" name:"OutputParams"`
	PlaybookUuid string `position:"Body" name:"PlaybookUuid"`
	ExeConfig    string `position:"Body" name:"ExeConfig"`
	RoleType     string `position:"Query" name:"RoleType"`
	Lang         string `position:"Body" name:"Lang"`
}

// ModifyPlaybookInputOutputResponse is the response struct for api ModifyPlaybookInputOutput
type ModifyPlaybookInputOutputResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPlaybookInputOutputRequest creates a request to invoke ModifyPlaybookInputOutput API
func CreateModifyPlaybookInputOutputRequest() (request *ModifyPlaybookInputOutputRequest) {
	request = &ModifyPlaybookInputOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "ModifyPlaybookInputOutput", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPlaybookInputOutputResponse creates a response to parse from ModifyPlaybookInputOutput response
func CreateModifyPlaybookInputOutputResponse() (response *ModifyPlaybookInputOutputResponse) {
	response = &ModifyPlaybookInputOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
