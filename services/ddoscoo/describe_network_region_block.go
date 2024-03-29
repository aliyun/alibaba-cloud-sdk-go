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

// DescribeNetworkRegionBlock invokes the ddoscoo.DescribeNetworkRegionBlock API synchronously
func (client *Client) DescribeNetworkRegionBlock(request *DescribeNetworkRegionBlockRequest) (response *DescribeNetworkRegionBlockResponse, err error) {
	response = CreateDescribeNetworkRegionBlockResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkRegionBlockWithChan invokes the ddoscoo.DescribeNetworkRegionBlock API asynchronously
func (client *Client) DescribeNetworkRegionBlockWithChan(request *DescribeNetworkRegionBlockRequest) (<-chan *DescribeNetworkRegionBlockResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkRegionBlockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkRegionBlock(request)
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

// DescribeNetworkRegionBlockWithCallback invokes the ddoscoo.DescribeNetworkRegionBlock API asynchronously
func (client *Client) DescribeNetworkRegionBlockWithCallback(request *DescribeNetworkRegionBlockRequest, callback func(response *DescribeNetworkRegionBlockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkRegionBlockResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkRegionBlock(request)
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

// DescribeNetworkRegionBlockRequest is the request struct for api DescribeNetworkRegionBlock
type DescribeNetworkRegionBlockRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
}

// DescribeNetworkRegionBlockResponse is the response struct for api DescribeNetworkRegionBlock
type DescribeNetworkRegionBlockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Config    Config `json:"Config" xml:"Config"`
}

// CreateDescribeNetworkRegionBlockRequest creates a request to invoke DescribeNetworkRegionBlock API
func CreateDescribeNetworkRegionBlockRequest() (request *DescribeNetworkRegionBlockRequest) {
	request = &DescribeNetworkRegionBlockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeNetworkRegionBlock", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkRegionBlockResponse creates a response to parse from DescribeNetworkRegionBlock response
func CreateDescribeNetworkRegionBlockResponse() (response *DescribeNetworkRegionBlockResponse) {
	response = &DescribeNetworkRegionBlockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
