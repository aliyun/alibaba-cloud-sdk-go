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

// ListLogstash invokes the elasticsearch.ListLogstash API synchronously
func (client *Client) ListLogstash(request *ListLogstashRequest) (response *ListLogstashResponse, err error) {
	response = CreateListLogstashResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogstashWithChan invokes the elasticsearch.ListLogstash API asynchronously
func (client *Client) ListLogstashWithChan(request *ListLogstashRequest) (<-chan *ListLogstashResponse, <-chan error) {
	responseChan := make(chan *ListLogstashResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogstash(request)
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

// ListLogstashWithCallback invokes the elasticsearch.ListLogstash API asynchronously
func (client *Client) ListLogstashWithCallback(request *ListLogstashRequest, callback func(response *ListLogstashResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogstashResponse
		var err error
		defer close(result)
		response, err = client.ListLogstash(request)
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

// ListLogstashRequest is the request struct for api ListLogstash
type ListLogstashRequest struct {
	*requests.RoaRequest
	ResourceGroupId string           `position:"Query" name:"resourceGroupId"`
	InstanceId      string           `position:"Query" name:"instanceId"`
	Size            requests.Integer `position:"Query" name:"size"`
	Description     string           `position:"Query" name:"description"`
	Page            requests.Integer `position:"Query" name:"page"`
	OwnerId         string           `position:"Query" name:"ownerId"`
	Version         string           `position:"Query" name:"version"`
	Tags            string           `position:"Query" name:"tags"`
}

// ListLogstashResponse is the response struct for api ListLogstash
type ListLogstashResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Headers   HeadersInListLogstash    `json:"Headers" xml:"Headers"`
	Result    []InstanceInListLogstash `json:"Result" xml:"Result"`
}

// CreateListLogstashRequest creates a request to invoke ListLogstash API
func CreateListLogstashRequest() (request *ListLogstashRequest) {
	request = &ListLogstashRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListLogstash", "/openapi/logstashes", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListLogstashResponse creates a response to parse from ListLogstash response
func CreateListLogstashResponse() (response *ListLogstashResponse) {
	response = &ListLogstashResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
