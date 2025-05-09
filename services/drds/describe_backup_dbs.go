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

// DescribeBackupDbs invokes the drds.DescribeBackupDbs API synchronously
func (client *Client) DescribeBackupDbs(request *DescribeBackupDbsRequest) (response *DescribeBackupDbsResponse, err error) {
	response = CreateDescribeBackupDbsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupDbsWithChan invokes the drds.DescribeBackupDbs API asynchronously
func (client *Client) DescribeBackupDbsWithChan(request *DescribeBackupDbsRequest) (<-chan *DescribeBackupDbsResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupDbsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupDbs(request)
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

// DescribeBackupDbsWithCallback invokes the drds.DescribeBackupDbs API asynchronously
func (client *Client) DescribeBackupDbsWithCallback(request *DescribeBackupDbsRequest, callback func(response *DescribeBackupDbsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupDbsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupDbs(request)
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

// DescribeBackupDbsRequest is the request struct for api DescribeBackupDbs
type DescribeBackupDbsRequest struct {
	*requests.RpcRequest
	PreferredRestoreTime string `position:"Query" name:"PreferredRestoreTime"`
	BackupId             string `position:"Query" name:"BackupId"`
	DrdsInstanceId       string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeBackupDbsResponse is the response struct for api DescribeBackupDbs
type DescribeBackupDbsResponse struct {
	*responses.BaseResponse
	Success   bool    `json:"Success" xml:"Success"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	DbNames   DbNames `json:"DbNames" xml:"DbNames"`
}

// CreateDescribeBackupDbsRequest creates a request to invoke DescribeBackupDbs API
func CreateDescribeBackupDbsRequest() (request *DescribeBackupDbsRequest) {
	request = &DescribeBackupDbsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeBackupDbs", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupDbsResponse creates a response to parse from DescribeBackupDbs response
func CreateDescribeBackupDbsResponse() (response *DescribeBackupDbsResponse) {
	response = &DescribeBackupDbsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
