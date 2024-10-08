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

// DescribeUserConnectionRecords invokes the ecd.DescribeUserConnectionRecords API synchronously
func (client *Client) DescribeUserConnectionRecords(request *DescribeUserConnectionRecordsRequest) (response *DescribeUserConnectionRecordsResponse, err error) {
	response = CreateDescribeUserConnectionRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserConnectionRecordsWithChan invokes the ecd.DescribeUserConnectionRecords API asynchronously
func (client *Client) DescribeUserConnectionRecordsWithChan(request *DescribeUserConnectionRecordsRequest) (<-chan *DescribeUserConnectionRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserConnectionRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserConnectionRecords(request)
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

// DescribeUserConnectionRecordsWithCallback invokes the ecd.DescribeUserConnectionRecords API asynchronously
func (client *Client) DescribeUserConnectionRecordsWithCallback(request *DescribeUserConnectionRecordsRequest, callback func(response *DescribeUserConnectionRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserConnectionRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserConnectionRecords(request)
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

// DescribeUserConnectionRecordsRequest is the request struct for api DescribeUserConnectionRecords
type DescribeUserConnectionRecordsRequest struct {
	*requests.RpcRequest
	ConnectEndTimeFrom   requests.Integer `position:"Query" name:"ConnectEndTimeFrom"`
	ConnectDurationFrom  requests.Integer `position:"Query" name:"ConnectDurationFrom"`
	ConnectDurationTo    requests.Integer `position:"Query" name:"ConnectDurationTo"`
	EndUserType          string           `position:"Query" name:"EndUserType"`
	DesktopGroupId       string           `position:"Query" name:"DesktopGroupId"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ConnectStartTimeFrom requests.Integer `position:"Query" name:"ConnectStartTimeFrom"`
	EndUserId            string           `position:"Query" name:"EndUserId"`
	DesktopId            string           `position:"Query" name:"DesktopId"`
	ConnectEndTimeTo     requests.Integer `position:"Query" name:"ConnectEndTimeTo"`
	ConnectStartTimeTo   requests.Integer `position:"Query" name:"ConnectStartTimeTo"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeUserConnectionRecordsResponse is the response struct for api DescribeUserConnectionRecords
type DescribeUserConnectionRecordsResponse struct {
	*responses.BaseResponse
	NextToken         string             `json:"NextToken" xml:"NextToken"`
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	ConnectionRecords []ConnectionRecord `json:"ConnectionRecords" xml:"ConnectionRecords"`
}

// CreateDescribeUserConnectionRecordsRequest creates a request to invoke DescribeUserConnectionRecords API
func CreateDescribeUserConnectionRecordsRequest() (request *DescribeUserConnectionRecordsRequest) {
	request = &DescribeUserConnectionRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeUserConnectionRecords", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserConnectionRecordsResponse creates a response to parse from DescribeUserConnectionRecords response
func CreateDescribeUserConnectionRecordsResponse() (response *DescribeUserConnectionRecordsResponse) {
	response = &DescribeUserConnectionRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
