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

// ReleaseBackupPlan invokes the dbs.ReleaseBackupPlan API synchronously
func (client *Client) ReleaseBackupPlan(request *ReleaseBackupPlanRequest) (response *ReleaseBackupPlanResponse, err error) {
	response = CreateReleaseBackupPlanResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseBackupPlanWithChan invokes the dbs.ReleaseBackupPlan API asynchronously
func (client *Client) ReleaseBackupPlanWithChan(request *ReleaseBackupPlanRequest) (<-chan *ReleaseBackupPlanResponse, <-chan error) {
	responseChan := make(chan *ReleaseBackupPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseBackupPlan(request)
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

// ReleaseBackupPlanWithCallback invokes the dbs.ReleaseBackupPlan API asynchronously
func (client *Client) ReleaseBackupPlanWithCallback(request *ReleaseBackupPlanRequest, callback func(response *ReleaseBackupPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseBackupPlanResponse
		var err error
		defer close(result)
		response, err = client.ReleaseBackupPlan(request)
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

// ReleaseBackupPlanRequest is the request struct for api ReleaseBackupPlan
type ReleaseBackupPlanRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	BackupPlanId string `position:"Query" name:"BackupPlanId"`
	OwnerId      string `position:"Query" name:"OwnerId"`
}

// ReleaseBackupPlanResponse is the response struct for api ReleaseBackupPlan
type ReleaseBackupPlanResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
}

// CreateReleaseBackupPlanRequest creates a request to invoke ReleaseBackupPlan API
func CreateReleaseBackupPlanRequest() (request *ReleaseBackupPlanRequest) {
	request = &ReleaseBackupPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ReleaseBackupPlan", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseBackupPlanResponse creates a response to parse from ReleaseBackupPlan response
func CreateReleaseBackupPlanResponse() (response *ReleaseBackupPlanResponse) {
	response = &ReleaseBackupPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
