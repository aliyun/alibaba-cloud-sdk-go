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

// DescribeEnsCommodityCode invokes the ens.DescribeEnsCommodityCode API synchronously
func (client *Client) DescribeEnsCommodityCode(request *DescribeEnsCommodityCodeRequest) (response *DescribeEnsCommodityCodeResponse, err error) {
	response = CreateDescribeEnsCommodityCodeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnsCommodityCodeWithChan invokes the ens.DescribeEnsCommodityCode API asynchronously
func (client *Client) DescribeEnsCommodityCodeWithChan(request *DescribeEnsCommodityCodeRequest) (<-chan *DescribeEnsCommodityCodeResponse, <-chan error) {
	responseChan := make(chan *DescribeEnsCommodityCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnsCommodityCode(request)
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

// DescribeEnsCommodityCodeWithCallback invokes the ens.DescribeEnsCommodityCode API asynchronously
func (client *Client) DescribeEnsCommodityCodeWithCallback(request *DescribeEnsCommodityCodeRequest, callback func(response *DescribeEnsCommodityCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnsCommodityCodeResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnsCommodityCode(request)
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

// DescribeEnsCommodityCodeRequest is the request struct for api DescribeEnsCommodityCode
type DescribeEnsCommodityCodeRequest struct {
	*requests.RpcRequest
	CommodityCode string `position:"Query" name:"CommodityCode"`
}

// DescribeEnsCommodityCodeResponse is the response struct for api DescribeEnsCommodityCode
type DescribeEnsCommodityCodeResponse struct {
	*responses.BaseResponse
	RequestId         string                  `json:"RequestId" xml:"RequestId"`
	CommodityCodeInfo []CommodityCodeInfoItem `json:"CommodityCodeInfo" xml:"CommodityCodeInfo"`
}

// CreateDescribeEnsCommodityCodeRequest creates a request to invoke DescribeEnsCommodityCode API
func CreateDescribeEnsCommodityCodeRequest() (request *DescribeEnsCommodityCodeRequest) {
	request = &DescribeEnsCommodityCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEnsCommodityCode", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEnsCommodityCodeResponse creates a response to parse from DescribeEnsCommodityCode response
func CreateDescribeEnsCommodityCodeResponse() (response *DescribeEnsCommodityCodeResponse) {
	response = &DescribeEnsCommodityCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
