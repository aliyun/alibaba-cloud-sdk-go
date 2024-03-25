package fnf

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

// UpdateSchedule invokes the fnf.UpdateSchedule API synchronously
func (client *Client) UpdateSchedule(request *UpdateScheduleRequest) (response *UpdateScheduleResponse, err error) {
	response = CreateUpdateScheduleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateScheduleWithChan invokes the fnf.UpdateSchedule API asynchronously
func (client *Client) UpdateScheduleWithChan(request *UpdateScheduleRequest) (<-chan *UpdateScheduleResponse, <-chan error) {
	responseChan := make(chan *UpdateScheduleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSchedule(request)
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

// UpdateScheduleWithCallback invokes the fnf.UpdateSchedule API asynchronously
func (client *Client) UpdateScheduleWithCallback(request *UpdateScheduleRequest, callback func(response *UpdateScheduleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateScheduleResponse
		var err error
		defer close(result)
		response, err = client.UpdateSchedule(request)
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

// UpdateScheduleRequest is the request struct for api UpdateSchedule
type UpdateScheduleRequest struct {
	*requests.RpcRequest
	CronExpression string           `position:"Body" name:"CronExpression"`
	Description    string           `position:"Body" name:"Description"`
	ScheduleName   string           `position:"Body" name:"ScheduleName"`
	Payload        string           `position:"Body" name:"Payload"`
	Enable         requests.Boolean `position:"Body" name:"Enable"`
	FlowName       string           `position:"Body" name:"FlowName"`
}

// UpdateScheduleResponse is the response struct for api UpdateSchedule
type UpdateScheduleResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Description      string `json:"Description" xml:"Description"`
	ScheduleId       string `json:"ScheduleId" xml:"ScheduleId"`
	Payload          string `json:"Payload" xml:"Payload"`
	ScheduleName     string `json:"ScheduleName" xml:"ScheduleName"`
	CreatedTime      string `json:"CreatedTime" xml:"CreatedTime"`
	LastModifiedTime string `json:"LastModifiedTime" xml:"LastModifiedTime"`
	CronExpression   string `json:"CronExpression" xml:"CronExpression"`
	Enable           bool   `json:"Enable" xml:"Enable"`
}

// CreateUpdateScheduleRequest creates a request to invoke UpdateSchedule API
func CreateUpdateScheduleRequest() (request *UpdateScheduleRequest) {
	request = &UpdateScheduleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "UpdateSchedule", "fnf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateScheduleResponse creates a response to parse from UpdateSchedule response
func CreateUpdateScheduleResponse() (response *UpdateScheduleResponse) {
	response = &UpdateScheduleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
