package vod

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

// DescribeVodEditingUsageData invokes the vod.DescribeVodEditingUsageData API synchronously
func (client *Client) DescribeVodEditingUsageData(request *DescribeVodEditingUsageDataRequest) (response *DescribeVodEditingUsageDataResponse, err error) {
	response = CreateDescribeVodEditingUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodEditingUsageDataWithChan invokes the vod.DescribeVodEditingUsageData API asynchronously
func (client *Client) DescribeVodEditingUsageDataWithChan(request *DescribeVodEditingUsageDataRequest) (<-chan *DescribeVodEditingUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodEditingUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodEditingUsageData(request)
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

// DescribeVodEditingUsageDataWithCallback invokes the vod.DescribeVodEditingUsageData API asynchronously
func (client *Client) DescribeVodEditingUsageDataWithCallback(request *DescribeVodEditingUsageDataRequest, callback func(response *DescribeVodEditingUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodEditingUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodEditingUsageData(request)
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

// DescribeVodEditingUsageDataRequest is the request struct for api DescribeVodEditingUsageData
type DescribeVodEditingUsageDataRequest struct {
	*requests.RpcRequest
	StartTime     string           `position:"Query" name:"StartTime"`
	SplitBy       string           `position:"Query" name:"SplitBy"`
	Product       string           `position:"Query" name:"Product"`
	EndTime       string           `position:"Query" name:"EndTime"`
	Specification string           `position:"Query" name:"Specification"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	AppId         string           `position:"Query" name:"AppId"`
	Region        string           `position:"Query" name:"Region"`
}

// DescribeVodEditingUsageDataResponse is the response struct for api DescribeVodEditingUsageData
type DescribeVodEditingUsageDataResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	StartTime   string       `json:"StartTime" xml:"StartTime"`
	EndTime     string       `json:"EndTime" xml:"EndTime"`
	EditingData []DataModule `json:"EditingData" xml:"EditingData"`
}

// CreateDescribeVodEditingUsageDataRequest creates a request to invoke DescribeVodEditingUsageData API
func CreateDescribeVodEditingUsageDataRequest() (request *DescribeVodEditingUsageDataRequest) {
	request = &DescribeVodEditingUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodEditingUsageData", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodEditingUsageDataResponse creates a response to parse from DescribeVodEditingUsageData response
func CreateDescribeVodEditingUsageDataResponse() (response *DescribeVodEditingUsageDataResponse) {
	response = &DescribeVodEditingUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
