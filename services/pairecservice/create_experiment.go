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

// CreateExperiment invokes the pairecservice.CreateExperiment API synchronously
func (client *Client) CreateExperiment(request *CreateExperimentRequest) (response *CreateExperimentResponse, err error) {
	response = CreateCreateExperimentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateExperimentWithChan invokes the pairecservice.CreateExperiment API asynchronously
func (client *Client) CreateExperimentWithChan(request *CreateExperimentRequest) (<-chan *CreateExperimentResponse, <-chan error) {
	responseChan := make(chan *CreateExperimentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateExperiment(request)
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

// CreateExperimentWithCallback invokes the pairecservice.CreateExperiment API asynchronously
func (client *Client) CreateExperimentWithCallback(request *CreateExperimentRequest, callback func(response *CreateExperimentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateExperimentResponse
		var err error
		defer close(result)
		response, err = client.CreateExperiment(request)
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

// CreateExperimentRequest is the request struct for api CreateExperiment
type CreateExperimentRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateExperimentResponse is the response struct for api CreateExperiment
type CreateExperimentResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ExperimentId string `json:"ExperimentId" xml:"ExperimentId"`
}

// CreateCreateExperimentRequest creates a request to invoke CreateExperiment API
func CreateCreateExperimentRequest() (request *CreateExperimentRequest) {
	request = &CreateExperimentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateExperiment", "/api/v1/experiments", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateExperimentResponse creates a response to parse from CreateExperiment response
func CreateCreateExperimentResponse() (response *CreateExperimentResponse) {
	response = &CreateExperimentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
