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

// DescribeVodDomainRealTimeQpsData invokes the vod.DescribeVodDomainRealTimeQpsData API synchronously
func (client *Client) DescribeVodDomainRealTimeQpsData(request *DescribeVodDomainRealTimeQpsDataRequest) (response *DescribeVodDomainRealTimeQpsDataResponse, err error) {
	response = CreateDescribeVodDomainRealTimeQpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainRealTimeQpsDataWithChan invokes the vod.DescribeVodDomainRealTimeQpsData API asynchronously
func (client *Client) DescribeVodDomainRealTimeQpsDataWithChan(request *DescribeVodDomainRealTimeQpsDataRequest) (<-chan *DescribeVodDomainRealTimeQpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainRealTimeQpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainRealTimeQpsData(request)
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

// DescribeVodDomainRealTimeQpsDataWithCallback invokes the vod.DescribeVodDomainRealTimeQpsData API asynchronously
func (client *Client) DescribeVodDomainRealTimeQpsDataWithCallback(request *DescribeVodDomainRealTimeQpsDataRequest, callback func(response *DescribeVodDomainRealTimeQpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainRealTimeQpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainRealTimeQpsData(request)
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

// DescribeVodDomainRealTimeQpsDataRequest is the request struct for api DescribeVodDomainRealTimeQpsData
type DescribeVodDomainRealTimeQpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVodDomainRealTimeQpsDataResponse is the response struct for api DescribeVodDomainRealTimeQpsData
type DescribeVodDomainRealTimeQpsDataResponse struct {
	*responses.BaseResponse
	RequestId string                                 `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeVodDomainRealTimeQpsData `json:"Data" xml:"Data"`
}

// CreateDescribeVodDomainRealTimeQpsDataRequest creates a request to invoke DescribeVodDomainRealTimeQpsData API
func CreateDescribeVodDomainRealTimeQpsDataRequest() (request *DescribeVodDomainRealTimeQpsDataRequest) {
	request = &DescribeVodDomainRealTimeQpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainRealTimeQpsData", "vod", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeVodDomainRealTimeQpsDataResponse creates a response to parse from DescribeVodDomainRealTimeQpsData response
func CreateDescribeVodDomainRealTimeQpsDataResponse() (response *DescribeVodDomainRealTimeQpsDataResponse) {
	response = &DescribeVodDomainRealTimeQpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
