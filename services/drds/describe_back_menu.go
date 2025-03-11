package drds

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

// DescribeBackMenu invokes the drds.DescribeBackMenu API synchronously
func (client *Client) DescribeBackMenu(request *DescribeBackMenuRequest) (response *DescribeBackMenuResponse, err error) {
	response = CreateDescribeBackMenuResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackMenuWithChan invokes the drds.DescribeBackMenu API asynchronously
func (client *Client) DescribeBackMenuWithChan(request *DescribeBackMenuRequest) (<-chan *DescribeBackMenuResponse, <-chan error) {
	responseChan := make(chan *DescribeBackMenuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackMenu(request)
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

// DescribeBackMenuWithCallback invokes the drds.DescribeBackMenu API asynchronously
func (client *Client) DescribeBackMenuWithCallback(request *DescribeBackMenuRequest, callback func(response *DescribeBackMenuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackMenuResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackMenu(request)
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

// DescribeBackMenuRequest is the request struct for api DescribeBackMenu
type DescribeBackMenuRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeBackMenuResponse is the response struct for api DescribeBackMenu
type DescribeBackMenuResponse struct {
	*responses.BaseResponse
	Success   bool                   `json:"Success" xml:"Success"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	List      ListInDescribeBackMenu `json:"List" xml:"List"`
}

// CreateDescribeBackMenuRequest creates a request to invoke DescribeBackMenu API
func CreateDescribeBackMenuRequest() (request *DescribeBackMenuRequest) {
	request = &DescribeBackMenuRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeBackMenu", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackMenuResponse creates a response to parse from DescribeBackMenu response
func CreateDescribeBackMenuResponse() (response *DescribeBackMenuResponse) {
	response = &DescribeBackMenuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
