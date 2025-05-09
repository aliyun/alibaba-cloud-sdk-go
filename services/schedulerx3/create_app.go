package schedulerx3

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

// CreateApp invokes the schedulerx3.CreateApp API synchronously
func (client *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
	response = CreateCreateAppResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAppWithChan invokes the schedulerx3.CreateApp API asynchronously
func (client *Client) CreateAppWithChan(request *CreateAppRequest) (<-chan *CreateAppResponse, <-chan error) {
	responseChan := make(chan *CreateAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApp(request)
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

// CreateAppWithCallback invokes the schedulerx3.CreateApp API asynchronously
func (client *Client) CreateAppWithCallback(request *CreateAppRequest, callback func(response *CreateAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAppResponse
		var err error
		defer close(result)
		response, err = client.CreateApp(request)
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

// CreateAppRequest is the request struct for api CreateApp
type CreateAppRequest struct {
	*requests.RpcRequest
	MseSessionId   string                     `position:"Query" name:"MseSessionId"`
	AccessToken    string                     `position:"Body" name:"AccessToken"`
	Title          string                     `position:"Body" name:"Title"`
	AppName        string                     `position:"Body" name:"AppName"`
	NoticeContacts *[]CreateAppNoticeContacts `position:"Body" name:"NoticeContacts"  type:"Json"`
	NoticeConfig   CreateAppNoticeConfig      `position:"Body" name:"NoticeConfig"  type:"Struct"`
	MaxJobs        requests.Integer           `position:"Body" name:"MaxJobs"`
	Calender       string                     `position:"Body" name:"Calender"`
	ClusterId      string                     `position:"Body" name:"ClusterId"`
	EnableLog      requests.Boolean           `position:"Body" name:"EnableLog"`
	MaxConcurrency requests.Integer           `position:"Body" name:"MaxConcurrency"`
}

// CreateAppNoticeContacts is a repeated param struct in CreateAppRequest
type CreateAppNoticeContacts struct {
	Webhook  string `name:"Webhook"`
	Mail     string `name:"Mail"`
	Phone    string `name:"Phone"`
	SmsCode  string `name:"SmsCode"`
	MailCode string `name:"MailCode"`
	UserName string `name:"UserName"`
}

// CreateAppNoticeConfig is a repeated param struct in CreateAppRequest
type CreateAppNoticeConfig struct {
	WebhookIsAtAll string `name:"WebhookIsAtAll"`
	NoticeChannel  string `name:"NoticeChannel"`
}

// CreateAppResponse is the response struct for api CreateApp
type CreateAppResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateAppRequest creates a request to invoke CreateApp API
func CreateCreateAppRequest() (request *CreateAppRequest) {
	request = &CreateAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SchedulerX3", "2024-06-24", "CreateApp", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAppResponse creates a response to parse from CreateApp response
func CreateCreateAppResponse() (response *CreateAppResponse) {
	response = &CreateAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
