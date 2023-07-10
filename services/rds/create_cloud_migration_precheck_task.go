package rds

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

// CreateCloudMigrationPrecheckTask invokes the rds.CreateCloudMigrationPrecheckTask API synchronously
func (client *Client) CreateCloudMigrationPrecheckTask(request *CreateCloudMigrationPrecheckTaskRequest) (response *CreateCloudMigrationPrecheckTaskResponse, err error) {
	response = CreateCreateCloudMigrationPrecheckTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCloudMigrationPrecheckTaskWithChan invokes the rds.CreateCloudMigrationPrecheckTask API asynchronously
func (client *Client) CreateCloudMigrationPrecheckTaskWithChan(request *CreateCloudMigrationPrecheckTaskRequest) (<-chan *CreateCloudMigrationPrecheckTaskResponse, <-chan error) {
	responseChan := make(chan *CreateCloudMigrationPrecheckTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCloudMigrationPrecheckTask(request)
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

// CreateCloudMigrationPrecheckTaskWithCallback invokes the rds.CreateCloudMigrationPrecheckTask API asynchronously
func (client *Client) CreateCloudMigrationPrecheckTaskWithCallback(request *CreateCloudMigrationPrecheckTaskRequest, callback func(response *CreateCloudMigrationPrecheckTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCloudMigrationPrecheckTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateCloudMigrationPrecheckTask(request)
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

// CreateCloudMigrationPrecheckTaskRequest is the request struct for api CreateCloudMigrationPrecheckTask
type CreateCloudMigrationPrecheckTaskRequest struct {
	*requests.RpcRequest
	DBInstanceName       string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TaskName             string           `position:"Query" name:"TaskName"`
	SourceAccount        string           `position:"Query" name:"SourceAccount"`
	SourcePort           requests.Integer `position:"Query" name:"SourcePort"`
	SourcePassword       string           `position:"Query" name:"SourcePassword"`
	SourceIpAddress      string           `position:"Query" name:"SourceIpAddress"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SourceCategory       string           `position:"Query" name:"SourceCategory"`
}

// CreateCloudMigrationPrecheckTaskResponse is the response struct for api CreateCloudMigrationPrecheckTask
type CreateCloudMigrationPrecheckTaskResponse struct {
	*responses.BaseResponse
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	TaskId         int64  `json:"TaskId" xml:"TaskId"`
	TaskName       string `json:"TaskName" xml:"TaskName"`
}

// CreateCreateCloudMigrationPrecheckTaskRequest creates a request to invoke CreateCloudMigrationPrecheckTask API
func CreateCreateCloudMigrationPrecheckTaskRequest() (request *CreateCloudMigrationPrecheckTaskRequest) {
	request = &CreateCloudMigrationPrecheckTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateCloudMigrationPrecheckTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCloudMigrationPrecheckTaskResponse creates a response to parse from CreateCloudMigrationPrecheckTask response
func CreateCreateCloudMigrationPrecheckTaskResponse() (response *CreateCloudMigrationPrecheckTaskResponse) {
	response = &CreateCloudMigrationPrecheckTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
