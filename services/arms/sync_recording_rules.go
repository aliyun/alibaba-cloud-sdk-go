package arms

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

// SyncRecordingRules invokes the arms.SyncRecordingRules API synchronously
func (client *Client) SyncRecordingRules(request *SyncRecordingRulesRequest) (response *SyncRecordingRulesResponse, err error) {
	response = CreateSyncRecordingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// SyncRecordingRulesWithChan invokes the arms.SyncRecordingRules API asynchronously
func (client *Client) SyncRecordingRulesWithChan(request *SyncRecordingRulesRequest) (<-chan *SyncRecordingRulesResponse, <-chan error) {
	responseChan := make(chan *SyncRecordingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncRecordingRules(request)
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

// SyncRecordingRulesWithCallback invokes the arms.SyncRecordingRules API asynchronously
func (client *Client) SyncRecordingRulesWithCallback(request *SyncRecordingRulesRequest, callback func(response *SyncRecordingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncRecordingRulesResponse
		var err error
		defer close(result)
		response, err = client.SyncRecordingRules(request)
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

// SyncRecordingRulesRequest is the request struct for api SyncRecordingRules
type SyncRecordingRulesRequest struct {
	*requests.RpcRequest
	ClusterId      string `position:"Query" name:"ClusterId"`
	TargetClusters string `position:"Query" name:"TargetClusters"`
}

// SyncRecordingRulesResponse is the response struct for api SyncRecordingRules
type SyncRecordingRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateSyncRecordingRulesRequest creates a request to invoke SyncRecordingRules API
func CreateSyncRecordingRulesRequest() (request *SyncRecordingRulesRequest) {
	request = &SyncRecordingRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SyncRecordingRules", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSyncRecordingRulesResponse creates a response to parse from SyncRecordingRules response
func CreateSyncRecordingRulesResponse() (response *SyncRecordingRulesResponse) {
	response = &SyncRecordingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
