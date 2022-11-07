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

// DisableBackupLog invokes the dbs.DisableBackupLog API synchronously
func (client *Client) DisableBackupLog(request *DisableBackupLogRequest) (response *DisableBackupLogResponse, err error) {
	response = CreateDisableBackupLogResponse()
	err = client.DoAction(request, response)
	return
}

// DisableBackupLogWithChan invokes the dbs.DisableBackupLog API asynchronously
func (client *Client) DisableBackupLogWithChan(request *DisableBackupLogRequest) (<-chan *DisableBackupLogResponse, <-chan error) {
	responseChan := make(chan *DisableBackupLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableBackupLog(request)
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

// DisableBackupLogWithCallback invokes the dbs.DisableBackupLog API asynchronously
func (client *Client) DisableBackupLogWithCallback(request *DisableBackupLogRequest, callback func(response *DisableBackupLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableBackupLogResponse
		var err error
		defer close(result)
		response, err = client.DisableBackupLog(request)
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

// DisableBackupLogRequest is the request struct for api DisableBackupLog
type DisableBackupLogRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	BackupPlanId string `position:"Query" name:"BackupPlanId"`
	OwnerId      string `position:"Query" name:"OwnerId"`
}

// DisableBackupLogResponse is the response struct for api DisableBackupLog
type DisableBackupLogResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	NeedPrecheck   bool   `json:"NeedPrecheck" xml:"NeedPrecheck"`
}

// CreateDisableBackupLogRequest creates a request to invoke DisableBackupLog API
func CreateDisableBackupLogRequest() (request *DisableBackupLogRequest) {
	request = &DisableBackupLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DisableBackupLog", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableBackupLogResponse creates a response to parse from DisableBackupLog response
func CreateDisableBackupLogResponse() (response *DisableBackupLogResponse) {
	response = &DisableBackupLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
