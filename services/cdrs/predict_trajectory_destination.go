package cdrs

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

// PredictTrajectoryDestination invokes the cdrs.PredictTrajectoryDestination API synchronously
func (client *Client) PredictTrajectoryDestination(request *PredictTrajectoryDestinationRequest) (response *PredictTrajectoryDestinationResponse, err error) {
	response = CreatePredictTrajectoryDestinationResponse()
	err = client.DoAction(request, response)
	return
}

// PredictTrajectoryDestinationWithChan invokes the cdrs.PredictTrajectoryDestination API asynchronously
func (client *Client) PredictTrajectoryDestinationWithChan(request *PredictTrajectoryDestinationRequest) (<-chan *PredictTrajectoryDestinationResponse, <-chan error) {
	responseChan := make(chan *PredictTrajectoryDestinationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PredictTrajectoryDestination(request)
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

// PredictTrajectoryDestinationWithCallback invokes the cdrs.PredictTrajectoryDestination API asynchronously
func (client *Client) PredictTrajectoryDestinationWithCallback(request *PredictTrajectoryDestinationRequest, callback func(response *PredictTrajectoryDestinationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PredictTrajectoryDestinationResponse
		var err error
		defer close(result)
		response, err = client.PredictTrajectoryDestination(request)
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

// PredictTrajectoryDestinationRequest is the request struct for api PredictTrajectoryDestination
type PredictTrajectoryDestinationRequest struct {
	*requests.RpcRequest
	CorpId          string           `position:"Body" name:"CorpId"`
	IdValue         string           `position:"Body" name:"IdValue"`
	IdType          string           `position:"Body" name:"IdType"`
	PredictTimeSpan requests.Integer `position:"Body" name:"PredictTimeSpan"`
}

// PredictTrajectoryDestinationResponse is the response struct for api PredictTrajectoryDestination
type PredictTrajectoryDestinationResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePredictTrajectoryDestinationRequest creates a request to invoke PredictTrajectoryDestination API
func CreatePredictTrajectoryDestinationRequest() (request *PredictTrajectoryDestinationRequest) {
	request = &PredictTrajectoryDestinationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "PredictTrajectoryDestination", "", "")
	request.Method = requests.POST
	return
}

// CreatePredictTrajectoryDestinationResponse creates a response to parse from PredictTrajectoryDestination response
func CreatePredictTrajectoryDestinationResponse() (response *PredictTrajectoryDestinationResponse) {
	response = &PredictTrajectoryDestinationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
