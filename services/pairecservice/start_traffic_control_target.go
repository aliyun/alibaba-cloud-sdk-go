package pairecservice

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

// StartTrafficControlTarget invokes the pairecservice.StartTrafficControlTarget API synchronously
func (client *Client) StartTrafficControlTarget(request *StartTrafficControlTargetRequest) (response *StartTrafficControlTargetResponse, err error) {
	response = CreateStartTrafficControlTargetResponse()
	err = client.DoAction(request, response)
	return
}

// StartTrafficControlTargetWithChan invokes the pairecservice.StartTrafficControlTarget API asynchronously
func (client *Client) StartTrafficControlTargetWithChan(request *StartTrafficControlTargetRequest) (<-chan *StartTrafficControlTargetResponse, <-chan error) {
	responseChan := make(chan *StartTrafficControlTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartTrafficControlTarget(request)
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

// StartTrafficControlTargetWithCallback invokes the pairecservice.StartTrafficControlTarget API asynchronously
func (client *Client) StartTrafficControlTargetWithCallback(request *StartTrafficControlTargetRequest, callback func(response *StartTrafficControlTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartTrafficControlTargetResponse
		var err error
		defer close(result)
		response, err = client.StartTrafficControlTarget(request)
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

// StartTrafficControlTargetRequest is the request struct for api StartTrafficControlTarget
type StartTrafficControlTargetRequest struct {
	*requests.RoaRequest
	TrafficControlTargetId string `position:"Path" name:"TrafficControlTargetId"`
	Body                   string `position:"Body" name:"body"`
}

// StartTrafficControlTargetResponse is the response struct for api StartTrafficControlTarget
type StartTrafficControlTargetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartTrafficControlTargetRequest creates a request to invoke StartTrafficControlTarget API
func CreateStartTrafficControlTargetRequest() (request *StartTrafficControlTargetRequest) {
	request = &StartTrafficControlTargetRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "StartTrafficControlTarget", "/api/v1/trafficcontroltargets/[TrafficControlTargetId]/action/start", "", "")
	request.Method = requests.POST
	return
}

// CreateStartTrafficControlTargetResponse creates a response to parse from StartTrafficControlTarget response
func CreateStartTrafficControlTargetResponse() (response *StartTrafficControlTargetResponse) {
	response = &StartTrafficControlTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
