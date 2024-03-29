package oceanbasepro

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

// ModifyInstanceNodeNum invokes the oceanbasepro.ModifyInstanceNodeNum API synchronously
func (client *Client) ModifyInstanceNodeNum(request *ModifyInstanceNodeNumRequest) (response *ModifyInstanceNodeNumResponse, err error) {
	response = CreateModifyInstanceNodeNumResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceNodeNumWithChan invokes the oceanbasepro.ModifyInstanceNodeNum API asynchronously
func (client *Client) ModifyInstanceNodeNumWithChan(request *ModifyInstanceNodeNumRequest) (<-chan *ModifyInstanceNodeNumResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceNodeNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceNodeNum(request)
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

// ModifyInstanceNodeNumWithCallback invokes the oceanbasepro.ModifyInstanceNodeNum API asynchronously
func (client *Client) ModifyInstanceNodeNumWithCallback(request *ModifyInstanceNodeNumRequest, callback func(response *ModifyInstanceNodeNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceNodeNumResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceNodeNum(request)
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

// ModifyInstanceNodeNumRequest is the request struct for api ModifyInstanceNodeNum
type ModifyInstanceNodeNumRequest struct {
	*requests.RpcRequest
	DryRun     requests.Boolean `position:"Body" name:"DryRun"`
	NodeNum    string           `position:"Body" name:"NodeNum"`
	InstanceId string           `position:"Body" name:"InstanceId"`
}

// ModifyInstanceNodeNumResponse is the response struct for api ModifyInstanceNodeNum
type ModifyInstanceNodeNumResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateModifyInstanceNodeNumRequest creates a request to invoke ModifyInstanceNodeNum API
func CreateModifyInstanceNodeNumRequest() (request *ModifyInstanceNodeNumRequest) {
	request = &ModifyInstanceNodeNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ModifyInstanceNodeNum", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceNodeNumResponse creates a response to parse from ModifyInstanceNodeNum response
func CreateModifyInstanceNodeNumResponse() (response *ModifyInstanceNodeNumResponse) {
	response = &ModifyInstanceNodeNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
