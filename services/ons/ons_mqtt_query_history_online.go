package ons

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

// OnsMqttQueryHistoryOnline invokes the ons.OnsMqttQueryHistoryOnline API synchronously
func (client *Client) OnsMqttQueryHistoryOnline(request *OnsMqttQueryHistoryOnlineRequest) (response *OnsMqttQueryHistoryOnlineResponse, err error) {
	response = CreateOnsMqttQueryHistoryOnlineResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMqttQueryHistoryOnlineWithChan invokes the ons.OnsMqttQueryHistoryOnline API asynchronously
func (client *Client) OnsMqttQueryHistoryOnlineWithChan(request *OnsMqttQueryHistoryOnlineRequest) (<-chan *OnsMqttQueryHistoryOnlineResponse, <-chan error) {
	responseChan := make(chan *OnsMqttQueryHistoryOnlineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttQueryHistoryOnline(request)
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

// OnsMqttQueryHistoryOnlineWithCallback invokes the ons.OnsMqttQueryHistoryOnline API asynchronously
func (client *Client) OnsMqttQueryHistoryOnlineWithCallback(request *OnsMqttQueryHistoryOnlineRequest, callback func(response *OnsMqttQueryHistoryOnlineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttQueryHistoryOnlineResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttQueryHistoryOnline(request)
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

// OnsMqttQueryHistoryOnlineRequest is the request struct for api OnsMqttQueryHistoryOnline
type OnsMqttQueryHistoryOnlineRequest struct {
	*requests.RpcRequest
	GroupId    string           `position:"Query" name:"GroupId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	BeginTime  requests.Integer `position:"Query" name:"BeginTime"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

// OnsMqttQueryHistoryOnlineResponse is the response struct for api OnsMqttQueryHistoryOnline
type OnsMqttQueryHistoryOnlineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsMqttQueryHistoryOnlineRequest creates a request to invoke OnsMqttQueryHistoryOnline API
func CreateOnsMqttQueryHistoryOnlineRequest() (request *OnsMqttQueryHistoryOnlineRequest) {
	request = &OnsMqttQueryHistoryOnlineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMqttQueryHistoryOnline", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMqttQueryHistoryOnlineResponse creates a response to parse from OnsMqttQueryHistoryOnline response
func CreateOnsMqttQueryHistoryOnlineResponse() (response *OnsMqttQueryHistoryOnlineResponse) {
	response = &OnsMqttQueryHistoryOnlineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
