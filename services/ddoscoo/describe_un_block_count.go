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

// DescribeUnBlockCount invokes the ddoscoo.DescribeUnBlockCount API synchronously
func (client *Client) DescribeUnBlockCount(request *DescribeUnBlockCountRequest) (response *DescribeUnBlockCountResponse, err error) {
	response = CreateDescribeUnBlockCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUnBlockCountWithChan invokes the ddoscoo.DescribeUnBlockCount API asynchronously
func (client *Client) DescribeUnBlockCountWithChan(request *DescribeUnBlockCountRequest) (<-chan *DescribeUnBlockCountResponse, <-chan error) {
	responseChan := make(chan *DescribeUnBlockCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUnBlockCount(request)
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

// DescribeUnBlockCountWithCallback invokes the ddoscoo.DescribeUnBlockCount API asynchronously
func (client *Client) DescribeUnBlockCountWithCallback(request *DescribeUnBlockCountRequest, callback func(response *DescribeUnBlockCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUnBlockCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeUnBlockCount(request)
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

// DescribeUnBlockCountRequest is the request struct for api DescribeUnBlockCount
type DescribeUnBlockCountRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeUnBlockCountResponse is the response struct for api DescribeUnBlockCount
type DescribeUnBlockCountResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	RemainCount int    `json:"RemainCount" xml:"RemainCount"`
}

// CreateDescribeUnBlockCountRequest creates a request to invoke DescribeUnBlockCount API
func CreateDescribeUnBlockCountRequest() (request *DescribeUnBlockCountRequest) {
	request = &DescribeUnBlockCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeUnBlockCount", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUnBlockCountResponse creates a response to parse from DescribeUnBlockCount response
func CreateDescribeUnBlockCountResponse() (response *DescribeUnBlockCountResponse) {
	response = &DescribeUnBlockCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
