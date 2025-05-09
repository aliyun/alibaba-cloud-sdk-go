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

// DescribeField invokes the sophonsoar.DescribeField API synchronously
func (client *Client) DescribeField(request *DescribeFieldRequest) (response *DescribeFieldResponse, err error) {
	response = CreateDescribeFieldResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFieldWithChan invokes the sophonsoar.DescribeField API asynchronously
func (client *Client) DescribeFieldWithChan(request *DescribeFieldRequest) (<-chan *DescribeFieldResponse, <-chan error) {
	responseChan := make(chan *DescribeFieldResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeField(request)
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

// DescribeFieldWithCallback invokes the sophonsoar.DescribeField API asynchronously
func (client *Client) DescribeFieldWithCallback(request *DescribeFieldRequest, callback func(response *DescribeFieldResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFieldResponse
		var err error
		defer close(result)
		response, err = client.DescribeField(request)
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

// DescribeFieldRequest is the request struct for api DescribeField
type DescribeFieldRequest struct {
	*requests.RpcRequest
	QueryKey string `position:"Query" name:"QueryKey"`
	RoleFor  string `position:"Query" name:"RoleFor"`
	RoleType string `position:"Query" name:"RoleType"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeFieldResponse is the response struct for api DescribeField
type DescribeFieldResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Name      string `json:"Name" xml:"Name"`
	Fields    string `json:"Fields" xml:"Fields"`
	FieldMode string `json:"FieldMode" xml:"FieldMode"`
}

// CreateDescribeFieldRequest creates a request to invoke DescribeField API
func CreateDescribeFieldRequest() (request *DescribeFieldRequest) {
	request = &DescribeFieldRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribeField", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeFieldResponse creates a response to parse from DescribeField response
func CreateDescribeFieldResponse() (response *DescribeFieldResponse) {
	response = &DescribeFieldResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
