package cloud_siem

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

// DescribeCustomizeRuleCount invokes the cloud_siem.DescribeCustomizeRuleCount API synchronously
func (client *Client) DescribeCustomizeRuleCount(request *DescribeCustomizeRuleCountRequest) (response *DescribeCustomizeRuleCountResponse, err error) {
	response = CreateDescribeCustomizeRuleCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomizeRuleCountWithChan invokes the cloud_siem.DescribeCustomizeRuleCount API asynchronously
func (client *Client) DescribeCustomizeRuleCountWithChan(request *DescribeCustomizeRuleCountRequest) (<-chan *DescribeCustomizeRuleCountResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomizeRuleCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomizeRuleCount(request)
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

// DescribeCustomizeRuleCountWithCallback invokes the cloud_siem.DescribeCustomizeRuleCount API asynchronously
func (client *Client) DescribeCustomizeRuleCountWithCallback(request *DescribeCustomizeRuleCountRequest, callback func(response *DescribeCustomizeRuleCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomizeRuleCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomizeRuleCount(request)
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

// DescribeCustomizeRuleCountRequest is the request struct for api DescribeCustomizeRuleCount
type DescribeCustomizeRuleCountRequest struct {
	*requests.RpcRequest
}

// DescribeCustomizeRuleCountResponse is the response struct for api DescribeCustomizeRuleCount
type DescribeCustomizeRuleCountResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeCustomizeRuleCountRequest creates a request to invoke DescribeCustomizeRuleCount API
func CreateDescribeCustomizeRuleCountRequest() (request *DescribeCustomizeRuleCountRequest) {
	request = &DescribeCustomizeRuleCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeCustomizeRuleCount", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCustomizeRuleCountResponse creates a response to parse from DescribeCustomizeRuleCount response
func CreateDescribeCustomizeRuleCountResponse() (response *DescribeCustomizeRuleCountResponse) {
	response = &DescribeCustomizeRuleCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
