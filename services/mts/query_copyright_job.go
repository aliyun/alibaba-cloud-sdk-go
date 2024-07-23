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

// QueryCopyrightJob invokes the mts.QueryCopyrightJob API synchronously
func (client *Client) QueryCopyrightJob(request *QueryCopyrightJobRequest) (response *QueryCopyrightJobResponse, err error) {
	response = CreateQueryCopyrightJobResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCopyrightJobWithChan invokes the mts.QueryCopyrightJob API asynchronously
func (client *Client) QueryCopyrightJobWithChan(request *QueryCopyrightJobRequest) (<-chan *QueryCopyrightJobResponse, <-chan error) {
	responseChan := make(chan *QueryCopyrightJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCopyrightJob(request)
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

// QueryCopyrightJobWithCallback invokes the mts.QueryCopyrightJob API asynchronously
func (client *Client) QueryCopyrightJobWithCallback(request *QueryCopyrightJobRequest, callback func(response *QueryCopyrightJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCopyrightJobResponse
		var err error
		defer close(result)
		response, err = client.QueryCopyrightJob(request)
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

// QueryCopyrightJobRequest is the request struct for api QueryCopyrightJob
type QueryCopyrightJobRequest struct {
	*requests.RpcRequest
	Level           requests.Integer `position:"Query" name:"Level"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	CreateTimeEnd   requests.Integer `position:"Query" name:"CreateTimeEnd"`
	JobId           string           `position:"Query" name:"JobId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CreateTimeStart requests.Integer `position:"Query" name:"CreateTimeStart"`
}

// QueryCopyrightJobResponse is the response struct for api QueryCopyrightJob
type QueryCopyrightJobResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Message    string     `json:"Message" xml:"Message"`
	StatusCode int64      `json:"StatusCode" xml:"StatusCode"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateQueryCopyrightJobRequest creates a request to invoke QueryCopyrightJob API
func CreateQueryCopyrightJobRequest() (request *QueryCopyrightJobRequest) {
	request = &QueryCopyrightJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryCopyrightJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryCopyrightJobResponse creates a response to parse from QueryCopyrightJob response
func CreateQueryCopyrightJobResponse() (response *QueryCopyrightJobResponse) {
	response = &QueryCopyrightJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
