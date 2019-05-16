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

package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteDevice invokes the iot.DeleteDevice API synchronously
// api document: https://help.aliyun.com/api/iot/deletedevice.html
func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
	response = CreateDeleteDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeviceWithChan invokes the iot.DeleteDevice API asynchronously
// api document: https://help.aliyun.com/api/iot/deletedevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDeviceWithChan(request *DeleteDeviceRequest) (<-chan *DeleteDeviceResponse, <-chan error) {
	responseChan := make(chan *DeleteDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDevice(request)
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

// DeleteDeviceWithCallback invokes the iot.DeleteDevice API asynchronously
// api document: https://help.aliyun.com/api/iot/deletedevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDeviceWithCallback(request *DeleteDeviceRequest, callback func(response *DeleteDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDeviceResponse
		var err error
		defer close(result)
		response, err = client.DeleteDevice(request)
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

// DeleteDeviceRequest is the request struct for api DeleteDevice
type DeleteDeviceRequest struct {
	*requests.RpcRequest
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	IotId         string `position:"Query" name:"IotId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// DeleteDeviceResponse is the response struct for api DeleteDevice
type DeleteDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteDeviceRequest creates a request to invoke DeleteDevice API
func CreateDeleteDeviceRequest() (request *DeleteDeviceRequest) {
	request = &DeleteDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteDevice", "iot", "openAPI")
	return
}

// CreateDeleteDeviceResponse creates a response to parse from DeleteDevice response
func CreateDeleteDeviceResponse() (response *DeleteDeviceResponse) {
	response = &DeleteDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
