package domain

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

// GetIntlDomainDownloadUrl invokes the domain.GetIntlDomainDownloadUrl API synchronously
func (client *Client) GetIntlDomainDownloadUrl(request *GetIntlDomainDownloadUrlRequest) (response *GetIntlDomainDownloadUrlResponse, err error) {
	response = CreateGetIntlDomainDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetIntlDomainDownloadUrlWithChan invokes the domain.GetIntlDomainDownloadUrl API asynchronously
func (client *Client) GetIntlDomainDownloadUrlWithChan(request *GetIntlDomainDownloadUrlRequest) (<-chan *GetIntlDomainDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *GetIntlDomainDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIntlDomainDownloadUrl(request)
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

// GetIntlDomainDownloadUrlWithCallback invokes the domain.GetIntlDomainDownloadUrl API asynchronously
func (client *Client) GetIntlDomainDownloadUrlWithCallback(request *GetIntlDomainDownloadUrlRequest, callback func(response *GetIntlDomainDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIntlDomainDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetIntlDomainDownloadUrl(request)
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

// GetIntlDomainDownloadUrlRequest is the request struct for api GetIntlDomainDownloadUrl
type GetIntlDomainDownloadUrlRequest struct {
	*requests.RpcRequest
}

// GetIntlDomainDownloadUrlResponse is the response struct for api GetIntlDomainDownloadUrl
type GetIntlDomainDownloadUrlResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string   `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string   `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrorMsg       string   `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode      string   `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool     `json:"Success" xml:"Success"`
	AllowRetry     bool     `json:"AllowRetry" xml:"AllowRetry"`
	AppName        string   `json:"AppName" xml:"AppName"`
	Url            string   `json:"Url" xml:"Url"`
	ErrorArgs      []string `json:"ErrorArgs" xml:"ErrorArgs"`
}

// CreateGetIntlDomainDownloadUrlRequest creates a request to invoke GetIntlDomainDownloadUrl API
func CreateGetIntlDomainDownloadUrlRequest() (request *GetIntlDomainDownloadUrlRequest) {
	request = &GetIntlDomainDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "GetIntlDomainDownloadUrl", "", "")
	request.Method = requests.POST
	return
}

// CreateGetIntlDomainDownloadUrlResponse creates a response to parse from GetIntlDomainDownloadUrl response
func CreateGetIntlDomainDownloadUrlResponse() (response *GetIntlDomainDownloadUrlResponse) {
	response = &GetIntlDomainDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
