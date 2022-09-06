package sas

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

// GetSuspiciousStatistics invokes the sas.GetSuspiciousStatistics API synchronously
func (client *Client) GetSuspiciousStatistics(request *GetSuspiciousStatisticsRequest) (response *GetSuspiciousStatisticsResponse, err error) {
	response = CreateGetSuspiciousStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// GetSuspiciousStatisticsWithChan invokes the sas.GetSuspiciousStatistics API asynchronously
func (client *Client) GetSuspiciousStatisticsWithChan(request *GetSuspiciousStatisticsRequest) (<-chan *GetSuspiciousStatisticsResponse, <-chan error) {
	responseChan := make(chan *GetSuspiciousStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSuspiciousStatistics(request)
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

// GetSuspiciousStatisticsWithCallback invokes the sas.GetSuspiciousStatistics API asynchronously
func (client *Client) GetSuspiciousStatisticsWithCallback(request *GetSuspiciousStatisticsRequest, callback func(response *GetSuspiciousStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSuspiciousStatisticsResponse
		var err error
		defer close(result)
		response, err = client.GetSuspiciousStatistics(request)
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

// GetSuspiciousStatisticsRequest is the request struct for api GetSuspiciousStatistics
type GetSuspiciousStatisticsRequest struct {
	*requests.RpcRequest
	GroupIdList string `position:"Query" name:"GroupIdList"`
	SourceIp    string `position:"Query" name:"SourceIp"`
}

// GetSuspiciousStatisticsResponse is the response struct for api GetSuspiciousStatistics
type GetSuspiciousStatisticsResponse struct {
	*responses.BaseResponse
	SuspiciousCount int    `json:"SuspiciousCount" xml:"SuspiciousCount"`
	RemindCount     int    `json:"RemindCount" xml:"RemindCount"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	TemindCount     int    `json:"TemindCount" xml:"TemindCount"`
	SeriousCount    int    `json:"SeriousCount" xml:"SeriousCount"`
	TotalCount      int    `json:"TotalCount" xml:"TotalCount"`
}

// CreateGetSuspiciousStatisticsRequest creates a request to invoke GetSuspiciousStatistics API
func CreateGetSuspiciousStatisticsRequest() (request *GetSuspiciousStatisticsRequest) {
	request = &GetSuspiciousStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "GetSuspiciousStatistics", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSuspiciousStatisticsResponse creates a response to parse from GetSuspiciousStatistics response
func CreateGetSuspiciousStatisticsResponse() (response *GetSuspiciousStatisticsResponse) {
	response = &GetSuspiciousStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
