package mts

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

// QueryMediaCensorJobList invokes the mts.QueryMediaCensorJobList API synchronously
func (client *Client) QueryMediaCensorJobList(request *QueryMediaCensorJobListRequest) (response *QueryMediaCensorJobListResponse, err error) {
	response = CreateQueryMediaCensorJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMediaCensorJobListWithChan invokes the mts.QueryMediaCensorJobList API asynchronously
func (client *Client) QueryMediaCensorJobListWithChan(request *QueryMediaCensorJobListRequest) (<-chan *QueryMediaCensorJobListResponse, <-chan error) {
	responseChan := make(chan *QueryMediaCensorJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaCensorJobList(request)
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

// QueryMediaCensorJobListWithCallback invokes the mts.QueryMediaCensorJobList API asynchronously
func (client *Client) QueryMediaCensorJobListWithCallback(request *QueryMediaCensorJobListRequest, callback func(response *QueryMediaCensorJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaCensorJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaCensorJobList(request)
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

// QueryMediaCensorJobListRequest is the request struct for api QueryMediaCensorJobList
type QueryMediaCensorJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextPageToken              string           `position:"Query" name:"NextPageToken"`
	StartOfJobCreatedTimeRange string           `position:"Query" name:"StartOfJobCreatedTimeRange"`
	State                      string           `position:"Query" name:"State"`
	EndOfJobCreatedTimeRange   string           `position:"Query" name:"EndOfJobCreatedTimeRange"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize            requests.Integer `position:"Query" name:"MaximumPageSize"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId                 string           `position:"Query" name:"PipelineId"`
	JobIds                     string           `position:"Query" name:"JobIds"`
}

// QueryMediaCensorJobListResponse is the response struct for api QueryMediaCensorJobList
type QueryMediaCensorJobListResponse struct {
	*responses.BaseResponse
	RequestId          string                               `json:"RequestId" xml:"RequestId"`
	NextPageToken      string                               `json:"NextPageToken" xml:"NextPageToken"`
	NonExistIds        NonExistIdsInQueryMediaCensorJobList `json:"NonExistIds" xml:"NonExistIds"`
	MediaCensorJobList MediaCensorJobList                   `json:"MediaCensorJobList" xml:"MediaCensorJobList"`
}

// CreateQueryMediaCensorJobListRequest creates a request to invoke QueryMediaCensorJobList API
func CreateQueryMediaCensorJobListRequest() (request *QueryMediaCensorJobListRequest) {
	request = &QueryMediaCensorJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaCensorJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMediaCensorJobListResponse creates a response to parse from QueryMediaCensorJobList response
func CreateQueryMediaCensorJobListResponse() (response *QueryMediaCensorJobListResponse) {
	response = &QueryMediaCensorJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
