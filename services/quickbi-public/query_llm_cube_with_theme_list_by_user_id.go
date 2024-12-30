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

// QueryLlmCubeWithThemeListByUserId invokes the quickbi_public.QueryLlmCubeWithThemeListByUserId API synchronously
func (client *Client) QueryLlmCubeWithThemeListByUserId(request *QueryLlmCubeWithThemeListByUserIdRequest) (response *QueryLlmCubeWithThemeListByUserIdResponse, err error) {
	response = CreateQueryLlmCubeWithThemeListByUserIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryLlmCubeWithThemeListByUserIdWithChan invokes the quickbi_public.QueryLlmCubeWithThemeListByUserId API asynchronously
func (client *Client) QueryLlmCubeWithThemeListByUserIdWithChan(request *QueryLlmCubeWithThemeListByUserIdRequest) (<-chan *QueryLlmCubeWithThemeListByUserIdResponse, <-chan error) {
	responseChan := make(chan *QueryLlmCubeWithThemeListByUserIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryLlmCubeWithThemeListByUserId(request)
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

// QueryLlmCubeWithThemeListByUserIdWithCallback invokes the quickbi_public.QueryLlmCubeWithThemeListByUserId API asynchronously
func (client *Client) QueryLlmCubeWithThemeListByUserIdWithCallback(request *QueryLlmCubeWithThemeListByUserIdRequest, callback func(response *QueryLlmCubeWithThemeListByUserIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryLlmCubeWithThemeListByUserIdResponse
		var err error
		defer close(result)
		response, err = client.QueryLlmCubeWithThemeListByUserId(request)
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

// QueryLlmCubeWithThemeListByUserIdRequest is the request struct for api QueryLlmCubeWithThemeListByUserId
type QueryLlmCubeWithThemeListByUserIdRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
}

// QueryLlmCubeWithThemeListByUserIdResponse is the response struct for api QueryLlmCubeWithThemeListByUserId
type QueryLlmCubeWithThemeListByUserIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryLlmCubeWithThemeListByUserIdRequest creates a request to invoke QueryLlmCubeWithThemeListByUserId API
func CreateQueryLlmCubeWithThemeListByUserIdRequest() (request *QueryLlmCubeWithThemeListByUserIdRequest) {
	request = &QueryLlmCubeWithThemeListByUserIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryLlmCubeWithThemeListByUserId", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryLlmCubeWithThemeListByUserIdResponse creates a response to parse from QueryLlmCubeWithThemeListByUserId response
func CreateQueryLlmCubeWithThemeListByUserIdResponse() (response *QueryLlmCubeWithThemeListByUserIdResponse) {
	response = &QueryLlmCubeWithThemeListByUserIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
