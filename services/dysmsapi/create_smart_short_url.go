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

// CreateSmartShortUrl invokes the dysmsapi.CreateSmartShortUrl API synchronously
func (client *Client) CreateSmartShortUrl(request *CreateSmartShortUrlRequest) (response *CreateSmartShortUrlResponse, err error) {
	response = CreateCreateSmartShortUrlResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSmartShortUrlWithChan invokes the dysmsapi.CreateSmartShortUrl API asynchronously
func (client *Client) CreateSmartShortUrlWithChan(request *CreateSmartShortUrlRequest) (<-chan *CreateSmartShortUrlResponse, <-chan error) {
	responseChan := make(chan *CreateSmartShortUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSmartShortUrl(request)
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

// CreateSmartShortUrlWithCallback invokes the dysmsapi.CreateSmartShortUrl API asynchronously
func (client *Client) CreateSmartShortUrlWithCallback(request *CreateSmartShortUrlRequest, callback func(response *CreateSmartShortUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSmartShortUrlResponse
		var err error
		defer close(result)
		response, err = client.CreateSmartShortUrl(request)
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

// CreateSmartShortUrlRequest is the request struct for api CreateSmartShortUrl
type CreateSmartShortUrlRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	RouteName            string           `position:"Query" name:"RouteName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PhoneNumbers         string           `position:"Query" name:"PhoneNumbers"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SourceUrl            string           `position:"Query" name:"SourceUrl"`
	OutId                string           `position:"Query" name:"OutId"`
}

// CreateSmartShortUrlResponse is the response struct for api CreateSmartShortUrl
type CreateSmartShortUrlResponse struct {
	*responses.BaseResponse
	Message   string      `json:"Message" xml:"Message"`
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Code      string      `json:"Code" xml:"Code"`
	Model     []ModelItem `json:"Model" xml:"Model"`
}

// CreateCreateSmartShortUrlRequest creates a request to invoke CreateSmartShortUrl API
func CreateCreateSmartShortUrlRequest() (request *CreateSmartShortUrlRequest) {
	request = &CreateSmartShortUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "CreateSmartShortUrl", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSmartShortUrlResponse creates a response to parse from CreateSmartShortUrl response
func CreateCreateSmartShortUrlResponse() (response *CreateSmartShortUrlResponse) {
	response = &CreateSmartShortUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
