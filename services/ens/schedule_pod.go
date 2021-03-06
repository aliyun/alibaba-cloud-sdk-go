package ens

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

// SchedulePod invokes the ens.SchedulePod API synchronously
func (client *Client) SchedulePod(request *SchedulePodRequest) (response *SchedulePodResponse, err error) {
	response = CreateSchedulePodResponse()
	err = client.DoAction(request, response)
	return
}

// SchedulePodWithChan invokes the ens.SchedulePod API asynchronously
func (client *Client) SchedulePodWithChan(request *SchedulePodRequest) (<-chan *SchedulePodResponse, <-chan error) {
	responseChan := make(chan *SchedulePodResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SchedulePod(request)
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

// SchedulePodWithCallback invokes the ens.SchedulePod API asynchronously
func (client *Client) SchedulePodWithCallback(request *SchedulePodRequest, callback func(response *SchedulePodResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SchedulePodResponse
		var err error
		defer close(result)
		response, err = client.SchedulePod(request)
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

// SchedulePodRequest is the request struct for api SchedulePod
type SchedulePodRequest struct {
	*requests.RpcRequest
	AreaCodes         string           `position:"Query" name:"AreaCodes"`
	GroupUuid         string           `position:"Query" name:"GroupUuid"`
	Isps              string           `position:"Query" name:"Isps"`
	Tenant            string           `position:"Query" name:"Tenant"`
	WorkloadUuid      string           `position:"Query" name:"WorkloadUuid"`
	Labels            string           `position:"Query" name:"Labels"`
	Regions           string           `position:"Query" name:"Regions"`
	ResourceAttribute string           `position:"Body" name:"ResourceAttribute"`
	AliUid            requests.Integer `position:"Query" name:"AliUid"`
	Requirements      string           `position:"Body" name:"Requirements"`
}

// SchedulePodResponse is the response struct for api SchedulePod
type SchedulePodResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int64  `json:"Code" xml:"Code"`
	Msg       string `json:"Msg" xml:"Msg"`
	Desc      string `json:"Desc" xml:"Desc"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateSchedulePodRequest creates a request to invoke SchedulePod API
func CreateSchedulePodRequest() (request *SchedulePodRequest) {
	request = &SchedulePodRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "SchedulePod", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSchedulePodResponse creates a response to parse from SchedulePod response
func CreateSchedulePodResponse() (response *SchedulePodResponse) {
	response = &SchedulePodResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
