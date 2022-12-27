package ens

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

// DescribeResourceTimeline invokes the ens.DescribeResourceTimeline API synchronously
func (client *Client) DescribeResourceTimeline(request *DescribeResourceTimelineRequest) (response *DescribeResourceTimelineResponse, err error) {
	response = CreateDescribeResourceTimelineResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceTimelineWithChan invokes the ens.DescribeResourceTimeline API asynchronously
func (client *Client) DescribeResourceTimelineWithChan(request *DescribeResourceTimelineRequest) (<-chan *DescribeResourceTimelineResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceTimelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceTimeline(request)
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

// DescribeResourceTimelineWithCallback invokes the ens.DescribeResourceTimeline API asynchronously
func (client *Client) DescribeResourceTimelineWithCallback(request *DescribeResourceTimelineRequest, callback func(response *DescribeResourceTimelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceTimelineResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceTimeline(request)
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

// DescribeResourceTimelineRequest is the request struct for api DescribeResourceTimeline
type DescribeResourceTimelineRequest struct {
	*requests.RpcRequest
	Uuid         string `position:"Query" name:"Uuid"`
	EndTime      string `position:"Query" name:"EndTime"`
	BeginTime    string `position:"Query" name:"BeginTime"`
	OpenapiCheck string `position:"Query" name:"OpenapiCheck"`
}

// DescribeResourceTimelineResponse is the response struct for api DescribeResourceTimeline
type DescribeResourceTimelineResponse struct {
	*responses.BaseResponse
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	Msg             string                `json:"Msg" xml:"Msg"`
	Desc            string                `json:"Desc" xml:"Desc"`
	AvailableEvents []AvailableEventsItem `json:"AvailableEvents" xml:"AvailableEvents"`
	BizEvents       []BizEventsItem       `json:"BizEvents" xml:"BizEvents"`
	InventoryEvents []InventoryEventsItem `json:"InventoryEvents" xml:"InventoryEvents"`
	ReserveEvents   []ReserveEventsItem   `json:"ReserveEvents" xml:"ReserveEvents"`
}

// CreateDescribeResourceTimelineRequest creates a request to invoke DescribeResourceTimeline API
func CreateDescribeResourceTimelineRequest() (request *DescribeResourceTimelineRequest) {
	request = &DescribeResourceTimelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeResourceTimeline", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeResourceTimelineResponse creates a response to parse from DescribeResourceTimeline response
func CreateDescribeResourceTimelineResponse() (response *DescribeResourceTimelineResponse) {
	response = &DescribeResourceTimelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
