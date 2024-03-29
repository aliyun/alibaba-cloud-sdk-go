package ddoscoo

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

// DescribeInstanceStatistics invokes the ddoscoo.DescribeInstanceStatistics API synchronously
func (client *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (response *DescribeInstanceStatisticsResponse, err error) {
	response = CreateDescribeInstanceStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceStatisticsWithChan invokes the ddoscoo.DescribeInstanceStatistics API asynchronously
func (client *Client) DescribeInstanceStatisticsWithChan(request *DescribeInstanceStatisticsRequest) (<-chan *DescribeInstanceStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceStatistics(request)
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

// DescribeInstanceStatisticsWithCallback invokes the ddoscoo.DescribeInstanceStatistics API asynchronously
func (client *Client) DescribeInstanceStatisticsWithCallback(request *DescribeInstanceStatisticsRequest, callback func(response *DescribeInstanceStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceStatistics(request)
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

// DescribeInstanceStatisticsRequest is the request struct for api DescribeInstanceStatistics
type DescribeInstanceStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp    string    `position:"Query" name:"SourceIp"`
	InstanceIds *[]string `position:"Query" name:"InstanceIds"  type:"Repeated"`
}

// DescribeInstanceStatisticsResponse is the response struct for api DescribeInstanceStatistics
type DescribeInstanceStatisticsResponse struct {
	*responses.BaseResponse
	RequestId          string              `json:"RequestId" xml:"RequestId"`
	InstanceStatistics []InstanceStatistic `json:"InstanceStatistics" xml:"InstanceStatistics"`
}

// CreateDescribeInstanceStatisticsRequest creates a request to invoke DescribeInstanceStatistics API
func CreateDescribeInstanceStatisticsRequest() (request *DescribeInstanceStatisticsRequest) {
	request = &DescribeInstanceStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeInstanceStatistics", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceStatisticsResponse creates a response to parse from DescribeInstanceStatistics response
func CreateDescribeInstanceStatisticsResponse() (response *DescribeInstanceStatisticsResponse) {
	response = &DescribeInstanceStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
