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

// BatchCheckImportDevice invokes the iot.BatchCheckImportDevice API synchronously
func (client *Client) BatchCheckImportDevice(request *BatchCheckImportDeviceRequest) (response *BatchCheckImportDeviceResponse, err error) {
	response = CreateBatchCheckImportDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// BatchCheckImportDeviceWithChan invokes the iot.BatchCheckImportDevice API asynchronously
func (client *Client) BatchCheckImportDeviceWithChan(request *BatchCheckImportDeviceRequest) (<-chan *BatchCheckImportDeviceResponse, <-chan error) {
	responseChan := make(chan *BatchCheckImportDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchCheckImportDevice(request)
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

// BatchCheckImportDeviceWithCallback invokes the iot.BatchCheckImportDevice API asynchronously
func (client *Client) BatchCheckImportDeviceWithCallback(request *BatchCheckImportDeviceRequest, callback func(response *BatchCheckImportDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchCheckImportDeviceResponse
		var err error
		defer close(result)
		response, err = client.BatchCheckImportDevice(request)
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

// BatchCheckImportDeviceRequest is the request struct for api BatchCheckImportDevice
type BatchCheckImportDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string                              `position:"Query" name:"IotInstanceId"`
	ProductKey    string                              `position:"Query" name:"ProductKey"`
	DeviceList    *[]BatchCheckImportDeviceDeviceList `position:"Query" name:"DeviceList"  type:"Repeated"`
	ApiProduct    string                              `position:"Body" name:"ApiProduct"`
	ApiRevision   string                              `position:"Body" name:"ApiRevision"`
}

// BatchCheckImportDeviceDeviceList is a repeated param struct in BatchCheckImportDeviceRequest
type BatchCheckImportDeviceDeviceList struct {
	DeviceSecret string `name:"DeviceSecret"`
	DeviceName   string `name:"DeviceName"`
	Sn           string `name:"Sn"`
}

// BatchCheckImportDeviceResponse is the response struct for api BatchCheckImportDevice
type BatchCheckImportDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	Code         string                       `json:"Code" xml:"Code"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInBatchCheckImportDevice `json:"Data" xml:"Data"`
}

// CreateBatchCheckImportDeviceRequest creates a request to invoke BatchCheckImportDevice API
func CreateBatchCheckImportDeviceRequest() (request *BatchCheckImportDeviceRequest) {
	request = &BatchCheckImportDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchCheckImportDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchCheckImportDeviceResponse creates a response to parse from BatchCheckImportDevice response
func CreateBatchCheckImportDeviceResponse() (response *BatchCheckImportDeviceResponse) {
	response = &BatchCheckImportDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
