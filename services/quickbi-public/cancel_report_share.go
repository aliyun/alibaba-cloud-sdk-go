package quickbi_public

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

// CancelReportShare invokes the quickbi_public.CancelReportShare API synchronously
func (client *Client) CancelReportShare(request *CancelReportShareRequest) (response *CancelReportShareResponse, err error) {
	response = CreateCancelReportShareResponse()
	err = client.DoAction(request, response)
	return
}

// CancelReportShareWithChan invokes the quickbi_public.CancelReportShare API asynchronously
func (client *Client) CancelReportShareWithChan(request *CancelReportShareRequest) (<-chan *CancelReportShareResponse, <-chan error) {
	responseChan := make(chan *CancelReportShareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelReportShare(request)
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

// CancelReportShareWithCallback invokes the quickbi_public.CancelReportShare API asynchronously
func (client *Client) CancelReportShareWithCallback(request *CancelReportShareRequest, callback func(response *CancelReportShareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelReportShareResponse
		var err error
		defer close(result)
		response, err = client.CancelReportShare(request)
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

// CancelReportShareRequest is the request struct for api CancelReportShare
type CancelReportShareRequest struct {
	*requests.RpcRequest
	ReportId    string           `position:"Query" name:"ReportId"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	ShareToIds  string           `position:"Query" name:"ShareToIds"`
	SignType    string           `position:"Query" name:"SignType"`
	ShareToType requests.Integer `position:"Query" name:"ShareToType"`
}

// CancelReportShareResponse is the response struct for api CancelReportShare
type CancelReportShareResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCancelReportShareRequest creates a request to invoke CancelReportShare API
func CreateCancelReportShareRequest() (request *CancelReportShareRequest) {
	request = &CancelReportShareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "CancelReportShare", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelReportShareResponse creates a response to parse from CancelReportShare response
func CreateCancelReportShareResponse() (response *CancelReportShareResponse) {
	response = &CancelReportShareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
