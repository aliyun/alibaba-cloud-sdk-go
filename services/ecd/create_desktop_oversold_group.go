package ecd

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

// CreateDesktopOversoldGroup invokes the ecd.CreateDesktopOversoldGroup API synchronously
func (client *Client) CreateDesktopOversoldGroup(request *CreateDesktopOversoldGroupRequest) (response *CreateDesktopOversoldGroupResponse, err error) {
	response = CreateCreateDesktopOversoldGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDesktopOversoldGroupWithChan invokes the ecd.CreateDesktopOversoldGroup API asynchronously
func (client *Client) CreateDesktopOversoldGroupWithChan(request *CreateDesktopOversoldGroupRequest) (<-chan *CreateDesktopOversoldGroupResponse, <-chan error) {
	responseChan := make(chan *CreateDesktopOversoldGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDesktopOversoldGroup(request)
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

// CreateDesktopOversoldGroupWithCallback invokes the ecd.CreateDesktopOversoldGroup API asynchronously
func (client *Client) CreateDesktopOversoldGroupWithCallback(request *CreateDesktopOversoldGroupRequest, callback func(response *CreateDesktopOversoldGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDesktopOversoldGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateDesktopOversoldGroup(request)
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

// CreateDesktopOversoldGroupRequest is the request struct for api CreateDesktopOversoldGroup
type CreateDesktopOversoldGroupRequest struct {
	*requests.RpcRequest
	ImageId                string           `position:"Query" name:"ImageId"`
	OversoldUserCount      requests.Integer `position:"Query" name:"OversoldUserCount"`
	Description            string           `position:"Query" name:"Description"`
	OversoldWarn           requests.Integer `position:"Query" name:"OversoldWarn"`
	IdleDisconnectDuration requests.Integer `position:"Query" name:"IdleDisconnectDuration"`
	SystemDiskSize         requests.Integer `position:"Query" name:"SystemDiskSize"`
	DirectoryId            string           `position:"Query" name:"DirectoryId"`
	DesktopType            string           `position:"Query" name:"DesktopType"`
	Period                 requests.Integer `position:"Query" name:"Period"`
	StopDuration           requests.Integer `position:"Query" name:"StopDuration"`
	KeepDuration           requests.Integer `position:"Query" name:"KeepDuration"`
	PeriodUnit             string           `position:"Query" name:"PeriodUnit"`
	DataDiskSize           requests.Integer `position:"Query" name:"DataDiskSize"`
	Name                   string           `position:"Query" name:"Name"`
	ConcurrenceCount       requests.Integer `position:"Query" name:"ConcurrenceCount"`
	PolicyGroupId          string           `position:"Query" name:"PolicyGroupId"`
}

// CreateDesktopOversoldGroupResponse is the response struct for api CreateDesktopOversoldGroup
type CreateDesktopOversoldGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateDesktopOversoldGroupRequest creates a request to invoke CreateDesktopOversoldGroup API
func CreateCreateDesktopOversoldGroupRequest() (request *CreateDesktopOversoldGroupRequest) {
	request = &CreateDesktopOversoldGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateDesktopOversoldGroup", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDesktopOversoldGroupResponse creates a response to parse from CreateDesktopOversoldGroup response
func CreateCreateDesktopOversoldGroupResponse() (response *CreateDesktopOversoldGroupResponse) {
	response = &CreateDesktopOversoldGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
