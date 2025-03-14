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

// CreateSampleConsistencyJob invokes the pairecservice.CreateSampleConsistencyJob API synchronously
func (client *Client) CreateSampleConsistencyJob(request *CreateSampleConsistencyJobRequest) (response *CreateSampleConsistencyJobResponse, err error) {
	response = CreateCreateSampleConsistencyJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSampleConsistencyJobWithChan invokes the pairecservice.CreateSampleConsistencyJob API asynchronously
func (client *Client) CreateSampleConsistencyJobWithChan(request *CreateSampleConsistencyJobRequest) (<-chan *CreateSampleConsistencyJobResponse, <-chan error) {
	responseChan := make(chan *CreateSampleConsistencyJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSampleConsistencyJob(request)
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

// CreateSampleConsistencyJobWithCallback invokes the pairecservice.CreateSampleConsistencyJob API asynchronously
func (client *Client) CreateSampleConsistencyJobWithCallback(request *CreateSampleConsistencyJobRequest, callback func(response *CreateSampleConsistencyJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSampleConsistencyJobResponse
		var err error
		defer close(result)
		response, err = client.CreateSampleConsistencyJob(request)
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

// CreateSampleConsistencyJobRequest is the request struct for api CreateSampleConsistencyJob
type CreateSampleConsistencyJobRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateSampleConsistencyJobResponse is the response struct for api CreateSampleConsistencyJob
type CreateSampleConsistencyJobResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	SampleConsistencyJobId string `json:"SampleConsistencyJobId" xml:"SampleConsistencyJobId"`
}

// CreateCreateSampleConsistencyJobRequest creates a request to invoke CreateSampleConsistencyJob API
func CreateCreateSampleConsistencyJobRequest() (request *CreateSampleConsistencyJobRequest) {
	request = &CreateSampleConsistencyJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateSampleConsistencyJob", "/api/v1/sampleconsistencyjobs", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSampleConsistencyJobResponse creates a response to parse from CreateSampleConsistencyJob response
func CreateCreateSampleConsistencyJobResponse() (response *CreateSampleConsistencyJobResponse) {
	response = &CreateSampleConsistencyJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
