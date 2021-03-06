package dbfs

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

// StopUpgradeTask invokes the dbfs.StopUpgradeTask API synchronously
func (client *Client) StopUpgradeTask(request *StopUpgradeTaskRequest) (response *StopUpgradeTaskResponse, err error) {
	response = CreateStopUpgradeTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopUpgradeTaskWithChan invokes the dbfs.StopUpgradeTask API asynchronously
func (client *Client) StopUpgradeTaskWithChan(request *StopUpgradeTaskRequest) (<-chan *StopUpgradeTaskResponse, <-chan error) {
	responseChan := make(chan *StopUpgradeTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopUpgradeTask(request)
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

// StopUpgradeTaskWithCallback invokes the dbfs.StopUpgradeTask API asynchronously
func (client *Client) StopUpgradeTaskWithCallback(request *StopUpgradeTaskRequest, callback func(response *StopUpgradeTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopUpgradeTaskResponse
		var err error
		defer close(result)
		response, err = client.StopUpgradeTask(request)
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

// StopUpgradeTaskRequest is the request struct for api StopUpgradeTask
type StopUpgradeTaskRequest struct {
	*requests.RpcRequest
	PageNumber        requests.Integer `position:"Query" name:"PageNumber"`
	BatchStrategyList string           `position:"Query" name:"BatchStrategyList"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
}

// StopUpgradeTaskResponse is the response struct for api StopUpgradeTask
type StopUpgradeTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopUpgradeTaskRequest creates a request to invoke StopUpgradeTask API
func CreateStopUpgradeTaskRequest() (request *StopUpgradeTaskRequest) {
	request = &StopUpgradeTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "StopUpgradeTask", "", "")
	request.Method = requests.POST
	return
}

// CreateStopUpgradeTaskResponse creates a response to parse from StopUpgradeTask response
func CreateStopUpgradeTaskResponse() (response *StopUpgradeTaskResponse) {
	response = &StopUpgradeTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
