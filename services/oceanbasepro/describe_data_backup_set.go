package oceanbasepro

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

// DescribeDataBackupSet invokes the oceanbasepro.DescribeDataBackupSet API synchronously
func (client *Client) DescribeDataBackupSet(request *DescribeDataBackupSetRequest) (response *DescribeDataBackupSetResponse, err error) {
	response = CreateDescribeDataBackupSetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataBackupSetWithChan invokes the oceanbasepro.DescribeDataBackupSet API asynchronously
func (client *Client) DescribeDataBackupSetWithChan(request *DescribeDataBackupSetRequest) (<-chan *DescribeDataBackupSetResponse, <-chan error) {
	responseChan := make(chan *DescribeDataBackupSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataBackupSet(request)
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

// DescribeDataBackupSetWithCallback invokes the oceanbasepro.DescribeDataBackupSet API asynchronously
func (client *Client) DescribeDataBackupSetWithCallback(request *DescribeDataBackupSetRequest, callback func(response *DescribeDataBackupSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataBackupSetResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataBackupSet(request)
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

// DescribeDataBackupSetRequest is the request struct for api DescribeDataBackupSet
type DescribeDataBackupSetRequest struct {
	*requests.RpcRequest
	StartTime        string           `position:"Body" name:"StartTime"`
	PageNumber       requests.Integer `position:"Body" name:"PageNumber"`
	PageSize         requests.Integer `position:"Body" name:"PageSize"`
	BackupObjectType string           `position:"Body" name:"BackupObjectType"`
	EndTime          string           `position:"Body" name:"EndTime"`
	InstanceId       string           `position:"Body" name:"InstanceId"`
	Status           string           `position:"Body" name:"Status"`
}

// DescribeDataBackupSetResponse is the response struct for api DescribeDataBackupSet
type DescribeDataBackupSetResponse struct {
	*responses.BaseResponse
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeDataBackupSetRequest creates a request to invoke DescribeDataBackupSet API
func CreateDescribeDataBackupSetRequest() (request *DescribeDataBackupSetRequest) {
	request = &DescribeDataBackupSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeDataBackupSet", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataBackupSetResponse creates a response to parse from DescribeDataBackupSet response
func CreateDescribeDataBackupSetResponse() (response *DescribeDataBackupSetResponse) {
	response = &DescribeDataBackupSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
