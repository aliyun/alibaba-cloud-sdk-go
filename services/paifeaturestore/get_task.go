package paifeaturestore

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

// GetTask invokes the paifeaturestore.GetTask API synchronously
func (client *Client) GetTask(request *GetTaskRequest) (response *GetTaskResponse, err error) {
	response = CreateGetTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskWithChan invokes the paifeaturestore.GetTask API asynchronously
func (client *Client) GetTaskWithChan(request *GetTaskRequest) (<-chan *GetTaskResponse, <-chan error) {
	responseChan := make(chan *GetTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTask(request)
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

// GetTaskWithCallback invokes the paifeaturestore.GetTask API asynchronously
func (client *Client) GetTaskWithCallback(request *GetTaskRequest, callback func(response *GetTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskResponse
		var err error
		defer close(result)
		response, err = client.GetTask(request)
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

// GetTaskRequest is the request struct for api GetTask
type GetTaskRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	TaskId     string `position:"Path" name:"TaskId"`
}

// GetTaskResponse is the response struct for api GetTask
type GetTaskResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ProjectId       string `json:"ProjectId" xml:"ProjectId"`
	ProjectName     string `json:"ProjectName" xml:"ProjectName"`
	Type            string `json:"Type" xml:"Type"`
	ObjectType      string `json:"ObjectType" xml:"ObjectType"`
	ObjectId        string `json:"ObjectId" xml:"ObjectId"`
	Status          string `json:"Status" xml:"Status"`
	Config          string `json:"Config" xml:"Config"`
	RunningConfig   string `json:"RunningConfig" xml:"RunningConfig"`
	GmtCreateTime   string `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime string `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	GmtExecutedTime string `json:"GmtExecutedTime" xml:"GmtExecutedTime"`
	GmtFinishedTime string `json:"GmtFinishedTime" xml:"GmtFinishedTime"`
}

// CreateGetTaskRequest creates a request to invoke GetTask API
func CreateGetTaskRequest() (request *GetTaskRequest) {
	request = &GetTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetTask", "/api/v1/instances/[InstanceId]/tasks/[TaskId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetTaskResponse creates a response to parse from GetTask response
func CreateGetTaskResponse() (response *GetTaskResponse) {
	response = &GetTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
