package mseap

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

// PushRpaTaskDetail invokes the mseap.PushRpaTaskDetail API synchronously
func (client *Client) PushRpaTaskDetail(request *PushRpaTaskDetailRequest) (response *PushRpaTaskDetailResponse, err error) {
	response = CreatePushRpaTaskDetailResponse()
	err = client.DoAction(request, response)
	return
}

// PushRpaTaskDetailWithChan invokes the mseap.PushRpaTaskDetail API asynchronously
func (client *Client) PushRpaTaskDetailWithChan(request *PushRpaTaskDetailRequest) (<-chan *PushRpaTaskDetailResponse, <-chan error) {
	responseChan := make(chan *PushRpaTaskDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushRpaTaskDetail(request)
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

// PushRpaTaskDetailWithCallback invokes the mseap.PushRpaTaskDetail API asynchronously
func (client *Client) PushRpaTaskDetailWithCallback(request *PushRpaTaskDetailRequest, callback func(response *PushRpaTaskDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushRpaTaskDetailResponse
		var err error
		defer close(result)
		response, err = client.PushRpaTaskDetail(request)
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

// PushRpaTaskDetailRequest is the request struct for api PushRpaTaskDetail
type PushRpaTaskDetailRequest struct {
	*requests.RpcRequest
	UserCallerParentId requests.Integer `position:"Query" name:"UserCallerParentId"`
	ApiType            string           `position:"Query" name:"ApiType"`
	ModelDetailId      requests.Integer `position:"Query" name:"ModelDetailId"`
	UserKp             string           `position:"Query" name:"UserKp"`
	Lang               string           `position:"Query" name:"Lang"`
	TaskId             requests.Integer `position:"Query" name:"TaskId"`
	UserCallerType     string           `position:"Query" name:"UserCallerType"`
	UserSecurityToken  string           `position:"Query" name:"UserSecurityToken"`
	UserAccessKeyId    string           `position:"Query" name:"UserAccessKeyId"`
	AliyunKp           string           `position:"Query" name:"AliyunKp"`
	UserBid            string           `position:"Query" name:"UserBid"`
	OriginalRequest    string           `position:"Query" name:"OriginalRequest"`
	InputScreenshot    string           `position:"Query" name:"InputScreenshot"`
	InputData          string           `position:"Query" name:"InputData"`
	OutputData         string           `position:"Query" name:"OutputData"`
	InputHtml          string           `position:"Query" name:"InputHtml"`
	TaskDetailId       requests.Integer `position:"Query" name:"TaskDetailId"`
	OutputHtml         string           `position:"Query" name:"OutputHtml"`
	UserClientIp       string           `position:"Query" name:"UserClientIp"`
	Bid                string           `position:"Query" name:"Bid"`
	OutputScreenshot   string           `position:"Query" name:"OutputScreenshot"`
	Status             requests.Integer `position:"Query" name:"Status"`
}

// PushRpaTaskDetailResponse is the response struct for api PushRpaTaskDetail
type PushRpaTaskDetailResponse struct {
	*responses.BaseResponse
	AllowRetry     bool   `json:"AllowRetry" xml:"AllowRetry"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMsg       string `json:"ErrorMsg" xml:"ErrorMsg"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Module         string `json:"Module" xml:"Module"`
	Success        bool   `json:"Success" xml:"Success"`
	AppName        string `json:"AppName" xml:"AppName"`
}

// CreatePushRpaTaskDetailRequest creates a request to invoke PushRpaTaskDetail API
func CreatePushRpaTaskDetailRequest() (request *PushRpaTaskDetailRequest) {
	request = &PushRpaTaskDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mseap", "2021-01-18", "PushRpaTaskDetail", "", "")
	request.Method = requests.POST
	return
}

// CreatePushRpaTaskDetailResponse creates a response to parse from PushRpaTaskDetail response
func CreatePushRpaTaskDetailResponse() (response *PushRpaTaskDetailResponse) {
	response = &PushRpaTaskDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
