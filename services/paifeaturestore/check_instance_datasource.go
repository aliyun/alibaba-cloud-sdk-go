package paifeaturestore

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

// CheckInstanceDatasource invokes the paifeaturestore.CheckInstanceDatasource API synchronously
func (client *Client) CheckInstanceDatasource(request *CheckInstanceDatasourceRequest) (response *CheckInstanceDatasourceResponse, err error) {
	response = CreateCheckInstanceDatasourceResponse()
	err = client.DoAction(request, response)
	return
}

// CheckInstanceDatasourceWithChan invokes the paifeaturestore.CheckInstanceDatasource API asynchronously
func (client *Client) CheckInstanceDatasourceWithChan(request *CheckInstanceDatasourceRequest) (<-chan *CheckInstanceDatasourceResponse, <-chan error) {
	responseChan := make(chan *CheckInstanceDatasourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckInstanceDatasource(request)
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

// CheckInstanceDatasourceWithCallback invokes the paifeaturestore.CheckInstanceDatasource API asynchronously
func (client *Client) CheckInstanceDatasourceWithCallback(request *CheckInstanceDatasourceRequest, callback func(response *CheckInstanceDatasourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckInstanceDatasourceResponse
		var err error
		defer close(result)
		response, err = client.CheckInstanceDatasource(request)
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

// CheckInstanceDatasourceRequest is the request struct for api CheckInstanceDatasource
type CheckInstanceDatasourceRequest struct {
	*requests.RoaRequest
	Body       string `position:"Body" name:"body"`
	InstanceId string `position:"Path" name:"InstanceId"`
}

// CheckInstanceDatasourceResponse is the response struct for api CheckInstanceDatasource
type CheckInstanceDatasourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateCheckInstanceDatasourceRequest creates a request to invoke CheckInstanceDatasource API
func CreateCheckInstanceDatasourceRequest() (request *CheckInstanceDatasourceRequest) {
	request = &CheckInstanceDatasourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "CheckInstanceDatasource", "/api/v1/instances/[InstanceId]/action/checkdatasource", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckInstanceDatasourceResponse creates a response to parse from CheckInstanceDatasource response
func CreateCheckInstanceDatasourceResponse() (response *CheckInstanceDatasourceResponse) {
	response = &CheckInstanceDatasourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
