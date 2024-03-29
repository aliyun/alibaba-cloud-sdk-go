package qualitycheck

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

// UploadDataSync invokes the qualitycheck.UploadDataSync API synchronously
func (client *Client) UploadDataSync(request *UploadDataSyncRequest) (response *UploadDataSyncResponse, err error) {
	response = CreateUploadDataSyncResponse()
	err = client.DoAction(request, response)
	return
}

// UploadDataSyncWithChan invokes the qualitycheck.UploadDataSync API asynchronously
func (client *Client) UploadDataSyncWithChan(request *UploadDataSyncRequest) (<-chan *UploadDataSyncResponse, <-chan error) {
	responseChan := make(chan *UploadDataSyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadDataSync(request)
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

// UploadDataSyncWithCallback invokes the qualitycheck.UploadDataSync API asynchronously
func (client *Client) UploadDataSyncWithCallback(request *UploadDataSyncRequest, callback func(response *UploadDataSyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadDataSyncResponse
		var err error
		defer close(result)
		response, err = client.UploadDataSync(request)
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

// UploadDataSyncRequest is the request struct for api UploadDataSync
type UploadDataSyncRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// UploadDataSyncResponse is the response struct for api UploadDataSync
type UploadDataSyncResponse struct {
	*responses.BaseResponse
	Code      string               `json:"Code" xml:"Code"`
	Message   string               `json:"Message" xml:"Message"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Success   bool                 `json:"Success" xml:"Success"`
	Data      DataInUploadDataSync `json:"Data" xml:"Data"`
}

// CreateUploadDataSyncRequest creates a request to invoke UploadDataSync API
func CreateUploadDataSyncRequest() (request *UploadDataSyncRequest) {
	request = &UploadDataSyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UploadDataSync", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadDataSyncResponse creates a response to parse from UploadDataSync response
func CreateUploadDataSyncResponse() (response *UploadDataSyncResponse) {
	response = &UploadDataSyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
