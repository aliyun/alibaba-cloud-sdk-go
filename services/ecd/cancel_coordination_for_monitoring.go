package ecd

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

// CancelCoordinationForMonitoring invokes the ecd.CancelCoordinationForMonitoring API synchronously
func (client *Client) CancelCoordinationForMonitoring(request *CancelCoordinationForMonitoringRequest) (response *CancelCoordinationForMonitoringResponse, err error) {
	response = CreateCancelCoordinationForMonitoringResponse()
	err = client.DoAction(request, response)
	return
}

// CancelCoordinationForMonitoringWithChan invokes the ecd.CancelCoordinationForMonitoring API asynchronously
func (client *Client) CancelCoordinationForMonitoringWithChan(request *CancelCoordinationForMonitoringRequest) (<-chan *CancelCoordinationForMonitoringResponse, <-chan error) {
	responseChan := make(chan *CancelCoordinationForMonitoringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelCoordinationForMonitoring(request)
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

// CancelCoordinationForMonitoringWithCallback invokes the ecd.CancelCoordinationForMonitoring API asynchronously
func (client *Client) CancelCoordinationForMonitoringWithCallback(request *CancelCoordinationForMonitoringRequest, callback func(response *CancelCoordinationForMonitoringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelCoordinationForMonitoringResponse
		var err error
		defer close(result)
		response, err = client.CancelCoordinationForMonitoring(request)
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

// CancelCoordinationForMonitoringRequest is the request struct for api CancelCoordinationForMonitoring
type CancelCoordinationForMonitoringRequest struct {
	*requests.RpcRequest
	UserType  string    `position:"Query" name:"UserType"`
	CoIds     *[]string `position:"Query" name:"CoIds"  type:"Repeated"`
	EndUserId string    `position:"Query" name:"EndUserId"`
}

// CancelCoordinationForMonitoringResponse is the response struct for api CancelCoordinationForMonitoring
type CancelCoordinationForMonitoringResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelCoordinationForMonitoringRequest creates a request to invoke CancelCoordinationForMonitoring API
func CreateCancelCoordinationForMonitoringRequest() (request *CancelCoordinationForMonitoringRequest) {
	request = &CancelCoordinationForMonitoringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CancelCoordinationForMonitoring", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelCoordinationForMonitoringResponse creates a response to parse from CancelCoordinationForMonitoring response
func CreateCancelCoordinationForMonitoringResponse() (response *CancelCoordinationForMonitoringResponse) {
	response = &CancelCoordinationForMonitoringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
