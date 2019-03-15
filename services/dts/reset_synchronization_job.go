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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ResetSynchronizationJob invokes the dts.ResetSynchronizationJob API synchronously
// api document: https://help.aliyun.com/api/dts/resetsynchronizationjob.html
func (client *Client) ResetSynchronizationJob(request *ResetSynchronizationJobRequest) (response *ResetSynchronizationJobResponse, err error) {
	response = CreateResetSynchronizationJobResponse()
	err = client.DoAction(request, response)
	return
}

// ResetSynchronizationJobWithChan invokes the dts.ResetSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/resetsynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetSynchronizationJobWithChan(request *ResetSynchronizationJobRequest) (<-chan *ResetSynchronizationJobResponse, <-chan error) {
	responseChan := make(chan *ResetSynchronizationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetSynchronizationJob(request)
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

// ResetSynchronizationJobWithCallback invokes the dts.ResetSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/resetsynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetSynchronizationJobWithCallback(request *ResetSynchronizationJobRequest, callback func(response *ResetSynchronizationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetSynchronizationJobResponse
		var err error
		defer close(result)
		response, err = client.ResetSynchronizationJob(request)
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

// ResetSynchronizationJobRequest is the request struct for api ResetSynchronizationJob
type ResetSynchronizationJobRequest struct {
	*requests.RpcRequest
	SynchronizationJobId     string `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
}

// ResetSynchronizationJobResponse is the response struct for api ResetSynchronizationJob
type ResetSynchronizationJobResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateResetSynchronizationJobRequest creates a request to invoke ResetSynchronizationJob API
func CreateResetSynchronizationJobRequest() (request *ResetSynchronizationJobRequest) {
	request = &ResetSynchronizationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ResetSynchronizationJob", "dts", "openAPI")
	return
}

// CreateResetSynchronizationJobResponse creates a response to parse from ResetSynchronizationJob response
func CreateResetSynchronizationJobResponse() (response *ResetSynchronizationJobResponse) {
	response = &ResetSynchronizationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
