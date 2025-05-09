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

// DescribeSoarRecordActionOutputList invokes the sophonsoar.DescribeSoarRecordActionOutputList API synchronously
func (client *Client) DescribeSoarRecordActionOutputList(request *DescribeSoarRecordActionOutputListRequest) (response *DescribeSoarRecordActionOutputListResponse, err error) {
	response = CreateDescribeSoarRecordActionOutputListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSoarRecordActionOutputListWithChan invokes the sophonsoar.DescribeSoarRecordActionOutputList API asynchronously
func (client *Client) DescribeSoarRecordActionOutputListWithChan(request *DescribeSoarRecordActionOutputListRequest) (<-chan *DescribeSoarRecordActionOutputListResponse, <-chan error) {
	responseChan := make(chan *DescribeSoarRecordActionOutputListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSoarRecordActionOutputList(request)
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

// DescribeSoarRecordActionOutputListWithCallback invokes the sophonsoar.DescribeSoarRecordActionOutputList API asynchronously
func (client *Client) DescribeSoarRecordActionOutputListWithCallback(request *DescribeSoarRecordActionOutputListRequest, callback func(response *DescribeSoarRecordActionOutputListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSoarRecordActionOutputListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSoarRecordActionOutputList(request)
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

// DescribeSoarRecordActionOutputListRequest is the request struct for api DescribeSoarRecordActionOutputList
type DescribeSoarRecordActionOutputListRequest struct {
	*requests.RpcRequest
	ActionUuid string           `position:"Query" name:"ActionUuid"`
	RoleFor    string           `position:"Query" name:"RoleFor"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	RoleType   string           `position:"Query" name:"RoleType"`
	QueryValue string           `position:"Query" name:"QueryValue"`
	Lang       string           `position:"Query" name:"Lang"`
}

// DescribeSoarRecordActionOutputListResponse is the response struct for api DescribeSoarRecordActionOutputList
type DescribeSoarRecordActionOutputListResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ActionOutputs string `json:"ActionOutputs" xml:"ActionOutputs"`
	PageNumber    int    `json:"PageNumber" xml:"PageNumber"`
	PageSize      int    `json:"PageSize" xml:"PageSize"`
	TotalCount    int    `json:"TotalCount" xml:"TotalCount"`
}

// CreateDescribeSoarRecordActionOutputListRequest creates a request to invoke DescribeSoarRecordActionOutputList API
func CreateDescribeSoarRecordActionOutputListRequest() (request *DescribeSoarRecordActionOutputListRequest) {
	request = &DescribeSoarRecordActionOutputListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "DescribeSoarRecordActionOutputList", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeSoarRecordActionOutputListResponse creates a response to parse from DescribeSoarRecordActionOutputList response
func CreateDescribeSoarRecordActionOutputListResponse() (response *DescribeSoarRecordActionOutputListResponse) {
	response = &DescribeSoarRecordActionOutputListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
