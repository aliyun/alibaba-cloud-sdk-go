package quickbi_public

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

// QueryEmbeddedStatus invokes the quickbi_public.QueryEmbeddedStatus API synchronously
func (client *Client) QueryEmbeddedStatus(request *QueryEmbeddedStatusRequest) (response *QueryEmbeddedStatusResponse, err error) {
	response = CreateQueryEmbeddedStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEmbeddedStatusWithChan invokes the quickbi_public.QueryEmbeddedStatus API asynchronously
func (client *Client) QueryEmbeddedStatusWithChan(request *QueryEmbeddedStatusRequest) (<-chan *QueryEmbeddedStatusResponse, <-chan error) {
	responseChan := make(chan *QueryEmbeddedStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEmbeddedStatus(request)
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

// QueryEmbeddedStatusWithCallback invokes the quickbi_public.QueryEmbeddedStatus API asynchronously
func (client *Client) QueryEmbeddedStatusWithCallback(request *QueryEmbeddedStatusRequest, callback func(response *QueryEmbeddedStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEmbeddedStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryEmbeddedStatus(request)
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

// QueryEmbeddedStatusRequest is the request struct for api QueryEmbeddedStatus
type QueryEmbeddedStatusRequest struct {
	*requests.RpcRequest
	WorksId     string `position:"Query" name:"WorksId"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryEmbeddedStatusResponse is the response struct for api QueryEmbeddedStatus
type QueryEmbeddedStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryEmbeddedStatusRequest creates a request to invoke QueryEmbeddedStatus API
func CreateQueryEmbeddedStatusRequest() (request *QueryEmbeddedStatusRequest) {
	request = &QueryEmbeddedStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryEmbeddedStatus", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEmbeddedStatusResponse creates a response to parse from QueryEmbeddedStatus response
func CreateQueryEmbeddedStatusResponse() (response *QueryEmbeddedStatusResponse) {
	response = &QueryEmbeddedStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
