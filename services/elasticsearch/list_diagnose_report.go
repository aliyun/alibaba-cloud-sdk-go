package elasticsearch

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

// ListDiagnoseReport invokes the elasticsearch.ListDiagnoseReport API synchronously
func (client *Client) ListDiagnoseReport(request *ListDiagnoseReportRequest) (response *ListDiagnoseReportResponse, err error) {
	response = CreateListDiagnoseReportResponse()
	err = client.DoAction(request, response)
	return
}

// ListDiagnoseReportWithChan invokes the elasticsearch.ListDiagnoseReport API asynchronously
func (client *Client) ListDiagnoseReportWithChan(request *ListDiagnoseReportRequest) (<-chan *ListDiagnoseReportResponse, <-chan error) {
	responseChan := make(chan *ListDiagnoseReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDiagnoseReport(request)
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

// ListDiagnoseReportWithCallback invokes the elasticsearch.ListDiagnoseReport API asynchronously
func (client *Client) ListDiagnoseReportWithCallback(request *ListDiagnoseReportRequest, callback func(response *ListDiagnoseReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDiagnoseReportResponse
		var err error
		defer close(result)
		response, err = client.ListDiagnoseReport(request)
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

// ListDiagnoseReportRequest is the request struct for api ListDiagnoseReport
type ListDiagnoseReportRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Size       requests.Integer `position:"Query" name:"size"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
	Page       requests.Integer `position:"Query" name:"page"`
	Detail     requests.Boolean `position:"Query" name:"detail"`
	Trigger    string           `position:"Query" name:"trigger"`
	Lang       string           `position:"Query" name:"lang"`
}

// ListDiagnoseReportResponse is the response struct for api ListDiagnoseReport
type ListDiagnoseReportResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Headers   HeadersInListDiagnoseReport      `json:"Headers" xml:"Headers"`
	Result    []ResultItemInListDiagnoseReport `json:"Result" xml:"Result"`
}

// CreateListDiagnoseReportRequest creates a request to invoke ListDiagnoseReport API
func CreateListDiagnoseReportRequest() (request *ListDiagnoseReportRequest) {
	request = &ListDiagnoseReportRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListDiagnoseReport", "/openapi/diagnosis/instances/[InstanceId]/reports", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDiagnoseReportResponse creates a response to parse from ListDiagnoseReport response
func CreateListDiagnoseReportResponse() (response *ListDiagnoseReportResponse) {
	response = &ListDiagnoseReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
