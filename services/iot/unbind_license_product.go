package iot

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

// UnbindLicenseProduct invokes the iot.UnbindLicenseProduct API synchronously
func (client *Client) UnbindLicenseProduct(request *UnbindLicenseProductRequest) (response *UnbindLicenseProductResponse, err error) {
	response = CreateUnbindLicenseProductResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindLicenseProductWithChan invokes the iot.UnbindLicenseProduct API asynchronously
func (client *Client) UnbindLicenseProductWithChan(request *UnbindLicenseProductRequest) (<-chan *UnbindLicenseProductResponse, <-chan error) {
	responseChan := make(chan *UnbindLicenseProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindLicenseProduct(request)
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

// UnbindLicenseProductWithCallback invokes the iot.UnbindLicenseProduct API asynchronously
func (client *Client) UnbindLicenseProductWithCallback(request *UnbindLicenseProductRequest, callback func(response *UnbindLicenseProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindLicenseProductResponse
		var err error
		defer close(result)
		response, err = client.UnbindLicenseProduct(request)
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

// UnbindLicenseProductRequest is the request struct for api UnbindLicenseProduct
type UnbindLicenseProductRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	LicenseCode   string `position:"Query" name:"LicenseCode"`
}

// UnbindLicenseProductResponse is the response struct for api UnbindLicenseProduct
type UnbindLicenseProductResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         bool   `json:"Data" xml:"Data"`
}

// CreateUnbindLicenseProductRequest creates a request to invoke UnbindLicenseProduct API
func CreateUnbindLicenseProductRequest() (request *UnbindLicenseProductRequest) {
	request = &UnbindLicenseProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UnbindLicenseProduct", "", "")
	request.Method = requests.POST
	return
}

// CreateUnbindLicenseProductResponse creates a response to parse from UnbindLicenseProduct response
func CreateUnbindLicenseProductResponse() (response *UnbindLicenseProductResponse) {
	response = &UnbindLicenseProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
