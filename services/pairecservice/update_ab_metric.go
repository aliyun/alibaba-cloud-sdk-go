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

// UpdateABMetric invokes the pairecservice.UpdateABMetric API synchronously
func (client *Client) UpdateABMetric(request *UpdateABMetricRequest) (response *UpdateABMetricResponse, err error) {
	response = CreateUpdateABMetricResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateABMetricWithChan invokes the pairecservice.UpdateABMetric API asynchronously
func (client *Client) UpdateABMetricWithChan(request *UpdateABMetricRequest) (<-chan *UpdateABMetricResponse, <-chan error) {
	responseChan := make(chan *UpdateABMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateABMetric(request)
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

// UpdateABMetricWithCallback invokes the pairecservice.UpdateABMetric API asynchronously
func (client *Client) UpdateABMetricWithCallback(request *UpdateABMetricRequest, callback func(response *UpdateABMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateABMetricResponse
		var err error
		defer close(result)
		response, err = client.UpdateABMetric(request)
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

// UpdateABMetricRequest is the request struct for api UpdateABMetric
type UpdateABMetricRequest struct {
	*requests.RoaRequest
	ABMetricId string `position:"Path" name:"ABMetricId"`
	Body       string `position:"Body" name:"body"`
}

// UpdateABMetricResponse is the response struct for api UpdateABMetric
type UpdateABMetricResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateABMetricRequest creates a request to invoke UpdateABMetric API
func CreateUpdateABMetricRequest() (request *UpdateABMetricRequest) {
	request = &UpdateABMetricRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateABMetric", "/api/v1/abmetrics/[ABMetricId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateABMetricResponse creates a response to parse from UpdateABMetric response
func CreateUpdateABMetricResponse() (response *UpdateABMetricResponse) {
	response = &UpdateABMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
