package oceanbasepro

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

// DescribeProject invokes the oceanbasepro.DescribeProject API synchronously
func (client *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
	response = CreateDescribeProjectResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectWithChan invokes the oceanbasepro.DescribeProject API asynchronously
func (client *Client) DescribeProjectWithChan(request *DescribeProjectRequest) (<-chan *DescribeProjectResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProject(request)
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

// DescribeProjectWithCallback invokes the oceanbasepro.DescribeProject API asynchronously
func (client *Client) DescribeProjectWithCallback(request *DescribeProjectRequest, callback func(response *DescribeProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectResponse
		var err error
		defer close(result)
		response, err = client.DescribeProject(request)
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

// DescribeProjectRequest is the request struct for api DescribeProject
type DescribeProjectRequest struct {
	*requests.RpcRequest
	Id string `position:"Body" name:"Id"`
}

// DescribeProjectResponse is the response struct for api DescribeProject
type DescribeProjectResponse struct {
	*responses.BaseResponse
}

// CreateDescribeProjectRequest creates a request to invoke DescribeProject API
func CreateDescribeProjectRequest() (request *DescribeProjectRequest) {
	request = &DescribeProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeProject", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProjectResponse creates a response to parse from DescribeProject response
func CreateDescribeProjectResponse() (response *DescribeProjectResponse) {
	response = &DescribeProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
