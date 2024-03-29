package dbs

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

// ModifyBackupPlanName invokes the dbs.ModifyBackupPlanName API synchronously
func (client *Client) ModifyBackupPlanName(request *ModifyBackupPlanNameRequest) (response *ModifyBackupPlanNameResponse, err error) {
	response = CreateModifyBackupPlanNameResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupPlanNameWithChan invokes the dbs.ModifyBackupPlanName API asynchronously
func (client *Client) ModifyBackupPlanNameWithChan(request *ModifyBackupPlanNameRequest) (<-chan *ModifyBackupPlanNameResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupPlanNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupPlanName(request)
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

// ModifyBackupPlanNameWithCallback invokes the dbs.ModifyBackupPlanName API asynchronously
func (client *Client) ModifyBackupPlanNameWithCallback(request *ModifyBackupPlanNameRequest, callback func(response *ModifyBackupPlanNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupPlanNameResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupPlanName(request)
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

// ModifyBackupPlanNameRequest is the request struct for api ModifyBackupPlanName
type ModifyBackupPlanNameRequest struct {
	*requests.RpcRequest
	ClientToken    string `position:"Query" name:"ClientToken"`
	BackupPlanId   string `position:"Query" name:"BackupPlanId"`
	OwnerId        string `position:"Query" name:"OwnerId"`
	BackupPlanName string `position:"Query" name:"BackupPlanName"`
}

// ModifyBackupPlanNameResponse is the response struct for api ModifyBackupPlanName
type ModifyBackupPlanNameResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
}

// CreateModifyBackupPlanNameRequest creates a request to invoke ModifyBackupPlanName API
func CreateModifyBackupPlanNameRequest() (request *ModifyBackupPlanNameRequest) {
	request = &ModifyBackupPlanNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ModifyBackupPlanName", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBackupPlanNameResponse creates a response to parse from ModifyBackupPlanName response
func CreateModifyBackupPlanNameResponse() (response *ModifyBackupPlanNameResponse) {
	response = &ModifyBackupPlanNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
