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

package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryAliases invokes the push.QueryAliases API synchronously
// api document: https://help.aliyun.com/api/push/queryaliases.html
func (client *Client) QueryAliases(request *QueryAliasesRequest) (response *QueryAliasesResponse, err error) {
	response = CreateQueryAliasesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAliasesWithChan invokes the push.QueryAliases API asynchronously
// api document: https://help.aliyun.com/api/push/queryaliases.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAliasesWithChan(request *QueryAliasesRequest) (<-chan *QueryAliasesResponse, <-chan error) {
	responseChan := make(chan *QueryAliasesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAliases(request)
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

// QueryAliasesWithCallback invokes the push.QueryAliases API asynchronously
// api document: https://help.aliyun.com/api/push/queryaliases.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAliasesWithCallback(request *QueryAliasesRequest, callback func(response *QueryAliasesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAliasesResponse
		var err error
		defer close(result)
		response, err = client.QueryAliases(request)
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

// QueryAliasesRequest is the request struct for api QueryAliases
type QueryAliasesRequest struct {
	*requests.RpcRequest
	AccessKeyId string           `position:"Query" name:"AccessKeyId"`
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
	DeviceId    string           `position:"Query" name:"DeviceId"`
}

// QueryAliasesResponse is the response struct for api QueryAliases
type QueryAliasesResponse struct {
	*responses.BaseResponse
	RequestId  string                  `json:"RequestId" xml:"RequestId"`
	AliasInfos QueryAliasesAliasInfos0 `json:"AliasInfos" xml:"AliasInfos"`
}

type QueryAliasesAliasInfos0 struct {
	AliasInfo []QueryAliasesAliasInfo1 `json:"AliasInfo" xml:"AliasInfo"`
}

type QueryAliasesAliasInfo1 struct {
	AliasName string `json:"AliasName" xml:"AliasName"`
}

// CreateQueryAliasesRequest creates a request to invoke QueryAliases API
func CreateQueryAliasesRequest() (request *QueryAliasesRequest) {
	request = &QueryAliasesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryAliases", "push", "openAPI")
	return
}

// CreateQueryAliasesResponse creates a response to parse from QueryAliases response
func CreateQueryAliasesResponse() (response *QueryAliasesResponse) {
	response = &QueryAliasesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
