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

// DescribeCustomizeRuleTest invokes the cloud_siem.DescribeCustomizeRuleTest API synchronously
func (client *Client) DescribeCustomizeRuleTest(request *DescribeCustomizeRuleTestRequest) (response *DescribeCustomizeRuleTestResponse, err error) {
	response = CreateDescribeCustomizeRuleTestResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomizeRuleTestWithChan invokes the cloud_siem.DescribeCustomizeRuleTest API asynchronously
func (client *Client) DescribeCustomizeRuleTestWithChan(request *DescribeCustomizeRuleTestRequest) (<-chan *DescribeCustomizeRuleTestResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomizeRuleTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomizeRuleTest(request)
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

// DescribeCustomizeRuleTestWithCallback invokes the cloud_siem.DescribeCustomizeRuleTest API asynchronously
func (client *Client) DescribeCustomizeRuleTestWithCallback(request *DescribeCustomizeRuleTestRequest, callback func(response *DescribeCustomizeRuleTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomizeRuleTestResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomizeRuleTest(request)
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

// DescribeCustomizeRuleTestRequest is the request struct for api DescribeCustomizeRuleTest
type DescribeCustomizeRuleTestRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// DescribeCustomizeRuleTestResponse is the response struct for api DescribeCustomizeRuleTest
type DescribeCustomizeRuleTestResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeCustomizeRuleTestRequest creates a request to invoke DescribeCustomizeRuleTest API
func CreateDescribeCustomizeRuleTestRequest() (request *DescribeCustomizeRuleTestRequest) {
	request = &DescribeCustomizeRuleTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeCustomizeRuleTest", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCustomizeRuleTestResponse creates a response to parse from DescribeCustomizeRuleTest response
func CreateDescribeCustomizeRuleTestResponse() (response *DescribeCustomizeRuleTestResponse) {
	response = &DescribeCustomizeRuleTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
