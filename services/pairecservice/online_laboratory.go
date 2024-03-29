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

// OnlineLaboratory invokes the pairecservice.OnlineLaboratory API synchronously
func (client *Client) OnlineLaboratory(request *OnlineLaboratoryRequest) (response *OnlineLaboratoryResponse, err error) {
	response = CreateOnlineLaboratoryResponse()
	err = client.DoAction(request, response)
	return
}

// OnlineLaboratoryWithChan invokes the pairecservice.OnlineLaboratory API asynchronously
func (client *Client) OnlineLaboratoryWithChan(request *OnlineLaboratoryRequest) (<-chan *OnlineLaboratoryResponse, <-chan error) {
	responseChan := make(chan *OnlineLaboratoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnlineLaboratory(request)
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

// OnlineLaboratoryWithCallback invokes the pairecservice.OnlineLaboratory API asynchronously
func (client *Client) OnlineLaboratoryWithCallback(request *OnlineLaboratoryRequest, callback func(response *OnlineLaboratoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnlineLaboratoryResponse
		var err error
		defer close(result)
		response, err = client.OnlineLaboratory(request)
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

// OnlineLaboratoryRequest is the request struct for api OnlineLaboratory
type OnlineLaboratoryRequest struct {
	*requests.RoaRequest
	LaboratoryId string `position:"Path" name:"LaboratoryId"`
	Body         string `position:"Body" name:"body"`
}

// OnlineLaboratoryResponse is the response struct for api OnlineLaboratory
type OnlineLaboratoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOnlineLaboratoryRequest creates a request to invoke OnlineLaboratory API
func CreateOnlineLaboratoryRequest() (request *OnlineLaboratoryRequest) {
	request = &OnlineLaboratoryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "OnlineLaboratory", "/api/v1/laboratories/[LaboratoryId]/action/online", "", "")
	request.Method = requests.POST
	return
}

// CreateOnlineLaboratoryResponse creates a response to parse from OnlineLaboratory response
func CreateOnlineLaboratoryResponse() (response *OnlineLaboratoryResponse) {
	response = &OnlineLaboratoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
