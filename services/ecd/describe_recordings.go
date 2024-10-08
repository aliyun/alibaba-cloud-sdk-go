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

// DescribeRecordings invokes the ecd.DescribeRecordings API synchronously
func (client *Client) DescribeRecordings(request *DescribeRecordingsRequest) (response *DescribeRecordingsResponse, err error) {
	response = CreateDescribeRecordingsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordingsWithChan invokes the ecd.DescribeRecordings API asynchronously
func (client *Client) DescribeRecordingsWithChan(request *DescribeRecordingsRequest) (<-chan *DescribeRecordingsResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordings(request)
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

// DescribeRecordingsWithCallback invokes the ecd.DescribeRecordings API asynchronously
func (client *Client) DescribeRecordingsWithCallback(request *DescribeRecordingsRequest, callback func(response *DescribeRecordingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordingsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordings(request)
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

// DescribeRecordingsRequest is the request struct for api DescribeRecordings
type DescribeRecordingsRequest struct {
	*requests.RpcRequest
	NeedSignedUrl          requests.Boolean `position:"Query" name:"NeedSignedUrl"`
	SignedUrlExpireMinutes requests.Integer `position:"Query" name:"SignedUrlExpireMinutes"`
	StartTime              string           `position:"Query" name:"StartTime"`
	NextToken              string           `position:"Query" name:"NextToken"`
	EndUserId              string           `position:"Query" name:"EndUserId"`
	DesktopId              string           `position:"Query" name:"DesktopId"`
	EndTime                string           `position:"Query" name:"EndTime"`
	MaxResults             requests.Integer `position:"Query" name:"MaxResults"`
	PolicyGroupId          string           `position:"Query" name:"PolicyGroupId"`
}

// DescribeRecordingsResponse is the response struct for api DescribeRecordings
type DescribeRecordingsResponse struct {
	*responses.BaseResponse
	NextToken  string      `json:"NextToken" xml:"NextToken"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Recordings []Recording `json:"Recordings" xml:"Recordings"`
}

// CreateDescribeRecordingsRequest creates a request to invoke DescribeRecordings API
func CreateDescribeRecordingsRequest() (request *DescribeRecordingsRequest) {
	request = &DescribeRecordingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeRecordings", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordingsResponse creates a response to parse from DescribeRecordings response
func CreateDescribeRecordingsResponse() (response *DescribeRecordingsResponse) {
	response = &DescribeRecordingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
