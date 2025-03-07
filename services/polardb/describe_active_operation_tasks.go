package polardb

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

// DescribeActiveOperationTasks invokes the polardb.DescribeActiveOperationTasks API synchronously
func (client *Client) DescribeActiveOperationTasks(request *DescribeActiveOperationTasksRequest) (response *DescribeActiveOperationTasksResponse, err error) {
	response = CreateDescribeActiveOperationTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActiveOperationTasksWithChan invokes the polardb.DescribeActiveOperationTasks API asynchronously
func (client *Client) DescribeActiveOperationTasksWithChan(request *DescribeActiveOperationTasksRequest) (<-chan *DescribeActiveOperationTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeActiveOperationTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActiveOperationTasks(request)
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

// DescribeActiveOperationTasksWithCallback invokes the polardb.DescribeActiveOperationTasks API asynchronously
func (client *Client) DescribeActiveOperationTasksWithCallback(request *DescribeActiveOperationTasksRequest, callback func(response *DescribeActiveOperationTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActiveOperationTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeActiveOperationTasks(request)
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

// DescribeActiveOperationTasksRequest is the request struct for api DescribeActiveOperationTasks
type DescribeActiveOperationTasksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ChangeLevel          string           `position:"Query" name:"ChangeLevel"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ProductName          string           `position:"Query" name:"ProductName"`
	EngineType           string           `position:"Query" name:"EngineType"`
	TaskType             string           `position:"Query" name:"TaskType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AllowCancel          requests.Integer `position:"Query" name:"AllowCancel"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBType               string           `position:"Query" name:"DBType"`
	AllowChange          requests.Integer `position:"Query" name:"AllowChange"`
	Region               string           `position:"Query" name:"Region"`
	Status               requests.Integer `position:"Query" name:"Status"`
}

// DescribeActiveOperationTasksResponse is the response struct for api DescribeActiveOperationTasks
type DescribeActiveOperationTasksResponse struct {
	*responses.BaseResponse
	TotalRecordCount int         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	RequestId        string      `json:"RequestId" xml:"RequestId"`
	PageSize         int         `json:"PageSize" xml:"PageSize"`
	PageNumber       int         `json:"PageNumber" xml:"PageNumber"`
	Items            []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeActiveOperationTasksRequest creates a request to invoke DescribeActiveOperationTasks API
func CreateDescribeActiveOperationTasksRequest() (request *DescribeActiveOperationTasksRequest) {
	request = &DescribeActiveOperationTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeActiveOperationTasks", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeActiveOperationTasksResponse creates a response to parse from DescribeActiveOperationTasks response
func CreateDescribeActiveOperationTasksResponse() (response *DescribeActiveOperationTasksResponse) {
	response = &DescribeActiveOperationTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
