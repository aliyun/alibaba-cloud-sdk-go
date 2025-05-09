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

// DescribeBackupTimes invokes the drds.DescribeBackupTimes API synchronously
func (client *Client) DescribeBackupTimes(request *DescribeBackupTimesRequest) (response *DescribeBackupTimesResponse, err error) {
	response = CreateDescribeBackupTimesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupTimesWithChan invokes the drds.DescribeBackupTimes API asynchronously
func (client *Client) DescribeBackupTimesWithChan(request *DescribeBackupTimesRequest) (<-chan *DescribeBackupTimesResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupTimesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupTimes(request)
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

// DescribeBackupTimesWithCallback invokes the drds.DescribeBackupTimes API asynchronously
func (client *Client) DescribeBackupTimesWithCallback(request *DescribeBackupTimesRequest, callback func(response *DescribeBackupTimesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupTimesResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupTimes(request)
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

// DescribeBackupTimesRequest is the request struct for api DescribeBackupTimes
type DescribeBackupTimesRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeBackupTimesResponse is the response struct for api DescribeBackupTimes
type DescribeBackupTimesResponse struct {
	*responses.BaseResponse
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	RestoreTime RestoreTime `json:"RestoreTime" xml:"RestoreTime"`
}

// CreateDescribeBackupTimesRequest creates a request to invoke DescribeBackupTimes API
func CreateDescribeBackupTimesRequest() (request *DescribeBackupTimesRequest) {
	request = &DescribeBackupTimesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeBackupTimes", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupTimesResponse creates a response to parse from DescribeBackupTimes response
func CreateDescribeBackupTimesResponse() (response *DescribeBackupTimesResponse) {
	response = &DescribeBackupTimesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
