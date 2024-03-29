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

// ClearDeviceDesiredProperty invokes the iot.ClearDeviceDesiredProperty API synchronously
func (client *Client) ClearDeviceDesiredProperty(request *ClearDeviceDesiredPropertyRequest) (response *ClearDeviceDesiredPropertyResponse, err error) {
	response = CreateClearDeviceDesiredPropertyResponse()
	err = client.DoAction(request, response)
	return
}

// ClearDeviceDesiredPropertyWithChan invokes the iot.ClearDeviceDesiredProperty API asynchronously
func (client *Client) ClearDeviceDesiredPropertyWithChan(request *ClearDeviceDesiredPropertyRequest) (<-chan *ClearDeviceDesiredPropertyResponse, <-chan error) {
	responseChan := make(chan *ClearDeviceDesiredPropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClearDeviceDesiredProperty(request)
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

// ClearDeviceDesiredPropertyWithCallback invokes the iot.ClearDeviceDesiredProperty API asynchronously
func (client *Client) ClearDeviceDesiredPropertyWithCallback(request *ClearDeviceDesiredPropertyRequest, callback func(response *ClearDeviceDesiredPropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClearDeviceDesiredPropertyResponse
		var err error
		defer close(result)
		response, err = client.ClearDeviceDesiredProperty(request)
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

// ClearDeviceDesiredPropertyRequest is the request struct for api ClearDeviceDesiredProperty
type ClearDeviceDesiredPropertyRequest struct {
	*requests.RpcRequest
	RealTenantId      string    `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string    `position:"Query" name:"RealTripartiteKey"`
	IotId             string    `position:"Query" name:"IotId"`
	IotInstanceId     string    `position:"Query" name:"IotInstanceId"`
	Identifies        *[]string `position:"Body" name:"Identifies"  type:"Repeated"`
	ProductKey        string    `position:"Query" name:"ProductKey"`
	ApiProduct        string    `position:"Body" name:"ApiProduct"`
	ApiRevision       string    `position:"Body" name:"ApiRevision"`
	DeviceName        string    `position:"Query" name:"DeviceName"`
}

// ClearDeviceDesiredPropertyResponse is the response struct for api ClearDeviceDesiredProperty
type ClearDeviceDesiredPropertyResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateClearDeviceDesiredPropertyRequest creates a request to invoke ClearDeviceDesiredProperty API
func CreateClearDeviceDesiredPropertyRequest() (request *ClearDeviceDesiredPropertyRequest) {
	request = &ClearDeviceDesiredPropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ClearDeviceDesiredProperty", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateClearDeviceDesiredPropertyResponse creates a response to parse from ClearDeviceDesiredProperty response
func CreateClearDeviceDesiredPropertyResponse() (response *ClearDeviceDesiredPropertyResponse) {
	response = &ClearDeviceDesiredPropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
