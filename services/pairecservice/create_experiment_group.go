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

// CreateExperimentGroup invokes the pairecservice.CreateExperimentGroup API synchronously
func (client *Client) CreateExperimentGroup(request *CreateExperimentGroupRequest) (response *CreateExperimentGroupResponse, err error) {
	response = CreateCreateExperimentGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateExperimentGroupWithChan invokes the pairecservice.CreateExperimentGroup API asynchronously
func (client *Client) CreateExperimentGroupWithChan(request *CreateExperimentGroupRequest) (<-chan *CreateExperimentGroupResponse, <-chan error) {
	responseChan := make(chan *CreateExperimentGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateExperimentGroup(request)
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

// CreateExperimentGroupWithCallback invokes the pairecservice.CreateExperimentGroup API asynchronously
func (client *Client) CreateExperimentGroupWithCallback(request *CreateExperimentGroupRequest, callback func(response *CreateExperimentGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateExperimentGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateExperimentGroup(request)
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

// CreateExperimentGroupRequest is the request struct for api CreateExperimentGroup
type CreateExperimentGroupRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateExperimentGroupResponse is the response struct for api CreateExperimentGroup
type CreateExperimentGroupResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ExperimentGroupId string `json:"ExperimentGroupId" xml:"ExperimentGroupId"`
}

// CreateCreateExperimentGroupRequest creates a request to invoke CreateExperimentGroup API
func CreateCreateExperimentGroupRequest() (request *CreateExperimentGroupRequest) {
	request = &CreateExperimentGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateExperimentGroup", "/api/v1/experimentgroups", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateExperimentGroupResponse creates a response to parse from CreateExperimentGroup response
func CreateCreateExperimentGroupResponse() (response *CreateExperimentGroupResponse) {
	response = &CreateExperimentGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
