package avatar

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

// Update2dAvatar invokes the avatar.Update2dAvatar API synchronously
func (client *Client) Update2dAvatar(request *Update2dAvatarRequest) (response *Update2dAvatarResponse, err error) {
	response = CreateUpdate2dAvatarResponse()
	err = client.DoAction(request, response)
	return
}

// Update2dAvatarWithChan invokes the avatar.Update2dAvatar API asynchronously
func (client *Client) Update2dAvatarWithChan(request *Update2dAvatarRequest) (<-chan *Update2dAvatarResponse, <-chan error) {
	responseChan := make(chan *Update2dAvatarResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Update2dAvatar(request)
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

// Update2dAvatarWithCallback invokes the avatar.Update2dAvatar API asynchronously
func (client *Client) Update2dAvatarWithCallback(request *Update2dAvatarRequest, callback func(response *Update2dAvatarResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Update2dAvatarResponse
		var err error
		defer close(result)
		response, err = client.Update2dAvatar(request)
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

// Update2dAvatarRequest is the request struct for api Update2dAvatar
type Update2dAvatarRequest struct {
	*requests.RpcRequest
	Image             string           `position:"Query" name:"Image"`
	Orientation       requests.Integer `position:"Query" name:"Orientation"`
	Code              string           `position:"Query" name:"Code"`
	Description       string           `position:"Query" name:"Description"`
	Video             string           `position:"Query" name:"Video"`
	Portrait          string           `position:"Query" name:"Portrait"`
	ExtParams         string           `position:"Query" name:"ExtParams"`
	Transparent       requests.Boolean `position:"Query" name:"Transparent"`
	CallbackParams    string           `position:"Query" name:"CallbackParams"`
	TenantId          requests.Integer `position:"Query" name:"TenantId"`
	Name              string           `position:"Query" name:"Name"`
	ExtParamsCLS      string           `position:"Query" name:"ExtParams_CLS"`
	Callback          requests.Boolean `position:"Query" name:"Callback"`
	CallbackParamsCLS string           `position:"Query" name:"CallbackParams_CLS"`
}

// Update2dAvatarResponse is the response struct for api Update2dAvatar
type Update2dAvatarResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdate2dAvatarRequest creates a request to invoke Update2dAvatar API
func CreateUpdate2dAvatarRequest() (request *Update2dAvatarRequest) {
	request = &Update2dAvatarRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "Update2dAvatar", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdate2dAvatarResponse creates a response to parse from Update2dAvatar response
func CreateUpdate2dAvatarResponse() (response *Update2dAvatarResponse) {
	response = &Update2dAvatarResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
