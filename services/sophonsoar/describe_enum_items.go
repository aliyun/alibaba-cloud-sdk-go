package sophonsoar

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

// DescribeEnumItems invokes the sophonsoar.DescribeEnumItems API synchronously
func (client *Client) DescribeEnumItems(request *DescribeEnumItemsRequest) (response *DescribeEnumItemsResponse, err error) {
	response = CreateDescribeEnumItemsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnumItemsWithChan invokes the sophonsoar.DescribeEnumItems API asynchronously
func (client *Client) DescribeEnumItemsWithChan(request *DescribeEnumItemsRequest) (<-chan *DescribeEnumItemsResponse, <-chan error) {
	responseChan := make(chan *DescribeEnumItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnumItems(request)
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

// DescribeEnumItemsWithCallback invokes the sophonsoar.DescribeEnumItems API asynchronously
func (client *Client) DescribeEnumItemsWithCallback(request *DescribeEnumItemsRequest, callback func(response *DescribeEnumItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnumItemsResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnumItems(request)
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

// DescribeEnumItemsRequest is the request struct for api DescribeEnumItems
type DescribeEnumItemsRequest struct {
	*requests.RpcRequest
	RoleFor  string `position:"Query" name:"RoleFor"`
	EnumType string `position:"Query" name:"EnumType"`
	RoleType string `position:"Query" name:"RoleType"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeEnumItemsResponse is the response struct for api DescribeEnumItems
type DescribeEnumItemsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeEnumItemsRequest creates a request to invoke DescribeEnumItems API
func CreateDescribeEnumItemsRequest() (request *DescribeEnumItemsRequest) {
	request = &DescribeEnumItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribeEnumItems", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeEnumItemsResponse creates a response to parse from DescribeEnumItems response
func CreateDescribeEnumItemsResponse() (response *DescribeEnumItemsResponse) {
	response = &DescribeEnumItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
