package adb

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

// ModifyMaintenanceAction invokes the adb.ModifyMaintenanceAction API synchronously
func (client *Client) ModifyMaintenanceAction(request *ModifyMaintenanceActionRequest) (response *ModifyMaintenanceActionResponse, err error) {
	response = CreateModifyMaintenanceActionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMaintenanceActionWithChan invokes the adb.ModifyMaintenanceAction API asynchronously
func (client *Client) ModifyMaintenanceActionWithChan(request *ModifyMaintenanceActionRequest) (<-chan *ModifyMaintenanceActionResponse, <-chan error) {
	responseChan := make(chan *ModifyMaintenanceActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMaintenanceAction(request)
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

// ModifyMaintenanceActionWithCallback invokes the adb.ModifyMaintenanceAction API asynchronously
func (client *Client) ModifyMaintenanceActionWithCallback(request *ModifyMaintenanceActionRequest, callback func(response *ModifyMaintenanceActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMaintenanceActionResponse
		var err error
		defer close(result)
		response, err = client.ModifyMaintenanceAction(request)
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

// ModifyMaintenanceActionRequest is the request struct for api ModifyMaintenanceAction
type ModifyMaintenanceActionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SwitchTime           string           `position:"Query" name:"SwitchTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ids                  string           `position:"Query" name:"Ids"`
}

// ModifyMaintenanceActionResponse is the response struct for api ModifyMaintenanceAction
type ModifyMaintenanceActionResponse struct {
	*responses.BaseResponse
	Ids       string `json:"Ids" xml:"Ids"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMaintenanceActionRequest creates a request to invoke ModifyMaintenanceAction API
func CreateModifyMaintenanceActionRequest() (request *ModifyMaintenanceActionRequest) {
	request = &ModifyMaintenanceActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ModifyMaintenanceAction", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyMaintenanceActionResponse creates a response to parse from ModifyMaintenanceAction response
func CreateModifyMaintenanceActionResponse() (response *ModifyMaintenanceActionResponse) {
	response = &ModifyMaintenanceActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
