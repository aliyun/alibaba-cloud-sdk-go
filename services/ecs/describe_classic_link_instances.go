package ecs

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

// DescribeClassicLinkInstances invokes the ecs.DescribeClassicLinkInstances API synchronously
// api document: https://help.aliyun.com/api/ecs/describeclassiclinkinstances.html
func (client *Client) DescribeClassicLinkInstances(request *DescribeClassicLinkInstancesRequest) (response *DescribeClassicLinkInstancesResponse, err error) {
	response = CreateDescribeClassicLinkInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClassicLinkInstancesWithChan invokes the ecs.DescribeClassicLinkInstances API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeclassiclinkinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClassicLinkInstancesWithChan(request *DescribeClassicLinkInstancesRequest) (<-chan *DescribeClassicLinkInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeClassicLinkInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClassicLinkInstances(request)
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

// DescribeClassicLinkInstancesWithCallback invokes the ecs.DescribeClassicLinkInstances API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeclassiclinkinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClassicLinkInstancesWithCallback(request *DescribeClassicLinkInstancesRequest, callback func(response *DescribeClassicLinkInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClassicLinkInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeClassicLinkInstances(request)
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

// DescribeClassicLinkInstancesRequest is the request struct for api DescribeClassicLinkInstances
type DescribeClassicLinkInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
	PageSize             string           `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
}

// DescribeClassicLinkInstancesResponse is the response struct for api DescribeClassicLinkInstances
type DescribeClassicLinkInstancesResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Links      Links  `json:"Links" xml:"Links"`
}

// CreateDescribeClassicLinkInstancesRequest creates a request to invoke DescribeClassicLinkInstances API
func CreateDescribeClassicLinkInstancesRequest() (request *DescribeClassicLinkInstancesRequest) {
	request = &DescribeClassicLinkInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeClassicLinkInstances", "ecs", "openAPI")
	return
}

// CreateDescribeClassicLinkInstancesResponse creates a response to parse from DescribeClassicLinkInstances response
func CreateDescribeClassicLinkInstancesResponse() (response *DescribeClassicLinkInstancesResponse) {
	response = &DescribeClassicLinkInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
