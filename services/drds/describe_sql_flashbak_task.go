package drds

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

// DescribeSqlFlashbakTask invokes the drds.DescribeSqlFlashbakTask API synchronously
func (client *Client) DescribeSqlFlashbakTask(request *DescribeSqlFlashbakTaskRequest) (response *DescribeSqlFlashbakTaskResponse, err error) {
	response = CreateDescribeSqlFlashbakTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSqlFlashbakTaskWithChan invokes the drds.DescribeSqlFlashbakTask API asynchronously
func (client *Client) DescribeSqlFlashbakTaskWithChan(request *DescribeSqlFlashbakTaskRequest) (<-chan *DescribeSqlFlashbakTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeSqlFlashbakTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSqlFlashbakTask(request)
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

// DescribeSqlFlashbakTaskWithCallback invokes the drds.DescribeSqlFlashbakTask API asynchronously
func (client *Client) DescribeSqlFlashbakTaskWithCallback(request *DescribeSqlFlashbakTaskRequest, callback func(response *DescribeSqlFlashbakTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSqlFlashbakTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeSqlFlashbakTask(request)
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

// DescribeSqlFlashbakTaskRequest is the request struct for api DescribeSqlFlashbakTask
type DescribeSqlFlashbakTaskRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeSqlFlashbakTaskResponse is the response struct for api DescribeSqlFlashbakTask
type DescribeSqlFlashbakTaskResponse struct {
	*responses.BaseResponse
	Success           bool              `json:"Success" xml:"Success"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	SqlFlashbackTasks SqlFlashbackTasks `json:"SqlFlashbackTasks" xml:"SqlFlashbackTasks"`
}

// CreateDescribeSqlFlashbakTaskRequest creates a request to invoke DescribeSqlFlashbakTask API
func CreateDescribeSqlFlashbakTaskRequest() (request *DescribeSqlFlashbakTaskRequest) {
	request = &DescribeSqlFlashbakTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeSqlFlashbakTask", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSqlFlashbakTaskResponse creates a response to parse from DescribeSqlFlashbakTask response
func CreateDescribeSqlFlashbakTaskResponse() (response *DescribeSqlFlashbakTaskResponse) {
	response = &DescribeSqlFlashbakTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
