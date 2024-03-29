package oms

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

// CheckReadyFlag invokes the oms.CheckReadyFlag API synchronously
func (client *Client) CheckReadyFlag(request *CheckReadyFlagRequest) (response *CheckReadyFlagResponse, err error) {
	response = CreateCheckReadyFlagResponse()
	err = client.DoAction(request, response)
	return
}

// CheckReadyFlagWithChan invokes the oms.CheckReadyFlag API asynchronously
func (client *Client) CheckReadyFlagWithChan(request *CheckReadyFlagRequest) (<-chan *CheckReadyFlagResponse, <-chan error) {
	responseChan := make(chan *CheckReadyFlagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckReadyFlag(request)
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

// CheckReadyFlagWithCallback invokes the oms.CheckReadyFlag API asynchronously
func (client *Client) CheckReadyFlagWithCallback(request *CheckReadyFlagRequest, callback func(response *CheckReadyFlagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckReadyFlagResponse
		var err error
		defer close(result)
		response, err = client.CheckReadyFlag(request)
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

// CheckReadyFlagRequest is the request struct for api CheckReadyFlag
type CheckReadyFlagRequest struct {
	*requests.RpcRequest
	Period     string           `position:"Query" name:"Period"`
	DomainCode string           `position:"Query" name:"DomainCode"`
	DataType   string           `position:"Query" name:"DataType"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	Cycles     requests.Integer `position:"Query" name:"Cycles"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
}

// CheckReadyFlagResponse is the response struct for api CheckReadyFlag
type CheckReadyFlagResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckReadyFlagRequest creates a request to invoke CheckReadyFlag API
func CreateCheckReadyFlagRequest() (request *CheckReadyFlagRequest) {
	request = &CheckReadyFlagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "CheckReadyFlag", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckReadyFlagResponse creates a response to parse from CheckReadyFlag response
func CreateCheckReadyFlagResponse() (response *CheckReadyFlagResponse) {
	response = &CheckReadyFlagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
