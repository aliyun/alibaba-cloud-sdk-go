package sddp

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

// DescribeDataLimitDetail invokes the sddp.DescribeDataLimitDetail API synchronously
func (client *Client) DescribeDataLimitDetail(request *DescribeDataLimitDetailRequest) (response *DescribeDataLimitDetailResponse, err error) {
	response = CreateDescribeDataLimitDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataLimitDetailWithChan invokes the sddp.DescribeDataLimitDetail API asynchronously
func (client *Client) DescribeDataLimitDetailWithChan(request *DescribeDataLimitDetailRequest) (<-chan *DescribeDataLimitDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeDataLimitDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataLimitDetail(request)
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

// DescribeDataLimitDetailWithCallback invokes the sddp.DescribeDataLimitDetail API asynchronously
func (client *Client) DescribeDataLimitDetailWithCallback(request *DescribeDataLimitDetailRequest, callback func(response *DescribeDataLimitDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataLimitDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataLimitDetail(request)
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

// DescribeDataLimitDetailRequest is the request struct for api DescribeDataLimitDetail
type DescribeDataLimitDetailRequest struct {
	*requests.RpcRequest
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	NetworkType requests.Integer `position:"Query" name:"NetworkType"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DescribeDataLimitDetailResponse is the response struct for api DescribeDataLimitDetail
type DescribeDataLimitDetailResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	DataLimit DataLimit `json:"DataLimit" xml:"DataLimit"`
}

// CreateDescribeDataLimitDetailRequest creates a request to invoke DescribeDataLimitDetail API
func CreateDescribeDataLimitDetailRequest() (request *DescribeDataLimitDetailRequest) {
	request = &DescribeDataLimitDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeDataLimitDetail", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataLimitDetailResponse creates a response to parse from DescribeDataLimitDetail response
func CreateDescribeDataLimitDetailResponse() (response *DescribeDataLimitDetailResponse) {
	response = &DescribeDataLimitDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
