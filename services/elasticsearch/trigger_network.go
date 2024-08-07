package elasticsearch

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

// TriggerNetwork invokes the elasticsearch.TriggerNetwork API synchronously
func (client *Client) TriggerNetwork(request *TriggerNetworkRequest) (response *TriggerNetworkResponse, err error) {
	response = CreateTriggerNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// TriggerNetworkWithChan invokes the elasticsearch.TriggerNetwork API asynchronously
func (client *Client) TriggerNetworkWithChan(request *TriggerNetworkRequest) (<-chan *TriggerNetworkResponse, <-chan error) {
	responseChan := make(chan *TriggerNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TriggerNetwork(request)
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

// TriggerNetworkWithCallback invokes the elasticsearch.TriggerNetwork API asynchronously
func (client *Client) TriggerNetworkWithCallback(request *TriggerNetworkRequest, callback func(response *TriggerNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TriggerNetworkResponse
		var err error
		defer close(result)
		response, err = client.TriggerNetwork(request)
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

// TriggerNetworkRequest is the request struct for api TriggerNetwork
type TriggerNetworkRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	Body        string `position:"Body" name:"body"`
}

// TriggerNetworkResponse is the response struct for api TriggerNetwork
type TriggerNetworkResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTriggerNetworkRequest creates a request to invoke TriggerNetwork API
func CreateTriggerNetworkRequest() (request *TriggerNetworkRequest) {
	request = &TriggerNetworkRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "TriggerNetwork", "/openapi/instances/[InstanceId]/actions/network-trigger", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTriggerNetworkResponse creates a response to parse from TriggerNetwork response
func CreateTriggerNetworkResponse() (response *TriggerNetworkResponse) {
	response = &TriggerNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
