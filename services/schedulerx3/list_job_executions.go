package schedulerx3

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

// ListJobExecutions invokes the schedulerx3.ListJobExecutions API synchronously
func (client *Client) ListJobExecutions(request *ListJobExecutionsRequest) (response *ListJobExecutionsResponse, err error) {
	response = CreateListJobExecutionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobExecutionsWithChan invokes the schedulerx3.ListJobExecutions API asynchronously
func (client *Client) ListJobExecutionsWithChan(request *ListJobExecutionsRequest) (<-chan *ListJobExecutionsResponse, <-chan error) {
	responseChan := make(chan *ListJobExecutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobExecutions(request)
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

// ListJobExecutionsWithCallback invokes the schedulerx3.ListJobExecutions API asynchronously
func (client *Client) ListJobExecutionsWithCallback(request *ListJobExecutionsRequest, callback func(response *ListJobExecutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobExecutionsResponse
		var err error
		defer close(result)
		response, err = client.ListJobExecutions(request)
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

// ListJobExecutionsRequest is the request struct for api ListJobExecutions
type ListJobExecutionsRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	StartTime      string           `position:"Query" name:"StartTime"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	JobId          requests.Integer `position:"Query" name:"JobId"`
	AppName        string           `position:"Query" name:"AppName"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	JobName        string           `position:"Query" name:"JobName"`
	JobExecutionId string           `position:"Query" name:"JobExecutionId"`
	EndTime        string           `position:"Query" name:"EndTime"`
	ClusterId      string           `position:"Query" name:"ClusterId"`
	Status         requests.Integer `position:"Query" name:"Status"`
}

// ListJobExecutionsResponse is the response struct for api ListJobExecutions
type ListJobExecutionsResponse struct {
	*responses.BaseResponse
	Code      int                     `json:"Code" xml:"Code"`
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Message   string                  `json:"Message" xml:"Message"`
	Success   bool                    `json:"Success" xml:"Success"`
	Data      DataInListJobExecutions `json:"Data" xml:"Data"`
}

// CreateListJobExecutionsRequest creates a request to invoke ListJobExecutions API
func CreateListJobExecutionsRequest() (request *ListJobExecutionsRequest) {
	request = &ListJobExecutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SchedulerX3", "2024-06-24", "ListJobExecutions", "", "")
	request.Method = requests.GET
	return
}

// CreateListJobExecutionsResponse creates a response to parse from ListJobExecutions response
func CreateListJobExecutionsResponse() (response *ListJobExecutionsResponse) {
	response = &ListJobExecutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
