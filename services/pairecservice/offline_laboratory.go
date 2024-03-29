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

// OfflineLaboratory invokes the pairecservice.OfflineLaboratory API synchronously
func (client *Client) OfflineLaboratory(request *OfflineLaboratoryRequest) (response *OfflineLaboratoryResponse, err error) {
	response = CreateOfflineLaboratoryResponse()
	err = client.DoAction(request, response)
	return
}

// OfflineLaboratoryWithChan invokes the pairecservice.OfflineLaboratory API asynchronously
func (client *Client) OfflineLaboratoryWithChan(request *OfflineLaboratoryRequest) (<-chan *OfflineLaboratoryResponse, <-chan error) {
	responseChan := make(chan *OfflineLaboratoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OfflineLaboratory(request)
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

// OfflineLaboratoryWithCallback invokes the pairecservice.OfflineLaboratory API asynchronously
func (client *Client) OfflineLaboratoryWithCallback(request *OfflineLaboratoryRequest, callback func(response *OfflineLaboratoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OfflineLaboratoryResponse
		var err error
		defer close(result)
		response, err = client.OfflineLaboratory(request)
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

// OfflineLaboratoryRequest is the request struct for api OfflineLaboratory
type OfflineLaboratoryRequest struct {
	*requests.RoaRequest
	LaboratoryId string `position:"Path" name:"LaboratoryId"`
	Body         string `position:"Body" name:"body"`
}

// OfflineLaboratoryResponse is the response struct for api OfflineLaboratory
type OfflineLaboratoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOfflineLaboratoryRequest creates a request to invoke OfflineLaboratory API
func CreateOfflineLaboratoryRequest() (request *OfflineLaboratoryRequest) {
	request = &OfflineLaboratoryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "OfflineLaboratory", "/api/v1/laboratories/[LaboratoryId]/action/offline", "", "")
	request.Method = requests.POST
	return
}

// CreateOfflineLaboratoryResponse creates a response to parse from OfflineLaboratory response
func CreateOfflineLaboratoryResponse() (response *OfflineLaboratoryResponse) {
	response = &OfflineLaboratoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
