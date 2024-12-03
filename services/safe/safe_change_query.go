package safe

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

// SafeChangeQuery invokes the safe.SafeChangeQuery API synchronously
func (client *Client) SafeChangeQuery(request *SafeChangeQueryRequest) (response *SafeChangeQueryResponse, err error) {
	response = CreateSafeChangeQueryResponse()
	err = client.DoAction(request, response)
	return
}

// SafeChangeQueryWithChan invokes the safe.SafeChangeQuery API asynchronously
func (client *Client) SafeChangeQueryWithChan(request *SafeChangeQueryRequest) (<-chan *SafeChangeQueryResponse, <-chan error) {
	responseChan := make(chan *SafeChangeQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SafeChangeQuery(request)
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

// SafeChangeQueryWithCallback invokes the safe.SafeChangeQuery API asynchronously
func (client *Client) SafeChangeQueryWithCallback(request *SafeChangeQueryRequest, callback func(response *SafeChangeQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SafeChangeQueryResponse
		var err error
		defer close(result)
		response, err = client.SafeChangeQuery(request)
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

// SafeChangeQueryRequest is the request struct for api SafeChangeQuery
type SafeChangeQueryRequest struct {
	*requests.RpcRequest
	NeedValidate  requests.Boolean `position:"Body" name:"NeedValidate"`
	AuthKey       string           `position:"Body" name:"AuthKey"`
	ReqTimestamp  requests.Integer `position:"Body" name:"ReqTimestamp"`
	ReturnType    requests.Boolean `position:"Query" name:"ReturnType"`
	SourceOrderId string           `position:"Body" name:"SourceOrderId"`
	AuthSign      string           `position:"Body" name:"AuthSign"`
	QueryType     string           `position:"Body" name:"QueryType"`
}

// SafeChangeQueryResponse is the response struct for api SafeChangeQuery
type SafeChangeQueryResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSafeChangeQueryRequest creates a request to invoke SafeChangeQuery API
func CreateSafeChangeQueryRequest() (request *SafeChangeQueryRequest) {
	request = &SafeChangeQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Safe", "2022-01-17", "SafeChangeQuery", "", "")
	request.Method = requests.POST
	return
}

// CreateSafeChangeQueryResponse creates a response to parse from SafeChangeQuery response
func CreateSafeChangeQueryResponse() (response *SafeChangeQueryResponse) {
	response = &SafeChangeQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
