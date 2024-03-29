package cloudapi

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

// DetachApiProduct invokes the cloudapi.DetachApiProduct API synchronously
func (client *Client) DetachApiProduct(request *DetachApiProductRequest) (response *DetachApiProductResponse, err error) {
	response = CreateDetachApiProductResponse()
	err = client.DoAction(request, response)
	return
}

// DetachApiProductWithChan invokes the cloudapi.DetachApiProduct API asynchronously
func (client *Client) DetachApiProductWithChan(request *DetachApiProductRequest) (<-chan *DetachApiProductResponse, <-chan error) {
	responseChan := make(chan *DetachApiProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachApiProduct(request)
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

// DetachApiProductWithCallback invokes the cloudapi.DetachApiProduct API asynchronously
func (client *Client) DetachApiProductWithCallback(request *DetachApiProductRequest, callback func(response *DetachApiProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachApiProductResponse
		var err error
		defer close(result)
		response, err = client.DetachApiProduct(request)
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

// DetachApiProductRequest is the request struct for api DetachApiProduct
type DetachApiProductRequest struct {
	*requests.RpcRequest
	Apis          *[]DetachApiProductApis `position:"Query" name:"Apis"  type:"Repeated"`
	SecurityToken string                  `position:"Query" name:"SecurityToken"`
	ApiProductId  string                  `position:"Query" name:"ApiProductId"`
}

// DetachApiProductApis is a repeated param struct in DetachApiProductRequest
type DetachApiProductApis struct {
	StageName string `name:"StageName"`
	ApiId     string `name:"ApiId"`
}

// DetachApiProductResponse is the response struct for api DetachApiProduct
type DetachApiProductResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachApiProductRequest creates a request to invoke DetachApiProduct API
func CreateDetachApiProductRequest() (request *DetachApiProductRequest) {
	request = &DetachApiProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DetachApiProduct", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachApiProductResponse creates a response to parse from DetachApiProduct response
func CreateDetachApiProductResponse() (response *DetachApiProductResponse) {
	response = &DetachApiProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
