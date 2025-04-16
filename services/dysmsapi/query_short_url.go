package dysmsapi

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

// QueryShortUrl invokes the dysmsapi.QueryShortUrl API synchronously
func (client *Client) QueryShortUrl(request *QueryShortUrlRequest) (response *QueryShortUrlResponse, err error) {
	response = CreateQueryShortUrlResponse()
	err = client.DoAction(request, response)
	return
}

// QueryShortUrlWithChan invokes the dysmsapi.QueryShortUrl API asynchronously
func (client *Client) QueryShortUrlWithChan(request *QueryShortUrlRequest) (<-chan *QueryShortUrlResponse, <-chan error) {
	responseChan := make(chan *QueryShortUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryShortUrl(request)
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

// QueryShortUrlWithCallback invokes the dysmsapi.QueryShortUrl API asynchronously
func (client *Client) QueryShortUrlWithCallback(request *QueryShortUrlRequest, callback func(response *QueryShortUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryShortUrlResponse
		var err error
		defer close(result)
		response, err = client.QueryShortUrl(request)
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

// QueryShortUrlRequest is the request struct for api QueryShortUrl
type QueryShortUrlRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ShortUrl             string           `position:"Body" name:"ShortUrl"`
	ProdCode             string           `position:"Body" name:"ProdCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryShortUrlResponse is the response struct for api QueryShortUrl
type QueryShortUrlResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryShortUrlRequest creates a request to invoke QueryShortUrl API
func CreateQueryShortUrlRequest() (request *QueryShortUrlRequest) {
	request = &QueryShortUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "QueryShortUrl", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryShortUrlResponse creates a response to parse from QueryShortUrl response
func CreateQueryShortUrlResponse() (response *QueryShortUrlResponse) {
	response = &QueryShortUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
