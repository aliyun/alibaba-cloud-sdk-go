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

// UpdateExperimentGroup invokes the pairecservice.UpdateExperimentGroup API synchronously
func (client *Client) UpdateExperimentGroup(request *UpdateExperimentGroupRequest) (response *UpdateExperimentGroupResponse, err error) {
	response = CreateUpdateExperimentGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateExperimentGroupWithChan invokes the pairecservice.UpdateExperimentGroup API asynchronously
func (client *Client) UpdateExperimentGroupWithChan(request *UpdateExperimentGroupRequest) (<-chan *UpdateExperimentGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateExperimentGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateExperimentGroup(request)
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

// UpdateExperimentGroupWithCallback invokes the pairecservice.UpdateExperimentGroup API asynchronously
func (client *Client) UpdateExperimentGroupWithCallback(request *UpdateExperimentGroupRequest, callback func(response *UpdateExperimentGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateExperimentGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateExperimentGroup(request)
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

// UpdateExperimentGroupRequest is the request struct for api UpdateExperimentGroup
type UpdateExperimentGroupRequest struct {
	*requests.RoaRequest
	Body              string `position:"Body" name:"body"`
	ExperimentGroupId string `position:"Path" name:"ExperimentGroupId"`
}

// UpdateExperimentGroupResponse is the response struct for api UpdateExperimentGroup
type UpdateExperimentGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateExperimentGroupRequest creates a request to invoke UpdateExperimentGroup API
func CreateUpdateExperimentGroupRequest() (request *UpdateExperimentGroupRequest) {
	request = &UpdateExperimentGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateExperimentGroup", "/api/v1/experimentgroups/[ExperimentGroupId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateExperimentGroupResponse creates a response to parse from UpdateExperimentGroup response
func CreateUpdateExperimentGroupResponse() (response *UpdateExperimentGroupResponse) {
	response = &UpdateExperimentGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
