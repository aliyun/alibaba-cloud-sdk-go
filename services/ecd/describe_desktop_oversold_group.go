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

// DescribeDesktopOversoldGroup invokes the ecd.DescribeDesktopOversoldGroup API synchronously
func (client *Client) DescribeDesktopOversoldGroup(request *DescribeDesktopOversoldGroupRequest) (response *DescribeDesktopOversoldGroupResponse, err error) {
	response = CreateDescribeDesktopOversoldGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDesktopOversoldGroupWithChan invokes the ecd.DescribeDesktopOversoldGroup API asynchronously
func (client *Client) DescribeDesktopOversoldGroupWithChan(request *DescribeDesktopOversoldGroupRequest) (<-chan *DescribeDesktopOversoldGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeDesktopOversoldGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDesktopOversoldGroup(request)
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

// DescribeDesktopOversoldGroupWithCallback invokes the ecd.DescribeDesktopOversoldGroup API asynchronously
func (client *Client) DescribeDesktopOversoldGroupWithCallback(request *DescribeDesktopOversoldGroupRequest, callback func(response *DescribeDesktopOversoldGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDesktopOversoldGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeDesktopOversoldGroup(request)
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

// DescribeDesktopOversoldGroupRequest is the request struct for api DescribeDesktopOversoldGroup
type DescribeDesktopOversoldGroupRequest struct {
	*requests.RpcRequest
	NextToken        string           `position:"Query" name:"NextToken"`
	MaxResults       requests.Integer `position:"Query" name:"MaxResults"`
	OversoldGroupIds *[]string        `position:"Query" name:"OversoldGroupIds"  type:"Repeated"`
}

// DescribeDesktopOversoldGroupResponse is the response struct for api DescribeDesktopOversoldGroup
type DescribeDesktopOversoldGroupResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Count     int        `json:"Count" xml:"Count"`
	NextToken string     `json:"NextToken" xml:"NextToken"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeDesktopOversoldGroupRequest creates a request to invoke DescribeDesktopOversoldGroup API
func CreateDescribeDesktopOversoldGroupRequest() (request *DescribeDesktopOversoldGroupRequest) {
	request = &DescribeDesktopOversoldGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeDesktopOversoldGroup", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDesktopOversoldGroupResponse creates a response to parse from DescribeDesktopOversoldGroup response
func CreateDescribeDesktopOversoldGroupResponse() (response *DescribeDesktopOversoldGroupResponse) {
	response = &DescribeDesktopOversoldGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
