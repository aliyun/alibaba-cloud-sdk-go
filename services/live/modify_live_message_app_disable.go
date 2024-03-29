package live

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

// ModifyLiveMessageAppDisable invokes the live.ModifyLiveMessageAppDisable API synchronously
func (client *Client) ModifyLiveMessageAppDisable(request *ModifyLiveMessageAppDisableRequest) (response *ModifyLiveMessageAppDisableResponse, err error) {
	response = CreateModifyLiveMessageAppDisableResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLiveMessageAppDisableWithChan invokes the live.ModifyLiveMessageAppDisable API asynchronously
func (client *Client) ModifyLiveMessageAppDisableWithChan(request *ModifyLiveMessageAppDisableRequest) (<-chan *ModifyLiveMessageAppDisableResponse, <-chan error) {
	responseChan := make(chan *ModifyLiveMessageAppDisableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLiveMessageAppDisable(request)
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

// ModifyLiveMessageAppDisableWithCallback invokes the live.ModifyLiveMessageAppDisable API asynchronously
func (client *Client) ModifyLiveMessageAppDisableWithCallback(request *ModifyLiveMessageAppDisableRequest, callback func(response *ModifyLiveMessageAppDisableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLiveMessageAppDisableResponse
		var err error
		defer close(result)
		response, err = client.ModifyLiveMessageAppDisable(request)
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

// ModifyLiveMessageAppDisableRequest is the request struct for api ModifyLiveMessageAppDisable
type ModifyLiveMessageAppDisableRequest struct {
	*requests.RpcRequest
	DataCenter string           `position:"Query" name:"DataCenter"`
	Disable    requests.Boolean `position:"Query" name:"Disable"`
	AppId      string           `position:"Query" name:"AppId"`
}

// ModifyLiveMessageAppDisableResponse is the response struct for api ModifyLiveMessageAppDisable
type ModifyLiveMessageAppDisableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AppId     string `json:"AppId" xml:"AppId"`
	AppSign   string `json:"AppSign" xml:"AppSign"`
	Disable   bool   `json:"Disable" xml:"Disable"`
}

// CreateModifyLiveMessageAppDisableRequest creates a request to invoke ModifyLiveMessageAppDisable API
func CreateModifyLiveMessageAppDisableRequest() (request *ModifyLiveMessageAppDisableRequest) {
	request = &ModifyLiveMessageAppDisableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyLiveMessageAppDisable", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLiveMessageAppDisableResponse creates a response to parse from ModifyLiveMessageAppDisable response
func CreateModifyLiveMessageAppDisableResponse() (response *ModifyLiveMessageAppDisableResponse) {
	response = &ModifyLiveMessageAppDisableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
