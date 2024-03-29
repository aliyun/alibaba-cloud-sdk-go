package ehpc

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

// DescribeJob invokes the ehpc.DescribeJob API synchronously
func (client *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
	response = CreateDescribeJobResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeJobWithChan invokes the ehpc.DescribeJob API asynchronously
func (client *Client) DescribeJobWithChan(request *DescribeJobRequest) (<-chan *DescribeJobResponse, <-chan error) {
	responseChan := make(chan *DescribeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeJob(request)
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

// DescribeJobWithCallback invokes the ehpc.DescribeJob API asynchronously
func (client *Client) DescribeJobWithCallback(request *DescribeJobRequest, callback func(response *DescribeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeJobResponse
		var err error
		defer close(result)
		response, err = client.DescribeJob(request)
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

// DescribeJobRequest is the request struct for api DescribeJob
type DescribeJobRequest struct {
	*requests.RpcRequest
	ClusterId string           `position:"Query" name:"ClusterId"`
	JobId     string           `position:"Query" name:"JobId"`
	Async     requests.Boolean `position:"Query" name:"Async"`
}

// DescribeJobResponse is the response struct for api DescribeJob
type DescribeJobResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Message   Message `json:"Message" xml:"Message"`
}

// CreateDescribeJobRequest creates a request to invoke DescribeJob API
func CreateDescribeJobRequest() (request *DescribeJobRequest) {
	request = &DescribeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DescribeJob", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeJobResponse creates a response to parse from DescribeJob response
func CreateDescribeJobResponse() (response *DescribeJobResponse) {
	response = &DescribeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
