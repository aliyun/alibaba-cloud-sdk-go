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

// ListComponentIndices invokes the elasticsearch.ListComponentIndices API synchronously
func (client *Client) ListComponentIndices(request *ListComponentIndicesRequest) (response *ListComponentIndicesResponse, err error) {
	response = CreateListComponentIndicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListComponentIndicesWithChan invokes the elasticsearch.ListComponentIndices API asynchronously
func (client *Client) ListComponentIndicesWithChan(request *ListComponentIndicesRequest) (<-chan *ListComponentIndicesResponse, <-chan error) {
	responseChan := make(chan *ListComponentIndicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListComponentIndices(request)
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

// ListComponentIndicesWithCallback invokes the elasticsearch.ListComponentIndices API asynchronously
func (client *Client) ListComponentIndicesWithCallback(request *ListComponentIndicesRequest, callback func(response *ListComponentIndicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListComponentIndicesResponse
		var err error
		defer close(result)
		response, err = client.ListComponentIndices(request)
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

// ListComponentIndicesRequest is the request struct for api ListComponentIndices
type ListComponentIndicesRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Size       requests.Integer `position:"Query" name:"size"`
	Name       string           `position:"Query" name:"name"`
	Page       requests.Integer `position:"Query" name:"page"`
}

// ListComponentIndicesResponse is the response struct for api ListComponentIndices
type ListComponentIndicesResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Headers   Headers                            `json:"Headers" xml:"Headers"`
	Result    []ResultItemInListComponentIndices `json:"Result" xml:"Result"`
}

// CreateListComponentIndicesRequest creates a request to invoke ListComponentIndices API
func CreateListComponentIndicesRequest() (request *ListComponentIndicesRequest) {
	request = &ListComponentIndicesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListComponentIndices", "/openapi/instances/[InstanceId]/component-index", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListComponentIndicesResponse creates a response to parse from ListComponentIndices response
func CreateListComponentIndicesResponse() (response *ListComponentIndicesResponse) {
	response = &ListComponentIndicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
