package elasticsearch

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

// DescribeComponentIndex invokes the elasticsearch.DescribeComponentIndex API synchronously
func (client *Client) DescribeComponentIndex(request *DescribeComponentIndexRequest) (response *DescribeComponentIndexResponse, err error) {
	response = CreateDescribeComponentIndexResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeComponentIndexWithChan invokes the elasticsearch.DescribeComponentIndex API asynchronously
func (client *Client) DescribeComponentIndexWithChan(request *DescribeComponentIndexRequest) (<-chan *DescribeComponentIndexResponse, <-chan error) {
	responseChan := make(chan *DescribeComponentIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeComponentIndex(request)
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

// DescribeComponentIndexWithCallback invokes the elasticsearch.DescribeComponentIndex API asynchronously
func (client *Client) DescribeComponentIndexWithCallback(request *DescribeComponentIndexRequest, callback func(response *DescribeComponentIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeComponentIndexResponse
		var err error
		defer close(result)
		response, err = client.DescribeComponentIndex(request)
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

// DescribeComponentIndexRequest is the request struct for api DescribeComponentIndex
type DescribeComponentIndexRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Name       string `position:"Path" name:"name"`
}

// DescribeComponentIndexResponse is the response struct for api DescribeComponentIndex
type DescribeComponentIndexResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeComponentIndexRequest creates a request to invoke DescribeComponentIndex API
func CreateDescribeComponentIndexRequest() (request *DescribeComponentIndexRequest) {
	request = &DescribeComponentIndexRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DescribeComponentIndex", "/openapi/instances/[InstanceId]/component-index/[name]", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeComponentIndexResponse creates a response to parse from DescribeComponentIndex response
func CreateDescribeComponentIndexResponse() (response *DescribeComponentIndexResponse) {
	response = &DescribeComponentIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
