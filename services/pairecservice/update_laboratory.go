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

// UpdateLaboratory invokes the pairecservice.UpdateLaboratory API synchronously
func (client *Client) UpdateLaboratory(request *UpdateLaboratoryRequest) (response *UpdateLaboratoryResponse, err error) {
	response = CreateUpdateLaboratoryResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLaboratoryWithChan invokes the pairecservice.UpdateLaboratory API asynchronously
func (client *Client) UpdateLaboratoryWithChan(request *UpdateLaboratoryRequest) (<-chan *UpdateLaboratoryResponse, <-chan error) {
	responseChan := make(chan *UpdateLaboratoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLaboratory(request)
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

// UpdateLaboratoryWithCallback invokes the pairecservice.UpdateLaboratory API asynchronously
func (client *Client) UpdateLaboratoryWithCallback(request *UpdateLaboratoryRequest, callback func(response *UpdateLaboratoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLaboratoryResponse
		var err error
		defer close(result)
		response, err = client.UpdateLaboratory(request)
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

// UpdateLaboratoryRequest is the request struct for api UpdateLaboratory
type UpdateLaboratoryRequest struct {
	*requests.RoaRequest
	LaboratoryId string `position:"Path" name:"LaboratoryId"`
	Body         string `position:"Body" name:"body"`
}

// UpdateLaboratoryResponse is the response struct for api UpdateLaboratory
type UpdateLaboratoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLaboratoryRequest creates a request to invoke UpdateLaboratory API
func CreateUpdateLaboratoryRequest() (request *UpdateLaboratoryRequest) {
	request = &UpdateLaboratoryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateLaboratory", "/api/v1/laboratories/[LaboratoryId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateLaboratoryResponse creates a response to parse from UpdateLaboratory response
func CreateUpdateLaboratoryResponse() (response *UpdateLaboratoryResponse) {
	response = &UpdateLaboratoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
