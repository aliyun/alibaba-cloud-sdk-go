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

// DescribeComponentAssetForm invokes the sophonsoar.DescribeComponentAssetForm API synchronously
func (client *Client) DescribeComponentAssetForm(request *DescribeComponentAssetFormRequest) (response *DescribeComponentAssetFormResponse, err error) {
	response = CreateDescribeComponentAssetFormResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeComponentAssetFormWithChan invokes the sophonsoar.DescribeComponentAssetForm API asynchronously
func (client *Client) DescribeComponentAssetFormWithChan(request *DescribeComponentAssetFormRequest) (<-chan *DescribeComponentAssetFormResponse, <-chan error) {
	responseChan := make(chan *DescribeComponentAssetFormResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeComponentAssetForm(request)
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

// DescribeComponentAssetFormWithCallback invokes the sophonsoar.DescribeComponentAssetForm API asynchronously
func (client *Client) DescribeComponentAssetFormWithCallback(request *DescribeComponentAssetFormRequest, callback func(response *DescribeComponentAssetFormResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeComponentAssetFormResponse
		var err error
		defer close(result)
		response, err = client.DescribeComponentAssetForm(request)
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

// DescribeComponentAssetFormRequest is the request struct for api DescribeComponentAssetForm
type DescribeComponentAssetFormRequest struct {
	*requests.RpcRequest
	RoleFor       string `position:"Query" name:"RoleFor"`
	ComponentName string `position:"Query" name:"ComponentName"`
	RoleType      string `position:"Query" name:"RoleType"`
	Lang          string `position:"Query" name:"Lang"`
}

// DescribeComponentAssetFormResponse is the response struct for api DescribeComponentAssetForm
type DescribeComponentAssetFormResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	ComponentAssetForm string `json:"ComponentAssetForm" xml:"ComponentAssetForm"`
}

// CreateDescribeComponentAssetFormRequest creates a request to invoke DescribeComponentAssetForm API
func CreateDescribeComponentAssetFormRequest() (request *DescribeComponentAssetFormRequest) {
	request = &DescribeComponentAssetFormRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribeComponentAssetForm", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeComponentAssetFormResponse creates a response to parse from DescribeComponentAssetForm response
func CreateDescribeComponentAssetFormResponse() (response *DescribeComponentAssetFormResponse) {
	response = &DescribeComponentAssetFormResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
